package button

import (
	msg "github.com/crescent-network/cre_tele_bot/msg"

	tele "gopkg.in/telebot.v3"
)

func HandleButtons(b *tele.Bot) {
	var (
		// Universal markup builders.
		menu = &tele.ReplyMarkup{ResizeKeyboard: true}
		// selector = &tele.ReplyMarkup{}

		btn_Claim        = menu.Text("Airdrop Claim")
		btn_Cal          = menu.Text("Airdrop Calculator")
		btn_airdropCEX   = menu.Text("CEX Vali.")
		btn_totalSupply  = menu.Text("Total Supply")
		btn_discord      = menu.Text("Discord")
		btn_distribution = menu.Text("Token Distribution")
		btn_launch       = menu.Text("Launch")
		btn_twitter      = menu.Text("Twitter")
		btn_governance   = menu.Text("Governance")
		// btn_minStake     = menu.Text("Min Amount Stake")
		//btn_blog         = menu.Text("Blog")

		// Inline buttons.
		// btnPrev = selector.Data("⬅", "prev")
		// btnNext = selector.Data("➡", "next")
	)

	// Locate buttons
	temp := menu.Split(5, []tele.Btn{
		btn_Claim, btn_Cal, btn_airdropCEX,
		btn_launch, btn_totalSupply, btn_distribution,
		btn_discord, btn_twitter, btn_governance,
	})

	menu.Reply(
		temp[0],
		temp[1],
	)

	// with "/faq" input, button pop-up
	b.Handle("/faq", func(c tele.Context) error {
		return c.Send("Choose a topic: ", menu)
	})

	// button-click messages
	b.Handle(&btn_Claim, func(c tele.Context) error {
		return c.Send(msg.Airdrop_claim)
	})

	b.Handle(&btn_Cal, func(c tele.Context) error {
		return c.Send(msg.Airdrop_cal)
	})

	b.Handle(&btn_airdropCEX, func(c tele.Context) error {
		return c.Send(msg.Airdrop_cex)
	})

	b.Handle(&btn_totalSupply, func(c tele.Context) error {
		return c.Send(msg.Totalsupply)
	})

	b.Handle(&btn_discord, func(c tele.Context) error {
		return c.Send(msg.Discord)
	})

	b.Handle(&btn_distribution, func(c tele.Context) error {
		return c.Send(msg.Distribution)
	})

	b.Handle(&btn_launch, func(c tele.Context) error {
		return c.Send(msg.Launch)
	})

	b.Handle(&btn_twitter, func(c tele.Context) error {
		return c.Send(msg.Twitter)
	})

	b.Handle(&btn_governance, func(c tele.Context) error {
		return c.Send(msg.Governance)
	})
}
