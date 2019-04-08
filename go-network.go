package main

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	v := url.Values{}
	v.Set("action", "login")

	v.Set("username", "")
	v.Set("password", "")
	v.Set("ac_id", "1")
	v.Set("user_ip", "")
	v.Set("nas_ip", "")
	v.Set("user_mac", "")
	v.Set("save_me", "1")
	v.Set("ajax", "1")
	//利用指定的method,url以及可选的body返回一个新的请求.如果body参数实现了io.Closer接口，Request返回值的Body 字段会被设置为body，并会被Client类型的Do、Post和PostFOrm方法以及Transport.RoundTrip方法关闭。
	//body := ioutil.NopCloser(strings.NewReader(v.Encode())) //把form数据编下码
	body := strings.NewReader(v.Encode())
	client := &http.Client{
		Transport: tr,
	} //客户端,被Get,Head以及Post使用
	reqest, err := http.NewRequest("POST", "https://10.108.255.249/include/auth_action.php", body)
	if err != nil {
		fmt.Println("Fatal error 1", err.Error())
	}
	//给一个key设定为响应的value.
	reqest.Header.Set("Content-Type", "application/x-www-form-urlencoded") //必须设定该参数,POST参数才能正常提交

	resp, err := client.Do(reqest) //发送请求
	if err != nil {
		fmt.Println("Fatal error 2", err.Error())
	}
	defer resp.Body.Close() //一定要关闭resp.Body
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Fatal error 3", err.Error())
	}

	fmt.Println(string(content))
}
