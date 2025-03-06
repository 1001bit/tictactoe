class RoomConn {
    socket: WebSocket;
    onmessage: (data: any) => void;

    constructor(roomId: string) {
        const protocol = window.location.protocol === "https:" ? "wss" : "ws";
        const host = window.location.host;

        this.socket = new WebSocket(`${protocol}://${host}/api/game/roomWS/${roomId}`);
        this.onmessage = (_data: any) => {}

        this.socket.onmessage = (event) => {
            this.onmessage(JSON.parse(event.data));
        }

        this.socket.onclose = () => {
            console.log("Connection closed");
        }

        this.socket.onopen = () => {
            console.log("Connection opened");
        }
    }
}