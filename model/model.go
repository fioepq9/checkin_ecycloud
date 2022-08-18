package model

type CheckinResponse struct {
	Msg           string           `json:"msg"`
	Ret           int              `json:"ret"`
	Traffic       string           `json:"traffic"`
	TrafficInfo   TrafficInfoModel `json:"trafficInfo"`
	UnFlowtraffic int64            `json:"unFlowtraffic"`
}

type TrafficInfoModel struct {
	LastUsedTraffic  string `json:"lastUsedTraffic"`
	TodayUsedTraffic string `json:"todayUsedTraffic"`
	UnUsedTraffic    string `json:"unUsedTraffic"`
}
