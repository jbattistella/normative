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
