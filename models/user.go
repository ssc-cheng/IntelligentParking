package models

import (
	_ "github.com/go-sql-driver/mysql"
)

// 用户个人信息
type User struct {
	UserId           int64   `json:"user_id" xorm:"user_id pk not null autoincr"` // 设置主键自增
	UserUniqueId     string  `json:"user_unique_id" xorm:"user_unique_id"`
	UserPhoneImei    string  `json:"user_phone_imei" xorm:"user_phone_imei"`
	UserNickname     string  `json:"user_nickname" xorm:"user_nickname"`
	UserWeixinId     string  `json:"user_weixin_id" xorm:"user_weixin_id"`
	UserQqId         string  `json:"user_qq_id" xorm:"user_qq_id"`
	UserAnhmayId     string  `json:"user_anhmay_id" xorm:"user_anhmay_id"`
	UserAccount      string  `json:"user_account" xorm:"user_account"`
	UserPassword     string  `json:"user_password" xorm:"user_password" validate:"required"` // required 表示不为空
	UserSex          int     `json:"user_sex" xorm:"user_sex"`
	UserAvatar       string  `json:"user_avatar" xorm:"user_avatar"`
	UserPhone        string  `json:"user_phone" xorm:"user_phone" valid:"Mobile"`
	UserCreditValue  float64 `json:"user_credit_value" xorm:"user_credit_value"`
	UserAddress      string  `json:"user_address" xorm:"user_address"`
	UserIdentify     string  `json:"user_identify" xorm:"user_identify"`
	UserDriver       string  `json:"user_driver" xorm:"user_driver"`
	UserEmail        string  `json:"user_email" xorm:"user_email" valid:"Email; MaxSize(50)"` // Email 字段需要符合邮箱格式，并且最大长度不能大于 50 个字符
	UserIntegral     int     `json:"user_integral" xorm:"user_integral"`
	UserBalance      float64 `json:"user_balance" xorm:"user_balance"`
	UserLogStatus    int64   `json:"user_log_status" xorm:"user_log_status"`
	UserRegisterType int     `json:"user_register_type" xorm:"user_register_type"`
	UserRegisterTime int64   `json:"user_register_time" xorm:"user_register_time"`
	UserUpdateTime   int64   `json:"user_update_time" xorm:"user_update_time"`
}

// 用户卡包
type UserCardbag struct {
	CardbagId        int64 `json:"cardbag_id" xorm:"cardbag_id pk not null autoincr"`
	CardbagVehicleId int   `json:"cardbag_vehicle_id" xorm:"cardbag_vehicle_id"`
	CardbagUserId    int   `json:"cardbag_user_id" xorm:"cardbag_user_id"`
	CardStartTime    int64 `json:"card_start_time" xorm:"card_start_time"`
	CardEndTime      int64 `json:"card_end_time" xorm:"card_end_time"`
	CardbagCardId    int   `json:"cardbag_card_id" xorm:"cardbag_card_id"`
}

// 用户消费记录
type UserConsume struct {
	UserConsumeId            int     `json:"user_consume_id" xorm:"user_consume_id pk not null autoincr"` // 设置主键
	UserConsumeUserId        int64   `json:"user_consume_user_id" xorm:"user_consume_user_id"`
	UserConsumeOrder         string  `json:"user_consume_order" xorm:"user_consume_order"`
	UserConsumeNum           float64 `json:"user_consume_num" xorm:"user_consume_num"`
	UserConsumePayType       int     `json:"user_consume_pay_type" xorm:"user_consume_pay_type"`
	UserConsumeTime          int64   `json:"user_consume_time" xorm:"user_consume_time"`
	UserConsumeTradingStatus int     `json:"user_consume_trading_status" xorm:"user_consume_trading_status"`
	UserConsumeUseType       int     `json:"user_consume_use_type" xorm:"user_consume_use_type"`
	UserConsumeDescription   string  `json:"user_consume_description" xorm:"user_consume_description"`
	UserConsumeRemarks       string  `json:"user_consume_remarks" xorm:"user_consume_remarks"`
}

