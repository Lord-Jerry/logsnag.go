package logsnag

const baseEndpoint = "https://api.logsnag.com"

var (
	logEndpoint      = baseEndpoint + "/v1/log"
	insightEndpoint  = baseEndpoint + "/v1/insight"
	identifyEndpoint = baseEndpoint + "/v1/identify"
	groupEndpoint    = baseEndpoint + "/v1/group"
)
