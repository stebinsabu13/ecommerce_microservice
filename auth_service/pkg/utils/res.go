package utils

type ResponseUsers struct {
	ID          uint   `json:"id"`
	FirstName   string `json:"firstname"`
	LastName    string `json:"lastname"`
	Email       string `json:"email"`
	MobileNum   string `json:"mobilenum"`
	Password    string `json:"password"`
	Block       bool   `json:"blocked"`
	Verified    bool   `json:"verified"`
	ReferalCode string `json:"referalcode"`
}
