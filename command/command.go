package command

import (
	msg "github.com/crescent-network/cre_tele_bot/msg"

	tele "gopkg.in/telebot.v3"
)

func Define_command(b *tele.Bot) {
	// Define commands and the responses of the commands
	com_launch := tele.Command{
		Text:        "/launch",
		Description: "Crescent launching",
	}

	b.Handle("/launch", func(c tele.Context) error {
		return c.Send(msg.Launch)
	})

	com_airdrop := tele.Command{
		Text:        "/airdrop_calculator",
		Description: "Crescent airdrop calculator",
	}
	b.Handle("/airdrop_calculator", func(c tele.Context) error {
		return c.Send(msg.Airdrop_cal)
	})

	// com_minstake := tele.Command{
	// 	Text:        "/airdrop_required_stake",
	// 	Description: "Max/Min ATOM staking for airdrop",
	// }
	// b.Handle("/airdrop_required_stake", func(c tele.Context) error {
	// 	return c.Send(msg.Airdrop_cex)
	// })

	com_totalsupply := tele.Command{
		Text:        "/supply",
		Description: "Total supply of Crescent",
	}
	b.Handle("/supply", func(c tele.Context) error {
		return c.Send(msg.Totalsupply)
	})

	com_Distribution := tele.Command{
		Text:        "/distribution",
		Description: "Token distribution",
	}
	b.Handle("/distribution", func(c tele.Context) error {
		return c.Send(msg.Distribution)
	})

	com_twitter := tele.Command{
		Text:        "/twitter",
		Description: "Crescent Twitter",
	}
	b.Handle("/twitter", func(c tele.Context) error {
		return c.Send(msg.Twitter)
	})

	com_governance := tele.Command{
		Text:        "/governance",
		Description: "Cosmos Hub Governance Proposal for Crescent",
	}
	b.Handle("/governance", func(c tele.Context) error {
		return c.Send(msg.Governance)
	})

	com_blog := tele.Command{
		Text:        "/blog",
		Description: "Crescent Blog",
	}
	b.Handle("/blog", func(c tele.Context) error {
		return c.Send(msg.Blog)
	})

	com_faq := tele.Command{
		Text:        "/faq",
		Description: "FAQ for Crescent",
	}
	b.Handle("/faq", func(c tele.Context) error {
		return c.Send("/faq")
	})

	com_discord := tele.Command{
		Text:        "/discord",
		Description: "Discord for Crescent",
	}
	b.Handle("/discord", func(c tele.Context) error {
		return c.Send(msg.Discord)
	})

	com_teaser := tele.Command{
		Text:        "/teaser",
		Description: "Crescent Teaser",
	}
	b.Handle("/teaser", func(c tele.Context) error {
		return c.Send(msg.Teaser_youtube)
	})

	// Register the commands
	coms := []tele.Command{com_faq, com_governance, com_teaser, com_twitter, com_blog, com_airdrop, com_launch, com_discord, com_totalsupply, com_Distribution}
	b.SetCommands(coms)

}
