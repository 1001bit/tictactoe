"use strict";
const roomIdElem = document.getElementById("room-id");
const roomId = new URLSearchParams(window.location.search).get("id");
roomIdElem.innerText = `Room "${roomId || "Unknown"}"`;
document.title = `Room "${roomId || "Unknown"}"`;
class Room {
    constructor() {
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
const room = new Room();
