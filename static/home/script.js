"use strict";
class RoomsLoader {
    constructor() {
        this.es = new EventSource("/api/game/roomsSSE");
        this.es.onmessage = (event) => this.handleMsg(event);
    }
    handleMsg(event) {
        console.log(event.data);
    }
}
new RoomsLoader();
