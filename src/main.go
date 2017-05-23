package main

import "flag"
import "fmt"
import "os"
import "os/exec"
import "strconv"
import "time"


const DEFAULT_SOUND = "/usr/share/sounds/freedesktop/stereo/complete.oga"


func main() {
    path := flag.String("plik", DEFAULT_SOUND, "ścieżka do pliku z ding")
    flag.Parse()
    args := flag.Args()

    fmt.Println(args)

    cmd := exec.Command("paplay", *path)
    cmd.Start()

    var interval int

    if len(args) >= 1 {
        tmp, err := strconv.Atoi(args[0])
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
        cmd := exec.Command("paplay", *path)
        cmd.Start()
        counter++
    }
}
