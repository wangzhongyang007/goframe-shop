package utility

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/crypto/gmd5"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gcharset"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/encoding/gurl"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"math/rand"
	"net"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

//密码加密
func EncryptPassword(password, salt string) string {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + gmd5.MustEncryptString(salt))
}

//时间戳转 yyyy-MM-dd HH:mm:ss
func TimeStampToDateTime(timeStamp int64) string {
	tm := gtime.NewFromTimeStamp(timeStamp)
	return tm.Format("Y-m-d H:i:s")
}

//时间戳转 yyyy-MM-dd
func TimeStampToDate(timeStamp int64) string {
	tm := gtime.NewFromTimeStamp(timeStamp)
	return tm.Format("Y-m-d")
}

//获取当前请求接口域名
func GetDomain(r *ghttp.Request) (string, error) {
	pathInfo, err := gurl.ParseURL(r.GetUrl(), -1)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("解析附件路径失败")
		return "", err
	}
	return fmt.Sprintf("%s://%s:%s/", pathInfo["scheme"], pathInfo["host"], pathInfo["port"]), nil
}

// GetUserAgent 获取user-agent
func GetUserAgent(ctx context.Context) string {
	return ghttp.RequestFromCtx(ctx).Header.Get("User-Agent")
}

//获取客户端IP
func GetClientIp(r *ghttp.Request) string {
	ip := r.Header.Get("X-Forwarded-For")
	if ip == "" {
		ip = r.GetClientIp()
	}
	return ip
}

//服务端ip
func GetLocalIP() (ip string, err error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, addr := range addrs {
		ipAddr, ok := addr.(*net.IPNet)
		if !ok {
			continue
		}
		if ipAddr.IP.IsLoopback() {
			continue
		}
		if !ipAddr.IP.IsGlobalUnicast() {
			continue
		}
		return ipAddr.IP.String(), nil
	}
	return
}

//获取ip所属城市
func GetCityByIp(ip string) string {
	if ip == "" {
		return ""
	}
	if ip == "[::1]" || ip == "127.0.0.1" {
		return "内网IP"
	}
	url := "http://whois.pconline.com.cn/ipJson.jsp?json=true&ip=" + ip
	bytes := g.Client().GetBytes(url)
	src := string(bytes)
	srcCharset := "GBK"
	tmp, _ := gcharset.ToUTF8(srcCharset, src)
	json, err := gjson.DecodeToJson(tmp)
	if err != nil {
		return ""
	}
	if json.GetInt("code") == 0 {
		city := fmt.Sprintf("%s %s", json.GetString("pro"), json.GetString("city"))
		return city
	} else {
		return ""
	}
}

//日期字符串转时间戳（秒）
func StrToTimestamp(dateStr string) int64 {
	tm, err := gtime.StrToTime(dateStr)
	if err != nil {
		g.Log().Error(err)
		return 0
	}
	return tm.Timestamp()
}

// GetDbConfig get db config
func GetDbConfig() (cfg *gdb.ConfigNode, err error) {
	cfg = g.DB().GetConfig()
	err = ParseDSN(cfg)
	return
}

// ParseDSN parses the DSN string to a Config
func ParseDSN(cfg *gdb.ConfigNode) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = gerror.New(r.(string))
		}
	}()
	dsn := cfg.Link
	if dsn == "" {
		return
	}
	foundSlash := false
	// gfast:123456@tcp(192.168.0.212:3306)/gfast-v2
	for i := len(dsn) - 1; i >= 0; i-- {
		if dsn[i] == '/' {
			foundSlash = true
			var j, k int

			// left part is empty if i <= 0
			if i > 0 {
				// [username[:password]@][protocol[(address)]]
				// Find the last '@' in dsn[:i]
				for j = i; j >= 0; j-- {
					if dsn[j] == '@' {
						// username[:password]
						// Find the first ':' in dsn[:j]
						for k = 0; k < j; k++ {
							if dsn[k] == ':' {
								cfg.Pass = dsn[k+1 : j]
								cfg.User = dsn[:k]
								break
							}
						}
						break
					}
				}

				// gfast:123456@tcp(192.168.0.212:3306)/gfast-v2
				// [protocol[(address)]]
				// Find the first '(' in dsn[j+1:i]
				var h int
				for k = j + 1; k < i; k++ {
					if dsn[k] == '(' {
						// dsn[i-1] must be == ')' if an address is specified
						if dsn[i-1] != ')' {
							if strings.ContainsRune(dsn[k+1:i], ')') {
								panic("invalid DSN: did you forget to escape a param value?")
							}
							panic("invalid DSN: network address not terminated (missing closing brace)")
						}
						for h = k + 1; h < i-1; h++ {
							if dsn[h] == ':' {
								cfg.Host = dsn[k+1 : h]
								cfg.Port = dsn[h+1 : i-1]
								break
							}
						}
						break
					}
				}
			}
			for j = i + 1; j < len(dsn); j++ {
				if dsn[j] == '?' {
					cfg.Name = dsn[i+1 : j]
					break
				} else {
					cfg.Name = dsn[i+1:]
				}
			}
			break
		}
	}
	if !foundSlash && len(dsn) > 0 {
		panic("invalid DSN: missing the slash separating the database name")
	}
	return
}

