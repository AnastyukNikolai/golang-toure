package models

import "golang-ture/ent"

type User struct {
	Id       int    `json:"id" db:"id"`
	Name     string `json:"name" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"-"`
}

type SignInput struct {
	Name     string `json:"name"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (i *User) FormatDBUserToModel(DBUser *ent.User) {
	i.Id = DBUser.ID
	i.Name = DBUser.Name
	i.Username = DBUser.Username
}
