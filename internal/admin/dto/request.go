package dto

type ApprovalAdminForm struct {
	UserId string `json:"user_id" form:"user_id"`
	Status uint   `json:"status" form:"status"`
}

type SignUpAdmin struct {
	Name     string `json:"name" form:"name"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type SignInAdmin struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}
