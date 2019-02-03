package dba

import (
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/mongo"
)

func Mongo(client *mongo.Client) gin.HandlerFunc {
	mongoClient := client
	return func(c *gin.Context) {
		c.Set("db", mongoClient.Database("wipu"))
		c.Next()
	}
}
