package handler

import (
	"net"
	"os"
	"strconv"

	"znci.dev/lamp-v2/packets"
)

type PacketHandlerFunction func(packet packets.Packet, conn net.Conn)

type GeneralPacketHandler interface {
	HandleWithConnection(PacketHandlerFunction, packets.Packet, net.Conn)
}

var AvailableHandlers chan GeneralPacketHandler

type PacketHandler struct {
	id   int64
	done chan bool
}

func (p PacketHandler) HandleWithConnection(function PacketHandlerFunction, packet packets.Packet, conn net.Conn) {
	function(packet, conn)
}

func init() {
	maxCPH, err := strconv.Atoi(os.Getenv("MAX_CONCURRENT_PACKET_HANDLERS"))
	if err != nil {
		maxCPH = 50
	}

	AvailableHandlers = make(chan GeneralPacketHandler, maxCPH)
	for i := 0; i < maxCPH; i++ {
		AvailableHandlers <- PacketHandler{
			id:   int64(i),
			done: make(chan bool),
		}
	}
}
