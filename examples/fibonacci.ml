export println

let i1 = 0
let i2 = 1

let turns = 10

loop turns > 1 {
    let tmp = i2
    i2 = i2 + i1
    i1 = tmp

    turns = turns - 1
}

println(i1)