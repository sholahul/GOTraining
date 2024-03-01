package model

type User struct {
	Id       int    `json:"id,omitempty"` //omitempty jika kosong tidak diberi balikan
	Username string `json:"username,omitempty"`
	Password string `json:"password,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}

type Users struct {
	Id       int    `json:"id,omitempty"` //omitempty jika kosong tidak diberi balikan
	Username string `json:"username,omitempty"`
	Email    string `json:"email,omitempty"`
	Role     string `json:"role,omitempty"`
}

type DataResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}
