	{{/*
	Trigger : When Someone Joins or Leaves
	Trigger Type : Welcome Message
	Made by Rockefeller street
#9307
	
	
	
	*/}}


	{{ $username := .User.Username}}
	{{ $avatar := .User.AvatarURL "256"}}
	{{ $message := "Welcome+to+server+name"}}
	{{ $background := "https://cdn.discordapp.com/attachments/755600572215722035/799132031610650684/Welcome_4.png"}}
	{{ $color := "black"}}
	{{ $humanage := currentUserAgeHuman}}
	{{ $advice := (execAdmin "advice") }}
	{{ $topic := (execAdmin "topic") }}
	{{ $catfact := (execAdmin "advice") }}
	{{ $avatar := (joinStr "" "https://cdn.discordapp.com/avatars/" (toString .User.ID) "/" .User.Avatar ".png") }}
	{{$url := print "https://api.no-api-key.com/api/v2/welcome?username=" 
	(urlescape .User.Username) 
	"&user_image=" (.User.AvatarURL "2048")
	$message 
		"&background=" $background 
		"&color=" $color
		"&text_heading=" (urlquery (print "Welcome to " .Server.Name "!" ))}}
	{{$embed := cembed
		"title" "<a:bluefire:785420344072601600>A New Comrade Appears!<a:bluefire:785420344072601600>"
		"color" (randInt 0 8645612) 
		"fields" (cslice 
			(sdict "name" "Advice" "value" $catfact "inline" false) 
			(sdict "name" "<a:babydance:785378327488757760> Account Age" "value" $humanage "inline" true) 
			(sdict "name" " <a:tick:785421896194981899> Joined at" "value" .Member.JoinedAt "inline" true) 
			(sdict "name" " <a:furrydance:785377918682660886> Member Count" "value" (toString .Guild.MemberCount) "inline" true) 
			
		) 
		"author" (sdict "name" " 🔥Server Salute!🔥 " "icon_url" "https://media.discordapp.net/attachments/755600572215722035/801474613535178752/logofull.png") 
		"image" (sdict "url" $url)
		"timestamp" .Guild.JoinedAt
	}}
	{{sendMessage nil .User.Mention}}
	{{sendMessage nil $embed}}
	{{addResponseReactions "pepeGunfire:737616060336963726"}}
	{{addMessageReactions  $embed "happyred:785378521933414420" "babydance:785378327488757760" "settings:785378242760146974" }}
