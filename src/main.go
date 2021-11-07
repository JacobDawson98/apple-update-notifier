package main

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/mmcdole/gofeed"
)

type ReleaseMessage struct {
	Title   string
	Link    string
	Published string
}

const appleDeveloperReleasesRssFeedUrl string = "https://developer.apple.com/news/releases/rss/releases.rss"

func main() {
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(appleDeveloperReleasesRssFeedUrl)
	releases := []ReleaseMessage{}

	for _, item := range feed.Items {
		if isItemRelatedToMobilePlatformRelease(item.Title) {
			releaseMessage := ReleaseMessage{
				Title:   item.Title,
				Link:    item.Link,
				Published: item.Published,
			}
			releases = append(releases, releaseMessage)
		}
	}

	for _, release := range releases {
		releaseMessage, _ := json.Marshal(release)
		fmt.Println(string(releaseMessage))
	}
}

const xcodeRelease string = "Xcode"
const iOSRelease string = "iOS"
const iPadOSRelease string = "iPadOS"

func isItemRelatedToMobilePlatformRelease(title string) bool {
	return strings.Contains(title, xcodeRelease) ||
		strings.Contains(title, iOSRelease) ||
		strings.Contains(title, iPadOSRelease)
}
