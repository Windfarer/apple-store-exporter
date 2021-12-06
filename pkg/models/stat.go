package models

type StatResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	TTL     int         `json:"ttl"`
	Data    Stat        `json:"data"`
}

type Stat struct {
	Mid       int `json:"mid"`
	Following int `json:"following"`
	Whisper   int `json:"whisper"`
	Black     int `json:"black"`
	Follower  int `json:"follower"`
}
