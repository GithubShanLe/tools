package models

import (
	"errors"
	"github.com/astaxie/beego/orm"
	"network/common"
	"network/logself"
)

type User struct {
	Index    int    `json:"index"`
	Id       string `json:"id"`
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Sex      bool   `json:"sex"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func AddUser(u *User, o orm.Ormer) (orm.Ormer, error) {
	if o != nil {
		user := UserTostream(u)
		_, err := o.Insert(user)
		if err != nil {
			logself.LogCenter.WriteLog(common.LogPath+"err.log", err.Error())
			o.Rollback()
			return o, err
		}
		return o, nil
	} else {
		return o, errors.New("orm.Ormer is nil")
	}
}

func UpdateUser(u *User, o orm.Ormer) (orm.Ormer, error) {
	if o != nil {
		user := UserTostream(u)
		_, err := o.Update(user)
		if err != nil {
			logself.LogCenter.WriteLog(common.LogPath+"err.log", err.Error())
			o.Rollback()
			return o, err
		}
		return o, nil
	} else {
		return o, errors.New("orm.Ormer is nil")
	}
}
func GetUser(u *UserModel) *User {
	o := orm.NewOrm()
	o.Using("default")
	qs := o.QueryTable("user_model")
	qs.All(u)
	hh := ModelToUser(u)
	return hh

}
