package service

type ServiceGroup struct {
	BaseService  BaseService
	JwtService   JwtService
	GaodeService GaodeService
	UserService  UserService
}

var ServiceGroupApp = new(ServiceGroup)
