package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Film struct {
	Title    string
	Director string
}

func HtmxIndex(c *gin.Context) {
	// var media []models.Media

	// initializers.DB.Find(&media)

	films := []Film{
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "Blade Runner", Director: "Ridley Scott"},
		{Title: "The Thing", Director: "John Carpenter"},
	}

	c.HTML(http.StatusOK, "sample/index.tmpl", gin.H{
		"Films": films,
	})
}

func HtmxCreate(c *gin.Context) {
	// get data off req body
	// title := c.Param("title")
	// director := c.Param("director")

	// hard coding for now (no DB interaction)
	films := []Film{
		{Title: "The Godfather", Director: "Francis Ford Coppola"},
		{Title: "Blade Runner", Director: "Ridley Scott"},
		{Title: "The Thing", Director: "John Carpenter"},
	}
	films = append(films, Film{Title: "Foo", Director: "Bar"})

	// create an entry
	// media := models.Media{MediaFormatID: 1, CategoryID: 1, GenreID: 14, MediaConditionID: 3, SleeveConditionID: 4, ArtistID: 1, AlbumLabelID: 1, Title: "Trilogy", PlayCounter: 1, LastPlayDate: time.Now(), ReleaseDate: 1986}
	// result := initializers.DB.Create(&media) // pass pointer of data to Create

	// // Get the data off request body
	// var body struct {
	// 	Title       string
	// 	PlayCounter int
	// }
	// c.Bind(&body)

	// user.ID             // returns inserted data's primary key
	// result.Error        // returns error
	// result.RowsAffected // returns inserted records count

	// return it
	// if result.Error != nil {
	// 	c.Status(400)
	// 	return
	// }

	c.HTML(http.StatusOK, "sample/index.tmpl", gin.H{
		"Films": films,
	})
}
