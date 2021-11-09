package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func main() {
	var api = "http://t.weather.sojson.com/api/weather/city/"

	var cityId int
	fmt.Print("请输入城市ID：")
	fmt.Scanln(&cityId)
	fmt.Println(cityId)

	res, err := http.Get(api + strconv.Itoa(cityId))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.Body, res.StatusCode)

}