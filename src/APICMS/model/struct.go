package model

type User struct {
	Id       int    `json:"id,omitempty"` //omitempty jika kosong tidak diberi balikan
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}

type DataResponse struct {
	ResponCode string `json:"responseCode"`
	Message    string `json:"message"`
	Data       any    `json:"data,omitempty"`
}
