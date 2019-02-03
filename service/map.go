package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson"
	"github.com/mongodb/mongo-go-driver/bson/objectid"
	"github.com/mongodb/mongo-go-driver/mongo"
	"github.com/mongodb/mongo-go-driver/options"
	"log"
	"time"
)

type HashPoint struct {
	DJIA      float64
	Location  bson.A
	Graticule bson.A
	Date      time.Time
}

type TrackPoint struct{
	Location bson.A
}

func TrackPointHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*mongo.Database)
		collection := db.Collection("track_points")
		ctx, _ := context.WithTimeout(context.Background(), 250*time.Millisecond)
		oid, err := objectid.FromHex(c.Param("id"))
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		projection := bson.D{
			{"location", 1},
			{"_id", 0},
		}
		findOptions := options.Find()
		findOptions.SetProjection(projection)

		cur, err := collection.Find(ctx, bson.D{{
			"route_id",
			oid,
		}}, findOptions)

		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		defer cur.Close(ctx)
		var results []*TrackPoint
		for cur.Next(ctx) {
			var result TrackPoint
			err := cur.Decode(&result)
			if err != nil {
				c.AbortWithError(500, err)
				return
			}
			results = append(results, &result)
		}
		if err := cur.Err(); err != nil {
			c.AbortWithError(500, err)
			return
		}
		c.JSON(200, results)
	}
}

func HashPointHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := c.MustGet("db").(*mongo.Database)
		collection := db.Collection("hash_points")
		ctx, _ := context.WithTimeout(context.Background(), 250*time.Millisecond)
		cur, err := collection.Find(ctx, bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cur.Close(ctx)
		var results []*HashPoint
		for cur.Next(ctx) {
			var result HashPoint
			err := cur.Decode(&result)
			if err != nil {
				log.Fatal(err)
			}
			results = append(results, &result)
		}
		if err := cur.Err(); err != nil {
			log.Fatal(err)
		}
		c.JSON(200, results)
	}
}
