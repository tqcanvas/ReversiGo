package main

import (
    "os"
    "golang.org/x/sys/windows"
)

//Function to enable ANSI escape codes in Windows
//Sourced from https://stackoverflow.com/questions/39627348/ansi-colours-on-windows-10-sort-of-not-working
func init() {
    stdout := windows.Handle(os.Stdout.Fd())
    var originalMode uint32

    windows.GetConsoleMode(stdout, &originalMode)
    windows.SetConsoleMode(stdout, originalMode|windows.ENABLE_VIRTUAL_TERMINAL_PROCESSING)
}