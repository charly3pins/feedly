package feedly

import "net/url"

// Feed stores the feed data
type Feed struct {
	// ID string the unique, immutable id of this feed.
	ID string
	// FeedID string same as id; for backward compatibility
	FeedID string
	// Subscribers integer number of Feedly Cloud subscribers who have this feed in their subscription list.
	Subscribers int
	// Title string the feed name.
	Title string
	// Description Optional string the feed description.
	Description string
	// Language Optional string this field is a combination of the language reported by the RSS feed, and the language automatically detected from the feed’s content. It might not be accurate, as many feeds misreport it.
	Language string
	// Velocity Optional float the average number of articles published weekly. This number is updated every few days.
	Velocity float64
	// Website Optional url the website for this feed.
	Website url.URL
	// Topics Optional string array an array of topics this feed covers. This list can be used in searches and mixes to build a list of related feeds and articles. E.g. if the list contains “productivity”, querying “productivity” in feed search will produce a list of related feeds.
	Topics []string
	// State Optional string only returned if the feed cannot be polled. Values include “dead” (cannot be polled), “dead.flooded” (if the feed produces too many articles per day), “dead.dropped” (if the feed has been removed), and “dormant” (if the feed hasn’t been updated in a few months).
	State string
}
