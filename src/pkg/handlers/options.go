package handlers

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
)

func Options(s *discordgo.Session, i *discordgo.InteractionCreate) {
	// Access options in the order provided by the user.
	options := i.ApplicationCommandData().Options

	// Or convert the slice into a map
	optionMap := make(map[string]*discordgo.ApplicationCommandInteractionDataOption, len(options))
	for _, opt := range options {
		optionMap[opt.Name] = opt
	}

	// This example stores the provided arguments in an []interface{}
	// which will be used to format the bot's response
	margs := make([]interface{}, 0, len(options))
	msgformat := "You learned how to use command options! " +
		"Take a look at the value(s) you entered:\n"

	// Get the value from the option map.
	// When the option exists, ok = true
	if option, ok := optionMap["string-option"]; ok {
		// Option values must be type asserted from interface{}.
		// Discordgo provides utility functions to make this simple.
		margs = append(margs, option.StringValue())
		msgformat += "> string-option: %s\n"
	}

	if opt, ok := optionMap["integer-option"]; ok {
		margs = append(margs, opt.IntValue())
		msgformat += "> integer-option: %d\n"
	}

	if opt, ok := optionMap["number-option"]; ok {
		margs = append(margs, opt.FloatValue())
		msgformat += "> number-option: %f\n"
	}

	if opt, ok := optionMap["bool-option"]; ok {
		margs = append(margs, opt.BoolValue())
		msgformat += "> bool-option: %v\n"
	}

	if opt, ok := optionMap["channel-option"]; ok {
		margs = append(margs, opt.ChannelValue(nil).ID)
		msgformat += "> channel-option: <#%s>\n"
	}

	if opt, ok := optionMap["user-option"]; ok {
		margs = append(margs, opt.UserValue(nil).ID)
		msgformat += "> user-option: <@%s>\n"
	}

	if opt, ok := optionMap["role-option"]; ok {
		margs = append(margs, opt.RoleValue(nil, "").ID)
		msgformat += "> role-option: <@&%s>\n"
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		// Ignore type for now, they will be discussed in "responses"
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: fmt.Sprintf(
				msgformat,
				margs...,
			),
		},
	})
}
