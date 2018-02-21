package main

import (
    "os"
    "log"
)

func main() {
    fileNames := os.Args[1:]
    for fileName := range fileNames {
        reader, err := mt940.ReaderFromFile(fileName)
        if err != nil {
            log.Printf("Error %s opening file [%s]", err, fileName)
            os.Exit(1)
        }
        log.Printf("MT940 contains %d lines\n", reader.CountLines())
    }
}
