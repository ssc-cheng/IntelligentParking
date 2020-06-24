package routers

import (
	"github.com/astaxie/beego"
	"intelligent_parking/controllers"
)

func init() {
	// 定义路由组
	user := beego.NewNamespace("/user",
		// 用户注册接口
		beego.NSRouter("/register", &controllers.UserController{}, "post:UserRegister"),
		// 用户登录接口
		beego.NSRouter("/login", &controllers.UserController{}, "post:UserLogin"),
		// 修改用户密码
		beego.NSRouter("/update",&controllers.UserController{},"post:UserUpdate"),
		// 绑定用户车辆数据
		beego.NSRouter("/bindcar",&controllers.UserController{},"post:UserBindCar"),
	)

	car := beego.NewNamespace("/vehicle",
		// 车辆消费记录
		beego.NSRouter("/expense",&controllers.VehicleController{},"post:VehicleExpense"),
		)

	//注册路由组
	beego.AddNamespace(user)
	beego.AddNamespace(car)

}
