/// <reference path="elems.ts" />

class Room {
    conn: RoomConn;
    topbar: TopBar;
    board: Board;

    sign: string;

    constructor(roomId: string) {
        this.conn = new RoomConn(roomId);
        this.conn.onmessage = (data: any) => {
            this.handleMessage(data);
        }

        this.topbar = new TopBar();
        this.topbar.setRoomId(roomId);

        this.board = new Board();

        this.sign = " ";
    }

    handleMessage(msg: any){
        if (!("type" in msg)) {
            return;
        }

        switch (msg.type) {
            case "start":
                this.handleStart(msg.you, msg.turn);
                break;
            case "stop":
                this.handleStop();
                break;
        }   
    }

    handleStart(sign: string, turn: string) {
        this.sign = sign;
        this.topbar.setTurn(turn == this.sign, turn);
    }

    handleStop() {
        this.topbar.stop();
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