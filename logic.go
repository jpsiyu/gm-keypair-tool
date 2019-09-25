package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
)

func CreateSeed() error {
	fmt.Println("钱包正在初始化...")
	err, res := GenSeed()
	if err != nil {
		return err
	}
	result := res.Result.(map[string]interface{})
	seed := result["seed"].(string)
	fmt.Printf("生成种子：%s\n", seed)
	fmt.Print("初始化密码(8位以上，包含字母数字)：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")
	err, res = SaveSeed(seed, input)
	if err != nil {
		return err
	}
	seedResult := res.Result.(map[string]interface{})
	if seedResult["isOK"] != true {
		return errors.New(seedResult["msg"].(string))
	}
	fmt.Println("钱包初始化完成!")
	return nil
}

func UnlockWallet() error {
	fmt.Print("解锁钱包，请输入密码：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")
	err, res := UnLock(input)
	if err != nil {
		return err
	}
	unlockResult := res.Result.(map[string]interface{})
	if unlockResult["isOK"] != true {
		return errors.New(unlockResult["msg"].(string))
	}
	return nil
}

func CreateAccount() error {
	fmt.Print("请给账号一个名字：")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")
	err, res := NewAccount(input)
	if err != nil {
		return err
	}

	if res.Error != "" {
		return errors.New(res.Error)
	}
	createResult := res.Result.(map[string]interface{})
	acc := createResult["acc"].(map[string]interface{})
	addr := acc["addr"].(string)

	err, res = DumpPrivkey(addr)
	dumpResult := res.Result.(map[string]interface{})
	privKey := dumpResult["data"]

	fmt.Printf("名字: %s\n", input)
	fmt.Printf("地址: %s\n", addr)
	fmt.Printf("私钥: %s\n", privKey)
	return nil
}
