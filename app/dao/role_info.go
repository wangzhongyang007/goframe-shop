// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"shop/app/dao/internal"
)

// roleInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type roleInfoDao struct {
	*internal.RoleInfoDao
}

var (
	// RoleInfo is globally public accessible object for table role_info operations.
	RoleInfo roleInfoDao
)

func init() {
	RoleInfo = roleInfoDao{
		internal.NewRoleInfoDao(),
	}
}

// Fill with you ideas below.
