package controllers

import (
	"github.com/astaxie/beego"
)

//MainController structure
type MainController struct {
	beego.Controller
}

//Get index from MainController
func (c *MainController) Get() {
	c.Data["Website"] = "NapBank.me"
	c.Data["Email"] = "karthik.suresh181@gmail.com"
	c.TplName = "index.html"
}

/*
//Post index from MainController
func (c *MainController) Post() {
	var signup models.SignUp
	var signupmodel = models.SignUp{}
	p := c.GetString("confpass")
	flash := beego.NewFlash()

	if p == "" {
		var s models.SignUp
		var a = models.SignUp{}
		s.Email = c.GetString("email")
		s.Pass = c.GetString("pass")
		a = a.Read(s)
	} else {
		if c.Ctx.Input.IsPost() {
			//c.Input().Add("id", "1")
			if err := c.ParseForm(&signupmodel); err != nil {
				c.Abort("500")
			}
			valid := validation.Validation{}
			b, err := valid.Valid(&signupmodel)
			fmt.Print(b)
			if err != nil {
				fmt.Print(err) // handle error
			}
			if !b {
				for _, err := range valid.Errors {
					flash.Error(err.Key, err.Message)
					//log.Println(err.Key, err.Message)
				}
			}
			signupmodel.InsertOrUpdate(signup)
		}
	}

	c.Data["SignUp"] = signup
	c.TplName = "index.html"
}*/
