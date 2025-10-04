package main

import (
    "fmt"
    "os"
    "hermawan-monitora/webserver"
    "hermawan-monitora/hmonglobal/lang"
)

func main() {
    // Check Arguments
    args := os.Args
    if (len(args) == 2) && ((args[1] == "-h") || (args[1] == "--help")) {
        fmt.Printf(lang.Help())
        return
    }
    // Process
    webserver.Run()
}
