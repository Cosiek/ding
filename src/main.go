package main

import "fmt"
import "os"
import "os/exec"
import "strconv"
import "time"


const DEFAULT_SOUND = "/usr/share/sounds/freedesktop/stereo/complete.oga"


func main() {
    cmd := exec.Command("paplay", DEFAULT_SOUND)
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
        cmd := exec.Command("paplay", DEFAULT_SOUND)
        cmd.Start()
        counter++
    }
}
