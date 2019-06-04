package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type UserModel struct {
	Index    int
	Id       string `orm:"pk"`
	Name     string
	Age      int
	Sex      bool
	Email    string
	Password string
}

func init() {
	orm.RegisterDriver("sqlite", orm.DRSqlite)
	err := orm.RegisterDataBase("default", "sqlite3", "database/mysqlite.db")
	if err != nil {
		log.Println(err)
	}
	orm.RegisterModel(new(UserModel))
	orm.RunSyncdb("default", false, true)

}

func ModelToUser(ue *UserModel) *User {
	var u User
	u.Name = ue.Name
	u.Index = ue.Index
	u.Id = ue.Id
	u.Age = ue.Age
	u.Email = ue.Email
	u.Password = ue.Password
	u.Sex = ue.Sex
	return &u
}

func UserTostream(ue *User) *UserModel {
	var u UserModel
	u.Name = ue.Name
	u.Index = ue.Index
	u.Id = ue.Id
	u.Age = ue.Age
	u.Email = ue.Email
	u.Password = ue.Password
	u.Sex = ue.Sex
	return &u
}
