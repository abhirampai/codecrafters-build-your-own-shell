package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func main() {
  for true {
    fmt.Fprint(os.Stdout, "$ ")

	  // Wait for user input
    userInput, err := bufio.NewReader(os.Stdin).ReadString('\n')

    if err != nil {
      fmt.Fprintln(os.Stderr, "Error reading input:", err)
      os.Exit(1)
    }

    command := strings.Split(strings.TrimSuffix(userInput, "\n"), " ")[0]
    switch command {
      case "exit":
          os.Exit(0)
      default:
        fmt.Println(command + ": not found")
    }
  }
}
