package repository

import (
	"novel_translate_indonesia/internal/database"
	"novel_translate_indonesia/internal/models"
)

type ChapterRepository struct{}

func NewChapterRepository() *ChapterRepository {
	return &ChapterRepository{}
}

func (r *ChapterRepository) GetAll() ([]models.Chapter, error) {
	var chapters []models.Chapter
	rows, err := database.DB.Query("SELECT id, title, COALESCE(slug, ''), COALESCE(status, 'pending') FROM chapters")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var ch models.Chapter
		err := rows.Scan(&ch.ID, &ch.Title, &ch.Slug, &ch.Status)
		if err != nil {
			return nil, err
		}
		chapters = append(chapters, ch)
	}
	return chapters, nil
}

func (r *ChapterRepository) GetByID(id string) (models.Chapter, error) {
	var ch models.Chapter
	err := database.DB.QueryRow("SELECT id, title, COALESCE(content_raw, ''), COALESCE(content_translated, ''), source_url FROM chapters WHERE id = ?", id).
		Scan(&ch.ID, &ch.Title, &ch.ContentRaw, &ch.ContentTranslated, &ch.SourceURL)
	return ch, err
}

func (r *ChapterRepository) Save(ch models.Chapter) error {
	_, err := database.DB.Exec("INSERT OR IGNORE INTO chapters (title, slug, source_url) VALUES (?, ?, ?)", ch.Title, ch.Title, ch.SourceURL)
	return err
}

func (r *ChapterRepository) UpdateContent(id string, content string) error {
	_, err := database.DB.Exec("UPDATE chapters SET content_raw = ? WHERE id = ?", content, id)
	return err
}

func (r *ChapterRepository) UpdateRefined(id string, content string) error {
	_, err := database.DB.Exec("UPDATE chapters SET content_translated = ?, status = 'refined' WHERE id = ?", content, id)
	return err
}
