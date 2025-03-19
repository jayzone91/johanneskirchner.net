function Render()
	local title = data.title
	local message = data.message

	return string.format(
		[[
        <html>
        <head><title>%s</title></head>
        <body>
            <h1>%s</h1>
            <p>%s</p>
        </body>
        </html>
    ]],
		title,
		title,
		message
	)
end
