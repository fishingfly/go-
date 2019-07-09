/**
*1 即开即得型
*2 双色球自选型
 */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
	"math/rand"
	"time"
)

type lotterController struct {
	Ctx iris.Context
}

func NewApp() *iris.Application {
	app := iris.New()
	mvc.New(app.Party("/")).Handle(&lotterController{})
	return app
}

func main() {
	app := NewApp()
	app.Run(iris.Addr(":8080"))

}

//即开即得型 http://localhost:8080/
func (c *lotterController) Get() string {
	var prize string
	seed := time.Now().UnixNano()
	code := rand.New(rand.NewSource(seed)).Intn(10)
	switch {
	case code == 1:
		prize = "一等奖"
	case code >= 2 && code <= 3:
		prize = "二等奖"
	case code >=4 && code <= 6:
		prize = "三等奖"
	default:
		return fmt.Sprintf("尾号为1获得一等奖<br/>" + "尾号为2,3获得二等奖<br/>" + "尾号为4,5,6获得三等奖<br/>"+ "code = %d<br/>"+"很遗憾，您未中奖<br/>", code)
	}
	return fmt.Sprintf("尾号为1获得一等奖<br/>" + "尾号为2,3获得二等奖<br/>" + "尾号为4,5,6获得三等奖<br/>"+ "code = %d<br/>"+"恭喜您，获得%s<br/>", code, prize)
}

// 双色球自选型 http://localhost:8080/prize
func (c *lotterController) GetPrize() string {
	seed := time.Now().UnixNano()
	r := rand.New(rand.NewSource(seed))
	var price [7]int
	for i := 0; i < 6; i++ {
		price[i] = r.Intn(33) + 1
	}
	// 最后一位的蓝色球, 1-16
	price[6] = r.Intn(16) + 1
	return fmt.Sprintf("今日开奖号码是: %v", price)
}

