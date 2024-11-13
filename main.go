package main

import (
	"os"
	"restaurant/database"
	"restaurant/middleware"
	"restaurant/routes"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var foodCollecion *mongo.Collection = database.OpenCollection(database.Client, "food")

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		port = "8000"
	}

	router := gin.New()                     //gin.New() creates a new Gin engine instance without any default middleware.
	router.Use(gin.Logger())                //router.Use(gin.Logger()) adds the logger middleware to log HTTP requests and responses. This helps in debugging and monitoring HTTP traffic.
	routes.UserRoutes(router)               //routes.UserRoutes(router): Registers routes related to user functionality (e.g., user login, registration) in the router.
	router.Use(middleware.Authentication()) //router.Use(middleware.Authentication()): Adds custom middleware for authentication, ensuring that routes defined after this line require authentication.

	routes.FoodRoutes(router)
	routes.MenuRoutes(router)
	routes.TableRoutes(router)
	routes.OrderRoutes(router)
	routes.OrderItemRoutes(router)
	routes.InvoiceRoutes(router)

	router.Run(":" + port)
}
