package main

import (
	"github.com/gin-gonic/gin"
	"github.com/olivere/elastic/v7"

	"context"
	"strconv"
	"fmt"
	"net/http"
	"math/rand"
	"time"
)

type Metrics struct {
	Id    int    `json:"id"`
	Title string `json:"title" binding:"required"`
	Value int    `json:"value" binding:"required"`
}

func main() {
	//Elastic connection
	ctx := context.Background()
	client, err := elastic.NewClient()
	if err != nil {
		// Handle error
		panic(err)
	}

	r := gin.Default()

	//POST
	r.POST("/api/metrics", func(c *gin.Context) {
		var m Metrics
		err := c.ShouldBindJSON(&m)
		if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
            return
		}
		rand.Seed(time.Now().UnixNano())
		succ, err := client.Index().
			Index("metrics").
			Type("metric").
			Id(strconv.Itoa(rand.Intn(10000))).
			BodyJson(m).
			Do(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		fmt.Printf("Indexed metrics %s to index %s, type %s\n", succ.Id, succ.Index, succ.Type)
	})

	//GET
	r.GET("/api/metrics/:id", func (c *gin.Context) {
		id := c.Param("id")
		// Get metrics with specified ID
		get1, err := client.Get().
			Index("metrics").
			Type("metric").
			Id(id).
			Do(ctx)
		if err != nil {
			// Handle error
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		if get1.Found {
			fmt.Println(get1.Fields)
			fmt.Printf("Got metrics %s in version %d from index %s, type %s\n", get1.Id, get1.Version, get1.Index, get1.Type)
		}
	})

	//DELETE
	r.DELETE("/api/metrics/:id", func(c *gin.Context) {
		id := c.Param("id")
		succ, err := client.Delete().
			Index("metrics").
			Type("metric").
			Id(id).
			Do(ctx)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		fmt.Println("Deleted with id: " + succ.Id)
		c.JSON(http.StatusAccepted, gin.H{"message": "deleted"})
	})

	//TODO: PUT
	r.PUT("/api/metrics/:id", func (c *gin.Context) {
		id := c.Param("id")
		var m Metrics
		err := c.ShouldBindJSON(&m)
		if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
            return
		}

		update, err := client.Update().Index("metrics").Type("metric").Id(id).
		Script(elastic.NewScriptInline("ctx._source.title = params.title; ctx._source.value = params.value").Lang("painless").Params(map[string]interface{}{"title": m.Title, "value": m.Value})).
		Do(ctx)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
			return
		}
		fmt.Printf("New version of metric %q is now %d\n", update.Id, update.Version)
	})
	
	r.Run() // localhost:8080
}