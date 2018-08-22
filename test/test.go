package main

import (
	"fmt"
	"log"

	"github.com/robfig/cron"
	uuid "github.com/satori/go.uuid"
)

var (
	years   = ""
	weeks   = "*"
	month   = "*"
	day     = "*"
	hours   = "*"
	points  = "*"
	seconds = "*"
)

func main() {

	// fmt.Sprintf("%s %s %s %s %s", seconds, points, hours, day, month, weeks, years)
	// fmt.Println(time.Now())

	// t1 := time.Now().Year() //年

	// t2 := time.Now().Month() //月

	// t3 := time.Now().Day() //日

	// t4 := time.Now().Hour() //小时

	// t5 := time.Now().Minute() //分钟

	// t6 := time.Now().Second() //秒

	// fmt.Println(t1, t2, t3, t4, t5, t6) //打印结果：2017 April 11 12 52 52
	// *  * 16 *  *  ?   *
	// 秒 分 时 日 月 星期 年
	// i := 0
	// c := cron.New()
	// spec := "* * 16 * * ?"
	// c.AddFunc(spec, func() {
	// 	i++
	// 	log.Println("cron running:", i)
	// })
	// c.Start()
	// select {}

	// EveryDay()
	// WorkDay()

	// MonthDay()

	// hdata := hashids.NewData()
	// hdata.MinLength = 6
	// hdata.Salt = time.Now().String() // 设置秘钥

	// hid, _ := hashids.NewWithData(hdata)

	// numbers := []int{20, 199, 260, 99}
	// hash, err := hid.Encode(numbers)
	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Printf("%v -> %v", numbers, hash)
	// fmt.Printf("%v ", NewId())

	// if hash != "7nnhzEsDkiYa" {
	// 	fmt.Println("hash does not match expected one")
	// }

}

func NewId() string {
	id, _ := uuid.NewV4()
	return id.String()
}

// 每天执行  早上六点十分执行
func EveryDay() {
	i := 0
	c := cron.New()
	// spec := "* * 6 * * ?"
	hours = "16"
	points = "35"
	spec := fmt.Sprintf("%s %s %s %s %s %s %s", seconds, points, hours, day, month, weeks, years)

	fmt.Println(spec)
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
	select {}
}

// 工作日 早上六点十分执行
func WorkDay() {

	i := 0
	c := cron.New()
	// 0 10 6 0 0 1/5 *
	// spec := "0 46 16 ? * MON,TUE,WED,THU,FRI *"
	seconds = "0"
	hours = "16"
	points = "46"
	weeks = "MON,TUE,WED,THU,FRI"
	spec := fmt.Sprintf("%s %s %s %s %s %s %s", seconds, points, hours, day, month, weeks, years)

	fmt.Println(spec)
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
	select {}

}

// 每周 例如 周四 下午 16 ： 55 分
func WeekDay() {
	i := 0
	c := cron.New()
	// spec := "0 46 16 ? * MON,TUE,WED,THU,FRI,SAT,SUN *"
	//
	seconds = "*"
	hours = "16"
	points = "55"
	weeks = "THU"
	spec := fmt.Sprintf("%s %s %s %s %s %s %s", seconds, points, hours, day, month, weeks, years)

	fmt.Println(spec)
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
	select {}
}

// 每月 每一天执行
func MonthDay() {
	i := 0
	c := cron.New()
	// spec := "0 12 17 9 * ? *"
	//
	seconds = "*"
	hours = "17"
	points = "12"
	day = "9"
	spec := fmt.Sprintf("%s %s %s %s %s %s %s", seconds, points, hours, day, month, weeks, years)

	fmt.Println(spec)
	c.AddFunc(spec, func() {
		i++
		log.Println("cron running:", i)
	})
	c.Start()
	select {}
}
