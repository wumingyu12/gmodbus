package main

import (
	"fmt"
	"github.com/wumingyu12/gmodbus"
)

func main() {
	pack := &gmodbus.PacketEnCode{}
	_, b := pack.WriteSingleCoil(0x0f, 1, true)
	printHex(b)

	//写多个线圈
	_, b = pack.WriteMultipleCoil(0x0f, 0, 5, []byte{0xff})
	printHex(b)
	// if 0x1241 != crc.value() {
	// 	t.Fatalf("crc expected %v, actual %v", 0x1241, crc.value())
	// }
}

func printHex(b []byte) {
	for _, v := range b {
		fmt.Printf("%02x ", v)
	}
	fmt.Printf("\r\n")
}
