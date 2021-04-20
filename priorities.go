package feedly

// Filter stores the filter data
type Filter struct {
	// Type filter type “matches” for entity/topic/phrases filtering; “likeBoard” for like-board filtering; “security” for vulnerability severity filtering
	Type string `json:"type"`
	// Parts array of ids an array of ids of the relevant type (“parts” for “matches”, “boards” for “likeBoard”, “severities” for “security”)
	Parts []string `json:"parts"`
	// Salience Optional about or mention for “entities” type only; indicate if the entries should be about the entities, or just mention the entities.
	Salience string `json:"salience,omitempty"`
}

// Priority stores the priority data
type Priority struct {
	// ID filter id the priority filter id.
	ID string `json:"id"`
	// Label string the priority filter label.
	Label string `json:"label"`
	// Layers list of filters the search filters used to find entries in the selected category or feed.
	Layers []Filter `json:"filters"`
	// StreamIDs list of category ids the personal category id on which this priority filter applies. Only one category is allowed.
	StreamIDs []string `json:"streamIds"`
	// Active Optional boolean is the importance filter active? (default: true).
	Active bool `json:"active,omitempty"`
	// ActiveUntil Optional timestamp time limit for this importance filter. After this date, the importance filter will not be refreshed.
	ActiveUntil Time `json:"activeUntil,omitempty"`
	// LastUpdated timestamp the last time the search query was run
	LastUpdated Time `json:"lastUpdated"`
	// LastEntryMatch timestamp the timestamp of the newest entry that matched the search query
	LastEntryMatch Time `json:"lastEntryMatch"`
	// NextRun Optional timestamp the next time the search query will run
	NextRun Time `json:"nextRun,omitempty"`
	// NumEntriesProcessed Optional number the number of entries that were processed by this priority filter over the past week.
	NumEntriesProcessed int `json:"numEntriesProcessed,omitempty"`
	// NumEntriesMatching Optional number the number of entries that were prioritized by this filter over the past week.
	NumEntriesMatching int `json:"numEntriesMatching,omitempty"`
}
