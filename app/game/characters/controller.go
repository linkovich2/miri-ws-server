package characters

import (
	"encoding/json"
	"github.com/jonathonharrell/miri-ws-server/app/content"
	"github.com/jonathonharrell/miri-ws-server/app/core"
	"strconv"
	// "github.com/jonathonharrell/miri-ws-server/app/database"
	"github.com/jonathonharrell/miri-ws-server/app/game"
	"github.com/jonathonharrell/miri-ws-server/app/logger"

	// "gopkg.in/mgo.v2/bson"
)

type (
	characterController struct{}
	optionsParams       struct {
		Get string `json:"get"`
	}
	deleteParams struct {
		Id string `json:"id"`
	}
	selectParams   deleteParams
	createResponse struct {
		Success bool     `json:"success"`
		Errors  []string `json:"errors"`
	}
)

var Controller = characterController{}

func (c *characterController) List(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	res, _ := json.Marshal(connection.Socket.User.Characters)
	connection.Socket.Send(res)
}

func (c *characterController) Delete(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	// @todo stub
}

func (c *characterController) Select(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	// @todo stub
}

func (c *characterController) Create(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	character := core.Character{}
	err := json.Unmarshal(*args, &character)
	if err != nil {
		logger.Write.Error(err.Error()) // @todo handle json malformed or something like that
	}

	// greater then 3 characters exist for this account, if it's not an admin user it should not be created
	if len(connection.Socket.User.Characters) >= 3 && !connection.Socket.User.IsAdmin() {
		logger.Write.Error("Character could not be saved for Connection [%s], character limit exceeded", connection.Socket.ID)
		handleCharacterCreationError(connection)
		return // stop execution here
	}

	if !validateCharacter(connection, &character) {
		handleCharacterCreationError(connection)
		return // stop execution
	}

	// @todo save the character in the database

	res, _ := json.Marshal(createResponse{true, []string{}})
	connection.Socket.Send(res)
}

func (c *characterController) Options(connection *game.Connection, game *game.Game, args *json.RawMessage) {
	var body interface{}
	params := optionsParams{}
	err := json.Unmarshal(*args, &params)
	if err != nil {
		logger.Write.Error(err.Error()) // @todo handle json malformed or something like that
	}

	switch params.Get {
	case "races":
		body = content.Races
	case "genders":
		body = content.Genders
	case "aesthetic_traits":
		body = content.AestheticTraits
	case "functional_traits":
		body = content.FunctionalTraits
	case "backgrounds":
		body = content.Backgrounds
	default:
		body = content.Races
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
	if _, raceExists := content.Races[character.Race]; raceExists {
		return true
	}

	logger.Write.Error("Character Creation Error (Connection [%s]): Provided Race [%s] does not exist.", connection.Socket.ID, character.Race)
	return false
}

func validateGender(connection *game.Connection, character *core.Character) bool {
	if gender, genderExists := content.Genders[character.Gender]; genderExists {
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
	for key, traits := range character.AestheticTraits {
		if _, categoryExists := content.AestheticTraits[key]; !categoryExists { // trait category does not exist
			logger.Write.Error("Character Creation Error (Connection [%s]): Unknown Trait Category [%s].", connection.Socket.ID, key)
			return false
		}
		category := content.AestheticTraits[key]

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
	for key, category := range content.AestheticTraits {
		if category.Minimum > 0 && category.IsAllowedForCharacter(character) {
			if len(character.AestheticTraits[key]) < category.Minimum {
				logger.Write.Error("Character Creation Error (Connection [%s]): Character doesn't have enough traits from Trait Category [%s]", connection.Socket.ID, key)
				return false
			}
		}
	}

	return true
}

func validateFunctionalTraits(connection *game.Connection, character *core.Character) (valid bool) {
	// @todo this should look the same as aesthetic trait validation except individual traits may be required within a category
	var points int

	for key, traits := range character.FunctionalTraits {
		if _, categoryExists := content.FunctionalTraits[key]; !categoryExists { // trait category does not exist
			logger.Write.Error("Character Creation Error (Connection [%s]): Unknown Trait Category [%s].", connection.Socket.ID, key)
			return false
		}
		category := content.FunctionalTraits[key]

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

	if points > 0 {
		logger.Write.Error("Character Creation Error (Connection [%s]): Point deficit too high.", connection.Socket.ID)
		return false
	}

	// loop through all aesthetic trait categories and check if one exists where one is required
	for key, category := range content.FunctionalTraits {
		if category.Minimum > 0 && category.IsAllowedForCharacter(character) {
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

func validateBackground(connection *game.Connection, character *core.Character) (valid bool) {
	// @todo all we need to do is check it exists and that it's allowed
	return
}

func validateName(connection *game.Connection, character *core.Character) (valid bool) {
	// @todo we'll do the same validation we do client side here, but we might also check for some disallowed names
	return
}
