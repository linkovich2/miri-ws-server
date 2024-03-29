package characters

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	"github.com/jonathonharrell/miri-ws-server/app/logger"
	"regexp"
	"strconv"
	"strings"
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	characterController struct{}
	characterResponse   struct {
		Character   core.Character `json:"character"`
		Location    string         `json:"location"`
		Description string         `json:"description"`
		Realm       string         `json:"realm"`
	}
	optionsParams struct {
		Get string `json:"get"`
	}
	deleteParams struct {
		Id bson.ObjectId `json:"id"`
	}
	selectParams   deleteParams
	createResponse struct {
		Success bool     `json:"success"`
		Errors  []string `json:"errors"`
	}
)

var Controller = characterController{}

func (c *characterController) List(connection *game.Connection, g *game.Game, args *json.RawMessage) {
	session, dbName := database.GetSession() // connect
	db := session.DB(dbName)
	defer session.Close()

	var characters []core.Character
	var response []characterResponse
	_ = db.C("characters").Find(bson.M{"user_id": connection.Socket.User.ID}).All(&characters)
	for _, c := range characters {
		response = append(response, characterResponse{
			Character:   c,
			Description: c.Describe(),
			Location:    g.LocationNameForCharacter(&c),
			Realm:       g.RealmNameForCharacter(&c),
		})
	}

	res, _ := json.Marshal(response)
	connection.Socket.Send(res)
}

