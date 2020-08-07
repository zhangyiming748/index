---
layout:     post                    # 使用的布局(不需要改)
title:     GetProgrammingWithGo             # 标题
subtitle:   我是一只被禁足的安小鸟 #副标题
date:       2020-04-17 00:00:01 GMT+0800             # 时间
author:     Zen                 # 作者
header-img: img/photo/birdAngle.webp    #这篇文章标题背景图片
catalog: Ture                   # 是否归档
tags:                               #标签
    - 读书
---

### 第2章 实验
```
package main

import (
	"fmt"
	_ "log"
)

const (
	DISTANCE   = 56000000
	HOURPERDAY = 24
	DAY        = 28
)

func myAnswer() {
	fmt.Println(DISTANCE / (DAY * HOURPERDAY))
}
func main() {
	myAnswer()
	answer()
}
func answer() {
	fmt.Printf("%v", 56000000/(24*28))
}
```
### 第3章 实验
```
func myAnswer() {
	num := rand.Intn(100) + 1
	var guess int


	for {
		log.Println("input a num")
		_,_=fmt.Scan(&guess)
		if guess>num {
			fmt.Println("more than answer")
			continue
		}else if guess<num {
			fmt.Println("less than answer")
			continue
		}else {
			fmt.Println("right answers")
			return
		}
	}

}
```
### 第4章 实验
```
package main

import (
	"fmt"
	"math/rand"
)

var era = "AD"

func myAnswer() {
	for i := 0; i < 10; i++ {
		year := rand.Intn(2020) + 1
		month := rand.Intn(12) + 1
		daysInMonth := 31
		leap := isLeapYear(year)
		switch leap {
		case true:
			if month == 2 {
				daysInMonth = 28
			}

		case false:
			daysInMonth = 30
		}

		day := rand.Intn(daysInMonth) + 1

		fmt.Println(era, year, month, day)
	}

}
func isLeapYear(year int) bool {
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}
```
### 第5章 实验
```
package main

import (
	"fmt"
	"math/rand"
)

const (
	secondsPerDay = 86400
)

func ticket() {

	var (
		distance int
		company  string
		trip     string
	)
	fmt.Println("Spaceline\tDays\tTrip\tType\tPrice")
	fmt.Println("==================================")
	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "Intel"
		case 1:
			company = "AMD"
		case 2:
			company = "Nvidia"
		}
	}
	speed := rand.Intn(15) + 16
	duration := distance / speed / secondsPerDay
	price := 20.0 + speed
	if rand.Intn(2) == 1 {
		trip = "Round-trip"
		price = price * 2
	}else {
		trip="One-way"
	}
	fmt.Println(company,duration,trip,price)
}
func main(){
	ticket()
}
```
### 第6章 实验
```
package main

import (
	"fmt"
	"math/rand"
)

func piggy(x, y float64) float64 {
	//存入
	fmt.Println(x)
	//现有
	fmt.Println(y)
	//存后
	return x + y
}
func main() {
	money := make(map[int]float64, 0)
	money[0] = 0.05
	money[1] = 0.10
	money[2] = 0.25
	var kind int
	//var once float64
	var totol float64
	for {
		kind = rand.Intn(3)

		if totol >= 20.00 {
			break
		} else {
			totol = totol + piggy(money[kind], totol)
		}
	}
	fmt.Printf("final%v", totol)
}

```
### 数据范围
|类型|范围|内存占用|
|:----:|:----:|:----:|
|int8|-128~127|8bit|
|uint8|0~255|8bit|
|int16|-32768~32767|16bit|
|uint16|0~65535|16bit|
|int32|-2147 4836 48~2147 4836 47|32bit|
|uint32|0~4194 9672 95|32bit|
|int64|-9223 3720 3685 4775 808~9223 3720 3685 4775 807|64bit|
|uint64|0~1844 6744 0737 0955 1615|64bit|
```
var crl int8 = 127
log.Println(crl+1)
```
### 第7章 实验
```
package main

import (
	"fmt"
	"log"
	"math/rand"
)

func piggy(x, y float64) float64 {
	//存入
	fmt.Println(x)
	//现有
	fmt.Println(y)
	//存后
	return x + y
}
func main() {
	money := make(map[int]float64, 0)
	money[0] = 0.05
	money[1] = 0.10
	money[2] = 0.25
	var kind int
	//var once float64
	var totol float64
	var count uint=0
	for {
		kind = rand.Intn(3)

		if totol >= 20.00 {
			break
		} else {
			totol = totol + piggy(money[kind], totol)
			log.Printf("$%v",money[kind])
			count++
		}
	}
	fmt.Printf("final%v\n", totol)
	fmt.Println("count=",count)
}
```
### 第8章 实验
```
package main

import (
	"fmt"
)

const (
	distance = 236000000000000000
	PerKM    = 9.4607e12
)

func bigNum() {
	lightYear := distance / PerKM
	fmt.Println(lightYear)
}
func main() {
	bigNum()
}
```
### 第9章 实验
```
package main

import (
	"fmt"
)

func caesar() {
	message := "L fdph,L vdz,L frqtxhuhg."
	for _, v := range message {
		if v >= 'a' && v < 'z' {
			v -= 3
			if v < 'a' {
				v += 26
			}
		} else if v >= 'A' && v <= 'Z' {
			v -= 3
			if v < 'A' {
				v = +26
			}
		}
		fmt.Printf("%c", v)
	}
}
func international() {
	message := "uv vagreangvbany fcnpr fgngvba"

	for _, c := range message {

		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		} else if c >= 'A' && c <= 'Z' {
			c = c + 13
			if c > 'Z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
}
func main() {
	caesar()

	international()
}
```

