import { TARGET_BOARD_3, TARGET_BOARD_4, TARGET_BOARD_5 } from '../const.js'

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
};

export const findPos = (value: number, puzzle: number[][]) => {
    for (let i = 0; i < puzzle.length; i++) {
        for (let j = 0; j < puzzle[i].length; j++) {
            if (puzzle[i][j] === value) {
                return i * puzzle.length + j;
            }
        }
    }
    return 0;
};

export const generateRandomBoard = (size: number) => {
    let tmp: number[][];
    if (size == 3) {
        tmp = TARGET_BOARD_3;
    }else if (size == 4) {
        tmp = TARGET_BOARD_4;
    } else {
        tmp = TARGET_BOARD_5;
    }
    console.log('taget_board', tmp)
    for (let i = 0; i < 100 + Math.floor(Math.random() * 100); i++) {
        var idx = findPos(0, tmp);
        var arr = [];
        if (idx % size > 0) {
            arr.push(idx - 1);
        }
        if (idx % size < size - 1) {
            arr.push(idx + 1);
        }
        if (Math.floor(idx / size) > 0 && idx > size) {
            arr.push(idx - size);
        }
        if (Math.floor(idx / size) < size - 1) {
            arr.push(idx + size);
        }
        var swapIndex = arr[Math.floor(Math.random() * arr.length)];
        tmp[Math.floor(idx / size)][idx % size] = tmp[Math.floor(swapIndex / size)][swapIndex % size];
        tmp[Math.floor(swapIndex / size)][swapIndex % size] = 0;
    }
    return tmp;
};