package service

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strings"
	"time"
)

type ScraperChapter struct {
	Title     string
	URL       string
	Content   string
}

func GetChapterList(novelURL string) ([]ScraperChapter, error) {
	// The /ajax/chapters/ endpoint works for this site
	ajaxURL := strings.TrimSuffix(novelURL, "/") + "/ajax/chapters/"

	client := &http.Client{Timeout: 30 * time.Second}
	req, _ := http.NewRequest("POST", ajaxURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	var chapters []ScraperChapter
	doc.Find(".wp-manga-chapter a").Each(func(i int, s *goquery.Selection) {
		href, _ := s.Attr("href")
		title := strings.TrimSpace(s.Text())
		if href != "" {
			chapters = append(chapters, ScraperChapter{
				Title: title,
				URL:   href,
			})
		}
	})

	// Reverse the list (from oldest to newest)
	for i, j := 0, len(chapters)-1; i < j; i, j = i+1, j-1 {
		chapters[i], chapters[j] = chapters[j], chapters[i]
	}

	return chapters, nil
}

func GetChapterContent(chapterURL string) (string, error) {
	client := &http.Client{
		Timeout: 30 * time.Second,
	}
	req, _ := http.NewRequest("GET", chapterURL, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/91.0.4472.124 Safari/537.36")

	res, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer res.Body.Close()

	if res.StatusCode != 200 {
		return "", fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return "", err
	}

	var content []string
	doc.Find(".reading-content .text-left p").Each(func(i int, s *goquery.Selection) {
		text := strings.TrimSpace(s.Text())
		if text != "" {
			content = append(content, text)
		}
	})

	return strings.Join(content, "\n\n"), nil
}
