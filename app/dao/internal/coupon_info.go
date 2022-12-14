// ==========================================================================
// This is auto-generated by gf cli tool. DO NOT EDIT THIS FILE MANUALLY.
// ==========================================================================

package internal

import (
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/frame/gmvc"
)

// CouponInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management.
type CouponInfoDao struct {
	gmvc.M                    // M is the core and embedded struct that inherits all chaining operations from gdb.Model.
	DB      gdb.DB            // DB is the raw underlying database management object.
	Table   string            // Table is the table name of the DAO.
	Columns couponInfoColumns // Columns contains all the columns of Table that for convenient usage.
}

// CouponInfoColumns defines and stores column names for table coupon_info.
type couponInfoColumns struct {
	Id         string //
	Name       string //
	Price      string // 优惠前面值 单位分
	GoodsIds   string // 关联使用的goods_ids  逗号分隔
	CategoryId string // 关联使用的分类id
	CreatedAt  string //
	UpdatedAt  string //
	DeletedAt  string //
}

func NewCouponInfoDao() *CouponInfoDao {
	return &CouponInfoDao{
		M:     g.DB("default").Model("coupon_info").Safe(),
		DB:    g.DB("default"),
		Table: "coupon_info",
		Columns: couponInfoColumns{
			Id:         "id",
			Name:       "name",
			Price:      "price",
			GoodsIds:   "goods_ids",
			CategoryId: "category_id",
			CreatedAt:  "created_at",
			UpdatedAt:  "updated_at",
			DeletedAt:  "deleted_at",
		},
	}
}
