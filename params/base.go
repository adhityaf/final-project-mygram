package params

type Response struct {
	Status         uint        `json:"status"`
	Message        string      `json:"message,omitempty"`
	Error          string      `json:"error,omitempty"`
	AdditionalInfo interface{} `json:"additional_info,omitempty"`
	Payload        interface{} `json:"payload,omitempty"`
}
