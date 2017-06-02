package filters

import (
    "github.com/astaxie/beego/context"
    "pybbs-go/models"
    "github.com/astaxie/beego"
    "strconv"
)

func IsLogin(ctx *context.Context) (bool, models.User) {
	var user models.User

	username := ctx.Input.CruSession.Get("username")
	if username == nil {
		return false, user
	}

	exist, user := models.FindUserByUserName(username.(string))
	return exist, user
}

var HasPermission = func(ctx *context.Context) {
    ok, user := IsLogin(ctx)
    if !ok {
        ctx.Redirect(302, "/login")
    } else {
        url := ctx.Request.RequestURI
        beego.Debug("url: ", url)
        flag := models.Enforcer.Enforce(strconv.Itoa(user.Id), url)
        if !flag {
            ctx.WriteString("你没有权限访问这个页面")
        }
    }
}

var FilterUser = func(ctx *context.Context) {
    ok, _ := IsLogin(ctx)
    if !ok {
        ctx.Redirect(302, "/login")
    }
}
