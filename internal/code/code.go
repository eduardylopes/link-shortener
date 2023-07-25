package code

type Code struct {
	Code string `json:"code"`
}

type Service interface {
	GenerateUniqueCode() string
}
