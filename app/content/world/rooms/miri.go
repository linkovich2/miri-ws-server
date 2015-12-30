package rooms

import (
	e "github.com/jonathonharrell/miri-ws-server/app/content/entities"
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var Miri = map[string]*core.Room{
	"1:0:1": &core.Room{
		Position:    core.Position{1, 0, 1},
		Name:        "Kra Jafic Traders' Outpost - Southern Bazaar",
		Description: "The edge of the bazaar - merchants come and go among the tents bartering goods and services. Bright orange and blue tents match the local dress. There is such a variety of smells from so many different locales - some pleasant, others foul. The chatter of trading crowds is oddly comforting.",
		Size:        core.RoomLarge,
		Details:     []string{"trading tents", "bazaar", "sand", "traders"},
		Entities:    []core.Entity{e.KraJaficMerchantCrowd.Copy(), e.KraJaficAmbiance.Copy()},
	},
	"0:1:1": &core.Room{
		Position:    core.Position{0, 1, 1},
		Name:        "Kra Jafic Traders' Outpost - Western Bazaar",
		Description: "The eastern area of this trading center seems quiet, but traders still come and go. There isn't a lot on display here - the trading must be going on inside the orange and blue tents.",
		Size:        core.RoomLarge,
		Details:     []string{"trading tents", "bazaar", "sand", "traders"},
		Entities:    []core.Entity{e.KraJaficMerchantCrowd.Copy(), e.KraJaficAmbiance.Copy()},
	},
	"1:1:1": &core.Room{
		Position:    core.Position{1, 1, 1},
		Name:        "Kra Jafic Traders' Outpost",
		Description: "Orange, blue and gold tents line the edges of this traders' hub on the northern rim of the Tulinne Desert. Traders seem to come and go at odd times. The air is hot and dry, and the wind stirs the sands.",
		Size:        core.RoomLarge,
		Details:     []string{"trading tents", "bazaar", "sand", "traders"},
		Entities:    []core.Entity{e.KraJaficMerchantCrowd.Copy(), e.KraJaficAmbiance.Copy()},
	},
	"1:2:1": &core.Room{
		Position:    core.Position{1, 2, 1},
		Name:        "Kra Jafic Traders' Outpost - Northern Edge",
		Description: "The tents of Kra Jafic are to the south, but at this edge of the outpost it is quiet. Here is where the stone road is no longer maintained, and the harsh desert surrounding this place seems endless.",
		Size:        core.RoomLarge,
		Details:     []string{"sand", "paved stone"},
	},
	"2:2:1": &core.Room{
		Position:    core.Position{2, 2, 1},
		Name:        "Kirenov Sands",
		Description: "To the north you can feintly make out the outline of the Kai Era mountains.",
		Size:        core.RoomMassive,
		Details:     []string{"sand"},
	},
	"2:1:1": &core.Room{
		Position:    core.Position{2, 1, 1},
		Name:        "Kirenov Sands",
		Description: "The sand seems to go on forever.",
		Size:        core.RoomMassive,
		Details:     []string{"sand"},
	},
	"1:-1:1": &core.Room{
		Position:    core.Position{1, -1, 1},
		Name:        "Kirenov Sands",
		Description: "This desert seems to go on forever.",
		Size:        core.RoomMassive,
		Details:     []string{"sand"},
	},
}
