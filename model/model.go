package model

type MyData struct {
	ID                    string `json:"ID"`
	CallingCallReference  string `json:"CallingCallReference"`
	CalledCallReference   string `json:"CalledCallReference"`
	CalledRecordType      string `json:"CalledRecordType"`
	CallingRecordType     string `json:"CallingRecordType"`
	CallingRecordTypeName string `json:"CallingRecordTypeName"`
	CalledRecordTypeName  string `json:"CalledRecordTypeName"`
	CallingNumber         string `json:"CallingNumber"`
	CalledNumber          string `json:"CalledNumber"`
	CallDuration          string `json:"CallDuration"`
	SetupTime             string `json:"SetupTime"`
	CallingIMSI           string `json:"CallingIMSI"`
}
