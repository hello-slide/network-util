package networkutil

import "github.com/rs/cors"

var CorsConfig = cors.New(cors.Options{
	AllowedOrigins: []string{
		"https://hello-slide.jp",
		"https://hello-slide.com",
		"https://hello-slide.net",
		"http://localhost:3000",
	},
})
