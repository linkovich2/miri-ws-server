package entities

import (
	"github.com/jonathonharrell/miri-ws-server/app/core"
)

var KraJaficMerchantCrowd = &core.ComponentBag{
	Name:      "Crowd of Merchants/Traders",
	Behaviors: []core.Behavior{core.CrowdConversationBehavior{}},
	Properties: []*core.Property{
		&core.Property{"conversation", "You overhear a disgruntled customer, \"<say>You cannot be serious! These prices are outrageous.</say>\";;The trader replies, \"<say>Bah, what would you know about furs as fine as these. The price is the price!</say>\""},
		&core.Property{"conversation", "A surly, stout man says to another trader, \"<say>The fabrics Al'tur sold us will fetch double in Kel Thayis</say>\";;The trader replies, \"<say>Yes, they are an excellent make. Perhaps I will create a robe from them for myself...</say>\";;The stout man responds, \"<say>Don't touch! You can't go digging into our supply like you did with the <em>mercer leaf</em>.</say>\";;The trader gasps, \"<say>Not so loud... There are people listening...</say>\" He looks around nervously."},
	},
}

var KraJaficAmbiance = &core.ComponentBag{
	Name:       "Ambiance",
	Behaviors:  []core.Behavior{core.AmbianceBehavior{}},
	NotVisible: true,
	Properties: []*core.Property{
		&core.Property{"ambiance", "<ambiance><em>The sands stir in the wind...</em></ambiance>"},
		&core.Property{"ambiance", "<ambiance><em>You can hear the tents flapping in the wind.</em></ambiance>"},
		&core.Property{"ambiance", "<ambiance><em>The sounds of haggling merchants surround you.</em></ambiance>"},
		&core.Property{"ambiance", "<ambiance><em>The sun beats down overhead...</em></ambiance>"},
	},
}
