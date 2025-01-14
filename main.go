package main

import (
	routes "example.com/Routes"
	"example.com/db"
	"github.com/gin-gonic/gin"
)

func main(){
    db.InitDB() 
	server := gin.Default()

	routes.Register(server)

    server.Run(":8080")
}
