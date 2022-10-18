package backend

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"shop/app/middleware"
	"shop/app/service/login"
	"shop/app/system/backend/admin"
	"shop/app/system/backend/article"
	"shop/app/system/backend/category"
	"shop/app/system/backend/comment"
	"shop/app/system/backend/consignee"
	"shop/app/system/backend/coupon"
	"shop/app/system/backend/data"
	"shop/app/system/backend/goods"
	"shop/app/system/backend/goodsOptions"
	"shop/app/system/backend/order"
	"shop/app/system/backend/permission"
	"shop/app/system/backend/refund"
	"shop/app/system/backend/role"
	"shop/app/system/backend/rotation"
	"shop/app/system/backend/upload"
	"shop/app/system/backend/user"
)

func Init(s *ghttp.Server) {
	backendLogin()
	s.Group("/backend/", func(group *ghttp.RouterGroup) {
		//不需要登录就能访问的接口
		group.Group("sso/", func(group *ghttp.RouterGroup) {
			group.POST("register/", admin.Admin.Add)
		})

		//以下需要登录
		group.Middleware(middleware.MiddlewareGToken.GetToken)
		//上传文件
		group.Group("upload/", func(group *ghttp.RouterGroup) {
			group.POST("img/", upload.Upload.Img)
		})
		//数据大屏
		group.Group("data/", func(group *ghttp.RouterGroup) {
			group.POST("head/", data.Data.HeadCard)
			group.POST("echarts/", data.Data.ECharts)
		})
		//轮播图管理
		group.Group("rotation/", func(group *ghttp.RouterGroup) {
			group.POST("add/", rotation.Rotation.Add)
			group.POST("update/", rotation.Rotation.Update)
			group.POST("delete/", rotation.Rotation.Delete)
			group.POST("list/", rotation.Rotation.List)
		})
		//管理员管理
		group.Group("admin/", func(group *ghttp.RouterGroup) {
			group.POST("add/", admin.Admin.Add)
			group.POST("update/", admin.Admin.Update)
			group.POST("delete/", admin.Admin.Delete)
			group.POST("list/", admin.Admin.List)
			group.POST("update/my/password", admin.Admin.UpdateMyPassword) //修改自己的密码
		})
		//角色管理
		group.Group("role/", func(group *ghttp.RouterGroup) {
			group.POST("add/", role.Role.Add)
			group.POST("update/", role.Role.Update)
			group.POST("delete/", role.Role.Delete)
			group.POST("list/", role.Role.List)
			group.POST("add/permission", role.Role.AddPermission)       //添加角色权限关联
			group.POST("delete/permission", role.Role.DeletePermission) //去掉角色权限关联
		})
		//权限管理
		group.Group("permission/", func(group *ghttp.RouterGroup) {
			group.POST("add/", permission.Permission.Add)
			group.POST("update/", permission.Permission.Update)
			group.POST("delete/", permission.Permission.Delete)
			group.POST("list/", permission.Permission.List)
		})
		//会员（用户）管理
		group.Group("user/", func(group *ghttp.RouterGroup) {
			//group.POST("add/", admin.Admin.Add)
			//group.POST("delete/", user.User.Delete)
			group.POST("update/", user.User.Update)
			group.POST("list/", user.User.List)
			group.POST("order/list/", order.Order.List)
			group.POST("comment/list/", comment.Comment.List)
			group.POST("article/list/", article.Article.List)
		})
		//评价管理
		group.Group("comment/", func(group *ghttp.RouterGroup) {
			group.POST("list/", comment.Comment.List)
			group.POST("delete/", comment.Comment.Delete)
		})
		//文章 种草
		group.Group("article/", func(group *ghttp.RouterGroup) {
			group.POST("add/", article.Article.Add)
			group.POST("update/", article.Article.Update)
			group.POST("delete/", article.Article.Delete)
			group.POST("list/", article.Article.List)
		})
		//优惠券
		group.Group("coupon/", func(group *ghttp.RouterGroup) {
			group.POST("add/", coupon.Coupon.Add)
			group.POST("update/", coupon.Coupon.Update)
			group.POST("delete/", coupon.Coupon.Delete)
			group.POST("list/", coupon.Coupon.List)
		})
		//分类
		group.Group("category/", func(group *ghttp.RouterGroup) {
			group.POST("add/", category.Category.Add)
			group.POST("update/", category.Category.Update)
			group.POST("delete/", category.Category.Delete)
			group.POST("list/", category.Category.List)
			group.POST("level/list/", category.Category.LevelList) //分级列表
		})
		//商品
		group.Group("goods/", func(group *ghttp.RouterGroup) {
			group.POST("add/", goods.Goods.Add)
			group.POST("update/", goods.Goods.Update)
			group.POST("delete/", goods.Goods.Delete)
			group.POST("list/", goods.Goods.List)
			group.POST("detail/", goods.Goods.Detail)
		})
		//商品规格 SKU
		group.Group("goods/sku/", func(group *ghttp.RouterGroup) {
			group.POST("add/", goodsOptions.GoodsOptions.Add)
			group.POST("update/", goodsOptions.GoodsOptions.Update)
			group.POST("delete/", goodsOptions.GoodsOptions.Delete)
			group.POST("list/", goodsOptions.GoodsOptions.List)
			group.POST("detail/", goodsOptions.GoodsOptions.Detail)
		})
		//订单
		group.Group("order/", func(group *ghttp.RouterGroup) {
			//group.POST("update/", order.Order.Update)
			//group.POST("delete/", order.Order.Delete)
			group.POST("list/", order.Order.List)
			group.POST("detail/", order.Order.Detail)
		})
		//售后
		group.Group("refund/", func(group *ghttp.RouterGroup) {
			//group.POST("update/", refund.Refund.Update)
			//group.POST("delete/", refund.Refund.Delete)
			group.POST("list/", refund.Refund.List)
			group.POST("detail/", refund.Refund.Detail)
		})
		//收货地址
		group.Group("consignee/", func(group *ghttp.RouterGroup) {
			group.POST("list/", consignee.Consignee.List)
		})
	})
}

func backendLogin() {
	// 启动gtoken
	middleware.GToken = &gtoken.GfToken{
		//Timeout:    gconv.Int(g.Cfg().Get("gtoken.timeout")) * 60 * 1000,
		//MaxRefresh: 60 * 1000, //单位毫秒 登录1分钟后有请求操作则主动刷新token有效期
		CacheMode:  1, //缓存模式 1 gcache 2 gredis 默认1
		LoginPath:  "/backend/sso/login",
		LogoutPath: "/backend/sso/logout",
		AuthPaths:  g.SliceStr{},
		//AuthPaths:        g.SliceStr{"/backend"},
		AuthExcludePaths: g.SliceStr{},
		GlobalMiddleware: true, // 开启全局拦截
		//MultiLogin:       g.Config().GetBool("gtoken.multi-login"),
		LoginBeforeFunc: login.Login.Login,
		LoginAfterFunc:  login.Login.LoginAfterFunc,
		LogoutAfterFunc: login.Login.Logout,
		AuthAfterFunc:   login.Login.AuthAfterFunc,
	}
	middleware.GToken.Start()
}
