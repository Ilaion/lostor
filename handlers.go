package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/Ilaion/lostor/internal/models"
)

func (app *App) getAllSeries(c *gin.Context) {
	series, err := app.series.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintln(err)})
		return
	}

	c.JSON(http.StatusOK, series)
}

func (app *App) getSeries(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintln(err)})
		return
	}

	s, err := app.series.Get(int(id))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "series not found"})
		return
	}

	c.JSON(http.StatusOK, s)
}

func (app *App) addSeries(c *gin.Context) {
	var s models.Series

	err := c.BindJSON(&s)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintln(err)})
		return
	}

	_, err = app.series.Add(s)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintln(err)})
		return
	}

	c.AbortWithStatus(http.StatusCreated)
}

func (app *App) deleteSeries(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintln(err)})
		return
	}

	err = app.series.Delete(id)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "series not found"})
		return
	}

	c.AbortWithStatus(http.StatusOK)
}

func (app *App) getAllTorrents(c *gin.Context) {
	ts, err := app.torrents.All()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": fmt.Sprintln(err)})
		return
	}

	c.JSON(http.StatusOK, ts)
}
