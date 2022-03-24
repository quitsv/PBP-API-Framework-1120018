package main

import (
	"framework/controllers"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()

	//cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
		AllowCredentials: true,
	}))

	//middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	//examples
	e.GET("/", welcome)

	//routes
	e.GET("/users", controllers.GetUsers)                                    // GET all user from the database
	e.GET("/users/:id", controllers.GetUser)                                 // :id is a placeholder for the user id
	e.PUT("/users/:id", controllers.UpdateUser)                              //update an user by id with queryParams of name and password
	e.POST("/users", controllers.CreateUser)                                 //create an user with form values of name and password
	e.DELETE("/users/:id", controllers.Authenticate(controllers.DeleteUser)) //delete an user by id with authentication
	e.POST("/users/login", controllers.Login)                                //login with form values of name and password
	e.POST("/users/logout", controllers.Logout)                              //logout user

	//start the server
	e.Logger.Fatal(e.Start(":8080")) //start the server on port 8080
}

//example for the welcome page that will return a string
func welcome(c echo.Context) error {
	return c.String(http.StatusOK, "Welcome to my API. Created by @quitsv (Axell Silvano)")
}
