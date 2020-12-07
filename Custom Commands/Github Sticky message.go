{{/*
	Trigger. .*
	Trigger Type: Regex

	Usage:
	Just type :)

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}

{{
	$message := cembed "title" "  Want to have your Github repo shown ? "
	 "description"  `
	   **Ask the mods**
	 Once deemed trustworthy, we'll give you access to our **webhook**. `
	
	 "thumbnail" (sdict "url" "https://www.pngitem.com/pimgs/m/128-1280162_github-logo-png-cat-transparent-png.png")
	
	
	 "color"  4437377
}}

{{/* do not edit below */}}
{{if $db := dbGet .Channel.ID "stickymessage"}}
	{{deleteMessage nil (toInt $db.Value) 0}}
{{end}}
{{$id := sendMessageRetID nil $message}}
{{dbSet .Channel.ID "stickymessage" (str $id)}}