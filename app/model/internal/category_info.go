// ==========================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/os/gtime"
)

// CategoryInfo is the golang structure for table category_info.
type CategoryInfo struct {
	Id        int         `orm:"id,primary" json:"id"`        //
	ParentId  int         `orm:"parent_id"  json:"parentId"`  // 父级id
	Name      string      `orm:"name"       json:"name"`      //
	PicUrl    string      `orm:"pic_url"    json:"picUrl"`    // icon
	DeletedAt *gtime.Time `orm:"deleted_at" json:"deletedAt"` //
	CreatedAt *gtime.Time `orm:"created_at" json:"createdAt"` //
	UpdatedAt *gtime.Time `orm:"updated_at" json:"updatedAt"` //
	Level     int         `orm:"level"      json:"level"`     // 等级 默认1级分类
	Sort      int         `orm:"sort"       json:"sort"`      //
}
