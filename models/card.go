package models

// 卡包
type Card struct {
	CardId int64 `json:"card_id" xorm:"card_id pk not null autoincr"`
	CardName int `json:"card_name" xorm:"card_name"`
	CardDay int `json:"card_day" xorm:"card_day"`
	CardParkingId int `json:"card_parking_id" xorm:"card_parking_id"`
}
