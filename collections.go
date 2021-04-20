package feedly

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

const collectionsEndpoint = "collections"

// Collection stores the collection data
type Collection struct {
	//ID String the collection id.
	ID string `json:"id"`
	// Created Timestamp the EPOCH timestamp when this collection was created.
	Created Time `json:"created"`
	//Label String the collection label. Default value is the collection name.
	Label string `json:"label"`
	// Description Optional String the description description, if defined.
	Description string `json:"description,omitempty"`
	// Cover Optional URL the URL of the cover image, if one was uploaded.
	Cover url.URL `json:"cover,omitempty"`
	// Feeds List of feeds the list of feeds in this collection.
	Feeds []Feed `json:"feeds"`
}

// ListCollections returns a list of collections.
// If withStats is true, it returns reading and tag stats for the past 31 days (default: false)
// If withEnterprise is true, it returns enterprise collections followed by the user as well as personal ones.
func (c Client) ListCollections(withStats, withEnterprise bool) ([]Collection, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint
	if withStats {
		url += "&withStats=true"
	}
	if withEnterprise {
		url += "&withEnterprise=true"
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var collections []Collection
	err = json.Unmarshal(body, &collections)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

// GetCollection returns details about a personal collection.
func (c Client) GetCollection(id string) (Collection, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint + "/" + id
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return Collection{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return Collection{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Collection{}, err
	}
	var collection Collection
	err = json.Unmarshal(body, &collection)
	if err != nil {
		return Collection{}, err
	}
	return collection, nil
}

// CreateOrUpdateCollectionRequest encapsulates the request payload for the CreateCollection and UpdateCollection methods.
type CreateOrUpdateCollectionRequest struct {
	// Label String the unique label for this collection; required for new categories, optional when editing an existing category.
	Label string `json:"label"`
	// ID Optional String the collection id. If missing, the server will generate one (new collection).
	ID string `json:"id,omitempty"`
	// Description Optional String a more detailed description for this collection.
	Description string `json:"description,omitempty"`
	// Feeds Optional list of feeds a list of feeds to be added to this collection.
	Feeds []Feed `json:"feeds,omitempty"`
	// DeleteCover Optional Boolean if true, the existing cover for this collection will be removed.
	DeleteCover bool `json:"deleteCover,omitempty"`
}

// CreateCollection creates a personal collection.
func (c Client) CreateCollection(col CreateOrUpdateCollectionRequest) (Collection, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint
	return c.createOrUpdateCollection(url, col)
}

// UpdateCollection updates a personal collection.
func (c Client) UpdateCollection(id string, col CreateOrUpdateCollectionRequest) (Collection, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint + "/" + id
	return c.createOrUpdateCollection(url, col)
}

func (c Client) createOrUpdateCollection(url string, col CreateOrUpdateCollectionRequest) (Collection, error) {
	payload, err := json.Marshal(col)
	if err != nil {
		return Collection{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return Collection{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return Collection{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Collection{}, err
	}
	var collection Collection
	err = json.Unmarshal(body, &collection)
	if err != nil {
		return Collection{}, err
	}
	return collection, nil
}

// UploadCollectionCoverImage uploads a new cover image into an existing peresonal collection.
func (c Client) UploadCollectionCoverImage(id string, coverImage io.Reader) (Collection, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint + "/" + id
	mpm, err := newMultiPartMIME(coverImage)
	if err != nil {
		return Collection{}, err
	}
	req, err := http.NewRequest("POST", url, mpm.bytes)
	if err != nil {
		return Collection{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return Collection{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Collection{}, err
	}
	var collection Collection
	err = json.Unmarshal(body, &collection)
	if err != nil {
		return Collection{}, err
	}
	return collection, nil
}

// AddFeedRequest encapsulates the request payload for the AddFeedToCollection and AddMultipleFeedToCollection methods.
type AddFeedRequest struct {
	// ID String the feed id.
	ID string `json:"id"`
	// Title Optional String the feed title; if missing, the default feed title will be used.
	Title string `json:"title,omitempty"`
}

// AddFeedToCollection adds a feed to a personal collection.
func (c Client) AddFeedToCollection(collectionID string, f AddFeedRequest) ([]Feed, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint + "/" + collectionID + "/feeds"
	payload, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feeds []Feed
	err = json.Unmarshal(body, &feeds)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

// AddMultipleFeedToCollection adss multiple feeds to a personal collection.
func (c Client) AddMultipleFeedToCollection(collectionID string, f []AddFeedRequest) ([]Feed, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint + "/" + collectionID + "/feeds/.mput"
	payload, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("PUT", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feeds []Feed
	err = json.Unmarshal(body, &feeds)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

// DeleteFeedFromCollection renives a feed from a personal collection.
func (c Client) DeleteFeedFromCollection(collectionID, feedID string) ([]Feed, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint + "/" + collectionID + "/feeds/" + feedID
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feeds []Feed
	err = json.Unmarshal(body, &feeds)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}

// DeleteFeedRequest encapsulates the request payload for the DeleteMultipleFeedFromCollection method.
type DeleteFeedRequest struct {
	// ID String the feed id.
	ID string `json:"id"`
}

// DeleteMultipleFeedFromCollection removes multiple feeds from a personal collection.
func (c Client) DeleteMultipleFeedFromCollection(collectionID string, f []DeleteFeedRequest) ([]Feed, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + collectionsEndpoint + "/" + collectionID + "/feeds/.mdelete"
	payload, err := json.Marshal(f)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("DELETE", url, bytes.NewReader(payload))
	if err != nil {
		return nil, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feeds []Feed
	err = json.Unmarshal(body, &feeds)
	if err != nil {
		return nil, err
	}
	return feeds, nil
}
