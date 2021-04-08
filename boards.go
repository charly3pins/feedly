package feedly

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
)

const boardsEndpoint = "boards"

// BoardResponse stores the response for the boards endpoints
type BoardResponse struct {
	//ID String the board id.
	ID string `json:"id"`
	// Created Timestamp the EPOCH timestamp when this board was created.
	Created Time `json:"created"`
	//Label String the board label.
	Label string `json:"label"`
	// Customizable Boolean if true, the label, description, cover etc can be changed by the user. If false, this board cannot be changed.
	Customizable bool `json:"customizable"`
	// Enterprise Boolean true for enterprise boards, false for personal boards.
	Enterprise bool `json:"enterprise"`
	// Description Optional String the board description.
	Description string `json:"description,omitempty"`
	// Cover Optional URL the URL of the cover image, if one was uploaded.
	Cover url.URL `json:"cover,omitempty"`
	// IsPublic Optional Boolean if true, this board is publicly shared.
	IsPublic bool `json:"isPublic,omitempty"`
	// ShowNotes Optional Boolean if true, notes are also visible to followers (public boards only).
	ShowNotes bool `json:"showNotes,omitempty"`
	// ShowHighlights Optional Boolean if true, highlights are also visible to followers (public boards only).
	ShowHighlights bool `json:"showHighlights,omitempty"`
	// HTMLURL Optional URL the public URL for this board (public boards only).
	HTMLURL url.URL `json:"htmlUrl,omitempty"`
	// StreamID Optional String the public feed id for this board (public boards only).
	StreamID string `json:"streamId,omitempty"`
}

// ListBoards returns a list of boards. If withEnterprise is true, it return enterprise boards
// followed by the user as well as personal ones.
func (c Client) ListBoards(withEnterprise bool) ([]BoardResponse, error) {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + boardsEndpoint
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
	var boards []BoardResponse
	err = json.Unmarshal(body, &boards)
	if err != nil {
		return nil, err
	}
	return boards, nil
}

// UpdateBoardRequest encapsulates the request payload for the UpdateBoard method
type UpdateBoardRequest struct {
	// ID Optional string
	ID string `json:"id"`
	// Label Optional string
	Label string `json:"label,omitempty"`
	// Description Optional string
	Description string `json:"description,omitempty"`
	// IsPublic Optional boolean
	IsPublic bool `json:"isPublic,omitempty"`
	// ShowNotes Optional Boolean if true, notes are also visible to followers (public boards only).
	ShowNotes bool `json:"showNotes,omitempty"`
	// ShowHighlights Optional Boolean if true, highlights are also visible to followers (public boards only).
	ShowHighlights bool `json:"showHighlights,omitempty"`
}

// UpdateBoard updates a board with the data given in the request
// Note: changing isPublic, showNotes or showHighlights requires a Feedly Pro subscription.
func (c Client) UpdateBoard(u UpdateBoardRequest) error {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + boardsEndpoint
	payload, err := json.Marshal(u)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, bytes.NewReader(payload))
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", "application/json")
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(http.StatusText(resp.StatusCode))
	}
	return nil
}

// UploadCoverImage uploads a new cover image into an existing board.
func (c Client) UploadCoverImage(id string, coverImage io.Reader) error {
	url := c.Config.BaseURL + "/" + c.Config.Version + "/" + boardsEndpoint + "/" + id

	mpm, err := newMultiPartMIME(coverImage)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", url, mpm.bytes)
	if err != nil {
		return err
	}
	req.Header.Add("Authorization", "Bearer "+c.Config.Token)
	req.Header.Add("Content-Type", mpm.contentType)
	resp, err := c.Client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return errors.New(http.StatusText(resp.StatusCode))
	}
	return nil
}

// multiPartMIME encapsulates the bytes and the content type of a file
type multiPartMIME struct {
	bytes       *bytes.Buffer
	contentType string
}

// newMultiPartMIME is a helper to create a multi-part MIME file with the name “cover”
func newMultiPartMIME(attachment io.Reader) (multiPartMIME, error) {
	body := bytes.Buffer{}
	writer := multipart.NewWriter(&body)
	part, err := writer.CreateFormFile("cover", "cover image")
	if err != nil {
		return multiPartMIME{}, err
	}
	if _, err := io.Copy(part, attachment); err != nil {
		return multiPartMIME{}, err
	}
	if err := writer.Close(); err != nil {
		return multiPartMIME{}, err
	}
	return multiPartMIME{&body, writer.FormDataContentType()}, nil
}
