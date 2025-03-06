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

	handleClick(x: number, y: number) {
		if (!this.allowPlace) {
			return;
		}

		let cell = document.getElementById(
			"cell-" + y + "-" + x,
		) as HTMLDivElement;
		cell.innerText = this.sign;
		this.allowPlace = false;
		this.placecallback(x, y);
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
