package controllers

import (
	"framework/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func GetItems(c echo.Context) error {
	db := Connect()
	defer db.Close()

	rows, err := db.Query("SELECT * FROM items")

	if err != nil {
		return err
	}

	defer rows.Close()

	var items []models.Item

	for rows.Next() {
		var item models.Item
		if err := rows.Scan(&item.Id, &item.Name, &item.Price); err != nil {
			return err
		}
		items = append(items, item)
	}

	return c.JSON(http.StatusOK, items)
}

func GetItem(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id := c.Param("id")

	row := db.QueryRow("SELECT id, name, price FROM items WHERE id = ?", id)

	var item models.Item

	if err := row.Scan(&item.Id, &item.Name, &item.Price); err != nil {
		return c.JSON(http.StatusOK, models.MessageResp{Message: "Item not found"})
	}
	return c.JSON(http.StatusOK, item)

}

func UpdateItem(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id := c.Param("id")

	name := c.QueryParam("name")
	price := c.QueryParam("price")

	_, err := db.Exec("UPDATE items SET name = ?, price = ? WHERE id = ?", name, price, id)

	if err != nil {
		return c.JSON(http.StatusOK, models.MessageResp{Message: "Item not found"})
	}

	resp := models.MessageResp{Message: "Item updated successfully"}
	return c.JSON(http.StatusOK, resp)
}

func CreateItem(c echo.Context) error {
	db := Connect()
	defer db.Close()

	name := c.FormValue("name")
	price := c.FormValue("price")

	_, err := db.Exec("INSERT INTO items(name, price) VALUES(?, ?)", name, price)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.MessageResp{Message: "Error creating item"})
	}

	resp := models.MessageResp{Message: "Item created successfully"}
	return c.JSON(http.StatusOK, resp)
}

func DeleteItem(c echo.Context) error {
	db := Connect()
	defer db.Close()

	id := c.Param("id")

	_, err := db.Exec("DELETE FROM items where id = ?", id)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.MessageResp{Message: "Error deleting item"})
	}

	resp := models.MessageResp{Message: "Item deleted succesfully"}
	return c.JSON(http.StatusOK, resp)
}
