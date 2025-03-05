class RoomConn {
    socket: WebSocket;

    constructor(roomId: string) {
        this.socket = new WebSocket("ws://localhost/api/game/roomWS/" + roomId);

        this.socket.onmessage = (event) => {
            console.log(event.data);
        }

        this.socket.onclose = () => {
            console.log("Connection closed");
        }

        this.socket.onopen = () => {
            console.log("Connection opened");
            this.socket.send("Hello");
        }
    }
}