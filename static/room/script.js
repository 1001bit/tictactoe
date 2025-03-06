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
    stop() {
        turnElem.innerText = "Waiting for another player";
    }
}
class Board {
    constructor() {
        for (let x = 0; x < 3; x++) {
            for (let y = 0; y < 3; y++) {
                let cell = document.getElementById("cell-" + y + "-" + x);
                cell.addEventListener("click", () => this.handleClick(x, y));
            }
        }
    }
    handleClick(x, y) {
        console.log(x, y);
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
        this.board = new Board();
        this.sign = " ";
    }
    handleMessage(msg) {
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
    handleStart(sign, turn) {
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
};
