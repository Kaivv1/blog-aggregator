package scrapper

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Language    string    `xml:"language"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func UrlToFeed(url string) (RSSFeed, error) {
	httpClient := http.Client{
		Timeout: 10 * time.Second,
	}

	resp, err := httpClient.Get(url)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to fetch url %s: %v", url, err)
	}
	defer resp.Body.Close()
	dat, err := io.ReadAll(resp.Body)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to read the resp body %s: %v", url, err)
	}
	rssFeed := RSSFeed{}
	err = xml.Unmarshal(dat, &rssFeed)
	if err != nil {
		return RSSFeed{}, fmt.Errorf("failed to unmarshal rss feed: %v", err)
	}
	return rssFeed, nil
}
