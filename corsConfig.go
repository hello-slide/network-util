package networkutil

import "github.com/rs/cors"

var CorsConfig = cors.New(cors.Options{
	AllowedOrigins: []string{
		"https://hello-slide.jp",
		"http://localhost:3000",
	},
})
