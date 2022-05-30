package main

import (
	"log"
	"rhodium"
	"rhodium/db"
	"rhodium/example/controllers"
	"rhodium/example/models"
	"rhodium/example/routes"
	"rhodium/example/rpc"
)

func main() {
	app := rhodium.New()
	dbConn := db.New()
	_ = dbConn.AutoMigrate(&models.Post{})

	app.Routes(
		routes.NewPageRoutes(controllers.NewPageController(dbConn)),
	)

	app.RPCRoutes(
		routes.NewPostRPCRoutes(rpc.NewPostRPC(dbConn)),
	)

	log.Panic(app.Run(":8080"))
}
