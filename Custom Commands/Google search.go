{{/*
	Trigger : ^-(google|youtube|porn|insta|facebook|i|ph|g|y|fb|twitter|t|quora|q)
	Trigger Type : Regex
	Made by Rockefeller street
#9307
	
	
	
	*/}}

{{if reFind "google|g" .Cmd}}
https://www.google.com/search?q={{urlquery (joinStr " " .CmdArgs)}}
{{else if  reFind "youtube|y" .Cmd }}
https://www.youtube.com/results?search_query={{urlquery (joinStr " " .CmdArgs)}}
{{else if reFind "porn|ph" .Cmd }}
https://www.pornhub.com/video/search?search={{urlquery (joinStr " " .CmdArgs)}}
{{else if reFind "insta|ig" .Cmd }}
https://www.instagram.com/{{urlquery (joinStr " " .CmdArgs)}}
{{else if reFind "facebook|fb" .Cmd }}
https://www.facebook.com/search/top?q={{urlquery (joinStr " " .CmdArgs)}}
{{else if reFind "twitter|t" .Cmd }}
https://twitter.com/search?q={{urlquery (joinStr " " .CmdArgs)}}&src=typed_query&f=user
{{else if reFind "quora|q" .Cmd }}
https://www.quora.com/search?q={{urlescape (joinStr " " .CmdArgs)}}
{{end}}
