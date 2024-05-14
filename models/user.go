package models

type User struct {
	Id                   string `json:"Id"`
	Nip                  int64  `json:"nip"`
	Name                 string `json:"name"`
	Password             string `json:"password"`
	IdentityCardScanning string `json:"identityCardScanning"`
	CreatedAt            string `json:"createdAt"`
	UpdatedAt            string `json:"updatedAt"`
	DeletedAt            string `json:"deletedAt"`
}