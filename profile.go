package feedly

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

const profileEndpoint = "profile"

// ProfileResponse stores the response for the profile endpoints
type ProfileResponse struct {
	//ID string the unique, immutable user id.
	ID string `json:"id"`
	// Email Optional string the email address extracted from the OAuth profile. Not always available, depending on the OAuth method used.
	Email string `json:"email,omitempty"`
	// GivenName Optional string the given (first) name. Not always available.
	GivenName string `json:"givenName,omitempty"`
	// FamilyName Optional string the family (last) name. Not always available.
	FamilyName string `json:"familyName,omitempty"`
	// Fullname Optional string the full name. Not always available.
	Fullname string `json:"fullName,omitempty"`
	// Picture Optional url a picture URL for this user, extracted from the OAuth profile.
	Picture string `json:"picture,omitempty"`
	// Gender Optional string “male” or “female”
	Gender string `json:"gender,omitempty"`
	// Locale Optional locale the locale, extracted from the OAuth profile.
	Locale string `json:"locale,omitempty"`
	// Google Optional string the Google user id, if the user went through Google’s OAuth flow.
	Google string `json:"google,omitempty"`
	// Reader Optional string the Google Reader user id. If present, this indicates a user who migrated from Google Reader.
	Reader string `json:"reader,omitempty"`
	// Twitter Optional string the Twitter handle (legacy).
	Twitter string `json:"twitter,omitempty"`
	// TwitterUserID Optional string the Twitter user id, if the user went through the Twitter OAuth flow.
	TwitterUserID string `json:"twitterUserId,omitempty"`
	// FacebookUserID Optional string the Facebook user id, if the user went through the Facebook OAuth flow.
	FacebookUserID string `json:"facebookUserId,omitempty"`
	// WordPressID Optional string the WordPress user id, if the user went through the WordPress OAuth flow.
	WordPressID string `json:"wordPressId,omitempty"`
	// WindowsLiveID Optional string the Windows Live user id, if the user went through the Windows Live OAuth flow.
	WindowsLiveID string `json:"windowsLiveIdessId,omitempty"`
	// Wave string the analytics “wave”. Format is: “yyyy.ww” where yyyy is the year, ww is the week number. E.g. “2014.02” means this user joined on the second week of 2014. See http://www.epochconverter.com/date-and-time/weeknumbers-by-year.php for week number definitions.
	Wave string `json:"wave,omitempty"`
	// Client string the client application used to create this account.
	Client string `json:"client"`
	// Source string the client name/version used to create this account.
	Source string `json:"source"`
	// Created Optional timestamp the timestamp, in ms, when this account was created. Not set for accounts created before 10/2/2013.
	Created Time `json:"created"`
	// Pro accounts only
	// Product Optional string the feedly pro subscription. Values include FeedlyProMonthly, FeedlyProYearly, FeedlyProLifetime etc.
	Product string `json:"product,omitempty"`
	// ProductExpiration Optional timestamp for expiring subscriptions only; the timestamp, in ms, when this subscription will expire.
	ProductExpiration string `json:"productExpiration,omitempty"`
	// SubscriptionStatus Optional string for expiring subscriptions only; values include Active, PastDue, Canceled, Unpaid, Deleted, Expired.
	SubscriptionStatus string `json:"subscriptionStatus,omitempty"`
	// IsEvernoteConnected Optional boolean true if the user has activated the Evernote integration.
	IsEvernoteConnected bool `json:"isEvernoteConnected,omitempty"`
	// IsPocketConnected Optional boolean true if the user has activated the Pocket integration.
	IsPocketConnected bool `json:"isPocketConnected,omitempty"`
}

// GetProfile returns the profile of the logged user from the Access Token
func (c Client) GetProfile() (ProfileResponse, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + profileEndpoint
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return ProfileResponse{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return ProfileResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ProfileResponse{}, err
	}
	var pResponse ProfileResponse
	err = json.Unmarshal(body, &pResponse)
	if err != nil {
		return ProfileResponse{}, err
	}
	return pResponse, nil
}

// UpdateProfileRequest encapsulates the request payload for the UpdateProfile endpoint
type UpdateProfileRequest struct {
	// Email Optional string
	Email string
	// GivenName Optional string
	GivenName string
	// FamilyName Optional string
	FamilyName string
	// Picture Optional string
	Picture string
	//Gender Optional boolean
	Gender bool
	// Locale Optional string
	Locale string
	// Twitter Optional string twitter handle. example: edwk
	Twitter string
	// Facebook Optional string facebook id
	Facebook string
}

// UpdateProfile updates the profile with the data given in the request
func (c Client) UpdateProfile(u UpdateProfileRequest) (ProfileResponse, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + profileEndpoint
	payload, err := json.Marshal(u)
	if err != nil {
		return ProfileResponse{}, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return ProfileResponse{}, err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return ProfileResponse{}, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ProfileResponse{}, err
	}
	var pResponse ProfileResponse
	err = json.Unmarshal(body, &pResponse)
	if err != nil {
		return ProfileResponse{}, err
	}
	return pResponse, nil
}