// 用户与车辆关系记录
type UserVehicle struct {
	UserVehicleId            int64 `json:"user_vehicle_id" xorm:"user_vehicle_id pk not null autoincr"` // 设置主键
	UserVehicleUserId        int64 `json:"user_vehicle_user_id" xorm:"user_vehicle_user_id"`
	UserVehicleCarId         int64 `json:"user_vehicle_car_id" xorm:"user_vehicle_car_id"`
	UserVehiclePermission    int   `json:"user_vehicle_permission" xorm:"user_vehicle_permission"`
	UserVehicleCreateTime    int64 `json:"user_vehicle_create_time" xorm:"user_vehicle_create_time"`
	UserVehicleUpdateTime    int64 `json:"user_vehicle_update_time" xorm:"user_vehicle_update_time"`
	UserVehicleBindingStatus int   `json:"user_vehicle_binding_status" xorm:"user_vehicle_binding_status"`
	AutoEduction             int   `json:"auto_eduction" xorm:"auto_eduction"`
}

// 用户积分记录
type UserIntegral struct {
	UserIntegralID     int64  `json:"user_integral_id" xorm:"user_integral_id pk not null autoincr"` // 设置主键
	UserIntegralUserId int    `json:"user_integral_user_id" xorm:"user_integral_user_id"`
	UserIntegralScene  int    `json:"user_integral_scene" xorm:"user_integral_scene"`
	UserIntegralNum    int    `json:"user_integral_num" xorm:"user_integral_num"`
	UserIntegralTime   uint32 `json:"user_integral_time" xorm:"user_integral_time"`
}

// 用户登录记录
type UserLogin struct {
	UserLoginId      int64  `json:"user_login_id" xorm:"user_login_id pk not null autoincr"` // 设置主键
	UserLoginUserId  int    `json:"user_login_user_id" xorm:"user_login_user_id"`
	UserLoginIp      string `json:"user_login_ip" xorm:"user_login_ip"`
	UserLoginTime    int64  `json:"user_login_time" xorm:"user_login_time"`
	UserLoginAddress string `json:"user_login_address" xorm:"user_login_address"`
	UserLoginApp     string `json:"user_login_app" xorm:"user_login_app"`
	UserLoginMethod  int    `json:"user_login_method" xorm:"user_login_method"`
}

// 用户消息
type UserNews struct {
	UserNewsId          int    `json:"user_news_id" xorm:"user_news_id pk not null autoincr"`
	UserNewsUserId      int    `json:"user_news_user_id" xorm:"user_news_user_id"`
	UserNewsType        int    `json:"user_news_type" xorm:"user_news_type"`
	UserNewsObject      int    `json:"user_news_object" xorm:"user_news_object"`
	UserNewsTheme       string `json:"user_news_theme" xorm:"user_news_theme"`
	UserNewsContent     string `json:"user_news_content" xorm:"user_news_content"`
	UserNewsLink        string `json:"user_news_link" xorm:"user_news_link"`
	UserNewsTime        int    `json:"user_news_time" xorm:"user_news_time"`
	UserNewsAssociation int    `json:"user_news_association" xorm:"user_news_association"`
	UserNewsHaveRead    int    `json:"user_news_have_read" xorm:"user_news_have_read"`
}

func (*User) TableName() string {
	return TableUser
}

func (*UserCardbag) TableName() string {
	return TableUserCardbag
}

func (*UserConsume) TableName() string {
	return TableUserConsume
}

func (*UserVehicle) TableName() string {
	return TableUserVehicle
}

func (*UserIntegral) TableName() string {
	return TableUserIntegral
}

func (*UserLogin) TableName() string {
	return TableUserLogin
}

func (*UserNews) TableName() string {
	return TableUserNews
}
