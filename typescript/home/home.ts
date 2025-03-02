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

    renderRooms(rooms: string[]) {
        roomsContainer.innerHTML = "";

        for (const roomId of rooms) {
            const roomElem = roomSample.cloneNode(true) as HTMLDivElement;
            roomElem.classList.remove("sample");
            const roomName = roomElem.getElementsByClassName("room-name")[0] as HTMLDivElement;
            const roomPlayers = roomElem.getElementsByClassName("room-players")[0] as HTMLDivElement;
            const roomJoin = roomElem.getElementsByClassName("room-join")[0] as HTMLAnchorElement;
            roomName.innerText = roomId;
            roomPlayers.innerText = "TODO/2 players";
            roomJoin.href = "/room?id=" + roomId;
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