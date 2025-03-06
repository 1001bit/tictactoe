class Board {
    constructor() {
        // handle click
        for (let x = 0; x < 3; x++) {
            for (let y = 0; y < 3; y++) {
                let cell = document.getElementById("cell-" + y + "-" + x) as HTMLDivElement;
                cell.addEventListener("click", () => this.handleClick(x, y));
            }
        }
    }

    handleClick(x: Number, y: Number) {
        console.log(x, y);
    }
}