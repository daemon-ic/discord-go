package handlers

import (
	"example/slash/src/pkg/banners"
	"example/slash/src/pkg/bot"
	"example/slash/src/pkg/db"
	"example/slash/src/shared"
	"log"
	"strconv"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func getNewPageNum(shop Shop_Instance) int {
	content := shop.I.Message.Content
	pageNum, err := strconv.Atoi(strings.Split(content, "/")[0])
	if err != nil {
		log.Fatal(err)
	}
	customId := shop.I.MessageComponentData().CustomID
	switch customId {
	case "banner_next":
		pageNum += 1
	case "banner_prev":
		pageNum -= 1
	}
	return pageNum
}

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

	pageNum := 1

	shop.InitialDisplay(shared.Shop_Display{
		CurrentPage: pageNum,
		TotalPages:  len(banners),
		ItemName:    banners[0].Name + "_banner",
		ImageUrl:    banners[0].ImageUrl,
		ItemPrice:   1000,
		PrevDisable: true,
		NextDisable: pageNum == len(banners),
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

	pageNum := getNewPageNum(shop)

	shop.ChangeDisplay(shared.Shop_Display{
		CurrentPage: pageNum,
		TotalPages:  len(banners),
		ItemName:    banners[pageNum-1].ImageUrl,
		ImageUrl:    banners[pageNum-1].ImageUrl,
		ItemPrice:   1000,
		PrevDisable: pageNum == 1,
		NextDisable: pageNum == len(banners),
	})
}
