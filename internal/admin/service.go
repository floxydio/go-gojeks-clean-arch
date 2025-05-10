package admin

import (
	"context"
	"gojeksrepo/ent"
	"gojeksrepo/ent/driverprofile"
	"gojeksrepo/ent/usersadmin"
	"gojeksrepo/internal/admin/dto"
	"gojeksrepo/pkg"

	"github.com/google/uuid"
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

func (r *Service) FindExistUser(form dto.SignInAdmin, ctx context.Context) (*ent.UsersAdmin, error) {
	data, err := r.db.UsersAdmin.Query().Where(usersadmin.Username(form.Username)).First(ctx)

	if err != nil {
		return nil, err
	}

	return data, nil

}

func (r *Service) SignupAdmin(form dto.SignUpAdmin, ctx context.Context) error {
	hash, errHash := pkg.HashPassword(form.Password)

	if errHash != nil {
		return errHash
	}

	_, err := r.db.UsersAdmin.Create().SetName(form.Name).SetUsername(form.Username).SetPassword(hash).Save(ctx)

	if err != nil {
		return err
	}

	return nil
}
