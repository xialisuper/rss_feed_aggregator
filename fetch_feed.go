package main

import (
	"context"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"sync"
	"time"

	"github.com/xialisuper/rss_feed_aggregator/internal/database"
)

// Worker represents a continuous feed fetching worker
type worker struct {
	Feeds []database.Feed
}

// RSS struct represents the entire RSS feed
type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Channel Channel  `xml:"channel"`
}

// Channel struct represents the channel element in RSS feed
type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

// Item struct represents each item in the channel
type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func StartFeedFetcher(cfg *apiConfig) {
	for {

		// 创建一个带有取消功能的context
		ctx := context.Background()

		// 从数据库中获取下一个要抓取的10个feed
		feeds, err := cfg.DB.GetNextFeedsToFetch(ctx, int32(10))

		if err != nil {
			panic(err)
		}

		w := worker{
			Feeds: feeds,
		}

		// 创建一个WaitGroup
		var wg sync.WaitGroup
		wg.Add(len(w.Feeds))

		w.fetchFeeds(&wg, cfg)

		wg.Wait()

		fmt.Println("All 10 feeds fetched. Waiting for 60 seconds...")
		time.Sleep(60 * time.Second)
	}
}

// fetchFeeds continuously fetches feeds
func (w *worker) fetchFeeds(wg *sync.WaitGroup, cfg *apiConfig) {
	for _, feed := range w.Feeds {

		fmt.Println("Fetching feed from:", feed)
		rss, err := fetchRSSFeed(feed.Url)
		wg.Done()

		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		// Process the RSS feed
		fmt.Println("Processing feed...")
		// logoutFeeds(rss)

		// Update the database with the latest feed data
		for _, item := range rss.Channel.Items {
			publishAt, err := parseRSSFeedDate(item.PubDate)
			if err != nil {
				fmt.Println("Error parsing date format:", rss.Channel.Items[0].PubDate, err)
				continue
			}
			err = cfg.DB.CreatePost(context.Background(), database.CreatePostParams{
				//todo
				Title:       item.Title,
				Url:         item.Link,
				Description: item.Description,
				PublishedAt: publishAt,
				FeedID:      feed.ID,
			})
	
			if err != nil {
				fmt.Println("Error updating database:", err)
				continue
			}
		}

		
	}
}

func parseRSSFeedDate(date string) (time.Time, error) {
	t, err := time.Parse(time.RFC1123, date)
	if err != nil {
		return time.Time{}, err
	}
	return t, nil
}

// func logoutFeeds(rss RSS) {

// 	fmt.Println("Channel Title:", rss.Channel.Title)
// 	// fmt.Println("Channel Link:", rss.Channel.Link)
// 	fmt.Println("Channel Description:", rss.Channel.Description)
// 	fmt.Println("Items:")
// 	for _, item := range rss.Channel.Items {
// 		fmt.Println("  Title:", item.Title)
// 		// fmt.Println("  Link:", item.Link)
// 		fmt.Println("  Description:", item.Description)
// 		// fmt.Println("  PubDate:", item.PubDate)
// 	}

// }

func fetchRSSFeed(url string) (RSS, error) {
	var rss RSS

	// Make an HTTP GET request to fetch the RSS feed
	response, err := http.Get(url)
	if err != nil {
		return rss, fmt.Errorf("failed to fetch RSS feed: %v", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return rss, fmt.Errorf("failed to read response body: %v", err)
	}

	// Unmarshal the XML data into the RSS struct
	err = xml.Unmarshal(body, &rss)
	if err != nil {
		return rss, fmt.Errorf("failed to unmarshal XML: %v", err)
	}

	return rss, nil
}
