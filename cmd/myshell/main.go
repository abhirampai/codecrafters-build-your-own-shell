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

    sanitizedUserInput := strings.Split(strings.TrimSuffix(userInput, "\n"), " ")
    
    command := sanitizedUserInput[0]
    args := sanitizedUserInput[1:]

    switch command {
      case "exit":
          os.Exit(0)
      case "echo":
          fmt.Println(strings.Join(args, " "))
      default:
        fmt.Println(command + ": not found")
    }
  }
}
