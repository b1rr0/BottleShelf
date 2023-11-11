package main

import (
	"context"
	"inventoryService/m/v2/ent"
	"log"
	"net/http"

	_ "github.com/lib/pq"

	controllers "inventoryService/m/v2/cmd/controllers"
	docs "inventoryService/m/v2/cmd/docs"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"     // swagger embed files
	ginSwagger "github.com/swaggo/gin-swagger" // gin-swagger middleware
)

type Server struct {
	db   *ent.Client
	http *gin.Engine
}

var svr Server

func main() {
	initDatabase()

	itemController := controllers.ItemController{}
	itemController.Client = svr.db

	runHttpServer(itemController)

	closeDatabase()
}

func initDatabase() {
	client, err := ent.Open("postgres", "host=0.0.0.0 port=5432 user=postgres password=Luntik228 dbname=bottleshelfInventory sslmode=disable")
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}

	svr.db = client
}

func runHttpServer(itemController controllers.ItemController) {
	r := gin.New()

	docs.SwaggerInfo.BasePath = "/api/v1"

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET(docs.SwaggerInfo.BasePath+"/ping", ping)
	r.GET(docs.SwaggerInfo.BasePath+"/inventory", itemController.GetIngridientsList)
	r.POST(docs.SwaggerInfo.BasePath+"/ingridient", itemController.AddIngridient)

	// api doc http://localhost:8080/swagger/index.html
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	_ = r.Run(":8080")
}

func closeDatabase() {
	defer svr.db.Close()
	// Run the auto migration tool.
	if err := svr.db.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
}

// @BasePath /api/v1

// PingExample godoc
// @Summary ping
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /ping [get]
func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
