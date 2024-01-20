package protocol

func VarInt(bytes []byte) (int64, int) {
	var value int64 = 0
	var bitOffset byte = 0
	var currIndx = 0
	var currentByte byte

	for {
		if bitOffset == 35 {
			return 0, 0
		}

		currentByte = bytes[currIndx]
		value |= (int64(currentByte&0x7F) << uint(bitOffset))

		bitOffset += 7
		currIndx++

		if (currentByte & 0x80) != 0x80 {
			break
		}
	}

	return int64(value), currIndx
}

func VarText(bytes []byte) (string, int) {
	length, lengthBytes := VarInt(bytes)
	if lengthBytes+int(length) > len(bytes) {
		return "", 0
	}
	return string(bytes[lengthBytes : lengthBytes+int(length)]), lengthBytes + int(length)
}

func writeVarInt(value int) []byte {
	data := make([]byte, 0)

	for {
		temp := byte(value & 0b01111111)
		value = value >> 7

		if value != 0 {
			temp |= 0b10000000
		}

		data = append(data, temp)

		if value == 0 {
			break
		}
	}

	return data
}

func writeString(value string) []byte {
	data := make([]byte, 0)

	length := len(value)
	data = append(data, writeVarInt(length)...)
	data = append(data, []byte(value)...)

	return data
}
