package controllers

import (
	"github.com/astaxie/beego"
	"intelligent_parking/models"
	"intelligent_parking/repo"
	"intelligent_parking/server"
	"intelligent_parking/utils"
	"intelligent_parking/utils/errors"
	"strconv"
)

type UserController struct {
	beego.Controller
}

func (u *UserController) UserRegister() {

	var (
		params models.User
		err    error
		count  int64
		ok     bool
	)

	defer func() {
		resp := utils.WriteJsonMsgWithError(count, err)
		u.Data["json"] = resp
		u.ServeJSON()
	}()

	if err = utils.ReadValidJSON(u.Ctx.Input.RequestBody, &params); err != nil {
		return
	}

	if _, ok, err = server.GUserRepo.SearchUserData("user_account", params.UserAccount); err != nil {
		return
	} else if ok {
		err = errors.New("该用户名已被注册")
		return
	}

	if count, err = server.GUserRepo.AddUserData(params); err != nil {
		return
	}

}

func (u *UserController) UserLogin() {
	var (
		params   models.User
		err      error
		resp_con struct {
			user_id int64
			Token   string
			Account string
			Phone   string
		}
		ok      bool
		hm_user models.User
	)

	defer func() {
		resp := utils.WriteJsonMsgWithError(resp_con, err)
		u.Data["json"] = resp
		u.ServeJSON()
	}()

	if err = utils.ReadValidJSON(u.Ctx.Input.RequestBody, &params); err != nil {
		return
	}

	if hm_user, ok, err = server.GUserRepo.SearchUserData("user_account", params.UserAccount); err != nil {
		return
	} else if !ok {
		err = errors.New("没有该用户数据")
		return
	}

	if hm_user.UserPassword != params.UserPassword {
		err = errors.New("用户密码不匹配")
		return
	}

	if resp_con.Token, err = server.GjwtToken.CreateToken([]byte(beego.AppConfig.String("secretKey")), strconv.FormatInt(hm_user.UserId, 10), hm_user.UserAccount); err != nil {
		err = errors.Wrap(err, "生成token失败")
		return
	}
	resp_con.Account = hm_user.UserAccount
	resp_con.Phone = hm_user.UserPhone
	resp_con.user_id = hm_user.UserId
}

// 前端需要传用户id
func (u *UserController) UserUpdate() {

	var (
		params models.User
		err    error
		con    string
	)

	defer func() {
		resp := utils.WriteJsonMsgWithError(con, err)
		u.Data["json"] = resp
		u.ServeJSON()
	}()

	if err = server.GjwtToken.Gtoken(u.Ctx.Input.Header("Authorization")); err != nil {
		return
	}

	if err = utils.ReadValidJSON(u.Ctx.Input.RequestBody, &params); err != nil {
		return
	}

	if err = repo.NewUserRepo().UpdateUserData(params.UserId, params); err != nil {
		return
	}
	con = "更新成功"

}

func (u *UserController) UserBindCar() {

	var (
		params struct {
			User_account string
			Carcard      string
		}
		err         error
		ok          bool
		user        models.User
		vehicle     models.Vehicle
		userVehicle models.UserVehicle
		respCon     int64
	)

	defer func() {
		resp := utils.WriteJsonMsgWithError(respCon, err)
		u.Data["json"] = resp
		u.ServeJSON()
	}()

	if err = server.GjwtToken.Gtoken(u.Ctx.Input.Header("Authorization")); err != nil {
		return
	}

	if err = utils.ReadValidJSON(u.Ctx.Input.RequestBody, &params); err != nil {
		return
	}
	// 查询用户
	if user, ok, err = repo.NewUserRepo().SearchUserData("user_account", params.User_account); err != nil {
		return
	} else if !ok {
		err = errors.New("没有该用户数据")
		return
	}

	// 查询车辆
	LOOP:
	if vehicle, ok, err = repo.NewVehicleRepo().SearchVehicleData("vehicle_card", params.Carcard); err != nil {
		return
	} else if !ok { // 插入车辆信息
		vehicle.VehicleCard = params.Carcard
		if _, err = repo.NewVehicleRepo().InsertVehicleData(vehicle); err != nil {
			return
		}
		goto LOOP
	}

	// 插入用户与车辆关系记录
	userVehicle.UserVehicleUserId = user.UserId
	userVehicle.UserVehicleCarId = vehicle.VehicleId
	if respCon, err = repo.NewUserVehicleRepo().InsertUserVehicleData(userVehicle); err != nil {
		return
	}

}
