package mt940

import (
    "os"
    "log"
    "strings"
    "fmt"
)

const (
    transactionReferenceNumber = ":20:"
    relatedReference           = ":21:"
    accountIdentification      = ":25:"
)

func ReaderFromFile(fileName string) (Document, error) {
    file, err := os.Open(fileName)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    document, message := Document{}, Message{}

    position := int(0)
    stat, err := file.Stat()
    if err != nil {
        log.Fatal(err)
    }
    size := int(stat.Size())
    bytesRead, data := make([]byte, 511), ""
    for position < size {
        file.Read(bytesRead)
        data = string(bytesRead)

        if data[0:1] == " " || data[0:1] == "\n" {
            position++
            continue
        }
        if data[0:1] == "-" {
            document.messages = append(document.messages, message)
            message = Message{}
        }
        if data[0:1] != ":" {
            return document, fmt.Errorf("Fuck parse error [%s:%d]", fileName, position)
        }

        if strings.Index(data, transactionReferenceNumber) == 0 {
            content, bytesRead := readTag(transactionReferenceNumber, data, 16, true)
            tag := Tag{
                field:   transactionReferenceNumber,
                content: content,
            }
            message.tags = append(message.tags, tag)
            position += bytesRead
            continue
        }
        if strings.Index(data, relatedReference) == 0 {

        }

        position++
    }

    return document, nil
}

type Tag struct {
    field   string
    content string
}

type Message struct {
    tags []Tag
}

type Document struct {
    messages []Message
}

func readTag(field string, data string, maxBytes int, skipOnNewLine bool) (string, int) {
    if skipOnNewLine {
        if newLinePosition := strings.Index(data, "\n"); newLinePosition >= 0 && newLinePosition <= maxBytes {
            maxBytes = newLinePosition
        }
    }
    return data[len(field):maxBytes], maxBytes
}

func NewDocument() *Document {
    return &Document{
        messages: nil,
    }
}
