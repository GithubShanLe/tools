package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"network/models"
)

type UserController struct {
	beego.Controller
}

func (s *UserController) Get() {
	var user models.UserModel
	us := models.GetUser(&user)
	s.Data["json"] = &us
	s.ServeJSON()
}

func (s *UserController) Post() {
	var user models.User
	err := json.Unmarshal(s.Ctx.Input.RequestBody, &user)

	if err != nil {
		s.Data["json"] = map[string]string{"Error": err.Error()}
		s.ServeJSON()
		return
	}
	o := orm.NewOrm()
	o.Using("default")
	o.Begin()
	_, err = models.UpdateUser(&user, o)
	if err != nil {
		o.Commit()
		s.Data["json"] = "Lee"
		s.ServeJSON()
	} else {
		o.Commit()
		s.Data["json"] = "OK"
		s.ServeJSON()
	}

}
