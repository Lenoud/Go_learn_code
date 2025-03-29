// @FileName : main.go
// @Time : 2025/3/23 下午12:39
// @Author : luobozi

package main

import (
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

// 此处开始编写代码

func main() {
	url_slis := []string{
		"https://p2.piqsels.com/preview/194/117/286/skyscraper-bridge-skyline-traffic-thumbnail.jpg",
		"https://p0.piqsels.com/preview/631/601/318/skyscraper-view-at-daytime-thumbnail.jpg",
	}
	n := 0
	for n < 5 {
		go download_img(url_slis[0], "go1_main_process_")
		go download_img(url_slis[1], "go2_main_process_")
		n++
	}
	time.Sleep(12 * time.Second)
}

func download_img(url, filename string) {
	client := http.Client{}
	r, err := http.NewRequest("GET", url, nil)
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println("图片下载错误！", err)
		return
	}
	//生成随机的图片名称
	url_name_sil := strings.Split(url, "/")
	url_name := url_name_sil[len(url_name_sil)-1]
	randmon_name := filename + strconv.Itoa(rand.Int()) + url_name
	file, _ := os.Create(randmon_name)

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("保存图片失败！", err)
		return
	}
	fmt.Println(randmon_name+".jpg"+"下载成功！", time.Now())

}
