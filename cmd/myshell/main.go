package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "slices"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint

func commandNotFound(command string) {
  fmt.Println(command + ": not found")
}

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
    builtInCommands := []string { "exit", "echo", "type" }

    switch command {
      case "exit":
          os.Exit(0)
      case "echo":
          fmt.Println(strings.Join(args, " "))
      case "type":
        typeCommand := args[0]
        pathEnv := os.Getenv("PATH")

        if slices.Contains(builtInCommands, typeCommand) {
          fmt.Println(typeCommand + " is a shell builtin")
        } else if len(pathEnv) > 0 {
          pathsToCheck := strings.Split(pathEnv, ":")
          pathFound := false

          for i := 0; i < len(pathsToCheck); i++ {
            executablePath := pathsToCheck[i] + "/" + typeCommand
            if _, err := os.Stat(executablePath); err == nil {
              fmt.Println(typeCommand + " is " + executablePath)
              pathFound = true
              break
            }
          }

          if !pathFound {
            commandNotFound(typeCommand)
          }
        } else {
          commandNotFound(typeCommand)
        } 
      default:
        commandNotFound(command)
    }
  }
}
