package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Product struct {
	Id          int    `json:"id"`
	Name        string `json:"name"`
	Price       int    `json:"price"`
	Description string `json:"description"`
}

func productsHandler(c *gin.Context) {
	products := []Product{
		{Id: 100, Name: "BassTune Headset 2.0", Price: 200, Description: "A headphone with a inbuilt high-quality microphone"},
		{Id: 101, Name: "Fastlane Toy Car", Price: 100, Description: "A toy car that comes with a free HD camera"},
		{Id: 102, Name: "ATV Gear Mouse", Price: 75, Description: "A high-quality mouse for office work and gaming"},
	}

	c.JSON(200, gin.H{
		"products": products,
	})
}

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	r.GET("/products", productsHandler)
	r.Run(":5000")
}
