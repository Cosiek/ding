package main

import "fmt"
import "os"
import "os/exec"
import "strconv"

func main() {
    fmt.Println("Ding!")
    cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
    cmd.Start()

    interval, _ := strconv.Atoi(os.Args[1])
    fmt.Println(interval)
}
