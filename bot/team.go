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
			Type:        discordgo.ApplicationCommandOptionString,
			Choices: []*discordgo.ApplicationCommandOptionChoice{
				{
					Name:  "MonSec",
					Value: "34111",
				},
				{
					Name:  "MISC",
					Value: "34111", //109523
				},
			},
		},
	},
}

func TeamCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}
	data := interaction.ApplicationCommandData()
	option := data.Options[0]
	if option.Name != TeamCommand.Options[0].Name {

	}
}
