package server

import (
	"context"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	db "github.com/jbattistella/normative/db/sqlc"
)

func Serve() {
	r := gin.Default()
	r.GET("/events/:limit", getEvents)
	r.GET("/events/:limit/:region", getEventsByRegion)
	r.Run("localhost:8080")
}

func getEvents(c *gin.Context) {

	dbQueries, err := db.ConnectDB()
	if err != nil {
		log.Println(err)
		// return http.StatusInternalServerError
	}

	limit := c.Param("limit")

	lm, _ := strconv.Atoi(limit)

	events, err := dbQueries.GetEventsList(context.Background(), int32(lm))
	if err != nil {
		log.Println(err)
	}

	c.IndentedJSON(http.StatusOK, events)

}

func getEventsByRegion(c *gin.Context) {

	dbQueries, err := db.ConnectDB()
	if err != nil {
		log.Println(err)
		// return http.StatusInternalServerError
	}

	region := c.Param("region")
	limit := c.Param("limit")

	lm, _ := strconv.Atoi(limit)

	args := db.GetEventsByRegionParams{Region: region, Limit: int32(lm)}

	events, err := dbQueries.GetEventsByRegion(context.Background(), args)
	if err != nil {
		log.Println(err)
	}

	c.IndentedJSON(http.StatusOK, events)

}
