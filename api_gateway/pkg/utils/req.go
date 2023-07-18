package utils

type BodySignUpuser struct {
	FirstName       string `json:"firstname" binding:"required"`
	LastName        string `json:"lastname" binding:"required"`
	Email           string `json:"email" binding:"required,email"`
	MobileNum       string `json:"mobilenum" binding:"required,min=10,max=10"`
	Password        string `json:"password" binding:"required,eqfield=ConfirmPassword"`
	ConfirmPassword string `json:"confirmpassword" binding:"required"`
	ReferalCode     string `json:"referalcode"`
}

type Otpverify struct {
	Otp         string `json:"otp" binding:"required"`
	OtpID       string `json:"otpid"`
	NewPassword string `json:"newpassword"`
	ReferalCode string `json:"referalcode"`
}

type BodyLogin struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type AddProduct struct {
	ModelName  string `json:"modelname" binding:"required"`
	Image      string `json:"image" binding:"required"`
	BrandID    uint   `json:"brandid"`
	CategoryID uint   `json:"categoryid" binding:"required"`
}

type AddProductDetail struct {
	Price             uint `json:"price" binding:"required"`
	Stock             uint `json:"stock" binding:"required"`
	AvailableSizeID   uint `json:"availablesizeid" binding:"required"`
	AvailableColourID uint `json:"availablecolourid" binding:"required"`
	ProductID         uint `json:"productid" binding:"required"`
	DiscountID        uint `json:"discountid"`
}

type Pagination struct {
	Offset uint `json:"offset"`
	Limit  uint `json:"limit"`
}
