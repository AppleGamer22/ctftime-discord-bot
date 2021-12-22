package bot

import (
	"github.com/bwmarrin/discordgo"
)

var TeamCommand = &discordgo.ApplicationCommand{
	Name:        "team",
	Type:        discordgo.ChatApplicationCommand,
	Description: "fetch information about a team",
	Options: []*discordgo.ApplicationCommandOption{
		{
			Name:        "id",
			Description: "CTFTime team ID",
			Required:    true,
			Type:        discordgo.ApplicationCommandOptionString,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "MonSec",
					Value: "34111",
				},
				{
					Name: "MISC",
					// Value: "34111",
					Value: "109523",
				},
				{
					Name:  "RISC",
					Value: "1510",
				},
			},
		},
	},
}
