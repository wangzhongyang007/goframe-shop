// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"shop/app/dao/internal"
)

// adminInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type adminInfoDao struct {
	*internal.AdminInfoDao
}

var (
	// AdminInfo is globally public accessible object for table admin_info operations.
	AdminInfo adminInfoDao
)

func init() {
	AdminInfo = adminInfoDao{
		internal.NewAdminInfoDao(),
	}
}

// Fill with you ideas below.
