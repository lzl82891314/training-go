package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 获取数据
	users := initMap()

	// 校验数据
	filteredUsers := filter(users, func(item map[string]string) bool {
		age, err := strconv.Atoi(item["age"])
		if err != nil {
			return false
		}
		return age >= 18 && age <= 35
	})

	// 筛选指定值
	ages := mapToStrings(filteredUsers, func(item map[string]string) string {
		return item["age"]
	})

	// 统计结果
	sumAge := fieldSum(ages, func(item string) int {
		age, err := strconv.Atoi(item)
		if err != nil {
			return 0
		}
		return age
	})

	fmt.Println(sumAge)
}

func initMap() []map[string]string {
	users := []map[string]string{
		{
			"name": "张三",
			"age":  "18",
		},
		{
			"name": "李四",
			"age":  "28",
		},
		{
			"name": "王五",
			"age":  "15",
		},
		{
			"name": "赵六",
			"age":  "61",
		},
		{
			"name": "雷七",
			"age":  "36",
		},
		{
			"name": "胡八",
			"age":  "35",
		},
	}
	return users
}

func filter(items []map[string]string, function func(item map[string]string) bool) []map[string]string {
	ans := make([]map[string]string, len(items))
	for _, item := range items {
		if function(item) {
			ans = append(ans, item)
		}
	}
	return ans
}

func mapToStrings(items []map[string]string, function func(item map[string]string) string) []string {
	result := make([]string, len(items))
	for _, item := range items {
		result = append(result, function(item))
	}
	return result
}

func fieldSum(items []string, function func(item string) int) int {
	var sum int
	for _, item := range items {
		sum += function(item)
	}
	return sum
}
