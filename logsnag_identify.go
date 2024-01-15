package logsnag

//type Insight struct {
//	LogSnag *LogSnag
//}

//func (config *Insight) Track(data *InsightTrackData) error {
//	if config.LogSnag.IsDisabled() {
//		return nil
//	}
//
//	data.Project = config.LogSnag.project
//	payload, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//
//	req, err := http.NewRequest("POST", insightEndpoint, bytes.NewBuffer(payload))
//	if err != nil {
//		return err
//	}
//	return config.LogSnag.initiateRequest(req)
//}
//
//func (config *Insight) Increment(data *InsightIncrementData) error {
//	if config.LogSnag.IsDisabled() {
//		return nil
//	}
//
//	data.Project = config.LogSnag.project
//	data.Value = map[string]interface{}{"$inc": data.Value}
//
//	payload, err := json.Marshal(data)
//	if err != nil {
//		return err
//	}
//
//	req, err := http.NewRequest("PATCH", insightEndpoint, bytes.NewBuffer(payload))
//	if err != nil {
//		return err
//	}
//	return config.LogSnag.initiateRequest(req)
//}
