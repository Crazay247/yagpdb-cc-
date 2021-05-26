
  
{{/*
	Trigger: (blast|kamehameha|negro|rtfm|fire)
	Trigger Type: Regex

	Usage:
	setup

	Made by Rockefeller street#9307
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}


{{if reFind "blast" .Cmd }}
https://i.imgur.com/FA3xbTD.mp4
{{else if  reFind "kamehameha" .Cmd }}
https://thumbs.gfycat.com/ThoughtfulEcstaticKitty.webp
{{else if  reFind "negro" .Cmd }}
https://thumbs.gfycat.com/AdorableJampackedArmadillo.webp
{{else if  reFind "rtfm" .Cmd }}
https://media1.giphy.com/media/dxxOef5Qie4WlST6QG/source.gif
{{deleteResponse (10)}}
{{else if  reFind "fire" .Cmd }}
<a:pepeGunfire:737616060336963726>
{{addResponseReactions "pepeGunfire:737616060336963726"}}

{{end}}

