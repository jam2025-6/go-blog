package service

import (
	"errors"
	"server/global"
	"server/model/appTypes"
	"server/model/database"
	"server/utils"

	"github.com/gofrs/uuid"
	"gorm.io/gorm"
)

type UserService struct {
}

func (userService *UserService) Register(user database.User) (database.User, error) {

	// errors.Is(错误A, 错误B) = 判断"错误A"是不是"错误B"这种类型
	// .First(&database.User{}) = 取符合条件的第一条记录，放到一个空的 User 结构体里
	if errors.Is(global.DB.Where("email = ?", user.Email).First(&database.User{}).Error, gorm.ErrRecordNotFound) {
		return database.User{}, errors.New("this email address is already registered, please check the information you filled in, or retrieve your password")
	}

	user.Password = utils.BcryptHash(user.Password)
	// 就是生成一个全球唯一的随机 ID（UUID），并且如果生成失败就直接 panic（程序崩溃），保证一定能拿到一个有效的 UUID。
	user.UUID = uuid.Must(uuid.NewV4())
	user.Avatar = "/image/avatar.jpg"
	user.RoleID = appTypes.User
	user.Register = appTypes.Email
	if err := global.DB.Create(&user).Error; err != nil {
		return database.User{}, nil
	}
	return user, nil
}
