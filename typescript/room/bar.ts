class TopBar {
    constructor() {
        
    }

    setRoomId(roomId: string) {
        roomIdElem.innerText = `Room "${roomId}"`;
        document.title = `Room "${roomId}"`;
    }

    setTurn(yours: boolean, sign: string) {
        turnElem.innerText = `Turn: ${yours ? "You" : "Opponent"} (${sign})`;
    }

    stop(){
        turnElem.innerText = "Waiting for another player";
    }
}