package main

import (
	"choiny/dba"
	"choiny/service"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"
	"log"
	"os"
	"time"
)

func main() {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, "mongodb://localhost:27017")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	r := gin.Default()
	r.Use(dba.Mongo(client))
	r.GET("/hash-points",service.HashPointHandler())
	r.GET("/map/track/:id", service.TrackPointHandler())
	err = r.Run()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
