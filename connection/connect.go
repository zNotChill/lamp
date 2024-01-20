package connection

import (
	"bufio"
	"io/ioutil"
	"net"

	"znci.dev/lamp-v2/packets"
	"znci.dev/lamp-v2/packets/handler"
	"znci.dev/lamp-v2/protocol"
	"znci.dev/lamp-v2/utils"
)

type ProtocolClient struct {
	protocol.ProtocolClient
}

func HandleConnection(conn net.Conn, c ProtocolClient) {
	defer func() {
		closeConnection(conn)
		recover()
	}()
	utils.Info("Accepted connection from " + conn.RemoteAddr().String())

	for {
		bytes := ReadBytes(conn)

		if bytes == nil || bytes[0] == 0xFE { // 0xFE is the ping packet
			continue
		}

		packetLength, readIndex := protocol.VarInt(bytes)
		packetID, _ := protocol.VarInt(bytes[readIndex:])
		packetData := bytes[(len(bytes) - int(packetLength) - 1):]

		packet := packets.New(packetLength, packetID, packetData, true)
		handlers := <-handler.AvailableHandlers

		go func() {
			funcForPacket := handler.GetFunctionFromPacket(*packet)
			if packet.Length != 0 {
				handlers.HandleWithConnection(funcForPacket, *packet, conn)
			}
			handler.AvailableHandlers <- handlers
		}()
	}
}

func closeConnection(conn net.Conn) {
	conn.Close()
	utils.Info("Closed connection from " + conn.RemoteAddr().String())
}

func ReadBytes(conn net.Conn) []byte {
	reader := bufio.NewReader(conn)
	bytes, err := ioutil.ReadAll(reader)
	if err != nil {
		utils.Error("Failed to read bytes from connection")
		return nil
	}
	return bytes
}
