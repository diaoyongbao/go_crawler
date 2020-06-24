package model
type  Proxy []struct {
	Proxy string `json:"proxy"`
	FailCount int `json:"fail_count"`
	Region  string `json:"region"`
	Type  string `json:"type "`
	Source  string `json:"source"`
	CheckCount int `json:"check_count"`
	LastStatus int `json:"last_status"`
	LastTime string `json:"last_time"`
}
