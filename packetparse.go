package gmodbus

import ()

const (
	// Bit access
	FuncCodeReadDiscreteInputs = 2
	FuncCodeReadCoils          = 1
	FuncCodeWriteSingleCoil    = 5
	FuncCodeWriteMultipleCoils = 15

	// 16-bit access
	FuncCodeReadInputRegisters         = 4
	FuncCodeReadHoldingRegisters       = 3
	FuncCodeWriteSingleRegister        = 6
	FuncCodeWriteMultipleRegisters     = 16
	FuncCodeReadWriteMultipleRegisters = 23
	FuncCodeMaskWriteRegister          = 22
	FuncCodeReadFIFOQueue              = 24
)

const (
	ExceptionCodeIllegalFunction                    = 1
	ExceptionCodeIllegalDataAddress                 = 2
	ExceptionCodeIllegalDataValue                   = 3
	ExceptionCodeServerDeviceFailure                = 4
	ExceptionCodeAcknowledge                        = 5
	ExceptionCodeServerDeviceBusy                   = 6
	ExceptionCodeMemoryParityError                  = 8
	ExceptionCodeGatewayPathUnavailable             = 10
	ExceptionCodeGatewayTargetDeviceFailedToRespond = 11
)

//组包
type PacketEnCode struct {
}

//
func uint16ToByte(in uint16) []byte {
	var out []byte
	var h, l uint8 = uint8(in >> 8), uint8(in & 0xff)
	out = append(out, h)
	out = append(out, l)
	return out
}

//打开单个线圈
//0F 05 00 00 FF 00 8D 14
//0F(从机地址) 05（功能码） 00 00（线圈地址） FF 00（开启，关闭为00 00） 8D 14（crc）
func (self *PacketEnCode) WriteSingleCoil(slaveAddr byte, addr uint16, ifopen bool) (error, []byte) {
	var packet []byte
	packet = append(packet, slaveAddr)
	packet = append(packet, byte(FuncCodeWriteSingleCoil))
	packet = append(packet, uint16ToByte(addr)...)
	//如果是开启设备
	if ifopen {
		packet = append(packet, []byte{0xff, 0x00}...)
	} else {
		packet = append(packet, []byte{0x00, 0x00}...)
	}
	//加上crc
	var crc crc
	crc.reset()
	crc.pushBytes(packet)
	packet = append(packet, crc.bytevalue()...)
	return nil, packet
}

//打开多个线圈
//0F 0f 00 00 00 03 01 ff
//0F(从机地址) 0f（功能码） 00 00（起始地址） 00 03（写线圈个数） 01（写字节个数） ff（要写的字节）
func (self *PacketEnCode) WriteMultipleCoil(slaveAddr byte, startaddr uint16, coilNum uint16, value []byte) (error, []byte) {
	var packet []byte
	packet = append(packet, slaveAddr)
	packet = append(packet, byte(FuncCodeWriteMultipleCoils))
	packet = append(packet, uint16ToByte(startaddr)...)
	packet = append(packet, uint16ToByte(coilNum)...)
	packet = append(packet, byte(len(value)))
	packet = append(packet, value...)
	//加上crc
	var crc crc
	crc.reset()
	crc.pushBytes(packet)
	packet = append(packet, crc.bytevalue()...)
	return nil, packet
}
