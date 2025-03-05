/// <reference path="elems.ts" />

class Room {
    conn: RoomConn;
    topbar: TopBar;

    constructor(roomId: string) {
        this.conn = new RoomConn(roomId);
        this.conn.onmessage = (data: any) => {
            this.handleMessage(data);
        }

        this.topbar = new TopBar();
        this.topbar.setRoomId(roomId);
    }

    handleMessage(msg: any){
        console.log(msg);
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