"use strict";
const createRoomLink = document.getElementById("create-room-link");
class RoomsLoader {
    constructor() {
        this.es = new EventSource("/api/game/roomsSSE");
        this.es.onmessage = (event) => this.handleMsg(event);
    }
    handleMsg(event) {
        console.log(event.data);
    }
}
class RoomCreator {
    constructor() {
        createRoomLink.onclick = () => {
            window.location.href = "/room?id=" + Math.random().toString(36).slice(2);
        };
    }
}
new RoomCreator();
new RoomsLoader();
