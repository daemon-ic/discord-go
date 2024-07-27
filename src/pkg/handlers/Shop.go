package handlers

import (
	"log"

	"example/slash/src/pkg/banners"
	"example/slash/src/pkg/bot"
	"example/slash/src/pkg/db"
	"example/slash/src/shared"

	"github.com/bwmarrin/discordgo"
)

func Shop(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	log.Println("init shop command")
	shop := Shop_Instance{session, interaction}

	ggmDB, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer ggmDB.Close()

	player, err := shop.ValidatePlayer(ggmDB)
	if err != nil {
		bot.Send("No account found for Player '"+player.Discord_Username+"'", shop.S, shop.I)
		return
	}

	banners, err := banners.GetAll(ggmDB)
	if err != nil {
		log.Fatal(err)
	}

	shared.PrettyLogJSON(banners)

	shop.InitialDisplay(shared.Shop_Display{
		CurrentPage: 1,
		TotalPages:  10,
		ItemName:    "test 1",
		ImageUrl:    "https://levelonegameshop.com/cdn/shop/products/d743a834353ab2c41a5358680a1270e6_720x.png",
		ItemPrice:   1000,
	})
}

func NavigateShop(session *discordgo.Session, interaction *discordgo.InteractionCreate) {
	log.Println("moving next page of shop")
	shop := Shop_Instance{session, interaction}

	ggmDB, err := db.Conn()
	if err != nil {
		log.Fatal(err)
	}
	defer ggmDB.Close()

	player, err := shop.ValidatePlayer(ggmDB)
	if err != nil {
		log.Println("invalid player attempted to navigate " + player.Discord_Username)
		return
	}

	banners, err := banners.GetAll(ggmDB)
	if err != nil {
		log.Fatal(err)
	}

	shared.PrettyLogJSON(banners)

	shop.ChangeDisplay(shared.Shop_Display{
		CurrentPage: 5,
		TotalPages:  10,
		ItemName:    "test 2",
		ImageUrl:    "https://levelonegameshop.com/cdn/shop/products/d743a834353ab2c41a5358680a1270e6_720x.png",
		ItemPrice:   1000,
	})
}
