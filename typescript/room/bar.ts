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

    setResult(result: string, sign: string){
        if(result == "D"){
            turnElem.innerText = "Draw";
        } else if (result == "W"){
            turnElem.innerText = `You win (${sign})`;
        } else {
            turnElem.innerText = `You lose (${sign})`;
        }
    }
}