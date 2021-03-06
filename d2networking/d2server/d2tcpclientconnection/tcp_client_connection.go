package d2tcpclientconnection

import (
	"encoding/json"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2client/d2clientconnectiontype"
	"github.com/OpenDiablo2/OpenDiablo2/d2networking/d2netpacket"
	"net"

	"github.com/OpenDiablo2/OpenDiablo2/d2game/d2player"
)

type TCPClientConnection struct {
	id            string
	tcpConnection net.Conn
	playerState   *d2player.PlayerState
}

func CreateTCPClientConnection(tcpConnection net.Conn, id string) *TCPClientConnection {
	return &TCPClientConnection{
		tcpConnection: tcpConnection,
		id:            id,
	}
}

func (t TCPClientConnection) GetUniqueID() string {
	return t.id
}

func (t *TCPClientConnection) SendPacketToClient(p d2netpacket.NetPacket) error {
	packet, err := json.Marshal(p)
	if err != nil {
		return err
	}

	_, err = t.tcpConnection.Write(packet)
	if err != nil {
		return err
	}

	return nil
}

func (t *TCPClientConnection) SetPlayerState(playerState *d2player.PlayerState) {
	t.playerState = playerState
}

func (t *TCPClientConnection) GetPlayerState() *d2player.PlayerState {
	return t.playerState
}

// GetConnectionType returns an enum representing the connection type.
// See: d2clientconnectiontype.
func (t TCPClientConnection) GetConnectionType() d2clientconnectiontype.ClientConnectionType {
	return d2clientconnectiontype.LANClient
}
