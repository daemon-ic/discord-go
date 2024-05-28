package main

import (
	"example/slash/src/pkg/handlers"

	"github.com/bwmarrin/discordgo"
)

var Handlers = map[string]func(s *discordgo.Session, i *discordgo.InteractionCreate){
	"basic-command":            handlers.BasicCommand,
	"basic-command-with-files": handlers.BasicCommandWithFiles,
	"localized-command":        handlers.LocalizedCommand,
	"options":                  handlers.Options,
	"permission-overview":      handlers.PermissionOverview,
	"subcommands":              handlers.Subcommands,
	"responses":                handlers.Responses,
	"followups":                handlers.FollowUps,
}
