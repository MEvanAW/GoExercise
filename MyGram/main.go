package main

import (
	"example.id/mygram/database"
	_ "example.id/mygram/docs"
	"example.id/mygram/routers"
)

// @title           MyGram API
// @version         1.0
// @description     API server for MyGram social media in "Scalable Webservice with Golang" course from Hacktiv8 × Kominfo.

// @contact.name   Muhammad Evan Anindya Wahyuaji
// @contact.email  m.evan.aw@google.com

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	database.StartDB()
	var port = ":8080"
	routers.StartServer().Run(port)
}
