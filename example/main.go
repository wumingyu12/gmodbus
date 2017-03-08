package main

import (
	"fmt"
	"github.com/wumingyu12/gmodbus"
)

func main() {
	pack := &gmodbus.PacketEnCode{}
	_, b := pack.WriteSingleCoil(0x0f, 1, true)
	fmt.Println(b)

	// if 0x1241 != crc.value() {
	// 	t.Fatalf("crc expected %v, actual %v", 0x1241, crc.value())
	// }
}
