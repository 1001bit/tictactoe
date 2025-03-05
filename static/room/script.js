"use strict";
class TopBar {
    constructor() {
    }
    setRoomId(roomId) {
        roomIdElem.innerText = `Room "${roomId}"`;
        document.title = `Room "${roomId}"`;
    }
    setTurn(yours, sign) {
        turnElem.innerText = `Turn: ${yours ? "You" : "Opponent"} (${sign})`;
    }
}
class RoomConn {
    constructor(roomId) {
        this.socket = new WebSocket("ws://localhost/api/game/roomWS/" + roomId);
        this.onmessage = (_data) => { };
        this.socket.onmessage = (event) => {
            this.onmessage(JSON.parse(event.data));
        };
        this.socket.onclose = () => {
            console.log("Connection closed");
        };
        this.socket.onopen = () => {
            console.log("Connection opened");
        };
    }
}
const roomIdElem = document.getElementById("room-id");
const turnElem = document.getElementById("turn");
class Room {
    constructor(roomId) {
        this.conn = new RoomConn(roomId);
        this.conn.onmessage = (data) => {
            this.handleMessage(data);
        };
        this.topbar = new TopBar();
        this.topbar.setRoomId(roomId);
    }
    handleMessage(msg) {
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
};
