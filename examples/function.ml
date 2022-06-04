export println

fun max(let x1, let x2) {
    if x1 > x2 {
        return x1
    }
    return x2
}

fun min(let x1, let x2) {
    if x1 < x2 {
        return x1
    }
    return x2
}

let mi = min(69, 420)
let ma = max(69, 420)

println("min (69, 420) =", mi)
println("max (69, 420) =", ma)