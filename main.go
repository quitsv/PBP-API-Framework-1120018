package main

import (
	"framework/controllers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	//examples
	e.GET("/", welcome)

	//functions for the controllers
	e.GET("/items", controllers.GetItems)          // GET all items from the database
	e.GET("/items/:id", controllers.GetItem)       // :id is a placeholder for the item id
	e.PUT("/items/:id", controllers.UpdateItem)    //update an item by id with queryParams of name and price
	e.POST("/items", controllers.CreateItem)       //create an item with form values of name and price
	e.DELETE("/items/:id", controllers.DeleteItem) //delete an item by id

	//start the server
	e.Logger.Fatal(e.Start(":8080")) //start the server on port 8080
}

//example for the welcome page that will return a string
func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to my API. Created by @quitsv (Axell Silvano)")
}