### 第10章 实验
```
func input(s string)(bool)  {
	var launch bool
	switch s {
	case "ture","yes","1":
		launch = true
	case "false","no","0":
		launch =false
	default:
		fmt.Println("wrong")
	}
	return launch
}
```
### 第11章 实验
```
package main

import (
	"fmt"
	"strings"
)

func decipher() {
	cipherText := "CSOITEUIWUIZNSROCNKFD"
	keyword := "GOLANG"
	message := ""
	keyIndex := 0
	for i := 0; i < len(cipherText); i++ {
		c := cipherText[i] - 'A'
		k := keyword[keyIndex] - 'A'

		c = (c-k+26)%26 + 'A'
		message += string(c)
		keyIndex++
		keyIndex %= len(keyword)
	}
	fmt.Println(message)

}
func cipher() {
	message := "WEDIGYOULUVTHEGOPHERS"
	keyword := "GOLANG"
	keyIndex := 0
	cipherText := ""
	message = strings.ToUpper(strings.Replace(message, " ", "", -1))
	keyword = strings.ToUpper(strings.Replace(message, " ", "", -1))
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'A' && c <= 'Z' {
			c -= 'A'
			k := keyword[keyIndex] - 'A'
			c = (c+k)%26 + 'A'
			keyIndex++
			keyIndex %= len(keyword)
		}
		cipherText += string(c)
	}
	fmt.Println(cipherText)
}
func main() {
	decipher()
	cipher()
}
```
### 第12章 实验
```
func kelvin2Celsius(k float64) float64 {
	return k-273.15
}

func Celsius2kelvin(c float64)float64{
	return (c*9.0/5.0)+32.0
}
```
### 第13章 实验
```
type celsius float64
type fahrenheit float64
type kelvin float64
func (c celsius) fahrenheit() fahrenheit {
	return fahrenheit((c*9.0/5.0)+32.0)
}
func (c celsius)kelvin()kelvin{
	return kelvin(c+273.15)
}
func(f fahrenheit)celsius()celsius{
	return celsius((f-32.0)*5.0/9.0)
}
func (f fahrenheit) kelvin()kelvin  {
	return f.celsius().kelvin()
}
func(k kelvin)celsius()celsius{
	return celsius(k-273.15)
}
func(k kelvin)fahrenheit()fahrenheit{
	return k.celsius().fahrenheit()
}
```
### 第14章 实验

### 第15章 实验
### 第16章 实验
### 第17章 实验
```
type StringSlice []string

func terraform(s StringSlice)StringSlice{
	sort.Strings(s)

	return s
}

func main() {
	s:=[]string{"mars","uranus","neptune"}
	s = terraform(s)
	fmt.Println(s)
}```
### 第18章 实验
