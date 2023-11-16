package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"errors"
)

type spart struct {
	ID       	string `json:"id"`
	ProductName string `json:"productname"`
	Price   	int `json:"price"`
	Quantity 	int `json:"quantity"`
}

var sparts = []spart {
	{ID: "1", ProductName: "Knalpot", Price: 500000, Quantity: 5},
	{ID: "2", ProductName: "Oli", Price: 50000, Quantity: 5},
	{ID: "3", ProductName: "Kaca Spion", Price: 125000, Quantity: 6},
}

func getSparts(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, sparts)
}

func spartById(c * gin.Context) {
	id := c.Param("id")
	spart, err := getSpartById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Spare Part Tidak Ditemukan."})
		return
	}
	c.IndentedJSON(http.StatusOK, spart)
}

func deleteSpart(c * gin.Context) {
	id, ok := c.GetQuery("id")
	
	if !ok {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID Tidak Ditemukan"})
	}

	spart, err := getSpartById(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Spare Part Tidak Ditemukan."})
		return
	}

	if spart.Quantity <= 0 {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Spare Part Sudah Habis."})
		return
	}

	spart.Quantity -= 1
	c.IndentedJSON(http.StatusOK, spart)
}

// func returnSpart(c * gin.Context) {
// 	id, ok := c.GetQuery("id")
	
// 	if !ok {
// 		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "ID Tidak Ditemukan"})
// 	}

// 	book, err := getSpartById(id)

// 	if err != nil {
// 		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Spare Part Tidak Ditemukan."})
// 		return
// 	}

// 	book.Quantity += 1 
// 	c.IndentedJSON(http.StatusOK, book)
// }

func getSpartById(id string) (*spart, error) {
	for i, b := range sparts {
		if b.ID == id {
			return &sparts[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func postSparts(c * gin.Context) {
	var newSpart spart

	if err := c.BindJSON(&newSpart); err != nil {
		return
	}

	sparts = append(sparts, newSpart)
	c.IndentedJSON(http.StatusCreated, newSpart)
}

func main() {
	router := gin.Default()
	router.GET("/sparts", getSparts)
	router.GET("/sparts/:id", spartById)
	router.POST("/sparts", postSparts)
	router.PATCH("/delete", deleteSpart)
	// router.PATCH("/return", returnSpart)

	router.Static("/css", "./css")
	router.Static("/assets", "./assets")
	router.StaticFile("/home", "homepage.html")
	router.StaticFile("/item1", "itempage1.html")
	router.StaticFile("/item2", "itempage2.html")
	router.StaticFile("/item3", "itempage3.html")
	router.StaticFile("/item4", "itempage4.html")
	router.StaticFile("/cart", "cart.html")

	router.Run("localhost:8080")
}