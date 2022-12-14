// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// UserInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type UserInfoDao struct {
	gmvc.M                  // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB          // DB is the raw underlying database management object.
	Table   string          // Table is the table name of the DAO.
	Columns userInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// UserInfoColumns defines and stores column names for table user_info.
type userInfoColumns struct {
	Id           string //
	Name         string // 用户名
	Avatar       string // 头像
	Password     string //
	UserSalt     string // 加密盐 生成密码用
	Sex          string // 1男 2女
	Status       string // 1正常 2拉黑冻结
	Sign         string // 个性签名
	SecretAnswer string // 密保问题的答案
	CreatedAt    string //
	UpdatedAt    string //
	DeletedAt    string //
}

func NewUserInfoDao() *UserInfoDao {
	return &UserInfoDao{
		M:     g.DB("default").Model("user_info").Safe(),
		DB:    g.DB("default"),
		Table: "user_info",
		Columns: userInfoColumns{
			Id:           "id",
			Name:         "name",
			Avatar:       "avatar",
			Password:     "password",
			UserSalt:     "user_salt",
			Sex:          "sex",
			Status:       "status",
			Sign:         "sign",
			SecretAnswer: "secret_answer",
			CreatedAt:    "created_at",
			UpdatedAt:    "updated_at",
			DeletedAt:    "deleted_at",
		},
	}
}
