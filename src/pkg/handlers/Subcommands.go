package handlers

import "github.com/bwmarrin/discordgo"

func Subcommands(s *discordgo.Session, i *discordgo.InteractionCreate) {
	options := i.ApplicationCommandData().Options
	content := ""

	// As you can see, names of subcommands (nested, top-level)
	// and subcommand groups are provided through the arguments.
	switch options[0].Name {
	case "subcommand":
		content = "The top-level subcommand is executed. Now try to execute the nested one."
	case "subcommand-group":
		options = options[0].Options
		switch options[0].Name {
		case "nested-subcommand":
			content = "Nice, now you know how to execute nested commands too"
		default:
			content = "Oops, something went wrong.\n" +
				"Hol' up, you aren't supposed to see this message."
		}
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
}
