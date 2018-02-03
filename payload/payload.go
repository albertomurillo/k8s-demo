package payload

// Message contains the payload transferred from backend app to frontend app
type Message struct {
	Message string `json:"message"`
}
