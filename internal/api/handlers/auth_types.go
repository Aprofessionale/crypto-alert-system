package handlers

// for POST /auth/subcsribe
type SubscribeRequest struct {
	Email string `json:"email" binding:"required,email"`
}

type GeneralResponse struct {
	Message string `json:"message"`
}

type ErrorResponse struct {
	Error string `json:"error"`
}
