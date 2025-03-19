// @FileName : shop.go
// @Time : 2025/3/18 下午7:11
// @Author : luobozi

package config

import (
	"fmt"
)

var login_user string
var userInfo map[string]interface{}
var shop_list map[string]interface{}

func Show_shop() {
	if login_user != "" {
		// 遍历 shop_list
		for k, v := range shop_list {
			// v 是 interface{}，需要类型断言为 map[string]interface{}
			if item, ok := v.(map[string]interface{}); ok {
				id, _ := item["id"].(string)
				name, _ := item["shop_name"].(string)
				price, _ := item["price"].(float64)
				fmt.Printf("位置：%s  商品编号：%s  商品名称：%s  商品价格：%.2f\n==================================\n",
					k, id, name, price)
			} else {
				fmt.Println("数据格式错误，无法解析商品信息:", v)
			}
		}
	} else {
		fmt.Println("请先登录!")
	}
}

// 购物车逻辑
var ShopCat []map[string]interface{}
var amount float64

func Shop_cat(select1 string) {
	if login_user != "" { // 确保用户已登录
		amount = 0.0
		for v, k := range shop_list {
			if v == select1 {
				ShopCat = append(ShopCat, k.(map[string]interface{})) // 追加商品到购物车
				fmt.Println("购物车内容: \n==================================")
				for ind, shop_map := range ShopCat {
					amount += shop_map["price"].(float64)
					fmt.Printf("编号：%d 商品名称：%s 价格：%.2f\n==================================\n", ind+1, shop_map["shop_name"], shop_map["price"])
				}
			}
		}
		fmt.Printf("目前订单总价格：%2.f\n==================\n", amount)
	} else {
		fmt.Println("请先登录!")
	}
}

func Overorder() bool {
	if login_user != "" && ShopCat != nil {
		fmt.Println("目前余额为：", userInfo["money"].(float64)-amount)
		ShopCat = nil
		return true
	} else {
		fmt.Println("请先登录，或者先添加购物车！")
		return false
	}
	return false
}
