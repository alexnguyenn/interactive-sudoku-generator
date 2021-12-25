import Cell from "./cell";

const EMPTY = (() => {
    let temp = [];
    for (let i = 0; i < 81; i++) {
        temp.push(0);
    }
    return temp.join("");
})();

export default class Grid {
    constructor(input = EMPTY) {
        let currentRow;
        this.rows = [];

        for (let idx = 0; idx < input.length; idx++) {
            if (idx % 9 === 0) {
                currentRow = [];
                this.rows.push(currentRow);
            }

            currentRow.push(
                new Cell(this.rows.length - 1, currentRow.length, input[idx])
            );
        }
    }

    toFlatString() {
        return this.rows
            .flat()
            .map(x => x.toString())
            .join("");
    }
}
