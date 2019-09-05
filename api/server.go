package api

import (
	"fmt"
	"golang-boilerplates/api/controllers"

	"github.com/gin-gonic/gin"
)

func Run() {
	port := 9000
	fmt.Printf("\n\tListening [::]:%d\n", port)
	listen(port)
}

func listen(port int) {
	r := gin.Default()
	fmt.Println("jancok router")

	r.POST("/products", controllers.GetProducts)
	r.Run(":3000")
}
