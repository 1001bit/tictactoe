/// <reference path="elems.ts" />

class Room {
	conn: RoomConn;
	topbar: TopBar;
	board: Board;

	constructor(roomId: string) {
		this.conn = new RoomConn(roomId);
		this.conn.onmessage = (data: any) => {
			this.handleMessage(data);
		};

		this.topbar = new TopBar();
		this.topbar.setRoomId(roomId);

		this.board = new Board();
		this.board.placecallback = (x: number, y: number) => {
			this.handlePlace(x, y);
		};
	}

	handleMessage(msg: any) {
		if (!("type" in msg)) {
			return;
		}

		switch (msg.type) {
			case "start":
				this.handleStart(msg.you, msg.turn);
				break;
			case "stop":
				this.handleStop();
				break;
			case "move":
				this.handleOpponentMove(msg.x, msg.y, msg.sign);
				break;
		}
	}

	handleStart(sign: string, turn: string) {
		this.board.setSign(sign);
		this.board.setAllowPlace(turn == sign);
		this.topbar.setTurn(turn == sign, turn);
	}

	handleStop() {
		this.topbar.stop();
		this.board.clear();
	}

	handleOpponentMove(x: number, y: number, sign: string) {
		if (sign == this.board.sign) {
			return;
		}

		this.board.handleOpponentMove(x, y, sign);
		this.topbar.setTurn(true, this.board.sign);
	}

	handlePlace(x: number, y: number) {
		const turn = this.board.sign == "X" ? "O" : "X";
		this.topbar.setTurn(false, turn);
		this.conn.sendPlaceMessage(x, y);
	}
}

window.onload = () => {
	const roomId = new URLSearchParams(window.location.search).get("id");
	if (!roomId) {
		window.location.href = "/";
		return;
	}
	new Room(roomId);
};
