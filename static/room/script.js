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
        this.socket.onmessage = (event) => {
            console.log(event.data);
        };
        this.socket.onclose = () => {
            console.log("Connection closed");
        };
        this.socket.onopen = () => {
            console.log("Connection opened");
            this.socket.send("Hello");
        };
    }
}
const roomIdElem = document.getElementById("room-id");
const turnElem = document.getElementById("turn");
class Room {
    constructor(roomId) {
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
};
