package common

import (
	"context"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"github.com/gogf/gf/crypto/gmd5"
	"github.com/gogf/gf/database/gdb"
	"github.com/gogf/gf/encoding/gurl"
	"github.com/gogf/gf/errors/gerror"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/os/gctx"
	"github.com/gogf/gf/os/gfile"
	"github.com/gogf/gf/os/gtime"
	"github.com/gogf/gf/text/gregex"
	"github.com/gogf/gf/util/gconv"
	"github.com/gogf/gf/util/grand"
	"io"
	"math"
	mrand "math/rand"
	"net"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"
)

func InArray(item int, items []int) bool {
	for _, v := range items {
		if v == item {
			return true
		}
	}
	return false
}

func MapIntStingInArray(item string, items map[int]string) int {
	for k, v := range items {
		if item == v {
			return k
		}
	}
	return 0
}
func Md5V(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
func UniqueId() string {
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return GetMd5(base64.URLEncoding.EncodeToString(b))
}

func GetMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

//Salt 生成
func Salt() string {
	builder := strings.Builder{}
	mrand.Seed(time.Now().UnixNano())
	for i := 0; i < 4; i++ {
		builder.WriteString(gconv.String(mrand.Intn(9)))
	}
	return builder.String()
}

//Base34 分享码
func Base34(num int64) (res string) {
	base := strings.Split("0123456789ABCDEFGHJKLMNPQRSTUVWXYZ", "")
	for num >= 34 {
		i := num % 34
		res = base[i] + res
		num = num / 34
	}
	res = base[num] + res
	return strings.Repeat("0", 6-len(res)) + res
}

//MapIntStingInArrayResult 判断这个键名是否在数组中是否在数组中并返回键值与键名
func MapIntStingKeyInArrayResult(item int, items map[int]string) (int, string) {
	for k, v := range items {
		if item == k {
			return k, v
		}
	}
	return 0, ""
}

func StrIsIn(str string, strs []string) bool {
	for _, v := range strs {
		if str == v {
			return true
		}
	}
	return false
}

/**
指定个数分割数组
*/
func Chunk(arr []gdb.Value, base int) (res [][]gdb.Value) {
	Len := len(arr)
	num := math.Ceil(float64(Len) / float64(base)) //分割次数
	var i int
	for i = 1; i <= int(num); i++ {
		low := ((i - 1) * base)
		height := ((i-1)*base + base)
		if height > Len {
			height = Len
		}
		tmp := arr[low:height:Len]
		res = append(res, tmp)
	}
	return
}

func ChunkInts(arr []int, base int) (res [][]int) {
	Len := len(arr)
	num := math.Ceil(float64(Len) / float64(base)) //分割次数
	var i int
	for i = 1; i <= int(num); i++ {
		low := ((i - 1) * base)
		height := ((i-1)*base + base)
		if height > Len {
			height = Len
		}
		tmp := arr[low:height:Len]
		res = append(res, tmp)
	}
	return
}

func ChunkInterface(arr []interface{}, base int) (res [][]interface{}) {
	Len := len(arr)
	num := math.Ceil(float64(Len) / float64(base)) //分割次数
	var i int
	for i = 1; i <= int(num); i++ {
		low := ((i - 1) * base)
		height := ((i-1)*base + base)
		if height > Len {
			height = Len
		}
		tmp := arr[low:height:Len]
		res = append(res, tmp)
	}
	return
}

/**
int去重
*/
func IntSliceUnique(data []int) (res []int) {
	if len(data) == 0 {
		return
	}
	dict := map[int]bool{}
	for _, v := range data {
		if dict[v] == true {
			continue
		}
		dict[v] = true
		res = append(res, v)
	}
	return
}

/**
string去重
*/
func StrSliceUnique(data []string) (res []string) {
	if len(data) == 0 {
		return
	}
	dict := map[string]bool{}
	for _, v := range data {
		if dict[v] == true {
			continue
		}
		dict[v] = true
		res = append(res, v)
	}
	return
}

/**
 * 数组去重 去空
 */
func RemoveDuplicatesAndEmpty(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

//GetDayTime i 1 昨天 7 7天前  15 15天前
// i天前的 00：00：00 与昨天的 23:59:59
func GetDayTime(i int) (startTime, endTime int64) {
	currentTime := time.Now()
	startTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-i, 00, 00, 00, 0, currentTime.Location()).Unix()
	endTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day()-1, 23, 59, 59, 0, currentTime.Location()).Unix()
	return
}

