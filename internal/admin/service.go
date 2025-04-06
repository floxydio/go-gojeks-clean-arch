package admin

import (
	"context"
	"github.com/google/uuid"
	"gojeksrepo/ent"
	"gojeksrepo/ent/driverprofile"
	"gojeksrepo/internal/admin/dto"
)

type Service struct {
	db *ent.Client
}

func NewServiceAdmin(dbService *ent.Client) *Service {
	return &Service{db: dbService}
}

func (r *Service) ApproveAdmin(form dto.ApprovalAdminForm, ctx context.Context) error {
	var statusUser driverprofile.Status
	var isActive bool

	if form.Status == 1 {
		statusUser = driverprofile.StatusApproved
		isActive = true
	} else if form.Status == 2 {
		statusUser = driverprofile.StatusReject
		isActive = false
	}
	userUUID, errParseUUID := uuid.Parse(form.UserId)
	if errParseUUID != nil {
		return errParseUUID
	}

	_, err := r.db.DriverProfile.UpdateOneID(userUUID).SetStatus(statusUser).SetIsActive(isActive).Save(ctx)

	if err != nil {
		return err
	}

	return nil

}
