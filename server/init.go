package server

import "intelligent_parking/repo"

var (
	GUserRepo *repo.UserRepo
	GUserVehicleRepo *repo.UserVehicleRepo
	GVehicleRepo *repo.VehicleRepo
	GjwtToken *jwtCustomClaims
)

//初始化repo实例
func NewRepo() {
	GUserRepo = repo.NewUserRepo()
	GUserVehicleRepo = repo.NewUserVehicleRepo()
	GVehicleRepo = repo.NewVehicleRepo()
	GjwtToken = NewJwtToken()
}
