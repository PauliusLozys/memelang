export println
let i = 0
loop i != 100 {

    let tmp = match(i) {
        == 100 => "One hundred",
        < 20 => "Below 20",
        > 50 => "Above 50",
        return "Between 20 and 50"
    }

    println("index is:", i, tmp)
    i = i + 5
}