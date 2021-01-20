{{$username := .User.Username}}
{{$avatar := .User.AvatarURL "256"}}
{{$message := "Welcome+to+server+name"}}
{{$background := "https://i.imgur.com/0PJZwzl.jpeg"}}
{{$color := "black"}}

{{$url := print "https://api.no-api-key.com/api/v2/welcome?username=" (urlescape .User.Username) "&user_image=" (.User.AvatarURL "1024") "&text_heading=" (urlquery (print "Welcome to " .Server.Name "!"))}}
{{sendMessage nil (cembed "image" (sdict "url" $url))}}





{{$username := .User.Username}}
{{$avatar := .User.AvatarURL "256"}}
{{$message := "Welcome+to+server+name"}}
{{$background := "https://i.imgur.com/0PJZwzl.jpeg"}}
{{$color := "black"}}

{{sendMessage nil (cembed "title" "Welcome!" "image" (sdict "url" (print "https://api.no-api-key.com/api/v2/welcome?username=" $username "&user_image=" $avatar "&text_heading=" $message "&background=" $background "&color=" $color)))}}