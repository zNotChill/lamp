package packets

type PacketID uint8

type Packet struct {
	Length     int64
	ID         int64
	Data       []byte
	FromClient bool
}

func New(length int64, id int64, data []byte, fromClient bool) *Packet {
	return &Packet{
		Length:     length,
		ID:         id,
		Data:       data,
		FromClient: fromClient,
	}
}
