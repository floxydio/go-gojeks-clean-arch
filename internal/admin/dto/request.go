package dto

type ApprovalAdminForm struct {
	UserId string `json:"user_id" form:"user_id"`
	Status uint   `json:"status" form:"status"`
}
