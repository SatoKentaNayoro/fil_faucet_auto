package main

import (
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func main() {
	// 填写要发送代币的地址
	destAddr := []string{"t3us7kvzxx5usee7vguzcb27rds7e7jo2o224i5p7tgw4o2mzpzeht2nnvj4ghoji3tvzmloy3lxh3ffbky3aa", "t16iz55rw4o4y7fxen3ek36nsj5wqn6vwohkrve7a", "t1j5zdhmbzf6mlvvgpgjr5y5nka2plg2s6pgezerq"}

	// 设置定时器，每隔一段时间发送一次代币
	ticker := time.NewTicker(5 * time.Minute)
	for range ticker.C {

		err := sendFunds(destAddr)
		if err != nil {
			log.Printf("Failed to send funds to %s: %v\n", destAddr, err)
			continue
		}
		log.Printf("Sent funds to %s\n", destAddr)
	}
}

func sendFunds(destAddrs []string) error {
	for _, destAddr := range destAddrs {
		// 构建要发送的数据
		data := url.Values{}
		data.Set("address", destAddr)

		// 构建请求
		req, err := http.NewRequest("POST", "https://faucet.calibration.fildev.network/send", strings.NewReader(data.Encode()))
		if err != nil {
			return err
		}
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
		req.Header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Safari/537.36")

		// 发送请求
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		// 检查响应状态码
		if resp.StatusCode != http.StatusOK {
			log.Printf("unexpected status code: %d, address: %s", resp.StatusCode, destAddr)
		}
	}
	return nil
}
