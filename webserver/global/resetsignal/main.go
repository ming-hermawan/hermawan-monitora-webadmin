package resetsignal


var resetSignal bool = false

func Get() bool {
    return resetSignal
}

func Set(param bool) {
    resetSignal = param
}
