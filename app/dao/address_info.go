// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"shop/app/dao/internal"
)

// addressInfoDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type addressInfoDao struct {
	*internal.AddressInfoDao
}

var (
	// AddressInfo is globally public accessible object for table address_info operations.
	AddressInfo addressInfoDao
)

func init() {
	AddressInfo = addressInfoDao{
		internal.NewAddressInfoDao(),
	}
}

// Fill with you ideas below.
