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
    setResult(result, sign) {
        if (result == "D") {
            turnElem.innerText = "Draw";
        }
        else if (result == "W") {
            turnElem.innerText = `You win (${sign})`;
        }
        else {
            turnElem.innerText = `You lose (${sign})`;
        }
    }
}
class Board {
    constructor() {
        this.allowPlace = false;
        this.sign = " ";
        this.placecallback = (_x, _y) => { };
        for (let x = 0; x < 3; x++) {
            for (let y = 0; y < 3; y++) {
                let cell = document.getElementById("cell-" + y + "-" + x);
                cell.addEventListener("click", () => this.handleClick(x, y));
            }
        }
    }
    setCell(x, y, sign) {
        let cell = document.getElementById("cell-" + y + "-" + x);
        cell.innerText = sign;
    }
    getCell(x, y) {
        let cell = document.getElementById("cell-" + y + "-" + x);
        return cell.innerText;
    }
    handleClick(x, y) {
        if (!this.allowPlace || this.getCell(x, y) != "") {
            return;
        }
        this.setCell(x, y, this.sign);
        this.allowPlace = false;
        this.placecallback(x, y);
    }
    handleOpponentMove(x, y, sign) {
        this.setCell(x, y, sign);
        this.allowPlace = true;
    }
    setAllowPlace(allow) {
        this.allowPlace = allow;
    }
    setSign(sign) {
        this.sign = sign;
    }
    clear() {
        this.sign = " ";
        this.allowPlace = false;
        for (let x = 0; x < 3; x++) {
            for (let y = 0; y < 3; y++) {
                let cell = document.getElementById("cell-" + y + "-" + x);
                cell.innerText = " ";
            }
        }
    }
}
class RoomConn {
    constructor(roomId) {
        const protocol = window.location.protocol === "https:" ? "wss" : "ws";
        const host = window.location.host;
        this.socket = new WebSocket(`${protocol}://${host}/api/game/roomWS/${roomId}`);
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
    sendPlaceMessage(x, y) {
        this.socket.send(JSON.stringify({
            type: "place",
            x: x,
            y: y,
        }));
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
        this.board.placecallback = (x, y) => {
            this.handleBoardPlace(x, y);
        };
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
            case "move":
                this.handleMove(msg.x, msg.y, msg.sign);
                if (msg.result != " ") {
                    this.handleEnd(msg.result);
                }
                break;
        }
    }
    handleStart(sign, turn) {
        this.board.clear();
        this.board.setSign(sign);
        this.board.setAllowPlace(turn == sign);
        this.topbar.setTurn(turn == sign, turn);
    }
    handleStop() {
        this.topbar.stop();
        this.board.clear();
    }
    handleMove(x, y, sign) {
        if (sign != this.board.sign) {
            this.board.handleOpponentMove(x, y, sign);
        }
        this.topbar.setTurn(sign != this.board.sign, sign == "O" ? "X" : "O");
    }
    handleEnd(result) {
        if (result == "D") {
            this.topbar.setResult("D", this.board.sign);
        }
        else {
            this.topbar.setResult(result == this.board.sign ? "W" : "L", this.board.sign);
        }
        this.board.setAllowPlace(false);
    }
    handleBoardPlace(x, y) {
        this.conn.sendPlaceMessage(x, y);
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
