package auth

import (
	"context"
	"fmt"
	"gojeksrepo/ent"
	"gojeksrepo/ent/user"
	"gojeksrepo/internal/auth/dto"
	"gojeksrepo/pkg"
)

type DbService struct {
	dbClient *ent.Client
}

func NewAuthService(db *ent.Client) *DbService {
	return &DbService{dbClient: db}
}

func (db *DbService) SignUpUser(form dto.SignUp, ctx context.Context) error {
	hash, errHash := pkg.HashPassword(form.Password)

	if errHash != nil {
		return errHash
	}

	_, errGenerate := db.dbClient.User.Create().SetName(form.Name).SetPassword(hash).SetPhone(form.Phone).SetEmail(form.Email).SetRole(user.RoleUser).Save(ctx)

	if errGenerate != nil {
		return errGenerate
	}

	return nil
}

func (db *DbService) SignUpDriver(form dto.SignUpDriver, ctx context.Context) error {
	hash, errHash := pkg.HashPassword(form.Password)
	tx, errTx := db.dbClient.Tx(ctx)
	if errTx != nil {
		return fmt.Errorf("starting a transaction: %w", errTx)
	}
	if errHash != nil {
		return errHash
	}
	data, errGenerate := tx.User.Create().SetName(form.Name).SetPassword(hash).SetEmail(form.Email).SetRole(user.RoleDriver).SetPhone(form.Phone).Save(ctx)
	if errGenerate != nil {
		return rollback(tx, errGenerate)
	}
	_, errDataDriverInsert := tx.DriverProfile.Create().SetLicenseNumber(form.Sim).SetKtpNumber(form.Ktp).SetVehicleInfo(form.VehicleType).SetUserID(data.ID).Save(ctx)

	if errDataDriverInsert != nil {
		return rollback(tx, errDataDriverInsert)
	}
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}

func (db *DbService) SignInUser(form dto.SignInForm, ctx context.Context) (dto.SignInResponse, error) {
	findExistingUser, err := db.dbClient.User.Query().Where(user.Email(form.Email)).Only(ctx)

	if ent.IsNotFound(err) {
		return dto.SignInResponse{
			Status:  400,
			Error:   true,
			Message: "This account is not registered",
		}, nil
	}

	if err != nil {
		return dto.SignInResponse{}, err
	}

	if !pkg.CheckPasswordHash(form.Password, findExistingUser.Password) {
		return dto.SignInResponse{
			Status:  400,
			Error:   true,
			Message: "Invalid email or password",
		}, nil
	}
	token, errToken := pkg.CreateToken(findExistingUser.Name, findExistingUser.ID.String())

	if errToken != nil {
		return dto.SignInResponse{}, errToken
	}
	return dto.SignInResponse{
		Status:  200,
		Error:   false,
		Token:   token,
		Message: "Successfully Sign In",
	}, nil
}

func (db *DbService) SignInUserDriver(form dto.SignInFormDriver, ctx context.Context) (dto.SignInResponse, error) {
	findExistingUser, err := db.dbClient.User.Query().Where(user.Phone(form.Phone)).Only(ctx)

	if ent.IsNotFound(err) {
		return dto.SignInResponse{
			Status:  400,
			Error:   true,
			Message: "This account is not registered",
		}, nil
	}

	if err != nil {
		return dto.SignInResponse{}, err
	}

	if !pkg.CheckPasswordHash(form.Password, findExistingUser.Password) {
		return dto.SignInResponse{
			Status:  400,
			Error:   true,
			Message: "Invalid phone or password",
		}, nil
	}

	if findExistingUser.IsVerified == false {
		return dto.SignInResponse{
			Status:  400,
			Error:   true,
			Message: "Your account is inactive - Request to admin",
		}, nil
	}

	token, errToken := pkg.CreateToken(findExistingUser.Name, findExistingUser.ID.String())

	if errToken != nil {
		return dto.SignInResponse{}, errToken
	}
	return dto.SignInResponse{
		Status:  200,
		Error:   false,
		Token:   token,
		Message: "Successfully Sign In",
	}, nil
}

func rollback(tx *ent.Tx, err error) error {
	if errRollback := tx.Rollback(); errRollback != nil {
		err = fmt.Errorf("%w: %v", err, errRollback)
	}
	return err
}
