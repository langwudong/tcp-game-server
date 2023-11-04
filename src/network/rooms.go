package network

import (
	"net"
	"sync"
)

type RoomList struct {
	Rooms []*Room `json:"rooms"`
}

func GetRooms() *RoomList {
	//使用单例模式创建rooms,注意需要初始化rooms.Rooms
	once.Do(func() {
		instance = &RoomList{}
		instance.Rooms = make([]*Room, 0)
	})
	return instance
}

func (rooms *RoomList) Contain(roomId string) (*Room, bool) {
	if rooms != nil {
		for _, room := range rooms.Rooms {
			if room.ID == roomId {
				return room, true
			}
		}
	}
	return nil, false
}

func (rooms *RoomList) QueryRoom(player *Player) *Room {
	if rooms != nil {
		for _, room := range rooms.Rooms {
			if room.ID == player.RoomID {
				return room
			}
		}
	}
	return nil
}

func (rooms *RoomList) RemoveRoom(room *Room) {
	for index, value := range rooms.Rooms {
		if value == room {
			rooms.Rooms = append((rooms.Rooms)[:index], (rooms.Rooms)[index+1:]...)
		}
	}
}

type Room struct {
	ID      string    `json:"room_id,omitempty"`
	Players []*Player `json:"players"`
}

type Player struct {
	Username  string   `json:"username,omitempty"`
	RoomID    string   `json:"room_id,omitempty"`
	Conn      net.Conn `json:"conn,omitempty"`
	GameState string   `json:"game_state,omitempty"`
	Index     int      `json:"index,omitempty"`
}

var (
	instance *RoomList
	once     sync.Once
	players  = make(map[net.Conn]*Player)
)
