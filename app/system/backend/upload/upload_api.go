package upload

import (
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"os"
	"shop/library/response"
)

var Upload = uploadApi{}

type uploadApi struct{}

// Upload uploads files to /tmp .
func (*uploadApi) Img(r *ghttp.Request) {
	files := r.GetUploadFiles("file")
	dirPath := "/tmp/"
	names, err := files.Save(dirPath, true)
	if err != nil {
		r.Response.WriteExit(err)
	}

	for _, name := range names {
		localFile := dirPath + name
		bucket := g.Cfg().GetString("qiniu.bucket")
		key := name
		accessKey := g.Cfg().GetString("qiniu.accessKey")
		secretKey := g.Cfg().GetString("qiniu.secretKey")

		putPolicy := storage.PutPolicy{
			Scope: bucket,
		}
		mac := qbox.NewMac(accessKey, secretKey)
		upToken := putPolicy.UploadToken(mac)

		cfg := storage.Config{}
		// 空间对应的机房
		cfg.Zone = &storage.ZoneHuabei
		// 是否使用https域名
		cfg.UseHTTPS = true
		// 上传是否使用CDN上传加速
		cfg.UseCdnDomains = false

		// 构建表单上传的对象
		formUploader := storage.NewFormUploader(&cfg)
		ret := storage.PutRet{}

		// 可选配置
		putExtra := storage.PutExtra{
			Params: map[string]string{},
		}

		err = formUploader.PutFile(r.GetCtx(), &ret, upToken, key, localFile, &putExtra)
		if err != nil {
			response.FailureWithData(r, 0, err, "")
		}

		fmt.Println(ret.Key, ret.Hash)

		//删除本地文件
		err = os.Remove(localFile)
		if err != nil {
			g.Dump("删除本地文件失败：", err)
		}
		fmt.Println("删除本地文件成功", localFile)

		//返回数据
		response.SuccessWithData(r, g.Map{
			"url": g.Cfg().GetString("qiniu.url") + ret.Key,
		})
	}
}
