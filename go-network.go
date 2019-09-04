package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	ef "./extfunc"
)

func goNetwork(usrName, passwd string) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
	}
	v := url.Values{}
	v.Set("action", "login")

	v.Set("username", usrName)
	v.Set("password", passwd)
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

func main() {
	var printVer bool
	var cmdConfig ef.Config
	//var usrname, passwd string
	var configFile string
	flag.BoolVar(&printVer, "version", false, "print version")
	flag.StringVar(&cmdConfig.Username, "u", "", "username")
	flag.StringVar(&cmdConfig.Password, "p", "", "password")
	flag.StringVar(&configFile, "c", "config.json", "config file path")
	flag.Parse()

	if printVer {
		fmt.Println("v1.0.0")
		os.Exit(0)
	}
	if cmdConfig.Username == "" || cmdConfig.Password == "" {
		exists, err := ef.IsFileExists(configFile)
		if !exists || err != nil {
			fmt.Println("There is something wrong about login account!")
			os.Exit(0)
		}
		config, err := ef.ParseConfig(configFile)
		if err != nil {
			config = &cmdConfig
			if !os.IsNotExist(err) {
				fmt.Fprintf(os.Stderr, "error reading %s: %v\n", configFile, err)
				os.Exit(1)
			}

		} else {
			// fmt.Println(config.Username)
			// fmt.Println(config.Password)
			// ef.UpdateConfig(config, &cmdConfig)
			goNetwork(config.Username, config.Password)
		}
	} else {
		goNetwork(cmdConfig.Username, cmdConfig.Password)
	}
}
