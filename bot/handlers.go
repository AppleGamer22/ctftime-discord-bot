package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/MonSec/ctftime-discord-bot/api"
	"github.com/bwmarrin/discordgo"
)

func CommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}
	log.Println("New Query")
	command := interaction.ApplicationCommandData()
	switch command.Name {
	case TeamCommand.Name:
		TeamCommandHandler(session, interaction)
	default:
		err := session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "I don't know this command, but I do know this: MonSec > MISC",
			},
		})
		if err != nil {
			log.Println(err)
		}
	}
}

func TeamCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	command := interaction.ApplicationCommandData()
	option := command.Options[0]
	team, err := api.TeamInfo(option.StringValue())
	if err != nil {
		errorEmbed := discordgo.MessageEmbed{
			Color:       0xff0000,
			Title:       "Error",
			Description: err.Error(),
		}
		err = session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Embeds: []*discordgo.MessageEmbed{
					&errorEmbed,
				},
			},
		})
		if err != nil {
			log.Println(err)
		}
	}
	yearString := fmt.Sprintf("%d", time.Now().Year())
	ratingString := fmt.Sprintf("%d", team.Rating[yearString].CountryPlace)
	teamEmbed := discordgo.MessageEmbed{
		Title: "Team Information",
		Image: &discordgo.MessageEmbedImage{
			URL:    team.Logo,
			Width:  128,
			Height: 128,
		},
		Fields: []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Team Name",
				Value: fmt.Sprintf("%s (%d)", team.Name, team.ID),
			},
			&discordgo.MessageEmbedField{
				Name: "Team Type",
				Value: func() string {
					if team.Academic {
						return "Academic"
					} else {
						return "Non-academic"
					}
				}(),
			},
			&discordgo.MessageEmbedField{
				Name:  fmt.Sprintf("Lastest Rating During %s (%s)", yearString, team.Country),
				Value: ratingString,
			},
		},
	}
	err = session.InteractionRespond(interaction.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf("/team %s", option.StringValue()),
			Embeds: []*discordgo.MessageEmbed{
				&teamEmbed,
			},
		},
	})
	if err != nil {
		log.Println(err)
	}
}

func UpcomingEventsCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	// command := interaction.ApplicationCommandData()
}
