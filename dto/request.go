package dto

type CreateITRequest struct {
	Nip      int64  `json:"nip"`
	Name     string `json:"name"`
	Password string `json:"password"`
}