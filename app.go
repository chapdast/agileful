package agileful

import (
	"encoding/json"
	"fmt"
	"github.com/chapdast/agileful/db"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"strconv"
)

var DB db.Database

func Run(data db.Database) (*fiber.App, error) {
	if data == nil {
		return nil, fmt.Errorf("nil db connection")
	}

	DB = data
	app := fiber.New()
	///QUERY :page, order=direction ,writer
	app.Get("/article", Articles)
	app.Post("/article", CreateArticle)
	app.Put("/article", EditArticle)
	app.Patch("/article/:id", UpdateViewCount)
	app.Delete("/article/:id", DeleteArticle)
	return app, nil
}

func Articles(c *fiber.Ctx) error {
	writer := c.Query("writer", "")
	page := c.Query("page", "1")
	order := c.Query("order", "")
	direction := c.Params("direction", "asc")
	pageSize := 3
	var opts []db.Option
	opts = append(opts, db.OptionLimit(pageSize))
	if writer != "" {
		opts = append(opts, db.OptionFilter("writer", db.Condition{
			Value:    writer,
			Operator: db.OprSubString,
		}))
	}
	if page != "" {
		p, err := strconv.Atoi(page)
		if err == nil {
			offset := p * pageSize
			opts = append(opts, db.OptionOffset(offset))
		}
	}
	if order != "" {
		desc := false
		if direction == "desc" {
			desc = true
		}
		opts = append(opts, db.OptionOrder(order, desc))
	}
	articles, err := DB.Read(c.Context(), opts...)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	if len(articles) == 0 {
		return c.SendString(`{}`)
	}
	articlesJson, err := json.Marshal(articles)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendString(string(articlesJson))
}

func CreateArticle(c *fiber.Ctx) error {
	article := db.Article{}
	if err := json.Unmarshal(c.Body(), &article); err != nil {
		return err
	}
	err := DB.Create(c.Context(), &article)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}
func EditArticle(c *fiber.Ctx) error {
	article := db.Article{}
	if err := json.Unmarshal(c.Body(), &article); err != nil {
		return err
	}
	err := DB.Update(c.Context(), &article)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}
func UpdateViewCount(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	err = DB.UpdateViewCount(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}
func DeleteArticle(c *fiber.Ctx) error {
	idStr := c.Params("id")
	if idStr == "" {
		return c.SendStatus(http.StatusBadRequest)
	}
	id, err := strconv.Atoi(idStr)
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}
	log.Println("ID", id)
	err = DB.Delete(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}
	return c.SendStatus(http.StatusOK)
}
