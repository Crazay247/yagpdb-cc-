{{/*
	Trigger: cc
	Trigger Type: command

	Note: If you enable this custom command make sure to disable the original CustomCommands command. If you do not know how to do so please watch this tutorial https://i.imgur.com/eAvOR2N.gif
	
	Made by TheHDCrafter#0001 
	Need help or you found a bug? Join https://discord.gg/GRns3fg or report it on GitHub.
*/}}
{{$perms := "ManageServer"}}
{{/*The bot will check if the user has this permission.
Permissions available: Administrator, ManageServer, ReadMessages, SendMessages, SendTTSMessages, ManageMessages, EmbedLinks, AttachFiles, ReadMessageHistory, MentionEveryone, VoiceConnect, VoiceSpeak, VoiceMuteMembers, VoiceDeafenMembers, VoiceMoveMembers, VoiceUseVAD, ManageNicknames, ManageRoles, ManageWebhooks, ManageEmojis, CreateInstantInvite, KickMembers, BanMembers, ManageChannels, AddReactions, ViewAuditLogs*/}}

{{if in (split (index (split (exec "viewperms") "\n") 2) ", ") $perms}}
	{{$prefix := index (reFindAllSubmatches `Prefix of \x60\d+\x60: \x60(.+)\x60` (exec "prefix")) 0 1}}
	{{$page := 1}}{{$grab := ""}}{{$max := 0}}{{$sdict := sdict}}{{$group := false}}{{$type := false}}{{$trigger := false}}{{$out := ""}}
	{{with .StrippedMsg}}
		{{with (reFindAllSubmatches `(?i)page (\d+)` .)}}
			{{$page = toInt (index . 0 1)}}
		{{end}}
		{{with (reFindAllSubmatches `(?i)group (.+)` .)}}
			{{$group = index . 0 1}}
		{{end}}
		{{with (reFindAllSubmatches `(?i)(?:type|trigger type) (None|n|Command|cmd?|Starts with|sw?|Contains|con?|Regex|reg|Exact Match|em?|Reaction|reac?t?|Interval|int|iv)` .)}}
			{{$t2 := lower (index . 0 1)}}
			
			{{if eq $t2 "none" "n"}}{{$type = "None"}}{{end}}
			{{if eq $t2 "command" "cm" "cmd"}}{{$type = "Command"}}{{end}}
			{{if eq $t2 "starts with" "s" "sw"}}{{$type = "Starts With"}}{{end}}
			{{if eq $t2 "contains" "co" "con"}}{{$type = "Contains"}}{{end}}
			{{if eq $t2 "regex" "reg"}}{{$type = "Regex"}}{{end}}
			{{if eq $t2 "exact match" "exact" "e" "em"}}{{$type = "Exact"}}{{end}}
			{{if eq $t2 "reaction" "rea" "reac" "react"}}{{$type = "Reaction"}}{{end}}
			{{if eq $t2 "interval" "int" "iv" "i"}}{{$type = "Interval"}}{{end}}
		{{end}}
		{{with (reFindAllSubmatches `(?i)(?:code|name|id) (.+)` .)}}
			{{$trigger = index . 0 1}}
		{{end}}
	{{end}}

	
	{{/*Guild Icon*/}}
	{{$ex := "png"}}{{if reFind "a_" .Guild.Icon}}{{$ex = "gif"}}{{end}}{{$icon := print "https://cdn.discordapp.com/icons/" .Guild.ID "/" .Guild.Icon "." $ex "?size=1024"}}

	{{/*Splitting exec cc output and filtering / compiling output*/}}
	{{if not $trigger}}
		{{range (split (exec "cc") "\n`#")}}
			{{- range (reFindAllSubmatches `\s*(\d+):\x60 (?:\x60(.+)\x60: )?(.*) - Group: \x60(.+)\x60` .) -}}
				{{- $out = ""}}{{$t4 := ""}}{{if ne (index . 2) ""}}{{$t4 = "`"}}{{end -}}
				{{- $t1 := print "<:cc_ID:780036580592189460>ID: **[" (index . 1) "](https://yagpdb.xyz/manage/" $.Guild.ID "/customcommands/commands/" (index . 1) "/)**\n<:sendmessage:704342789172887672>Trigger: " $t4 (or (index . 2) "<:cross:705738821110595607>") $t4 "\n<:cc_triggertype:779859347365822505>Trigger Type: **" (or (index . 3) "<:cross:705738821110595607>") "**\n<:cc_group:779859070122721320>Group: **" (index . 4) "**\n\n" -}}
				{{- if $group -}}
					{{- if eq (lower (index . 4)) (lower $group) -}}
						{{- $out = $t1 -}}
					{{- end -}}
				{{- else if $type -}}
					{{- if eq (or (index . 3) "None") $type -}}
						{{- $out = $t1 -}}
					{{- end -}}
				{{- else -}}
					{{- $out = $t1 -}}
				{{- end -}}
				{{- if ge (add (len $grab) (len $out)) 2000 -}}
					{{- $max = add $max 1 -}}
					{{- $sdict.Set (str $max) $grab -}}
					{{- $grab = $out -}}
				{{- else -}}
					{{- $grab = print $grab $out -}}
				{{- end -}}
			{{- end -}}
		{{end}}
		{{$max = add $max 1}}
		{{$sdict.Set (str $max) $grab}}
		{{$ccc := 100}}{{if .IsPremium}}{{$ccc = 250}}{{end}}

		{{/*Output*/}}
		{{$id := sendMessageRetID nil (cembed "thumbnail" (sdict "url" (print $icon (reReplace `\s` (print "?type=" $type "?group=" $group "?page=" $page "?max=" $max "?perms=" $perms) `%20`))) "author" (sdict "name" (print "Custom Commands - " .Guild.Name) "icon_url" "https://cdn.discordapp.com/emojis/780080308820901919.png" "url" (print "https://yagpdb.xyz/manage/" .Guild.ID "/customcommands/")) "description" ($sdict.Get (str $page)) "color" 7864168 "footer" (sdict "text" (print "Page: " $page "/" $max)) "fields" (cslice (sdict "name" "**• More Info**" "value" (print "Usage: " $prefix "cc page <Number>, " $prefix "cc group <GroupName>, " $prefix "cc type <Regex/cmd...>, " $prefix "cc name <CCID/Trigger/Name>\n\nCustom Commands: **" (len (split (exec "cc") "\n`#")) "/" $ccc "**"))))}}
		{{if and (ne $page $max) (eq $page 1)}}{{addMessageReactions nil $id "blank:656125423247425557" "▶️"}}{{end}}
		{{if and (ne $page 1) (eq $page $max)}}{{addMessageReactions nil $id "◀️" "blank:656125423247425557"}}{{end}}
		{{if and (ne $page $max) (ne $page 1)}}{{addMessageReactions nil $id "◀️" "▶️"}}{{end}}
	{{else}}
		{{$cc := exec "cc" $trigger}}
		{{$text := ""}}{{$t4 := ""}}
		{{if not (reFind `^No command by that name or id found, here is a list of them all:` $cc)}}
			{{$m := (reFindAllSubmatches `^#\s*(\d+) - (.*)(?:(?:: \x60(.*)\x60 - Case sensitive trigger: \x60(.+)\x60)|\b) - Group: \x60(.+)\x60\n\x60\x60\x60\n((?:.|\n){0,})\x60\x60\x60$` $cc)}}
			{{if $m}}{{$t67 := index $m 0 4}}{{if not $t67}}{{$t67 = "<:cross:705738821110595607>"}}{{end}}
				{{if (index $m 0 3)}}{{$t4 = "`"}}{{end}}
				{{$text = print "<:cc_ID:780036580592189460>ID: **[" (index $m 0 1) "](https://yagpdb.xyz/manage/" $.Guild.ID "/customcommands/commands/" (index $m 0 1) "/)**\n<:sendmessage:704342789172887672>Trigger: " $t4 (or (index $m 0 3) "<:cross:705738821110595607>") $t4 "\n<:cc_triggertype:779859347365822505>Trigger Type: **" (or (index $m 0 2) "<:cross:705738821110595607>") "**\n<:cc_Aa:780097353004482611>Case sensitive: **" (title $t67) "**\n<:cc_group:779859070122721320>Group: **" (index $m 0 5) "**\n\n" -}}
				{{if not (reFind `\x60\x60\x60` (index $m 0 6))}}
					{{$t3 := print $text "```go\n" (index $m 0 6) "\n```"}}
					{{if gt (len $t3) 2000}}
						{{$text = print $text "*The custom command code was too long to display. Please download the file to view the entire code.*"}}
					{{else}}
						{{$text = $t3}}
					{{end}}
				{{else}}
					{{$text = print $text "*The message couldn't be displayed because of the codeblock format. Please download the file to view the entire code.*"}}
				{{end}}
				{{sendMessage nil (complexMessage "file" (index $m 0 6) "embed" (cembed "description" $text "color" 7864168 "thumbnail" (sdict "url" $icon) "author" (sdict "name" (print "Custom Commands - " $.Guild.Name) "icon_url" "https://cdn.discordapp.com/emojis/780080308820901919.png" "url" (print "https://yagpdb.xyz/manage/" $.Guild.ID "/customcommands/"))))}}
			{{end}}
		{{else}}
			{{sendMessage nil (cembed "title" "Error" "description" (print "<:cross:705738821110595607> the custom command `" $trigger "` was not found.") "color" 0xDD2E44)}}
		{{end}}
	{{end}}
{{else}}
	{{sendMessage nil (cembed "title" "Missing permissions" "description" (print "<:cross:705738821110595607> You are missing the permission `" $perms "` to use this command!") "color" 0xDD2E44)}}
{{end}}