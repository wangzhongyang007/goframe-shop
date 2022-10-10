package response

//11XXXX 云仓服务
//|11|00|云仓服务 - 店铺模块错误|
//|11|01|云仓服务 - 商品模块错误|
//|11|02|云仓服务 - 订单模块错误|
//|11|03|云仓服务 - 售后模块错误|
//|11|04|云仓服务 - 其他模块错误|


//云仓: 售后模块错误
const (
	// ErrRefundCreate : 申请售后api error
	ErrRefundCreate int = iota + 110301

	// ErrRefundCheck : 获取售后类型 售后原因 最大金额失败
	ErrRefundCheck

	// ErrRefundBefore : 售后前置
	ErrRefundBefore

	// ErrRefundInfo : 售后详情
	ErrRefundInfo

	// ErrRefundCancel : 取消售后
	ErrRefundCancel

	// ErrRefundBackSend : 回填物流
	ErrRefundBackSend
)

//云仓: 商品模块错误
const (
	// ErrRefundCreate : 分组允许最大50个
	ErrGroupMax int = iota + 110101

	// ErrGroupNotFound : 分组不存在
	ErrGroupNotFound

	// ErrGroupExist : 分组已存在
	ErrGroupExist

	// ErrWarehouseNotFound : 仓库不存在
	ErrWarehouseNotFound

	// ErrWarehouseExist : 仓库已存在
	ErrWarehouseExist

	// ErrDefaultSendWarehouseExist : 默认发货仓已存在
	ErrDefaultSendWarehouseExist

	// ErrDefaultGetWarehouseExist : 默认收货仓已存在
	ErrDefaultGetWarehouseExist
)

//云仓: 其他模块错误
const (
	// ErrAddressNotFound : 地址不存在
	ErrAddressNotFound int = iota + 110401
)


//售后模块
func init()  {
	MustRegister(ErrRefundCreate, "%%0")
	MustRegister(ErrRefundCheck, "%%0")
	MustRegister(ErrRefundBefore, "%%0")
	MustRegister(ErrRefundCancel, "%%0")
	MustRegister(ErrRefundBackSend, "%%0")
}

//商品模块
func init()  {
	MustRegister(ErrGroupMax, "最多允许添加50个分组,请编辑分组或删除分组后再添加")
	MustRegister(ErrGroupNotFound, "分组不存在")
	MustRegister(ErrWarehouseNotFound, "仓库不存在")
	MustRegister(ErrGroupExist, "分组已存在")
	MustRegister(ErrWarehouseExist, "仓库已存在")
	MustRegister(ErrDefaultSendWarehouseExist, "默认发货仓已存在")
	MustRegister(ErrDefaultGetWarehouseExist, "默认收货仓已存在")
}

//其他模块
func init()  {
	MustRegister(ErrAddressNotFound, "地址不存在")
}



func MustRegister(code int, message string) {
	CodeService[code] = message
}