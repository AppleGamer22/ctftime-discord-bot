package bot

import (
	"fmt"
	"log"
	"time"

	"github.com/MonSec/ctftime-discord-bot/api"
	"github.com/bwmarrin/discordgo"
)

func TeamCommandHandler(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	if interaction.Type != discordgo.InteractionApplicationCommand {
		return
	}
	log.Println("Hi")
	command := interaction.ApplicationCommandData()
	switch command.Name {
	case TeamCommand.Name:
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
					Value: team.Name,
				},
				&discordgo.MessageEmbedField{
					Name:  "Team Country",
					Value: team.Country,
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
