package models

type User struct {
	Id         uint64 `form:"id"           db:"id"         json:"id,omitempty"`
	Name       string `form:"name"         db:"name"       json:"name,omitempty"`
	Surname    string `form:"surname"      db:"surname"    json:"surname,omitempty"`
	Patronymic string `form:"patronymic"   db:"patronymic" json:"patronymic,omitempty"`
	Username   string `form:"username"     db:"username"   json:"username,omitempty"`
	Password   string `form:"password"     db:"password"   json:"password,omitempty"`
	Email      string `form:"email"        db:"email"      json:"email,omitempty"`
}
