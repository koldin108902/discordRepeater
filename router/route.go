package router

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func Route(w *discordgo.Session, r *discordgo.MessageCreate) {
	if len(r.Content) < 1 {
		return
	}
	if r.Author.Bot {
		return
	}
	msg := strings.Split(r.Content, "")

	if msg[0] != "$" {
		return
	}
	cmd := strings.Split(strings.Join(msg[1:], ""), " ")

	if cmd[0] == "반복" {
		StartRepeat(w, r)
	}
	if cmd[0] == "종료" {
		StopRepeat(w, r)
	}
	if cmd[0] == "설명" {
		CmdHelp(w, r)
	}
}
