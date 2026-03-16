package handler

import (
	"log"
	"novel_translate_indonesia/internal/models"
	"novel_translate_indonesia/internal/repository"
	"novel_translate_indonesia/internal/service"

	"github.com/gofiber/fiber/v2"
)

type ChapterHandler struct {
	Repo       *repository.ChapterRepository
	Translator *service.Translator
}

func NewChapterHandler(repo *repository.ChapterRepository, translator *service.Translator) *ChapterHandler {
	return &ChapterHandler{
		Repo:       repo,
		Translator: translator,
	}
}

func (h *ChapterHandler) Index(c *fiber.Ctx) error {
	chapters, err := h.Repo.GetAll()
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	return c.Render("index", fiber.Map{
		"Chapters": chapters,
	}, "layout")
}

func (h *ChapterHandler) Sync(c *fiber.Ctx) error {
	novelURL := "https://meionovels.com/novel/kusuriya-no-hitorigoto-ln/"
	chapters, err := service.GetChapterList(novelURL)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	for _, ch := range chapters {
		err := h.Repo.Save(models.Chapter{
			Title:     ch.Title,
			SourceURL: ch.URL,
		})
		if err != nil {
			log.Println("Error inserting chapter:", err)
		}
	}

	return c.SendString("<div class='alert alert-success'>Chapters synced! Refreshing...</div>")
}

func (h *ChapterHandler) Show(c *fiber.Ctx) error {
	id := c.Params("id")
	ch, err := h.Repo.GetByID(id)
	if err != nil {
		return c.Status(404).SendString("Chapter not found")
	}

	if ch.ContentRaw == "" {
		content, err := service.GetChapterContent(ch.SourceURL)
		if err == nil {
			ch.ContentRaw = content
			h.Repo.UpdateContent(id, content)
		}
	}

	return c.Render("chapter", fiber.Map{
		"Chapter": ch,
	}, "layout")
}

func (h *ChapterHandler) Translate(c *fiber.Ctx) error {
	id := c.Params("id")
	ch, err := h.Repo.GetByID(id)
	if err != nil {
		return c.Status(404).SendString("Chapter not found")
	}

	translated, err := h.Translator.Translate(ch.ContentRaw)
	if err != nil {
		return c.Status(500).SendString(err.Error())
	}

	h.Repo.UpdateRefined(id, translated)

	return c.SendString(translated)
}
