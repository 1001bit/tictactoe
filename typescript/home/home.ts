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

new RoomsLoader();