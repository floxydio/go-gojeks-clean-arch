package dto

type SignUp struct {
	Name     string `json:"name" form:"name"`
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

type SignUpDriver struct {
	Name        string `json:"name" form:"name"`
	Email       string `json:"email" form:"email"`
	Phone       string `json:"phone" form:"phone"`
	Password    string `json:"password" form:"password"`
	Sim         string `json:"sim" form:"sim"`
	Ktp         string `json:"ktp" form:"ktp"`
	VehicleType string `json:"vehicle_type" form:"vehicle_type"`
}

type SignInForm struct {
	Email    string `json:"email" form:"email"`
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}

type SignInFormDriver struct {
	Phone    string `json:"phone" form:"phone"`
	Password string `json:"password" form:"password"`
}
