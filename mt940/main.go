package mt940

import (
    "os"
    "log"
    "bufio"
    "strings"
)

const (
    transactionReferenceNumber = "20"
    relatedReference           = "21"
    accountIdentification      = "25"
)

func ReaderFromFile(fileName string) Document {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    d, s, t := Document{}, Section{}, Tag{}
    scanner := bufio.NewScanner(file)
    for scanner.Scan() {
        line := strings.TrimSpace(scanner.Text())
        if line == "" {
            continue
        }
        if startWith
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }
    return document
}

type Tag struct {
    Number string
    Value  string
}

type Section struct {
    tags []Tag
}

type Document struct {
    sections []Section
}

func NewDocument() *Document {
    return &Document{
        sections: nil,
    }
}
