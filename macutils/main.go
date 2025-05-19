package main

import (
	"fmt"
	"macutils/macparser"
)

func main() {
	mac := " b022.7aea.bb6-d   "
	info, err := macparser.DetectAndNormalize(mac)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Tipo:", info.MACType)
	fmt.Println("Formato Linux: \n", info.LinuxFormat)
	fmt.Println("Formato Huawei: \n", info.HuaweiFormat)
	fmt.Println("Formato Cisco: \n", info.CiscoFormat)
}
