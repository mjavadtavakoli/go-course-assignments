package main

import (
    "github.com/gin-gonic/gin"
    "net/http"
)

type Movie struct {
    ID    int    `json:"id"`
    Title string `json:"title"`
    Year  int    `json:"year"`
}

func main() {
    r := gin.Default()
    
    movies := []Movie{
        {1, "The Matrix", 1999},
        {2, "Inception", 2010},
        {3, "Interstellar", 2014},
    }
    
    r.GET("/movies", func(c *gin.Context) {
        c.JSON(http.StatusOK, movies)
    })
    
    r.Run(":8080")
}