package models

type Vehicle struct {
	VehicleId                   int64 `json:"vehicle_id" xorm:"vehicle_id pk not null autoincr"` // 设置主键自增
	VehicleCard                 string `json:"vehicle_card" xorm:"vehicle_card"`
	VehicleDriver               string `json:"vehicle_driver" xorm:"vehicle_driver"`
	VehiclePicture              string `json:"vehicle_picture" xorm:"vehicle_picture"`
	VehicleModel                string `json:"vehicle_model" xorm:"vehicle_model"`
	VehicleColor                string `json:"vehicle_color" xorm:"vehicle_color"`
	VehicleType                 int64 `json:"vehicle_type" xorm:"vehicle_type"`
	VehicleDriverPicture        string `json:"vehicle_driver_picture" xorm:"vehicle_driver_picture"`
	VehicleParkStatus           int64 `json:"vehicle_park_status" xorm:"vehicle_park_status"`
	VehicleStatus               int64 `json:"vehicle_status" xorm:"vehicle_status"`
	VehicleConsumeid            int `json:"vehicle_consumeid" xorm:"vehicle_consumeid"`
	VehicleRegisterTime         int64 `json:"vehicle_register_time" xorm:"vehicle_register_time"`
	VehicleUpdateTime           int64 `json:"vehicle_update_time" xorm:"vehicle_update_time"`
	VehicleOwner                string `json:"vehicle_owner" xorm:"vehicle_owner"`
	VehicleIdentificationNumber string `json:"vehicle_identification_number" xorm:"vehicle_identification_number"`
	EngineNumber                string `json:"engine_number" xorm:"engine_number"`
	DrivingLicenseRegistDate    int `json:"driving_license_regist_date" xorm:"driving_license_regist_date"`
	DrivingLicenseIssueDate     int `json:"driving_license_issue_date" xorm:"driving_license_issue_date"`
	DrivingLicensePhoto         string `json:"driving_license_photo" xorm:"driving_license_photo"`
	DrivingApprovalStatus       int `json:"driving_approval_status" xorm:"driving_approval_status"`
}


func (*Vehicle) TableName() string {
	return TableVehicle
}