package main

import "fmt"

func main() {

	var countryCapitalMap map[string]string

	/*创建集合*/
	countryCapitalMap = make(map[string]string)

	countryCapitalMap["France"] = "Paris"
	countryCapitalMap["Italy"] = "Rome"

	/*使用Key 输出 map 值*/
	for country := range countryCapitalMap{
		fmt.Println("Capital of",country,"is",countryCapitalMap[country])
	}

	/*查看元素是否在集合中存在*/
	capital ,ok := countryCapitalMap["United States"]

	if ok {
		fmt.Println("Capital of United Stateds is",capital)
	}else{
		fmt.Println("Capital of United Stateds is not present")
	}
}
