package models

type EconomicEvents struct {
	Events    []EconomicEvent `json:"events"`
	Info      string          `json:"info"`
	NumEvents int             `json:"num-events"`
}

type EconomicEvent struct {
	Actual     string `json:"Actual"`
	Currency   string `json:"Currency"`
	Datetime   string `json:"Datetime"`
	Forecast   string `json:"Forecast"`
	Impact     string `json:"Impact"`
	LastUpdate int64  `json:"Last Update"`
	Name       string `json:"Name"`
	Previous   string `json:"Previous"`
	Region     string `json:"Region"`
	Timestamp  int64  `json:"Timestamp"`
}
