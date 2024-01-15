package logsnag

import (
	"net/http"
	"time"
)

type Insight struct {
	LogSnag *LogSnag
}

type LogSnag struct {
	token            string
	project          string
	trackingDisabled bool
	client           *http.Client
	Insight          *Insight
}

type TrackData struct {
	Channel     string          `json:"channel,omitempty"`
	UserId      string          `json:"user_id,omitempty"`
	Event       string          `json:"event"`
	Description string          `json:"description,omitempty"`
	Icon        string          `json:"icon,omitempty"`
	Tags        *map[string]any `json:"tags,omitempty"`
	Notify      bool            `json:"notify,omitempty"`
	Parser      string          `json:"parser,omitempty"`
	Timestamp   *time.Time      `json:"timestamp,omitempty"`
}

type IdentifyData struct {
	UserId     string          `json:"user_id,omitempty"`
	Properties *map[string]any `json:"properties,omitempty"`
}

type GroupData struct {
	UserId     string          `json:"user_id,omitempty"`
	GroupId    string          `json:"group_id,omitempty"`
	Properties *map[string]any `json:"properties,omitempty"`
}

type InitiateRequest struct {
	Payload  any
	Method   string
	Endpoint string
}

type InsightData struct {
	Value any    `json:"value,omitempty"`
	Title string `json:"title,omitempty"`
	Icon  string `json:"icon,omitempty"`
}
