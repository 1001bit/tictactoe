class RoomsLoader {
    es: EventSource;

    constructor() {
        this.es = new EventSource("/api/game/roomsSSE");
        this.es.onmessage = (event) => this.handleMsg(event);
    }

    handleMsg(event: MessageEvent) {
        console.log(event.data);
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