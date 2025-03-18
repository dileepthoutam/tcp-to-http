package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

const inputFilePath = "messages.txt"

func ReadMessages() {
    file, err := os.Open(inputFilePath)
    if err != nil {
        log.Fatalf("Could not open file %s with error %s\n", inputFilePath, err)
    }

    for {
        buff := make([]byte, 8, 8)
        n, err := file.Read(buff)
        if err != nil {
            if errors.Is(err, io.EOF) {
                break
            }
            fmt.Printf("error: %s\n", err.Error())
        }
        fmt.Printf("read: %s\n", string(buff[:n]))
    }
}

func ReadMessagesLineByLine() {
    file, err := os.Open("messages.txt")

    if err != nil {
        fmt.Println("File not found.")
        return
    }

    buff := make([]byte, 8)
    line := ""
    for {
        n, err := file.Read(buff)

        if err != nil {
            if line != "" {
                fmt.Printf("read: %s\n", line)
            }
            if err == io.EOF {
                break
            } else {
                fmt.Println("Error reading the file: ", err)
                break
            }
        }

        str := string(buff[:n])
        parts := strings.Split(str, "\n")
        for i := 0; i < len(parts) - 1; i++ {
            fmt.Printf("read: %s%s\n", line, parts[i])
            line = ""
        }

        line += parts[len(parts) - 1]

        /* if !bytes.Contains(buff[:n], []byte("\n")) {
            line += string(buff[:n])
        } else {
            fmt.Printf("read: %s", line)
            split := bytes.Split(buff, []byte("\n"))
            fmt.Printf("%s\n", split[0])
            line = ""
            line += string(split[1])
        } */

    }

}

func main() {
    ReadMessagesLineByLine()
}
