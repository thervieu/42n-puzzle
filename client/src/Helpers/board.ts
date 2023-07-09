import { TARGET_BOARD_3 } from "../const";
const findPos = (value: number, puzzle: number[][]) => {
    for (let i = 0; i < puzzle.length; i++) {
        for (let j = 0; j < puzzle[i].length; j++) {
            if (puzzle[i][j] === value) {
                return i * 3 + j;
            }
        }
    }
    return 0;
};
export const generateRandomBoard = () => {
    let tmp: number[][];
    tmp = TARGET_BOARD_3;
    for (let i = 0; i < 100 + Math.floor(Math.random() * 100); i++) {
        var idx = findPos(0, tmp);
        var arr = [];
        if (idx % 3 > 0) {
            arr.push(idx - 1);
        }
        if (idx % 3 < 3 - 1) {
            arr.push(idx + 1);
        }
        if (Math.floor(idx / 3) > 0 && idx > 3) {
            arr.push(idx - 3);
        }
        if (Math.floor(idx / 3) < 3 - 1) {
            arr.push(idx + 3);
        }
        var swapIndex = arr[Math.floor(Math.random() * arr.length)];
        tmp[Math.floor(idx / 3)][idx % 3] = tmp[Math.floor(swapIndex / 3)][swapIndex % 3];
        tmp[Math.floor(swapIndex / 3)][swapIndex % 3] = 0;
    }
    return tmp;
};
export async function solvePuzzle(puzzle: number[][], search_: string, heuristic_: string) {
    var puzzleArr: number[] = [];
    for (let i = 0; i < puzzle.length; i++) {
        for (let j = 0; j < puzzle[i].length; j++) {
            puzzleArr.push(puzzle[i][j]);
        }
    }
    const response = await fetch('http://localhost:3020/solve', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({
            puzzle: puzzleArr,
            search: search_,
            heuristic: heuristic_
        })
    });
    if (response.ok) {
        return await response.json();
    } else {
        throw new Error('help');
    }
}

export const isSolved = (board: number[][]) => {
    for (let i = 0; i < 3; i++) {
        for (let j = 0; j < 3; j++) {
            if (board[i][j] != TARGET_BOARD_3[i][j]) {
                return false;
            }
        }
    }
    return true;
};


export const isValidMove = (row: number, col: number, emptyRow: number, emptyCol: number) => {
    // Vérifier si la case cliquée est adjacente à la case vide
    if (Math.abs(row - emptyRow) + Math.abs(col - emptyCol) !== 1) {
        return false;
    }
    // Vérifier si la case cliquée est à une distance de 1 de la case vide
    if (row < 0 || row >= 3 || col < 0 || col >= 3) {
        return false;
    }
    // Mouvement valide
    return true;
};