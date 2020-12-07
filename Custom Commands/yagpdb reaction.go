{{/*
	Trigger Type: Reaction (add)

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}

{{$cid := .Reaction.ChannelID}}{{$mid := .Reaction.MessageID}}{{$en := .Reaction.Emoji.Name}}
{{if eq .Message.Author.ID 204255221017214977}}
	{{if (eq $en "❤️")}}
		{{editMessage $cid $mid (cembed
		"title" "❤️ Ever Thought about Dating me?"
		"description" "You know I'm single Right, we could go on some wild dates together 💋💋"
		"image" (sdict "url" "https://redgifs.com/watch/mindlessappropriategraywolf")
		"color" 0xde2e43)}}
		{{deleteAllMessageReactions $cid $mid}}
		{{addMessageReactions $cid $mid "❌"}}
	{{else if (eq $en "💡")}}
		{{editMessage $cid $mid (cembed
		"title" "💡Wanna have a one night stand?"
		"description" "Maybe I'm ready to bang you any day of the week, just hit me up and I'll come swinging to kiss your jiggly puffs goodnight🍑🍑"
		"color" 0xffd984)}}
		{{deleteAllMessageReactions $cid $mid}}
		{{addMessageReactions $cid $mid "❌"}}
	{{else if (eq $en "🛡️")}}
		{{editMessage $cid $mid (cembed
		"title" "🛡️ Can you send us some nudes"
		"description" "We certainly do appreaciate you sending us some of your nudes, that would be the most optimum of solutions during this period of isolation 👧😽"
		"color" 0x55acef)}}
		{{deleteAllMessageReactions $cid $mid}}
		{{addMessageReactions $cid $mid "❌"}}
	{{else if (eq $en "❌")}}
		{{editMessage $cid $mid (cembed
			"title" "Welcome To Bromance "
			"description" `❤️ Ever Thought about Dating me?
			💡 Wanna have a one night stand?
			🛡️ Can you send us some nudes`
			"color" 0xe67e22)}}
		{{deleteAllMessageReactions $cid $mid}}
		{{addMessageReactions $cid $mid "❤️" "💡" "🛡️"}}
	{{end}}
{{end}}