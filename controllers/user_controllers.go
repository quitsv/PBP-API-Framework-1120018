package controllers

import (
	"framework/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetUsers(c echo.Context) error {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT id, name, password, userType FROM users")

	if err != nil {
		return err
	}

	defer rows.Close()

	var users []models.User

	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Password, &user.UserType); err != nil {
			return err
		}
		users = append(users, user)
	}

	if len(users) < 1 {
		return c.JSON(http.StatusOK, models.MessageResp{Message: "no users found"})
	}
	return c.JSON(http.StatusOK, users)
}

func GetUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id := c.Param("id")

	row := db.QueryRow("SELECT id, name, password, userType FROM users WHERE id = ?", id)

	var user models.User

	if err := row.Scan(&user.Id, &user.Name, &user.Password, &user.UserType); err != nil {
		return c.JSON(http.StatusOK, models.MessageResp{Message: "user not found"})
	}
	return c.JSON(http.StatusOK, user)
}

func Login(c echo.Context) error {
	db := Connect()
	defer db.Close()

	name := c.FormValue("name")
	password := c.FormValue("password")

	row := db.QueryRow("SELECT id, name, password, userType FROM users WHERE name = ? AND password = ?", name, password)

	var user models.User

	if err := row.Scan(&user.Id, &user.Name, &user.Password, &user.UserType); err != nil {
		return c.JSON(http.StatusOK, models.MessageResp{Message: "user not found"})
	}

	generateToken(c, user.Name, user.Password, user.UserType)
	return c.JSON(http.StatusOK, user)
}

func Logout(c echo.Context) error {
	resetToken(c)
	return c.JSON(http.StatusOK, models.MessageResp{Message: "logged out successfully"})
}

func CreateUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	name := c.FormValue("name")
	password := c.FormValue("password")

	_, err := db.Exec("INSERT INTO users(name, password) VALUES(?, ?)", name, password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.MessageResp{Message: "Error creating user"})
	}

	resp := models.MessageResp{Message: "user created successfully"}
	return c.JSON(http.StatusOK, resp)
}

func UpdateUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id := c.Param("id")

	name := c.QueryParam("name")
	password := c.QueryParam("password")

	_, err := db.Exec("UPDATE users SET name = ?, password = ? WHERE id = ?", name, password, id)

	if err != nil {
		return c.JSON(http.StatusOK, models.MessageResp{Message: "user not found"})
	}

	resp := models.MessageResp{Message: "user updated successfully"}
	return c.JSON(http.StatusOK, resp)
}

func DeleteUser(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id := c.Param("id")

	_, err := db.Exec("DELETE FROM users where id = ?", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.MessageResp{Message: "Error deleting user"})
	}

	resp := models.MessageResp{Message: "user deleted succesfully"}
	return c.JSON(http.StatusOK, resp)
}
