package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)

type ProgrammingLanguage struct {
	Id       string `json:"id"`
	Language string `json:"language"`
	Creator  string `json:"creator"`
}

var languages = []ProgrammingLanguage{
	{Id: "1", Language: "C", Creator: "Dennis Ritchie"},
	{Id: "2", Language: "Java", Creator: "James Gosling"},
	{Id: "3", Language: "C++", Creator: " Bjarne Stroustrup"},
	{Id: "4", Language: "Python", Creator: "Guido van Rossum"},
}

func main() {

	app := fiber.New()
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(languages)
	})

	app.Get("/language", func(ctx *fiber.Ctx) error {
		return ctx.Status(http.StatusOK).JSON(languages)
	})

	app.Get("/language/:id", func(ctx *fiber.Ctx) error {
		Id := ctx.Params("id")

		for _, lang := range languages {
			if lang.Id == Id {
				return ctx.Status(http.StatusOK).JSON(lang)
			}
		}
		return ctx.Status(http.StatusNotFound).SendString("Language not found")
	})

	app.Post("/language", func(ctx *fiber.Ctx) error {
		var language ProgrammingLanguage
		if err := ctx.BodyParser(&language); err != nil {
			ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"Error": err.Error()})

			return err
		}

		languages = append(languages, language)
		return ctx.Status(http.StatusCreated).JSON(language)
	})

	app.Delete("/language/:id", func(ctx *fiber.Ctx) error {
		Id := ctx.Params("id")

		for i, lang := range languages {
			if lang.Id == Id {
				languages = append(languages[:i], languages[i+1:]...)
				break
			}
		}
		return ctx.Status(http.StatusNoContent).JSON(fiber.Map{"data": languages})
	})

	app.Put("/language/:id", func(ctx *fiber.Ctx) error {
		id := ctx.Params("id")

		if id == "" {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Please make sure id"})
		}

		var updateLanguage ProgrammingLanguage
		updateLanguage.Id = id

		if err := ctx.BodyParser(&updateLanguage); err != nil {
			ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
			return err
		}

		if updateLanguage.Language == "" || updateLanguage.Creator == "" {
			return ctx.Status(http.StatusBadRequest).JSON(fiber.Map{"error": "Language and creator is required."})
		}

		for i, lang := range languages {
			if lang.Id == updateLanguage.Id {
				languages = append(languages[:i], languages[i+1:]...)
				languages = append(languages, updateLanguage)
			}
		}

		return ctx.Status(http.StatusOK).JSON(updateLanguage)
	})

	PORT := os.Getenv("PORT")
	app.Listen(PORT)
}