//GetStartAndEndTime 获取时间戳当天的 00：00：00 与 23:59:59
func GetStartAndEndTime(t int64) (startTime, endTime int64) {
	currentTime := time.Unix(t, 0)
	startTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 00, 00, 00, 0, currentTime.Location()).Unix()
	endTime = time.Date(currentTime.Year(), currentTime.Month(), currentTime.Day(), 23, 59, 59, 0, currentTime.Location()).Unix()
	return
}

func CutTime(createTime, endTime int) (arr []int) {
	newTime := createTime
	for {
		if newTime <= endTime {
			arr = append(arr, newTime)
			newTime += 86400
		} else {
			return arr
		}
	}
}
func SotrTowMap(towMap []map[string]interface{}) []map[string]interface{} {
	var tmp []map[string]interface{}
	lens := len(towMap)
	var i int
	for i = lens - 1; i >= 0; i-- {
		tmp = append(tmp, towMap[i])
	}
	return tmp
}

////字符串获取首字母
//func StrInitial(hans string) (initial string) {
//	//数字开头
//	if unicode.IsDigit(int32(hans[0])) {
//		initial = strings.ToUpper(string([]byte(hans)[:1]))
//		return
//	}
//	//字母开头且不是汉字
//	if unicode.IsLetter(int32(hans[0])) && int32(hans[0]) >= 65 && int32(hans[0]) <= 122 {
//		initial = strings.ToUpper(string([]byte(hans)[:1]))
//		return
//	}
//
//	//第一个匹配的是全部元素
//	var hzRegexp, _ = regexp.Compile("([\u4e00-\u9fa5]+)")
//	sub := hzRegexp.FindSubmatch([]byte(hans))
//	if len(sub) == 0 {
//		return
//	}
//	pynew := pinyin.NewArgs()
//
//	pys := pinyin.Pinyin(string(sub[0]), pynew)[0][0]
//	initial = strings.ToUpper(string([]byte(pys)[:1]))
//	return
//}

