package handler

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/AdluAghnia/go-artic/db"
	"github.com/AdluAghnia/go-artic/models"
	"github.com/gofiber/fiber/v2"
)

func IndexHandler(c *fiber.Ctx) error {
	articles, err := models.GetArticles()
	if err != nil {
		log.Println(err)
		return c.SendString(err.Error())
	}
	return c.Render("index", fiber.Map{
		"Articles": articles,
	}, "layouts/main")
}

func ShowArticle(c *fiber.Ctx) error {
	article, err := models.GetArticleByID(c.Params("id"))
	if err != nil {
		log.Println(err)
		return c.SendString(err.Error())
	}
	return c.Render("show-article", article, "layouts/main")
}

func CreateArticleHandler(c *fiber.Ctx) error {
	return c.Render("create-article", nil)
}

func SaveArticleHandler(c *fiber.Ctx) error {
	title := c.FormValue("title")
	content := c.FormValue("content")
	picture, err := c.FormFile("picture")

	if err != nil {
		log.Println(err)
		return c.SendString(err.Error())
	}

	// Get the file extension
	extension := filepath.Ext(picture.Filename)

	// Hash the filename with SHA-256 and append timestamp for uniqueness
	hasher := sha256.New()
	hasher.Write([]byte(picture.Filename + time.Now().String()))
	encryptedFilename := hex.EncodeToString(hasher.Sum(nil))
	imagePath := fmt.Sprintf("./public/images/%s%s", encryptedFilename, extension)

	if err := c.SaveFile(picture, imagePath); err != nil {
		log.Println(err)
		return c.SendString(err.Error())
	}

	article := models.NewArticle(
		title,
		content,
		encryptedFilename+extension,
	)

	db, err := db.NewDB()
	if err != nil {
		log.Println(err)
		return c.SendString(err.Error())
	}

	if err := article.SaveArticle(db); err != nil {
		log.Println(err)
		return c.SendString(err.Error())
	}

	return c.Render("article", article)
}
