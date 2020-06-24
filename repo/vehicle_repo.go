package repo

import (
	"intelligent_parking/db"
	"intelligent_parking/models"
	"intelligent_parking/utils/errors"
)

type VehicleRepo struct {}

func NewVehicleRepo() *VehicleRepo {
	return new(VehicleRepo)
}

// 查询车辆信息
func (v *VehicleRepo) SearchVehicleData(input_params string,params interface{}) (vehicle models.Vehicle,ok bool,err error) {

	if ok,err = db.GetLocalDB().Table(models.TableVehicle).Where(input_params+"=?",params).Get(&vehicle);err != nil{
		return
	}
	return

}

// 新增车辆信息
func (v *VehicleRepo) InsertVehicleData(vehicle models.Vehicle)(count int64,err error) {

	if count,err = db.GetLocalDB().Table(models.TableVehicle).Insert(&vehicle);err != nil{
		errors.Wrap(err,"插入车辆信息失败")
		return
	}else if count == 0 {
		err = errors.Wrap(err,"插入了0条数据")
		return
	}
	return

}


