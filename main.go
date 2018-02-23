package main

import (
	"fmt"

	"github.com/PuerkitoBio/goquery"
)

func main() {
	// 対応するエリア番号
	// 2 : 北海道
	// 3 : 東北
	// 4 : 関東
	// 5 : 中部
	// 6 : 近畿
	// 8 : 中国
	// 9 : 四国
	// 7 : 九州
	var area = "4"

	doc, err := goquery.NewDocument("https://transit.yahoo.co.jp/traininfo/area/" + area + "/")
	if err != nil {
		fmt.Print("scrapping faile!!")
	}
	doc.Find("tr").Each(func(_ int, s *goquery.Selection) {
		_, exists := s.Find("span").First().Attr("class")
		if exists == true {
			fmt.Println("------------------")
			fmt.Println(s.Text())
		}
	})
}
