package model

type SignupData struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	RetryPassword string `json:"retry_password"`
}
