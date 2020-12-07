{{/*
	Trigger. .*
	Trigger Type: Regex

	Usage:
	Just type :)

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}

{{
	$message := cembed "title" "<a:furrydance:785377918682660886>  Have Suggestions? "
	 "description" ( joinStr "\n\n"
	 "  **Guidelines Below** -:"
	 "<a:bluefire:785420344072601600> Use **-suggest** command for ideas worth discussing " 
	 "<a:hypesquad:785377850403586100>  Use  **suggest** <text> to give ideas")
	 "thumbnail" (sdict "url" "https://thumbs.gfycat.com/HardtofindDelectableAsianconstablebutterfly-size_restricted.gif")
	
	
	 "color" (randInt 0 16777216)
}}

{{/* do not edit below */}}
{{if $db := dbGet .Channel.ID "stickymessage"}}
	{{deleteMessage nil (toInt $db.Value) 0}}
{{end}}
{{$id := sendMessageRetID nil $message}}
{{dbSet .Channel.ID "stickymessage" (str $id)}}