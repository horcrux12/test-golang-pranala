package out

type WebResponse struct {
	Data    interface{} `json:"data,omitempty"`
	Message string      `json:"message"`
}