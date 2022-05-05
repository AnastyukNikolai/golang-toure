package repositories

import (
	"context"
	"golang-ture/ent"
	"golang-ture/ent/user"
	"golang-ture/internal/models"
)

type UserRepositoryEnt struct {
	DBClient *ent.Client
}

func NewUserRepositoryEnt(DBClient *ent.Client) *UserRepositoryEnt {
	return &UserRepositoryEnt{DBClient: DBClient}
}

func (r *UserRepositoryEnt) CreateUser(user models.SignInput) (int, error) {
	userDB, err := r.DBClient.User.
		Create().
		SetName(user.Name).
		SetUsername(user.Username).
		SetPassword(user.Password).
		Save(context.Background())
	if err != nil {
		return 0, err
	}
	return userDB.ID, nil
}

func (r *UserRepositoryEnt) GetUser(username, password string) (models.User, error) {
	var userModel models.User
	userDB, err := r.DBClient.User.Query().Where(user.Username(username), user.Password(password)).Only(context.Background())
	if err != nil {
		return userModel, err
	}
	userModel.FormatDBUserToModel(userDB)
	return userModel, nil
}

func (r *UserRepositoryEnt) GetById(userId int) (models.User, error) {
	var userModel models.User
	userDB, err := r.DBClient.User.Query().Where(user.ID(userId)).Only(context.Background())
	if err != nil {
		return userModel, err
	}
	userModel.FormatDBUserToModel(userDB)
	return userModel, nil
}
