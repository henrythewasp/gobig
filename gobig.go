package main

import (
    "log"
    "flag"
    "fmt"
    "os"
)

func main() {
    namePtr := flag.String("f", "", "name of file to create")
    sizePtr := flag.Int64("b", 0, "size of file, in bytes")

    flag.Parse()

    if (len(*namePtr) == 0) {
        fmt.Print("Enter name of file to create: ")
        fmt.Scanf("%s\n", namePtr)
        fmt.Print("Enter size of file (in bytes): ")
        fmt.Scanf("%d\n", sizePtr)
    }

    f, err := os.Create(*namePtr)
    if err != nil {
        log.Fatal(err)
    }
    _, err = f.WriteString("The quick brown fox jumps over the lazy dog")
    if err != nil {
        log.Fatal(err)
    }

    if err := f.Truncate(*sizePtr); err != nil {
        log.Fatal(err)
    }

    fmt.Printf("Created file %v with %v bytes\n", *namePtr, *sizePtr)
}