//获取附件真实路径
func GetRealFilesUrl(r *ghttp.Request, path string) (realPath string, err error) {
	if gstr.ContainsI(path, "http") {
		realPath = path
		return
	}
	realPath, err = GetDomain(r)
	if err != nil {
		return
	}
	realPath = realPath + path
	return
}

//获取附件相对路径
func GetFilesPath(fileUrl string) (path string, err error) {
	upType := gstr.ToLower(g.Cfg().GetString("upload.type"))
	upPath := gstr.Trim(g.Cfg().GetString("upload.local.UpPath"), "/")
	if upType != "local" || (upType == "local" && !gstr.ContainsI(fileUrl, upPath)) {
		path = fileUrl
		return
	}
	pathInfo, err := gurl.ParseURL(fileUrl, 32)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("解析附件路径失败")
		return
	}
	pos := gstr.PosI(pathInfo["path"], upPath)
	if pos >= 0 {
		path = gstr.SubStr(pathInfo["path"], pos)
	}
	return
}

//货币转化为分
func CurrencyLong(currency interface{}) int64 {
	strArr := gstr.Split(gconv.String(currency), ".")
	switch len(strArr) {
	case 1:
		return gconv.Int64(strArr[0]) * 100
	case 2:
		if len(strArr[1]) == 1 {
			strArr[1] += "0"
		} else if len(strArr[1]) > 2 {
			strArr[1] = gstr.SubStr(strArr[1], 0, 2)
		}
		return gconv.Int64(strArr[0])*100 + gconv.Int64(strArr[1])
	}
	return 0
}

func GetExcPath() string {
	file, _ := exec.LookPath(os.Args[0])
	// 获取包含可执行文件名称的路径
	path, _ := filepath.Abs(file)
	// 获取可执行文件所在目录
	index := strings.LastIndex(path, string(os.PathSeparator))
	ret := path[:index]
	return strings.Replace(ret, "\\", "/", -1)
}

//流水号
func CreateLogSn(prefix string) string {
	rand.Seed(time.Now().UnixNano())
	return prefix + strings.Replace(time.Now().Format("20060102150405.000"), ".", "", -1) + strconv.Itoa(rand.Intn(899)+100)
}

//获取随机整数
func RandInt(max int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(max)
}

//获取今天的开始时间 0点
//gtime.New(time.Now()).StartOfDay()

//获取今天的结束时间 24点
//gtime.New(time.Now()).EndOfDay()

//日期范围查询
//whereCondition.Set(dao.UserInfo.Columns.CreatedAt+" >=", gtime.New(req.Date).StartOfDay())
//whereCondition.Set(dao.UserInfo.Columns.CreatedAt+" <=", gtime.New(req.Date).EndOfDay())

func GetOrderNum() (number string) {
	rand.Seed(time.Now().UnixNano())
	number = gconv.String(time.Now().UnixNano()) + gconv.String(rand.Intn(1000))
	return
}

func GetRefundNum() (number string) {
	rand.Seed(time.Now().UnixNano())
	number = "refund" + gconv.String(time.Now().UnixNano()) + gconv.String(rand.Intn(1000))
	return
}

//生成最近一周的日期
func GetRecent7Date() (dates []string) {
	gt := gtime.New(time.Now())
	dates = []string{
		gt.Format("Y-m-d"),
		gt.Add(-gtime.D * 1).Format("Y-m-d"),
		gt.Add(-gtime.D * 2).Format("Y-m-d"),
		gt.Add(-gtime.D * 3).Format("Y-m-d"),
		gt.Add(-gtime.D * 4).Format("Y-m-d"),
		gt.Add(-gtime.D * 5).Format("Y-m-d"),
		gt.Add(-gtime.D * 6).Format("Y-m-d"),
	}
	return
}

//获取一周前的日期
func GetBefore7Date() (date string) {
	gt := gtime.New(time.Now())
	date = gt.Add(-gtime.D * 6).Format("Y-m-d")
	return
}
