package content

import (
	"github.com/jonathonharrell/miri-ws-server/app/core/game"
)

var Races = map[string]game.Race{
	"HUMAN": game.Race{
		Name: "Human",
		ID:   "HUMAN",
		LongDescription: "<p><strong>Distinct Groups:</strong> The Eldorei, The Fox-Ear Clan, The Kai-Ara, and The Briree</p>" +
			"<p><strong>Appearance: </strong>Varying in size, skin and hair color, from slender to fat, brown-haired to blonde, blue-eyed to grey.</p>" +
			"<p><strong>Average Height: </strong>Between 5'2 and 6'4 on average</p>" +
			"<p><strong>Average Weight: </strong>Between 110 and 250lbs on average</p>" +
			"<p><strong>Average Lifespan: </strong>Between 50 and 100 years</p>" +
			"<p><strong>Personality: </strong>Varied, but commonly adaptable, ambitious, or self-centered.</p>" +
			"<p><strong>History: </strong>Humans have, since the days of Fae, been masters of conquering and migrating. Their civilizations are among the largest in the Miri, and their number is several times as great as that of the other races. They are masters of trade and diplomacy, also having a large number of allies. Whether that is due to their personable natures one cannot be certain. Whatever drives them, humans endeavor to succeed.</p>",
		Description: "The most diverse race in The Miri.",
	},
	"ELF": game.Race{
		Name:        "Elf",
		ID:          "ELF",
		Description: "A graceful people, typically well-mannered and empathetic.",
		LongDescription: "<p><strong>Distinct Groups:</strong> Dark Elf, Wood Elf, Blood Elf, High Elf</p>" +
			"<p><strong>Appearance: </strong>Slender and graceful, with fine features and pointed ears.</p>" +
			"<p><strong>Average Height: </strong>Between 4'10 and 6'0 on average</p>" +
			"<p><strong>Average Weight: </strong>Between 90 and 160lbs on average</p>" +
			"<p><strong>Average Lifespan: </strong>Between 550 and 650 years</p>" +
			"<p><strong>Personality: </strong>Varied, but commonly curious, polite, or haughty.</p>" +
			"<p><strong>History: </strong></p>",
	},
	"DWARF": game.Race{
		Name:        "Dwarf",
		ID:          "DWARF",
		Description: "Stout, stoic and enduring, like the mountains they adore.",
		LongDescription: "<p><strong>Distinct Groups:</strong> Hill Dwarf, Mountain Dwarf</p>" +
			"<p><strong>Appearance: </strong>Short, stout. Often have long, intricate or decorated beards</p>" +
			"<p><strong>Average Height: </strong>Between 4'2 and 4'9</p>" +
			"<p><strong>Average Weight: </strong>Between 150 and 250lbs</p>" +
			"<p><strong>Average Lifespan: </strong>Between 200 and 450 years</p>" +
			"<p><strong>Personality: </strong>Varied, but commonly proud, determined, loyal, decisive, or stubborn.</p>" +
			"<p><strong>History: </strong>Dwarves have been around for as long as the Miri itself, or at least as far back as the historians can see. They've played a part in each of the major wars, in fact some of the finest warriors in folklore were dwarves - especially according to the dwarven archives. Many have staked their claim in the Miri as artisans, something dwarves have proven to be especially adept at. Mountain Dwarves are known for their amazing weapon and armor smiths, while Hill Dwarves are known for their cuisine.</p>",
	},
}

// @todo FUTURE RACES
//
// "name":"Spriggan",
// "description":"[WIP]",
// "id":"SPRIGGAN",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Gnome",
// "description":"[WIP]",
// "id":"GNOME",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Catlike",
// "description":"[WIP]",
// "id":"CATLIKE",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Dragonborn",
// "description":"[WIP]",
// "id":"DRAGONBORN",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Minotaur",
// "description":"[WIP]",
// "id":"MINOTAUR",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Orc",
// "description":"[WIP]",
// "id":"ORC",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Half-Orc",
// "description":"[WIP]",
// "id":"HALFORC",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Half-Giant",
// "description":"[WIP]",
// "id":"HALFGIANT",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Goblin",
// "description":"[WIP]",
// "id":"GOBLIN",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Centaur",
// "description":"[WIP]",
// "id":"CENTAUR",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Half-Angel",
// "description":"[WIP]",
// "id":"HALFANGEL",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Tiefling",
// "description":"[WIP]",
// "id":"TIEFLING",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Half-Elf",
// "description":"[WIP]",
// "id":"HALFELF",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Automaton",
// "description":"[WIP]",
// "id":"AUTOMATON",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Satyr",
// "description":"[WIP]",
// "id":"SATYR",
// "long_description": "<p>This is a work in progress.</p>"
//
// "name":"Eutara",
// "description":"Hospitable ocean-dwellers",
// "id":"EUTARA",
// "long_description": "<p>A lot more about dwarves mother fuckers. We can put <strong>HTML</strong> in here.</p>"
//
// "name":"Halfling",
// "description":"Peaceful and practical. Halflings ",
// "id":"HALFLING",
// "long_description": "<p>This is a work in progress.</p>"
//
