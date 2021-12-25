import React, { useEffect, useState } from "react";
import ReactDOM from "react-dom";
import Grid from "./grid.js";
import axios from "axios";

import "./index.css";

const Square =({value, row, col, onCellValueChange}) => (
    <input
        type="text"
        value={value === 0 ? "" : value}
        maxLength="1"
        onChange={(evt) => {
            const value = evt.target.value;
            if (parseInt(value, 10) || value === "") {
                onCellValueChange(row, col, value);
            }
        }}
    />
);

const SudukoBoard = ({ puzzleGrid, onCellValueChange }) => (
    <table className="sudoku">
        <tbody>
        { puzzleGrid.rows.map((row, idx) => (
            <tr key={idx}>
                { row.map(cell => (
                    <td key={cell.col}>
                        <Square
                            value={cell.value}
                            row={cell.row}
                            col={cell.col}
                            onCellValueChange={onCellValueChange}
                        />
                    </td>
                )) }
            </tr>
        )) }
        </tbody>
    </table>
);

function SudokuGame() {
    const [puzzle, setPuzzle] = useState(new Grid());
    const [solution, setSolution] = useState(new Grid());
    const [solutionVisible, showSolution] = useState(false);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        getData()
    }, [])

    function getData() {
        setLoading(true);
        setError(null);
        axios("http://127.0.0.1:8000/api/sudoku/")
        .then((response) => {
            console.log(response.data);
            setPuzzle(new Grid(response.data.instance));
            setSolution(new Grid(response.data.solution));
        })
        .catch(error => {
            console.error("Error fetching data ", error);
            setError(error);
        })
        .finally(() => {
            setLoading(false);
        })
    }

    function onCellValueEdited (row, col, value) {
        if (!loading) {
            const newGrid = new Grid(puzzle.toFlatString());
            newGrid.rows[row][col].value = value;
            setPuzzle(newGrid);
        }
 
    }

    function changeSolutionVisible() {
        if (solutionVisible) {
            showSolution(false);
        } else {
            showSolution(true);
        }
    }
    
    return (
        <div className="game">
            {loading && <h1>Loading...</h1>}
            {error && <h1>Error fetching data from server</h1>}
            <SudukoBoard
                puzzleGrid={puzzle}
                onCellValueChange={onCellValueEdited}
            />
            {solutionVisible && <SudukoBoard
                puzzleGrid={solution}
                onCellValueChange={() => void 0}
            />}
            <div className="buttons">
                <button onClick={changeSolutionVisible}>Show/Hide Solution</button>
                <button onClick={getData}>Generate New</button>
            </div>
        </div>
    );
}

ReactDOM.render(
    <SudokuGame/>,
    document.getElementById("root")
);
