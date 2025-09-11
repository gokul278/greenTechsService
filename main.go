package main

import (
	"fmt"
	"greenstech/routes"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	r := gin.Default()

	// Load the DotENV
	err := godotenv.Load()
	if err != nil {
		log.Fatal("❌Error loading .env file")
	}

	// ⚠️ Trust only localhost proxy (or none if you want)
	r.SetTrustedProxies(nil)

	// ✅ CORS configuration to allow only one origin
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:3000"}, // Change to your allowed origin
	// 	AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// }))
	r.Use(cors.New(cors.Config{
		AllowOriginFunc: func(origin string) bool {
			return true // allow all origins dynamically
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		AllowCredentials: true,
	}))

	//API calls 🚀

	//Authentication
	fmt.Println("=================Authentication=================")
	fmt.Println()
	routes.InitLoginRoutes(r)
	fmt.Println()

	//Profile
	fmt.Println("=================Profile=================")
	fmt.Println()
	routes.InitProfileRoutes(r)
	fmt.Println()

	//Registration
	fmt.Println("=================Registration=================")
	fmt.Println()
	routes.InitRegistationRoutes(r)
	fmt.Println()

	//File Handler
	fmt.Println("=================File Handler=================")
	fmt.Println()
	routes.InitFileHandlerRoutes(r)
	fmt.Println()

	//File Handler
	fmt.Println("=================Sub Trainer=================")
	fmt.Println()
	routes.InitSubtrainerRoutes(r)

	fmt.Println()
	fmt.Println()
	//Ping 🎯API
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong from Greens Tech Service",
		})
	})

	//Run the Server and Log Message
	fmt.Println("✅Server is Running at Port:" + os.Getenv("PORT"))
	r.Run("0.0.0.0:" + os.Getenv("PORT"))
}
