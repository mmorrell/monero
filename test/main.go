package main

import (
	"fmt"
	"monero"
)

func main() {
	daemon := monero.NewDaemonClient("http://127.0.0.1:18081/json_rpc")
	blockCount, err := daemon.GetBlockCount()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Count:", blockCount)
}