package main

var pub = publisher{
	broadcast:  make(chan message),
	register:   make(chan subscriber),
	deregister: make(chan subscriber),
	rooms:      make(map[string]map[*connection]bool),
}

func (pub *publisher) publish() {
	for {
		select {
		case sub := <-pub.register:
			conns := pub.rooms[sub.roomId]
			if conns == nil {
				conns = make(map[*connection]bool)
				pub.rooms[sub.roomId] = conns
			}
			pub.rooms[sub.roomId][sub.conn] = true
		case sub := <-pub.deregister:
			conns := pub.rooms[sub.roomId]
			if conns != nil {
				if _, ok := conns[sub.conn]; ok {
					delete(conns, sub.conn)
					close(sub.conn.sendMessage)
					if len(conns) == 0 {
						delete(pub.rooms, sub.roomId)
					}
				}
			}
		case msg := <-pub.broadcast:
			conns := pub.rooms[msg.roomId]
			for conn := range conns {
				select {
				case conn.sendMessage <- msg.text:
				default:
					close(conn.sendMessage)
					delete(conns, conn)
					if len(conns) == 0 {
						delete(pub.rooms, msg.roomId)
					}
				}
			}
		}
	}
}