func (c *characterController) Delete(connection *game.Connection, g *game.Game, args *json.RawMessage) {
	params := deleteParams{}
	err := json.Unmarshal(*args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	session, dbName := database.GetSession() // connect
	db := session.DB(dbName)
	defer session.Close()

	character := &core.Character{}
	err = db.C("characters").Find(bson.M{"_id": params.Id, "user_id": connection.Socket.User.ID}).One(&character)
	if err != nil {
		logger.Write.Error("Connection [%s] tried to delete a character that either didn't belong to them or doesn't exist.", connection.Socket.ID)
		res, _ := json.Marshal(createResponse{false, []string{"That character either doesn't belong to you, or doesn't exist."}})
		connection.Socket.Send(res)
		return
	}

	if g.CurrentlyPlaying(character) {
		logger.Write.Error("Connection [%s] tried to select a character that was already in-game ([%s]).", connection.Socket.ID, character.Name)
		res, _ := json.Marshal(createResponse{false, []string{"This character is currently in-game."}})
		connection.Socket.Send(res)
		return
	}

	err = db.C("characters").Remove(bson.M{"_id": params.Id, "user_id": connection.Socket.User.ID})
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	res, _ := json.Marshal(createResponse{true, []string{}})
	connection.Socket.Send(res)
}

func (c *characterController) Select(connection *game.Connection, g *game.Game, args *json.RawMessage) {
	params := selectParams{}
	err := json.Unmarshal(*args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	session, dbName := database.GetSession() // connect
	db := session.DB(dbName)
	defer session.Close()

	character := &core.Character{}
	err = db.C("characters").Find(bson.M{"_id": params.Id, "user_id": connection.Socket.User.ID}).One(&character)
	if err != nil {
		logger.Write.Error("Connection [%s] tried to select a character that either didn't belong to them or doesn't exist.", connection.Socket.ID)
		res, _ := json.Marshal(createResponse{false, []string{"That character either doesn't belong to you, or doesn't exist."}})
		connection.Socket.Send(res)
		return
	}

	if g.CurrentlyPlaying(character) {
		logger.Write.Error("Connection [%s] tried to select a character that was already in-game ([%s]).", connection.Socket.ID, character.Name)
		res, _ := json.Marshal(createResponse{false, []string{"This character is currently in-game."}})
		connection.Socket.Send(res)
		return
	}

	connection.Character = character
	g.Connect <- connection
	res, _ := json.Marshal(createResponse{true, []string{}})
	connection.Socket.Send(res)
}

func (c *characterController) Create(connection *game.Connection, g *game.Game, args *json.RawMessage) {
	character := core.Character{}
	err := json.Unmarshal(*args, &character)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	session, dbName := database.GetSession() // connect
	db := session.DB(dbName)
	defer session.Close()

	var characters []core.Character
	_ = db.C("characters").Find(bson.M{"user_id": connection.Socket.User.ID}).All(&characters)

	// greater then 3 characters exist for this account, if it's not an admin user it should not be created
	if len(characters) >= 3 && !connection.Socket.User.IsAdmin() {
		logger.Write.Error("Character could not be saved for Connection [%s], character limit exceeded", connection.Socket.ID)
		handleCharacterCreationError(connection)
		return // stop execution here
	}

	if !validateCharacter(connection, &character) {
		handleCharacterCreationError(connection)
		return // stop execution
	}

	bg := core.GetBackground(character.Background)

	character.UserID = connection.Socket.User.ID
	character.Created = time.Now() // timestamp this bad boy
	character.Position = bg.StartPosition
	character.Realm = bg.StartRealm
	_ = db.C("characters").Insert(&character)

	res, _ := json.Marshal(createResponse{true, []string{}})
	connection.Socket.Send(res)
}

func (c *characterController) Options(connection *game.Connection, g *game.Game, args *json.RawMessage) {
	var body interface{}
	params := optionsParams{}
	err := json.Unmarshal(*args, &params)
	if err != nil {
		logger.Write.Error(err.Error())
		return
	}

	switch params.Get {
	case "races":
		body = core.GetRaces()
	case "genders":
		body = core.GetGenders()
	case "aesthetic_traits":
		body = core.GetAestheticTraits()
	case "functional_traits":
		body = core.GetFunctionalTraits()
	case "backgrounds":
		body = core.GetBackgrounds()
	default:
		body = core.GetRaces()
	}

	res, _ := json.Marshal(body)
	connection.Socket.Send(res)
}

func handleCharacterCreationError(connection *game.Connection) {
	errors, _ := json.Marshal(createResponse{false, []string{"Something went wrong with character creation."}})
	connection.Socket.Send(errors)
}

func validateCharacter(connection *game.Connection, character *core.Character) bool {
	if !validateRace(connection, character) {
		return false
	}
	if !validateGender(connection, character) {
		return false
	}
	if !validateAestheticTraits(connection, character) {
		return false
	}
	if !validateFunctionalTraits(connection, character) {
		return false
	}
	if !validateBackground(connection, character) {
		return false
	}
	if !validateName(connection, character) {
		return false
	}
	return true
}

func validateRace(connection *game.Connection, character *core.Character) bool {
	races := core.GetRaces()
	if _, raceExists := races[character.Race]; raceExists {
		return true
	}

	logger.Write.Error("Character Creation Error (Connection [%s]): Provided Race [%s] does not exist.", connection.Socket.ID, character.Race)
	return false
}

func validateGender(connection *game.Connection, character *core.Character) bool {
	genders := core.GetGenders()
	if gender, genderExists := genders[character.Gender]; genderExists {
		if gender.RaceAllowed(character.Race) {
			return true
		} else {
			logger.Write.Error(
				"Character Creation Error (Connection [%s]): Gender [%s] not allowed for Race [%s].",
				connection.Socket.ID,
				character.Gender,
				character.Race,
			)
			return false
		}
	}

	logger.Write.Error("Character Creation Error (Connection [%s]): Gender [%s] doesn't exist.", connection.Socket.ID, character.Gender)
	return false
}

func validateAestheticTraits(connection *game.Connection, character *core.Character) bool {
	list := core.GetAestheticTraits()
	for key, traits := range character.AestheticTraits {
		if _, categoryExists := list[key]; !categoryExists { // trait category does not exist
			logger.Write.Error("Character Creation Error (Connection [%s]): Unknown Trait Category [%s].", connection.Socket.ID, key)
			return false
		}
		category := list[key]

		if !category.IsAllowedForCharacter(character) && len(traits) > 0 { // not allowed for character and one exists
			logger.Write.Error("Character Creation Error (Connection [%s]): Trait Category [%s] not allowed for character.", connection.Socket.ID, key)
			return false
		}

		if len(traits) > 1 && category.Unique { // character contains more then one trait in a unique category
			logger.Write.Error("Character Creation Error (Connection [%s]): Trait Category [%s] is unique.", connection.Socket.ID, key)
			return false
		}

		for _, trait := range traits {
			if _, traitExists := category.Traits[trait]; !traitExists { // trait itself does not exist
				logger.Write.Error("Character Creation Error (Connection [%s]): Trait [%s] does not exist.", connection.Socket.ID, trait)
				return false
			}

			t := category.Traits[trait]
			if !t.IsAllowedForCharacter(character) { // individual trait is not allowed for this character
				logger.Write.Error("Character Creation Error (Connection [%s]): Trait [%s] is not allowed for character.", connection.Socket.ID, trait)
				return false
			}
		}
	}

	// loop through all aesthetic trait categories and check if one exists where one is required
	for key, category := range list {
		if category.Minimum > 0 && category.IsAllowedForCharacter(character) {
			if len(character.AestheticTraits[key]) < category.Minimum {
				logger.Write.Error("Character Creation Error (Connection [%s]): Character doesn't have enough traits from Trait Category [%s]", connection.Socket.ID, key)
				return false
			}
		}
	}

	return true
}

func validateFunctionalTraits(connection *game.Connection, character *core.Character) bool {
	var points int

	list := core.GetFunctionalTraits()

	for key, traits := range character.FunctionalTraits {
		if _, categoryExists := list[key]; !categoryExists { // trait category does not exist
			logger.Write.Error("Character Creation Error (Connection [%s]): Unknown Trait Category [%s].", connection.Socket.ID, key)
			return false
		}
		category := list[key]

		if !category.IsAllowedForCharacter(character) && len(traits) > 0 { // not allowed for character and one exists
			logger.Write.Error("Character Creation Error (Connection [%s]): Trait Category [%s] not allowed for character.", connection.Socket.ID, key)
			return false
		}

		if len(traits) > 1 && category.Unique { // character contains more then one trait in a unique category
			logger.Write.Error("Character Creation Error (Connection [%s]): Trait Category [%s] is unique.", connection.Socket.ID, key)
			return false
		}

		for _, trait := range traits {
			if _, traitExists := category.Traits[trait]; !traitExists { // trait itself does not exist
				logger.Write.Error("Character Creation Error (Connection [%s]): Trait [%s] does not exist.", connection.Socket.ID, trait)
				return false
			}

			t := category.Traits[trait]
			pointValue, err := strconv.Atoi(t.Points)

			if err != nil {
				logger.Write.Error("Weird conversion for point value for functional trait [%s]", t.ID)
				return false
			}

			points = points + pointValue

			if !t.IsAllowedForCharacter(character) { // individual trait is not allowed for this character
				logger.Write.Error("Character Creation Error (Connection [%s]): Trait [%s] is not allowed for character.", connection.Socket.ID, trait)
				return false
			}
		}
	}

	if points < 0 {
		logger.Write.Error("Character Creation Error (Connection [%s]): Point deficit too high.", connection.Socket.ID)
		return false
	}

	// loop through all aesthetic trait categories and check if one exists where one is required
	for key, category := range list {
		if !category.IsAllowedForCharacter(character) {
			continue
		}

		if category.Minimum > 0 {
			if len(character.FunctionalTraits[key]) < category.Minimum {
				logger.Write.Error("Character Creation Error (Connection [%s]): Character doesn't have enough traits from Trait Category [%s]", connection.Socket.ID, key)
				return false
			}
		}

		for id, trait := range category.Traits {
			if trait.Required && trait.IsAllowedForCharacter(character) {
				if _, categoryExists := character.FunctionalTraits[key]; !categoryExists {
					logger.Write.Error("Character Creation Error (Connection [%s]): Required Trait Category [%s] doesn't exist.", connection.Socket.ID, key)
					return false
				}

				exists := false
				for _, t := range character.FunctionalTraits[key] {
					if t == id {
						exists = true
						break
					}
				}

				if !exists { // trait is required but doesn't exist on the character object
					logger.Write.Error("Character Creation Error (Connection [%s]): Required Trait [%s] doesn't exist.", connection.Socket.ID, trait)
					return false
				}
			}
		}
	}

	return true
}

func validateBackground(connection *game.Connection, character *core.Character) bool {
	bgs := core.GetBackgrounds()
	if _, backgroundExists := bgs[character.Background]; !backgroundExists {
		logger.Write.Error("Character Creation Error (Connection [%s]): Provided Background [%s] doesn't exist.", connection.Socket.ID, character.Background)
		return false
	}

	b := bgs[character.Background]
	if !b.IsAllowedForCharacter(character) {
		logger.Write.Error("Character Creation Error (Connection [%s]): Background [%s] isn't allowed for character.", connection.Socket.ID, character.Background)
		return false
	}

	return true
}

func validateName(connection *game.Connection, character *core.Character) bool {
	if len(character.Name) < 5 { // name not long enough
		logger.Write.Error("Character Creation Error (Connection [%s]): Name [%s] isn't long enough.", connection.Socket.ID, character.Name)
		return false
	}

	splitName := strings.Split(character.Name, " ")
	if len(splitName) != 2 {
		logger.Write.Error("Character Creation Error (Connection [%s]): Name [%s] must have only one space (first and last name)", connection.Socket.ID, character.Name)
		return false
	}

	safe, err := regexp.Match(`^['a-zA-Z-\s]+$`, []byte(character.Name))
	if !safe || err != nil {
		logger.Write.Error("Character Creation Error (Connection [%s]): Name [%s] has invalid characters", connection.Socket.ID, character.Name)
		return false
	}

	if len(splitName[0]) < 2 {
		logger.Write.Error("Character Creation Error (Connection [%s]): First name must be two or more characters in Name [%s].", connection.Socket.ID, character.Name)
		return false
	}
	if len(splitName[1]) < 2 {
		logger.Write.Error("Character Creation Error (Connection [%s]): Last name must be two or more characters in Name [%s].", connection.Socket.ID, character.Name)
		return false
	}

	// @todo in the future we may want to also filter here through names we explicitly disallow
	return true
}
