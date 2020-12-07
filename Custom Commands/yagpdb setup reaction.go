  
{{/*
	Trigger: setup
	Trigger Type: Command

	Usage:
	setup

	Made by TheHDCrafter#0001
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}

{{$x := sendMessageRetID nil (cembed
"title" " <:logofull:767288516560158720> Welcome To Bromance  <:logofull:767288516560158720> "
"description" ` React with the following emojis to display the corresponding site 
<a:happyred:785378521933414420> **THE BROS**
<a:babydance:785378327488757760> **THE ROLES**
<a:settings:785378242760146974> **THE RULES**
<a:nitro:785377955739074580> **THE LEADERBOARD**`
"thumbnail" (sdict "url" "https://cdn.discordapp.com/attachments/780649758409621514/780684990765596682/Green_and_Black_Geometric_Techno_Technology_YouTube_Thumbnail_4.png")

"color" (randInt 0 8777216))}}
{{addMessageReactions nil $x "happyred:785378521933414420" "babydance:785378327488757760" "settings:785378242760146974" ":nitro:785377955739074580" }}