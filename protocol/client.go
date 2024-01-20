package protocol

import (
	"bytes"
	"compress/zlib"
	"crypto/cipher"
	"fmt"
	"net"
	"time"

	"github.com/google/uuid"
	"znci.dev/lamp-v2/interfaces"
	"znci.dev/lamp-v2/packets"
	"znci.dev/lamp-v2/utils"
)

type ProtocolClient struct {
	server      interfaces.Server `type:"Interface"`
	conn        net.Conn          `type:"Conn"`
	Ended       bool              `type:"Boolean"`
	state       string            `type:"String"`
	compression bool              `type:"Boolean"`
	encryption  bool              `type:"Boolean"`

	Username string    `type:"String"`
	UUID     uuid.UUID `type:"UUID"`
	Skin     string    `type:"String"`

	ProtocolVersion int            `type:"Int"`
	decrypter       *cipher.Stream `type:"CipherStream"`
	encrypter       *cipher.Stream `type:"CipherStream"`

	listeners  map[string]func(interface{}) `type:"Listeners"`
	endHandler func()                       `type:"EndHandler"`
}

func (client *ProtocolClient) WritePacket(packetId packets.PacketID, data []byte) error {
	packetIdVarInt := writeVarInt(int(packetId))
	packetLengthVarInt := writeVarInt(len(packetIdVarInt) + len(data))

	packetData := append(packetIdVarInt, data...)

	if client.compression {
		if len(packetData) > client.server.GetConfig().CompressionThreshold {
			var in bytes.Buffer
			w := zlib.NewWriter(&in)
			w.Write(packetData)
			w.Close()
			data = append(packetLengthVarInt, in.Bytes()...)
			data = append(writeVarInt(len(data)), data...)
		} else {
			data = append(append(writeVarInt(1+len(packetData)), 0), packetData...)
		}
	} else {
		data = append(packetLengthVarInt, packetData...)
	}

	client.conn.SetWriteDeadline(time.Now().Add(time.Second * 5))
	err := client.writeBytes(data)
	if err != nil {
		utils.Error("Failed to write packet to connection")
		return err
	}
	return nil
}

func (client *ProtocolClient) writeBytes(data []byte) error {
	if !client.encryption && client.encrypter != nil {
		(*client.encrypter).XORKeyStream(data, data)
	}

	fmt.Println(data)
	_, err := client.conn.Write(data)
	return err
}
