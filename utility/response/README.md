
## 使用方法

```golang  
package main

import (
	"github.com/gogf/gf/v2/net/ghttp"
    "stbz-private/library/response"

)
    //文件：library/response/response_code.go 定义错误码
    
    //文件：app/system/order/internal/service/order.go
    func (orderService) Detail(ctx context.Context, req *define.OrderServiceDetailReq) (res *define.OrderServiceDetailRes, err error) {
        //gcode参数1:code 码
        //gcode参数2:code码需要替换的文字
        //gcode参数3:错误返回值 支持任何类型
        err=  gerror.NewCode(gcode.New(1000, "错误信息", "错误返回信息")) 
	    return
    }
    
    func (orderAPI) Detail(r *ghttp.Request) {
	var req *define.OrderServiceDetailReq
	if err := r.Parse(&req); err != nil {
		response.ParamErr(r, err)
	}
    //文件：app/system/order/internal/api/order.go
	if res, err := service.Order.Detail(r.Context(), req); err != nil {
		response.Code(r, err)
	} else {
		response.SuccessWithData(r, res)
	}

```    

