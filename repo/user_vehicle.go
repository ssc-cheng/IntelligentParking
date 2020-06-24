package repo

import (
	"intelligent_parking/db"
	"intelligent_parking/models"
	"intelligent_parking/utils/errors"
)

type UserVehicleRepo struct {}

func NewUserVehicleRepo() *UserVehicleRepo {
	return new(UserVehicleRepo)
}

// 查询用户与车辆关系表
func (u *UserVehicleRepo) SearchUserVehicleData(input_params string,params interface{})(user_vehicle []*models.UserVehicle,err error) {
	if err = db.GetLocalDB().Table(models.TableVehicle).Where(input_params+"=?",params).Find(&user_vehicle);err != nil{
		err = errors.Wrap(err,"没有对应用户与车辆记录")
		return
	}
	return
}

// 插入用户与车辆关系表数据
func (u *UserVehicleRepo) InsertUserVehicleData(data models.UserVehicle)(count int64,err error) {

	if count,err = db.GetLocalDB().Table(models.TableUserVehicle).Insert(&data);err != nil{
		err = errors.Wrap(err,"插入用户与车辆关系表失败")
		return
	}else if count == 0 {
		err = errors.New("插入了0条数据")
		return
	}
	return
}
