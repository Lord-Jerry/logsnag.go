package logsnag

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func NewLogSnag(token string, project string, disableTracking bool) *LogSnag {
	var config = &LogSnag{
		token:            token,
		project:          project,
		trackingDisabled: disableTracking,
		client:           &http.Client{},
	}
	config.Insight = &Insight{LogSnag: config}

	return config
}

func (config *LogSnag) initiateRequest(v *InitiateRequest) error {
	if config.IsDisabled() {
		return nil
	}

	payload, err := json.Marshal(v.Payload)
	if err != nil {
		return err
	}

	req, err := http.NewRequest(v.Method, v.Endpoint, bytes.NewBuffer(payload))
	if err != nil {
		return err
	}

	req.Header.Add("Authorization", "Bearer "+config.token)
	req.Header.Add("Content-Type", "application/json")

	res, err := config.client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		body, err := io.ReadAll(res.Body)
		if err != nil {
			return fmt.Errorf("failed to read response body: %v", err)
		}
		return fmt.Errorf("request failed with status code %d: %s", res.StatusCode, string(body))
	}

	return nil

}

func (config *LogSnag) GetProject() *string { return &config.project }
func (config *LogSnag) IsDisabled() bool    { return config.trackingDisabled }
func (config *LogSnag) DisableTracking()    { config.trackingDisabled = true }
func (config *LogSnag) EnableTracking()     { config.trackingDisabled = false }

func (config *LogSnag) Track(data *TrackData) error {
	return config.initiateRequest(&InitiateRequest{
		Payload: &struct {
			*TrackData
			Project string `json:"project"`
		}{
			TrackData: data,
			Project:   config.project,
		},
		Method:   "POST",
		Endpoint: logEndpoint,
	})
}

func (config *LogSnag) Identify(data *IdentifyData) error {
	return config.initiateRequest(&InitiateRequest{
		Payload: &struct {
			*IdentifyData
			Project string `json:"project"`
		}{
			IdentifyData: data,
			Project:      config.project,
		},
		Method:   "POST",
		Endpoint: identifyEndpoint,
	})
}

func (config *LogSnag) Group(data *GroupData) error {
	return config.initiateRequest(&InitiateRequest{
		Payload: &struct {
			*GroupData
			Project string `json:"project"`
		}{
			GroupData: data,
			Project:   config.project,
		},
		Method:   "POST",
		Endpoint: groupEndpoint,
	})
}

func (config *Insight) Track(data *InsightData) error {
	return config.LogSnag.initiateRequest(&InitiateRequest{
		Payload: &struct {
			*InsightData
			Project string `json:"project"`
		}{
			InsightData: data,
			Project:     config.LogSnag.project,
		},
		Method:   "POST",
		Endpoint: insightEndpoint,
	})

}

func (config *Insight) Increment(data *InsightData) error {
	data.Value = map[string]interface{}{"$inc": data.Value}
	return config.LogSnag.initiateRequest(&InitiateRequest{
		Payload: &struct {
			*InsightData
			Project string `json:"project"`
		}{
			InsightData: data,
			Project:     config.LogSnag.project,
		},
		Method:   "PATCH",
		Endpoint: insightEndpoint,
	})
}
