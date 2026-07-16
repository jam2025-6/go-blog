package appTypes

type RoleID uint

const (
	Guest RoleID = iota // 游客
	User                // 普通用户
	Admin               // 管理员
)
