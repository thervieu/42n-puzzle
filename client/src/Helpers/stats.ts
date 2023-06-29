export const averageTime = (arr: any) => {
    let sum = 0;
    for (let i = 0; i < arr.length; i++) {
        console.log(arr[i].time)
        if (arr[i].time.endsWith("ms") === true) {
            sum += Number(arr[i].time.substring(0, arr[i].time.length-2))
        } else if (arr[i].time.endsWith("µs") === true) {
            sum += (Number(arr[i].time.substring(0, arr[i].time.length-2)) * 0.001)
        } else {
            console.log("seconds")
            sum += Number(arr[i].time.substring(0, arr[i].time.length-1) * 1000)
        }
    }
    console.log("averageTime", sum, arr.length, sum / arr.length)
    return sum / arr.length
};

export const averageMoves = (arr: any) => {
    let sum = 0;
    for (let i = 0; i < arr.length; i++) {
        console.log(i, arr[i].moves)
        sum += arr[i].moves
    }
    console.log("averageMoves", sum, arr.length, sum / arr.length)
    return sum / arr.length
};

export const averageTC = (arr: any) => {
    let sum = 0;
    for (let i = 0; i < arr.length; i++) {
        console.log(i, arr[i].time_complexity)
        sum += arr[i].time_complexity
    }
    console.log("averageTC", sum, arr.length, sum / arr.length)
    return sum / arr.length
};

export const averageSC = (arr: any) => {
    let sum = 0;
    for (let i = 0; i < arr.length; i++) {
        console.log(i, arr[i].space_complexity)
        sum += arr[i].space_complexity
    }
    console.log("averageSC", sum, arr.length, sum / arr.length)
    return sum / arr.length
};