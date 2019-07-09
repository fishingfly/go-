package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type gift struct {
	id       int      // 奖品Id
	name     string   //奖品名称
	pic      string   //奖品图片
	link     string   // 奖品的连接
	inuse    bool     //是否使用中
	rate     int      //中奖概率，万分之n，0-9999
	rateMin  int      // 大于等于最小中奖编码
	rateMax  int      // 小于中奖编码
}

const rateMax = 10
var logger *log.Logger

type lotteryController struct {
	Ctx iris.Context
}

func main() {
	app := newApp()
	app.Run(iris.Addr(":8080"))
}

//初始化日志
func initLog() {
	f, _ := os.Create("lottery_demo.log")
	logger = log.New(f, "", log.Ldate|log.Lmicroseconds)
}

func newApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotteryController{})
	initLog()
	return app
}

func newGift() *[5]gift {
	giftList := new([5]gift)
	g1 := gift{
		1,
		"富强福",
		"富强福.jpg",
		"",
		true,
		0,
		0,
		0,
	}
	g2 := gift{
		1,
		"和谐福",
		"和谐福.jpg",
		"",
		true,
		0,
		0,
		0,
	}
	g3 := gift{
		1,
		"友善福",
		"友善福.jpg",
		"",
		true,
		0,
		0,
		0,
	}
	g4 := gift{
		1,
		"敬业福",
		"敬业福.jpg",
		"",
		true,
		0,
		0,
		0,
	}
	g5 := gift{
		1,
		"爱国福",
		"爱国福.jpg",
		"",
		true,
		0,
		0,
		0,
	}
	giftList[0] = g1
	giftList[1] = g2
	giftList[2] = g3
	giftList[3] = g4
	giftList[4] = g5
	return giftList
}

func giftRage(rate string) *[5]gift {
	giftList := newGift()
	rates := strings.Split(rate, ",")
	ratesLen := len(rates)
	rateStart := 0
	for i, data := range giftList {
		if !data.inuse {
			continue
		}
		grate := 0
		if i < ratesLen {
			grate, _ = strconv.Atoi(rates[i])
		}
		giftList[i].rate = grate
		giftList[i].rateMin = rateStart
		giftList[i].rateMax = rateStart + grate
		if giftList[i].rateMax >= rateMax {
			giftList[i].rateMax = rateMax
			rateStart = 0
		} else {
			rateStart += grate
		}
	}
	fmt.Printf("giftlist=%v\n", giftList)
	return giftList
}

// GET http://localhost:8080/?rate=3,3,2,1,0
func (c *lotteryController) Get() string {
	rate := c.Ctx.URLParamDefault("rate", "3,3,2,1,0")
	giftList := giftRage(rate)
	return fmt.Sprintf("%v\n", giftList)
}

// http://localhost:8080/lucky
func (c *lotteryController) GetLucky() map[string]interface{} {
	uid , _ := c.Ctx.URLParamInt("uid")
	rate := c.Ctx.URLParamDefault("rate", "3,3,2,1,0")
	code := luckCode()
	ok := false
	result := make(map[string]interface{})
	result["success"] = ok
	giftList := giftRage(rate)
	for _, data := range giftList {
		if !data.inuse {
			continue
		}
		if data.rateMin <= int(code) && data.rateMax > int(code) {
			//中奖了，抽奖编码在奖品编码范围内
			//开始发奖
			sendData := data.pic
			ok = true
			if ok {
				// 中奖后，成功得到了奖品
				// 生成中奖纪录
				saveLuckyData(code, data.id, data.name, data.link,sendData)
				result["success"] = ok
				result["id"] = data.id
				result["name"] = data.name
				result["link"] = data.link
				result["data"] = sendData
				result["uid"] = uid
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

// 记录用户中奖信息
func saveLuckyData(code int32, id int, name, link,sendDate string) {
	logger.Printf("lucky, code = %d, id = %d,name = %s, link = %s, sendDate = %s\n",
		code, id, name, link,sendDate)
}