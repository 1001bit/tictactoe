/// <reference path="elems.ts" />

class Room {
    conn: RoomConn;
    topbar: TopBar;

    constructor(roomId: string) {
        this.conn = new RoomConn(roomId);
        this.topbar = new TopBar();
        this.topbar.setRoomId(roomId);
    }
}

window.onload = () => {
    const roomId = new URLSearchParams(window.location.search).get("id");
    if (!roomId) {
        window.location.href = "/";
        return;
        
    }
    new Room(roomId);
}