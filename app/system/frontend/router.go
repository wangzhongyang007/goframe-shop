package frontend

import (
	"github.com/goflyfox/gtoken/gtoken"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"shop/app/middleware"
	"shop/app/service/frontendLogin"
	"shop/app/system/backend/rotation"
	"shop/app/system/backend/upload"
	"shop/app/system/frontend/address"
	"shop/app/system/frontend/article"
	"shop/app/system/frontend/cart"
	"shop/app/system/frontend/category"
	"shop/app/system/frontend/collection"
	"shop/app/system/frontend/comment"
	"shop/app/system/frontend/consignee"
	"shop/app/system/frontend/goods"
	"shop/app/system/frontend/order"
	"shop/app/system/frontend/praise"
	"shop/app/system/frontend/refund"
	"shop/app/system/frontend/userCoupon"
)

func Init(s *ghttp.Server) {
	Login()
	s.Group("/frontend/", func(group *ghttp.RouterGroup) {
		//不需要登录的
		//上传文件
		group.Group("upload/", func(group *ghttp.RouterGroup) {
			group.POST("img/", upload.Upload.Img)
		})
		//注册及找回密码
		group.Group("sso/", func(group *ghttp.RouterGroup) {
			group.POST("register/", frontendLogin.FrontendLogin.Register)
		})
		//商品
		group.Group("goods/", func(group *ghttp.RouterGroup) {
			group.POST("list/", goods.Goods.List)
			group.POST("detail/", goods.Goods.Detail)
			group.POST("category/", goods.Goods.Category)
		})
		//轮播图
		group.Group("rotation/", func(group *ghttp.RouterGroup) {
			group.POST("list/", rotation.Rotation.List)
		})
		//分类
		group.Group("category/", func(group *ghttp.RouterGroup) {
			group.POST("list/", category.Category.List)
		})
		//收货地址
		group.Group("address/", func(group *ghttp.RouterGroup) {
			group.POST("list/", address.Address.List)
		})

		//以下是需要登录的
		group.Middleware(middleware.MiddlewareGToken.GetToken)
		//登录账号相关
		group.Group("sso/", func(group *ghttp.RouterGroup) {
			group.POST("password/update", frontendLogin.FrontendLogin.UpdatePassword)
		})
		//购物车
		group.Group("cart/", func(group *ghttp.RouterGroup) {
			group.POST("add/", cart.Cart.Add)
			group.POST("update/", cart.Cart.Update)
			group.POST("delete/", cart.Cart.Delete)
			group.POST("list/", cart.Cart.List)
		})
		//优惠券
		group.Group("user/coupon/", func(group *ghttp.RouterGroup) {
			group.POST("add/", userCoupon.UserCoupon.Add) //领券
			group.POST("list/", userCoupon.UserCoupon.List)
		})
		//订单
		group.Group("order/", func(group *ghttp.RouterGroup) {
			group.POST("add/", order.Order.Add) //生成订单
			group.POST("list/", order.Order.List)
		})
		//售后
		group.Group("refund/", func(group *ghttp.RouterGroup) {
			group.POST("add/", refund.Refund.Add) //申请售后
			group.POST("list/", refund.Refund.List)
		})
		//收货地址
		group.Group("consignee/", func(group *ghttp.RouterGroup) {
			group.POST("add/", consignee.Consignee.Add)
			group.POST("update/", consignee.Consignee.Update)
			group.POST("delete/", consignee.Consignee.Delete)
			group.POST("list/", consignee.Consignee.List)
		})
		//收藏
		group.Group("collection/", func(group *ghttp.RouterGroup) {
			group.POST("add/", collection.Collection.Add)
			group.POST("delete/", collection.Collection.Delete)
			group.POST("list/", collection.Collection.List)
		})
		//文章 种草
		group.Group("article/", func(group *ghttp.RouterGroup) {
			group.POST("add/", article.Article.Add)
			group.POST("update/", article.Article.Update)
			group.POST("delete/", article.Article.Delete)
			group.POST("list/", article.Article.List)      //全部文章列表
			group.POST("my/list/", article.Article.MyList) //我的文章列表
			group.POST("detail/", article.Article.Detail)  //文章详情
		})
		//点赞
		group.Group("praise/", func(group *ghttp.RouterGroup) {
			group.POST("add/", praise.Praise.Add)
			group.POST("delete/", praise.Praise.Delete)
			group.POST("list/", praise.Praise.List)
		})
		//评论
		group.Group("comment/", func(group *ghttp.RouterGroup) {
			group.POST("add/", comment.Comment.Add)
			group.POST("delete/", comment.Comment.Delete)
			group.POST("list/", comment.Comment.List)
		})
	})
}

//前端项目登录
func Login() {
	// 启动gtoken
	middleware.GToken = &gtoken.GfToken{
		//都用默认的
		//Timeout:    gconv.Int(g.Cfg().Get("gtoken.timeout")) * gconv.Int(gtime.M),
		//MaxRefresh: 60 * 1000, //单位毫秒 登录1分钟后有请求操作则主动刷新token有效期
		CacheMode:  2,
		LoginPath:  "/frontend/sso/login",
		LogoutPath: "/frontend/sso/logout",
		AuthPaths:  g.SliceStr{},
		//AuthPaths:        g.SliceStr{"/backend"},
		AuthExcludePaths: g.SliceStr{},
		GlobalMiddleware: true, // 开启全局拦截
		//MultiLogin:       g.Config().GetBool("gtoken.multi-login"),
		LoginBeforeFunc: frontendLogin.FrontendLogin.Login,
		LoginAfterFunc:  frontendLogin.FrontendLogin.LoginAfterFunc,
		LogoutAfterFunc: frontendLogin.FrontendLogin.Logout,
		AuthAfterFunc:   frontendLogin.FrontendLogin.AuthAfterFunc,
	}
	middleware.GToken.Start()
}
