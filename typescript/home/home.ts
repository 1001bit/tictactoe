interface RoomMsg {
    id: string;
    players: number;
}

class RoomsLoader {
    es: EventSource;

    constructor() {
        this.es = new EventSource("/api/game/roomsSSE");
        this.es.onmessage = (event) => this.handleMsg(event);
    }

    handleMsg(event: MessageEvent) {
        console.log(event.data);
        const data = JSON.parse(event.data);
        if ("rooms" in data) {
            this.renderRooms(data.rooms);
        }
    }

    renderRooms(rooms: RoomMsg[]) {
        roomsContainer.innerHTML = "";

        for (const roomMsg of rooms) {
            const roomElem = roomSample.cloneNode(true) as HTMLDivElement;
            roomElem.classList.remove("sample");
            const roomName = roomElem.getElementsByClassName("room-name")[0] as HTMLDivElement;
            const roomPlayers = roomElem.getElementsByClassName("room-players")[0] as HTMLDivElement;
            const roomJoin = roomElem.getElementsByClassName("room-join")[0] as HTMLAnchorElement;
            roomName.innerText = roomMsg.id;
            roomPlayers.innerText = roomMsg.players + "/2 players";
            roomJoin.href = "/room?id=" + roomMsg.id;
            roomsContainer.appendChild(roomElem);
        }
    }
}

class RoomCreator {
    constructor(){
        createRoomLink.onclick = () => {
            window.location.href = "/room?id=" + Math.random().toString(36).slice(2);
        }
    }
}

new RoomCreator();
new RoomsLoader();