// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// CategoryInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type CategoryInfoDao struct {
	gmvc.M                      // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB              // DB is the raw underlying database management object.
	Table   string              // Table is the table name of the DAO.
	Columns categoryInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// CategoryInfoColumns defines and stores column names for table category_info.
type categoryInfoColumns struct {
	Id        string //
	ParentId  string // 父级id
	Name      string //
	PicUrl    string // icon
	DeletedAt string //
	CreatedAt string //
	UpdatedAt string //
	Level     string // 等级 默认1级分类
	Sort      string //
}

func NewCategoryInfoDao() *CategoryInfoDao {
	return &CategoryInfoDao{
		M:     g.DB("default").Model("category_info").Safe(),
		DB:    g.DB("default"),
		Table: "category_info",
		Columns: categoryInfoColumns{
			Id:        "id",
			ParentId:  "parent_id",
			Name:      "name",
			PicUrl:    "pic_url",
			DeletedAt: "deleted_at",
			CreatedAt: "created_at",
			UpdatedAt: "updated_at",
			Level:     "level",
			Sort:      "sort",
		},
	}
}
