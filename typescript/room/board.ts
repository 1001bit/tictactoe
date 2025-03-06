class Board {
	allowPlace: boolean;
	sign: string;
	placecallback: (x: number, y: number) => void;

	constructor() {
		this.allowPlace = false;
		this.sign = " ";
		this.placecallback = (_x: number, _y: number) => {};

		// handle click
		for (let x = 0; x < 3; x++) {
			for (let y = 0; y < 3; y++) {
				let cell = document.getElementById(
					"cell-" + y + "-" + x,
				) as HTMLDivElement;
				cell.addEventListener("click", () => this.handleClick(x, y));
			}
		}
	}

	setCell(x: number, y: number, sign: string) {
		let cell = document.getElementById(
			"cell-" + y + "-" + x,
		) as HTMLDivElement;
		cell.innerText = sign;
	}

	getCell(x: number, y: number): string {
		let cell = document.getElementById(
			"cell-" + y + "-" + x,
		) as HTMLDivElement;
		return cell.innerText;
	}

	handleClick(x: number, y: number) {
		if (!this.allowPlace || this.getCell(x, y) != "") {
			return;
		}

		this.setCell(x, y, this.sign);
		this.allowPlace = false;
		this.placecallback(x, y);
	}

	handleOpponentMove(x: number, y: number, sign: string) {
		this.setCell(x, y, sign);
		this.allowPlace = true;
	}

	setAllowPlace(allow: boolean) {
		this.allowPlace = allow;
	}

	setSign(sign: string) {
		this.sign = sign;
	}

	clear() {
		this.sign = " ";
		this.allowPlace = false;

		for (let x = 0; x < 3; x++) {
			for (let y = 0; y < 3; y++) {
				let cell = document.getElementById(
					"cell-" + y + "-" + x,
				) as HTMLDivElement;
				cell.innerText = " ";
			}
		}
	}
}
