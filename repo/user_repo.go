package repo

import (
	"intelligent_parking/db"
	"intelligent_parking/models"
	"intelligent_parking/utils/errors"
)

type UserRepo struct {}

func NewUserRepo() *UserRepo {
	return new(UserRepo)
}

// 新增用户信息
func (u *UserRepo) AddUserData(data models.User) (count int64,err error) {

	if count,err = db.GetLocalDB().Table(models.TableUser).Insert(&data);err != nil{
		return
	}else if count == 0 {
		err = errors.Wrap(err,"插入了0条数据")
		return
	}
	return
}

// 查询用户信息
func (u *UserRepo) SearchUserData(input_params string,params interface{}) (user models.User,ok bool,err error) {
	if ok,err = db.GetLocalDB().Table(models.TableUser).Where(input_params+"=?",params).Get(&user);err != nil{
		return
	}
	return
}


// 判断记录是否存在
func (u *UserRepo) IsValue(userid int64)(ok bool ,err error) {
	if ok,err = db.GetLocalDB().Table(models.TableUser).Exist(&models.User{
		UserId: userid,
	});err != nil{
		return
	}
	return
}

// 更新用户表信息
func (u *UserRepo) UpdateUserData(userid int64,params models.User) (err error){
	var count int64
	if count,err = db.GetLocalDB().Table(models.TableUser).Where("user_id=?",userid).Update(&params);err != nil{
		return
	}else if count == 0 {
		err = errors.New("没有要更新的数据信息")
	}
	return
}


