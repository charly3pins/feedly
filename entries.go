package feedly

const entriesEndpoint = "entries"

// Content stores the content data
type Content struct {
	// Content string the content itself contains sanitized HTML markup
	Content string `json:"content"`
	// Direction “ltr” for left-to-right, “rtl” for right-to-left
	Direction string `json:"direction"`
}

// Link stores the link data
type Link struct {
	// HRef string that contains the URL
	HRef string `json:"href"`
	// Type string with the media type
	Type string `json:"type"`
}

// Origin stores the origin feed from which this article was crawled data
type Origin struct {
	// StreamID	string contain the feed id
	StreamID string `json:"streamId,omitempty"`
	// Title string the feed title
	Title string `json:"title,omitempty"`
	// HTMLURL string the feed's website
	HTMLURL string `json:"htmlUrl,omitempty"`
}

// Visual stores the visual data
type Visual struct {
	// URL string the image URL
	URL string `json:"url,omitempty"`
	// Width int the widht of the image
	Width int `json:"width,omitempty"`
	// Height int the height of the image
	Height int `json:"height,omitempty"`
	// ContentType string the MIME type of the image
	ContentType string `json:"contentType,omitempty"`
}

// Entry stores the entry data
type Entry struct {
	//ID string the unique, immutable ID for this particular article.
	ID string `json:"id"`
	// Title Optional string the article’s title. This string does not contain any HTML markup.
	Title string `json:"title,omitempty"`
	// Content Optional content object the article content. This object typically has two values: “content” for the content itself, and “direction” (“ltr” for left-to-right, “rtl” for right-to-left). The content itself contains sanitized HTML markup.
	Content Content `json:"content,omitempty"`
	// Summary Optional content object the article summary. See the content object above.
	Summary Content `json:"summary,omitempty"`
	// Author Optional string the author’s name
	Author string `json:"author,omitempty"`
	// Crawled timestamp the immutable timestamp, in ms, when this article was processed by the feedly Cloud servers.
	Crawled Time `json:"crawled"`
	// Recrawled Optional timestamp the timestamp, in ms, when this article was re-processed and updated by the feedly Cloud servers.
	Recrawled Time `json:"recrawled,omitempty"`
	// Published timestamp the timestamp, in ms, when this article was published, as reported by the RSS feed (often inaccurate).
	Published Time `json:"published"`
	// Updated Optional timestamp the timestamp, in ms, when this article was updated, as reported by the RSS feed
	Updated Time `json:"updated,omitempty"`
	// Alternate Optional link object array a list of alternate links for this article. Each link object contains a media type and a URL. Typically, a single object is present, with a link to the original web page.
	Alternate []Link `json:"alternate,omitempty"`
	// Origin Optional origin object the feed from which this article was crawled. If present, “streamId” will contain the feed id, “title” will contain the feed title, and “htmlUrl” will contain the feed’s website.
	Origin Origin `json:"origin,omitempty"`
	// Keywords Optional string array a list of keyword strings extracted from the RSS entry.
	Keywords []string `json:"keywords,omitempty"`
	// Visual Optional visual object an image URL for this entry. If present, “url” will contain the image URL, “width” and “height” its dimension, and “contentType” its MIME type.
	Visual Visual `json:"visual,omitempty"`
	// Unread boolean was this entry read by the user? If an Authorization header is not provided, this will always return false. If an Authorization header is provided, it will reflect if the user has read this entry or not.
	Unread bool `json:"unread"`
	// Tags Optional tag object array a list of tag objects (“id” and “label”) that the user added to this entry. This value is only returned if an Authorization header is provided, and at least one tag has been added. If the entry has been explicitly marked as read (not the feed itself), the “global.read” tag will be present.
	Tags []Tag `json:"tags,omitempty"`
	// Categories category object array a list of category objects (“id” and “label”) that the user associated with the feed of this entry. This value is only returned if an Authorization header is provided.
	Categories []Category `json:"categories,omitempty"`
	// Engagement Optional integer an indicator of how popular this entry is. The higher the number, the more readers have read, saved or shared this particular entry.
	Engagement int `json:"engagement,omitempty"`
	// ActionTimestamp Optional timestamp for tagged articles, contains the timestamp when the article was tagged by the user. This will only be returned when the entry is returned through the streams API.
	ActionTimestamp Time `json:"actionTimestamp,omitempty"`
	// Enclosure Optional link object array a list of media links (videos, images, sound etc) provided by the feed. Some entries do not have a summary or content, only a collection of media links.
	Enclosure []Link `json:"enclosure,omitempty"`
	// Fingerprint string the article fingerprint. This value might change if the article is updated.
	Fingerprint string `json:"fingerprint"`
	// OriginID string the unique id of this post in the RSS feed (not necessarily a URL!)
	OriginID string `json:"originId"`
	//SID Optional string an internal search id.
	SID string `json:"sid,omitempty"`
	// Priorities Optional priority object array a list of priority filters that match this entry (pro+ and team only).
	Priorities []Priority `json:"priorities,omitempty"`
}
