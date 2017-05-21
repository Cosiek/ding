package main

import "fmt"
import "os"
import "os/exec"
import "strconv"
import "time"


func main() {
    cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
    cmd.Start()

    var interval int

    if len(os.Args) == 2 {
        tmp, err := strconv.Atoi(os.Args[1])
        if err != nil{
            fmt.Println("Podaj prawidłowy argument (liczba)!")
            os.Exit(1)
        }
        interval = tmp
    } else {
        fmt.Println("Podaj prawidłową częstotliwość ding!")
        os.Exit(1)
    }

    counter := 0
    for true {
        fmt.Println("Ding!", counter)
        time.Sleep(time.Second * time.Duration(interval))
        cmd := exec.Command("paplay", "/usr/share/sounds/freedesktop/stereo/complete.oga")
        cmd.Start()
        counter++
    }
}
