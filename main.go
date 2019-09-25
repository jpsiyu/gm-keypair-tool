package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println("欢迎使用GM密钥对生成工具!")
	//判断钱包是否已经初始化
	err, res := GetWalletStatus()
	if err != nil {
		fmt.Println(err)
		return
	}
	walletStatus := res.Result.(map[string]interface{})
	if walletStatus["isHasSeed"] == false {
		err = CreateSeed()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	if walletStatus["isWalletLock"] == true {
		err = UnlockWallet()
		if err != nil {
			fmt.Println(err)
			return
		}
	}

	fmt.Print("是否生成新的密钥对？(y/n): ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimRight(input, "\n")
	if input != "y" {
		fmt.Println("Good bye!")
		return
	}
	err = CreateAccount()
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("程序运行结束")
	reader.ReadString('\n')
}
