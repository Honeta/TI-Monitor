package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gocolly/colly"
)

var target_list []string

func find(target string) {
	flag := 0
	for flag == 0 {
		c := colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.61 Safari/537.36"),
			colly.MaxDepth(0),
		)
		c.OnResponse(func(r *colly.Response) {
			flag = 1
			length := len(r.Body) - 1
			pos := 0
			for pos = length; r.Body[pos] != ':'; pos-- {
			}
			value_str := r.Body[pos+1 : length]
			value, _ := strconv.Atoi(string(value_str))
			if value > 0 {
				fmt.Printf("%s %d\n", time.Now().String()[:19]+" "+target, value)
			} else {
				fmt.Printf("%s Not Found\n", time.Now().String()[:19]+" "+target)
			}

		})
		c.OnError(func(r *colly.Response, err error) {
		})
		c.Visit("https://www.ti.com.cn/storeservices/cart/opninventory?opn=" + target)
	}
}

func main() {
	file, err := os.Open("list_example.txt")
	if err != nil {
		panic(err)
	}
	scanner := bufio.NewScanner(file)
	for {
		success := scanner.Scan()
		if !success {
			err = scanner.Err()
			if err == nil {
				break
			} else {
				panic(err)
			}
		}
		target_list = append(target_list, scanner.Text())
	}
	for {
		for i := 0; i < len(target_list); i++ {
			target := target_list[i]
			find(target)
			time.Sleep(time.Duration(200) * time.Millisecond)
		}
	}
}