//两数差集 只差slice1的差集
func Difference(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	inter := Intersect(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}

//并集
func Intersect(slice1, slice2 []string) []string {
	m := make(map[string]int)
	nn := make([]string, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

//数组平分
func splitArray(arr []int, num int64) [][]int {
	max := int64(len(arr))
	if max < num {
		return nil
	}
	var segmens = make([][]int, 0)
	quantity := max / num
	end := int64(0)
	for i := int64(1); i <= num; i++ {
		qu := i * quantity
		if i != num {
			segmens = append(segmens, arr[i-1+end:qu])
		} else {
			segmens = append(segmens, arr[i-1+end:])
		}
		end = qu - i
	}
	return segmens
}
func InArrayString(need string, needArr []string) bool {
	for _, v := range needArr {
		if need == v {
			return true
		}
	}
	return false
}

//Sn 流水号
func Sn(pre ...string) (res string) {

	if len(pre) > 0 {
		res = pre[0]
	}
	res += gtime.Now().Format("YmdHisu") + grand.Digits(3)
	return
}

func LocalPath() (path string, err error) {
	path, err = mkdir()
	if err != nil {
		return
	}
	path = uniqueFilename(path)
	return
}

func mkdir() (path string, err error) {
	path = g.Log().GetPath() + "/files/"
	err = gfile.Mkdir(path)
	return
}

func uniqueFilename(path string) string {
	return path + Sn() + ".xlsx"
}

//Split 字符串分割
func Split(s, sep string) []string {
	if s == "" {
		return []string{}
	}
	return strings.Split(s, sep)
}

//UserPhone 脱敏手机号
func UserPhone(phone string) (res string) {
	if phone == "" {
		return ""
	}
	res, _ = gregex.ReplaceString(`(\d{3})(\d{4})(\d{4})`, "$1****$3", phone)
	return
}

//UserName 脱敏文字
func UserName(name string) string {
	nameString := []rune(name)
	newName := ""
	reg := `^[A-Za-z]+$`
	regNum := `^[0-9]+$`
	rgx := regexp.MustCompile(reg)
	resAz := rgx.MatchString(name)
	rgxNum := regexp.MustCompile(regNum)
	resNum := rgxNum.MatchString(name)
	if resAz == true || resNum == true || utf8.RuneCountInString(name) <= 2 {
		rangeright := utf8.RuneCountInString(name) - 1
		if rangeright < 1 {
			rangeright = 1
		}
		numLeft := string(nameString[0:rangeright])
		newName = numLeft + "*"
	} else if utf8.RuneCountInString(name) == 3 {
		numLeft := string(nameString[0:1])
		numRight := string(nameString[2:3])
		newName = numLeft + "*" + numRight
	} else {
		numLeft := string(nameString[0:1])
		numRange := utf8.RuneCountInString(name) - 2
		Encryption := ""
		for i := 0; i <= numRange; i++ {
			Encryption += "*"
		}
		rangeright := utf8.RuneCountInString(name) - 1
		if rangeright < 1 {
			rangeright = 0
		}
		numRight := string(nameString[rangeright:utf8.RuneCountInString(name)])
		newName = numLeft + Encryption + numRight
	}

	return newName
}
func Take(a float64, b int) float64 {
	f1 := gconv.Float64(a) * gconv.Float64(b)
	return gconv.Float64(f1)
}

//业务层是uint 没有泛型只能固定
func IsRepeat(num []uint) bool {
	hash := make(map[interface{}]bool)
	for _, v := range num {
		if hash[v] == true {
			return true
		} else {
			hash[v] = true
		}
	}
	return false
}

func InArrayInt(search uint, values []uint) bool {
	for _, value := range values {
		if value == search {
			return true
		}
	}
	return false
}
func DelItems(vs []int, dels []int) []int {
	dMap := make(map[int]bool)
	for _, s := range dels {
		dMap[s] = true
	}

	for i := 0; i < len(vs); i++ {
		if _, ok := dMap[vs[i]]; ok {
			vs = append(vs[:i], vs[i+1:]...)
			i = i - 1
		}
	}
	return vs
}

/*
获取真实IP
*/

func ExternalIP() (net.IP, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return nil, err
		}
		for _, addr := range addrs {
			ip := getIpFromAddr(addr)
			if ip == nil {
				continue
			}
			return ip, nil
		}
	}
	return nil, errors.New("connected to the network?")
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

//两数差集 只差slice1的差集
func DifferenceInt(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	inter := IntersectInt(slice1, slice2)
	for _, v := range inter {
		m[v]++
	}
	for _, value := range slice1 {
		times, _ := m[value]
		if times == 0 {
			nn = append(nn, value)
		}
	}
	return nn
}
func IntersectInt(slice1, slice2 []int) []int {
	m := make(map[int]int)
	nn := make([]int, 0)
	for _, v := range slice1 {
		m[v]++
	}
	for _, v := range slice2 {
		times, _ := m[v]
		if times == 1 {
			nn = append(nn, v)
		}
	}
	return nn
}

func Password(password, Salt string) (res string) {
	return gmd5.MustEncryptString(gmd5.MustEncryptString(password) + Salt)
}

//判断文件夹是否存在
func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func Reverse(arr *[]int, length int) {
	var temp int
	for i := 0; i < length/2; i++ {
		temp = (*arr)[i]
		(*arr)[i] = (*arr)[length-1-i]
		(*arr)[length-1-i] = temp
	}
}

//获取模型
func GetModule(url string) (string, error) {
	pathInfo, err := gurl.ParseURL(url, -1)
	if err != nil {
		g.Log().Error(err)
		err = gerror.New("解析附件路径失败")
		return "", err
	}
	urls := strings.Split(gconv.String(pathInfo["host"]), ".")

	return urls[0], nil
}

func Ctx(pre ...string) context.Context {
	var ctx = context.WithValue(gctx.New(), "RequestId", Sn(pre...))
	return ctx
}

//获取当前时间戳 uint
func Timestamp() uint {
	return gconv.Uint(gtime.Timestamp())
}
