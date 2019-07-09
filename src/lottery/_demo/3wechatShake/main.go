/**
*微信摇一摇
*基础功能：
*lucky 只有一个抽奖的接口
 */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"sync"
	"time"
)

const (
	giftTypeCoin      = iota // 虚拟币
	giftTypeCoupon           // 不同券
	giftTypeCouponFix        // 相同的券
	giftTypeRealSmall        // 实物小奖
	giftTypeRealLarge        // 实物大奖
)

type gift struct {
	id       int      // 奖品Id
	name     string   //奖品名称
	pic      string   //奖品图片
	link     string   // 奖品的连接
	gtype    int      // 奖品类型
	data     string   // 奖品的数据
	dataList []string // 奖品数据集合（不同的优惠券的编码）
	total    int    //总数， 0 不限量
	left     int      //剩余数量
	inuse    bool     //是否使用中
	rate     int      //中奖概率，万分之n，0-9999
	rateMin  int      // 大于等于最小中奖编码
	rateMax  int      // 小于中奖编码
}

const rateMax = 10000

//最大的中奖号码
var logger *log.Logger

//奖品列表
var giftList []*gift

var mu sync.Mutex

type lotteryController struct {
	Ctx iris.Context
}

//初始化日志
func initLog() {
	f, _ := os.Create("lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

//初始化礼物
func initGift() {
	giftList = make([]*gift, 5)
	g1 := gift{
		1,
		"手机大奖",
		"",
		"",
		giftTypeRealLarge,
		"",
		nil,
		10000,
		10000,
		true,
		10000,
		0,
		10000,
	}
	g2 := gift{
		2,
		"充电器",
		"",
		"",
		giftTypeRealSmall,
		"",
		nil,
		1000,
		1000,
		false,
		200,
		0,
		20,
	}
	g3 := gift{
		3,
		"优惠券满200减50元",
		"",
		"",
		giftTypeCoupon,
		"mall-coupon-2019",
		nil,
		5,
		5,
		false,
		1000,
		0,
		1000,
	}
	g4 := gift{
		4,
		"直降优惠券50元",
		"",
		"",
		giftTypeCoupon,
		"",
		[]string{"c01","c02","c03","c04","c05"},
		5,
		5,
		false,
		2000,
		0,
		20000,
	}
	g5 := gift{
		5,
		"金币",
		"",
		"",
		giftTypeCoin,
		"10金币",
		nil,
		5,
		5,
		false,
		5000,
		0,
		5000,
	}
	giftList[0] = &g1
	giftList[1] = &g2
	giftList[2] = &g3
	giftList[3] = &g4
	giftList[4] = &g5
	// 数据整理，中奖区间数据
	rateStart := 0
	for _, data := range giftList {
		if !data.inuse {
			continue
		}
		data.rateMin = rateStart
		data.rateMax = rateStart + data.rate
		if data.rateMax >= rateMax {
			data.rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += data.rate
		}
	}
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	initLog()
	return app
}
func main() {
	app := newApp()
	initGift()
	initLog()
	mu = sync.Mutex{}
	app.Run(iris.Addr(":8080"))
}

func (c *lotteryController) Get() string {
	count := 0
	total := 0
	for _, data := range giftList {
		if data.inuse && (data.total == 0 ||
			(data.total > 0 && data.left > 0)) {
			count++
			total += data.left
		}
	}
	return fmt.Sprintf("当前有效奖品种类数量：%d, 限量奖品总数量= %d\n", count, total)
}
// http://localhost:8080/lucky
func (c *lotteryController) GetLucky() map[string]interface{} {
	mu.Lock()
	defer mu.Unlock()
	code := luckCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	result["data"] = "未抽中"
	for _, data := range giftList {
		if !data.inuse || (data.total > 0 && data.left <= 0) {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			//中奖了，抽奖编码在奖品编码范围内
			//开始发奖
			sendData := ""
			switch data.gtype {
			case giftTypeCoin:
				ok,sendData = sendCoin(data)
			case giftTypeCoupon:
				ok,sendData = sendCoupon(data)
			case giftTypeCouponFix:
				ok,sendData = sendCouponFix(data)
			case giftTypeRealSmall:
				ok,sendData = sendRealSmall(data)
			case giftTypeRealLarge:
				ok,sendData = sendRealLarge(data)
			}
			if ok {
				// 中奖后，成功得到了奖品
				// 生成中奖纪录
				saveLuckyData(code, data.id, data.name, data.link,sendData, data.left)
				result["success"] = ok
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData
				break
			}
		}
	}
	return result
}

func luckCode() int32 {
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Int31n(int32(rateMax))
	return code
}

 func sendCoin(data *gift) (bool, string) {
 	if data.total == 0 {
 		// 数量无限
 		return true,data.data
	} else if data.left > 0 {
		//还有剩余
		data.left--
		return true,data.data
	} else {
		return false, "奖品已发完！"
	}
 }

func sendCoupon(data * gift) (bool, string) {
	if data.left > 0 {
		// 还有剩余
		left := data.left - 1
		data.left = left
		return  true, data.data
	} else {
		return false, "奖品已发完"
	}
}


func sendCouponFix(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true,data.data
	} else if data.left > 0 {
		//还有剩余
		data.left--
		return true,data.data
	} else {
		return false, "奖品已发完！"
	}
}

func sendRealSmall(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true,data.data
	} else if data.left > 0 {
		//还有剩余
		data.left--
		return true,data.data
	} else {
		return false, "奖品已发完！"
	}
}

func sendRealLarge(data *gift) (bool, string) {
	if data.total == 0 {
		// 数量无限
		return true,data.data
	} else if data.left > 0 {
		//还有剩余
		data.left--
		return true,data.data
	} else {
		return false, "奖品已发完！"
	}
}

// 记录用户中奖信息
func saveLuckyData(code int32, id int, name, link,sendDate string, left int) {
	logger.Printf("lucky, code = %d, id = %d,name = %s, link = %s, sendDate = %s, left = %d",
		code, id, name, link,sendDate, left)
}