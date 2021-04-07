package feedly

import (
	"encoding/json"
	"time"
)

// Time creates a custom type for time.Time in order to marshal/unmarshal correctly the timestamps
type Time time.Time

// MarshalJSON implements the json.Marshaler interface
func (t Time) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(t).Unix() * 1000)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (t *Time) UnmarshalJSON(b []byte) error {
	var updated *int64
	err := json.Unmarshal(b, &updated)
	if err != nil {
		return err
	}
	if updated != nil {
		*t = Time(time.Unix(*updated/1000, 0))
	}
	return nil
}
