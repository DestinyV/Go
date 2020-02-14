package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

// 公用一个client适用于连接比较频繁
var (
	client = &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true, // 禁用keepalive
		},
	}
)

func main() {
	// resp, err := http.Get("http://127.0.0.1:8085/v1/f2/?name=lili&age=5")

	// if err != nil {
	// 	fmt.Println(err)
	// 	return
	// }
	// bin, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	fmt.Println("err")
	// 	return
	// }
	// fmt.Println(string(bin))

	urlObj, _ := url.Parse("http://127.0.0.1:8085/v1/f2/")
	prm := url.Values{}
	prm.Set("name", "limald")
	prm.Set("age", "9600")
	prmStr := prm.Encode()
	fmt.Println(prmStr)
	urlObj.RawQuery = prmStr
	req, err := http.NewRequest("GET", urlObj.String(), nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	// resp, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	// 适用于连接不是特别频繁用完就关闭连接
	client := &http.Client{
		Transport: &http.Transport{
			DisableKeepAlives: true, // 禁用keepalive
		},
	}
	resp, err := client.Do(req)
	defer resp.Body.Close()
	bin, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("err")
		return
	}
	fmt.Println(string(bin))

}
