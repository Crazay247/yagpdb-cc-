{{/*
	Trigger: https://(?:\w+\.)?discord(?:app)?\.com/channels\/(\d+)\/(\d+)\/(\d+)
	Trigger Type: Regex

	Usage:
	Right click on a message to copy the message link and send it in chat.

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}

{{$m := reFindAllSubmatches `https://(?:\w+\.)?discord(?:app)?\.com/channels\/(\d+)\/(\d+)\/(\d+)` .Message.Content}}{{$reminder := true}}
{{if not (eq (toInt64 (index $m 0 1)) .Guild.ID)}}
	{{sendMessage nil (cembed
		"author" (sdict
			"name" "Unknown User"
			"icon_url" "https://i.imgur.com/jNz2Dwp.png")
		"description" (print "\n\n**[Message](" (index $m 0 0) "/) in <#" (index $m 0 2) ">**\n" "<:excl:565142262401728512> Unknown Message")
		"color" 0xF04747
		"footer" (sdict
			"text" (print "Command triggered by " .Message.Author.String ". Message from"))
		"timestamp" (div (index $m 0 3) 4194304 | mult 1000000 | toDuration | .DiscordEpoch.Add))}}
{{else}}
	{{$msg := getMessage (index $m 0 2) (index $m 0 3)}}{{$content := $msg.Content}}
	{{$mc := reReplace `(?:https?://)?(?:www\.)?(discord(?:app)?\.gg(?:/|\\+/+)|discord(?:app)?\.com(?:/|\\+/+)(?:invite/))\w{2,}|(?:https?://)?(?:www\.)?discord(?:app)?\.(?:io|me|li)(?:/|\\+/+)[A-z+0-9]{2,}` $msg.Content "[Invite Removed](https://example.com/)"}}
	{{$e := sdict
		"author" (sdict
			"name" (print $msg.Author.String " (ID " $msg.Author.ID ")")
			"icon_url" ($msg.Author.AvatarURL "1024"))
		"color" 0x36393F
		"footer" (sdict
			"text" (print "Command triggered by " .User.String ". Message from"))
		"timestamp" $msg.Timestamp}}
	{{if and $msg.Content (not $msg.Embeds)}}
		{{$e.Set "description" (print "\n\n**[Message](" (index $m 0 0) ") in <#" $msg.ChannelID ">**\n" $mc)}}
	{{end}}
	{{$fo := $e.Get "footer"}}
	{{$ti := $msg.Timestamp}}
	{{if or $msg.Attachments $msg.Embeds}}
		{{range $c,$ma := (or $msg.Attachments $msg.Embeds)}}
			{{$filename := ""}}{{$count := ""}}
			{{if eq (printf "%T" .) "*discordgo.MessageAttachment"}}
				{{$filename = .Filename}}{{$count = len $msg.Attachments}}
			{{else if eq (printf "%T" .) "*discordgo.MessageEmbed"}}
				{{if eq .Type "image"}}
					{{$filename = .URL}}{{$count = len $msg.Embeds}}
				{{end}}
			{{end}}
			{{if ne $filename ""}}
				{{$e.Del "footer"}}
				{{$e.Del "timestamp"}}
				{{if eq $c 0}}
					{{$e.Set "description" (print "\n\n**[Message](" (index $m 0 0) ") in <#" $msg.ChannelID ">**\n" $mc)}}
				{{end}}
				{{if eq $count (add $c 1)}}
					{{$e.Set "footer" $fo}}
					{{$e.Set "timestamp" $ti}}
				{{end}}
				{{if (reFind `(?i)\.(jpg|jpeg|png|gif|tif|tiff|webp)$` $filename)}}
					{{$e.Set "image" (sdict "url" .URL)}}
					{{sendMessage nil (cembed $e)}}
				{{else}}
					{{$e.Set "fields" (cslice
						(sdict "name" "File Name" "value" (print "`" $filename "`") "inline" true)
						(sdict "name" "URL" "value" (print "[File Link](" .URL ")") "inline" true))}}
					{{sendMessage nil (cembed $e)}}
				{{end}}
				{{$e.Del "fields"}}
				{{$e.Del "image"}}
				{{$e.Del "description"}}
				{{$e.Del "author"}}
				{{$reminder = false}}{{$content = ""}}
			{{end}}
		{{end}}
	{{else}}
		{{if and $msg.Content (not $msg.Embeds)}}
			{{sendMessage nil (cembed $e)}}
		{{end}}
	{{end}}
	{{if $msg.Embeds}}
		{{if $reminder}}{{$embed2 := ""}}
			{{if eq (index $msg.Embeds 0).Type "rich"}}
			    {{$embed2 = index $msg.Embeds 0}}
			    {{sendMessage nil (complexMessage "embed" $embed2 "content" $content)}}
			{{else}}
	            {{sendMessage nil  $content}}
	        {{end}}
			{{$e.Set "description" (print "**[Embed Link](" (index $m 0 0) ") to <#" $msg.ChannelID ">**")}}
			{{sendMessage nil (cembed $e)}}
		{{end}}
	{{end}}
{{end}}
{{if eq (len (index $m 0 0)) (len .Message.Content)}}
	{{deleteTrigger 1}}
{{end}}