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

    counter := 1
    for true {
        fmt.Println("Ding!", counter)
        // check if file exists at given path
        if _, err := os.Stat(*path); err == nil {
            cmd := exec.Command("paplay", *path)
            err := cmd.Run()
            if err != nil {
                fmt.Println("Błąd odtwarzania!:\n" + err.Error())
                os.Exit(1)
            }
        } else {
            fmt.Println("Nie znaleziono takiego pliku z Ding!")
            os.Exit(1)
        }
        time.Sleep(time.Second * time.Duration(interval))
        counter++
    }
}
