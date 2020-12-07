{{/*
	Trigger: .*
	Trigger Type: Regex

	Description: YAGPDB will react on images and videos. You can optionally turn on an auto delete for none images. Watch this gif if you have no clue what I am talking about https://i.imgur.com/4xvBJeo.gif

	Note: Please configure the variables after pasting the code into your control panel and also make sure to limit this to a channel if needed!

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}
{{$delete := false}}
{{/*Set this to true if you want to delete all messages other than images or videos*/}}
{{$react := true}}
{{/*Set this to true if you want yagpdb to react with the emojis*/}}
{{$emojis := cslice "😂" "Hopeless:756140606522130542" "haku:781755715161161758" "OOF:781868908549242941"}}
{{/*Replace the emojis if you want. Default emojis will be the unicode like 👎 or 👍 and custom emojis will have to be emojiname:emojiid*/}}



{{$woulddelete1 := false}}{{$woulddelete2 := false}}{{$woulddelete3 := false}}
{{/*Attachment Media*/}}
{{range .Message.Attachments}}
	{{if not (reFind `(?i)\.(?:jpg|jpeg|png|gif|tif|webp|mp4|webm|mov)` .Filename)}}
		{{$woulddelete1 = true}}
	{{end}}
{{else}}
	{{$woulddelete1 = true}}
{{end}}

{{/*Media Link*/}}
{{range (reFindAll `https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)` .Message.Content)}}
	{{range (reFindAll `(?i)https?:(?://([^/?#]*))?(?:[^?\n#\s]*?/[^?\n#\s]+\.(jpg|jpeg|png|gif|tif|webp|mp4|webm|mov))(?:\?([^?\n#\s]*))?(?:#(.*))?` .)}}
	{{else}}
		{{$woulddelete2 = true}}
	{{end}}
{{else}}
	{{$woulddelete2 = true}}
{{end}}

{{/*Embed Media*/}}
{{sleep 5}}
{{range .Message.Embeds}}
	{{if ne .Type "video" "image" "gifv"}}
		{{$woulddelete3 = true}}
	{{end}}
{{else}}
	{{$woulddelete3 = true}}
{{end}}

{{/*Delete Message*/}}
{{if and $delete $woulddelete1 $woulddelete2 $woulddelete3}}
	{{deleteTrigger 0}}
{{end}}

{{/*Add Reactions*/}}
{{if and (or (not $woulddelete1) (not $woulddelete2) (not $woulddelete3)) $react}}
	{{range $emojis}}{{addReactions .}}{{end}}
{{end}}