package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}

	router := gin.New()
	router.Use(gin.Logger())
	router.LoadHTMLGlob("templates/*.tmpl.html")
	router.Static("/static", "static")

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.Run(":" + port)

	// slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// theMap := map[int]string{
	// 	1: "Dom",
	// 	2: "Tina",
	// 	3: "Mia",
	// 	4: "Christian",
	// 	5: "Luka",
	// }
	//
	// fmt.Println(theMap)
	//
	// fmt.Println("slice:", slice)
	//
	// sliceOfSlice25 := slice[2:5]
	// sliceOfSlice47 := slice[4:7]
	//
	// fmt.Println("sliceOfSlice25:", sliceOfSlice25)
	// fmt.Println("sliceOfSlice47:", sliceOfSlice47)
	//
	// sliceCopy := make([]int, 3)
	// copy(sliceCopy, slice[3:6])
	//
	// slice[4] = 11
	//
	// fmt.Println("slice:", slice)
	// fmt.Println("sliceOfSlice25:", sliceOfSlice25)
	// fmt.Println("sliceOfSlice47:", sliceOfSlice47)
	// fmt.Println("sliceCopy:", sliceCopy)
	//
}
