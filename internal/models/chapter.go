package models

type Chapter struct {
	ID                int    `json:"id"`
	NovelID           int    `json:"novel_id"`
	Title             string `json:"title"`
	Slug              string `json:"slug"`
	SourceURL         string `json:"source_url"`
	ContentRaw        string `json:"content_raw"`
	ContentTranslated string `json:"content_translated"`
	Status            string `json:"status"`
}
