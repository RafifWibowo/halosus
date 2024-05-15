package dto

type CreateITRequest struct {
	Nip      int64  `json:"nip"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

type CreateNurseRequest struct {
	Nip                 int64  `json:"nip"`
	Name                string `json:"name"`
	IdentityCardScanImg string `json:"identityCardScanImg"`
}

type UpdateNurseRequest struct {
	Nip  int64  `json:"nip"`
	Name string `json:"name"`
}

type GrantAccessRequest struct {
	Password string `json:"password"`
}

type LoginITRequest struct {
	Nip      int64  `json:"nip"`
	Password string `json:"password"`
}