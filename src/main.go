package main

import "fmt"
import "os/exec"

func main() {
    fmt.Println("Ding!")
    cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
    cmd.Start()
}
