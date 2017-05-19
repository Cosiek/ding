package main

import "fmt"
import "os"
import "os/exec"
import "strconv"
import "time"


func main() {
    cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
    cmd.Start()

    interval, _ := strconv.Atoi(os.Args[1])
    counter := 0

    for true {
        fmt.Println("Ding!", counter)
        time.Sleep(time.Second * time.Duration(interval))
        cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
        cmd.Start()
        counter++
    }
}
