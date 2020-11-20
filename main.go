package main

import (
	"fmt"
	"github.com/guancbo/test-demo/api"
)

func main() {
	api.SetMachineID(111)

	fmt.Println(api.GetSnowflakeID() >> 22) //右移22位，就可以得到时间戳(毫秒级)
}
