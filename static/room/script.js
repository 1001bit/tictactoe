"use strict";
class Room {
    constructor() {
        const id = new URLSearchParams(window.location.search).get("id");
        this.socket = new WebSocket("ws://localhost/api/game/roomWS/" + id);
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
