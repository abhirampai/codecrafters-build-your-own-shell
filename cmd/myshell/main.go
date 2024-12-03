package main

import (
	"bufio"
	"fmt"
	"os"
  "strings"
  "slices"
  "os/exec"
  "unicode"
)

// Ensures gofmt doesn't remove the "fmt" import in stage 1 (feel free to remove this!)
var _ = fmt.Fprint
var pathEnv = os.Getenv("PATH")

func commandNotFound(command string) {
  fmt.Println(command + ": not found")
}

func findExecutablePath(command string) (executablePath string, pathFound bool) {
  pathsToCheck := strings.Split(pathEnv, ":")
  pathFound = false

  for i := 0; i < len(pathsToCheck); i++ {
    executablePath = pathsToCheck[i] + "/" + command
    if _, err := os.Stat(executablePath); err == nil {
      pathFound = true
      break
    }
  }
  return
}

func splitString(s string) []string {
 	var result []string
	var current []rune
	var quote rune
  var nestedQuote rune
	escaped := false

	for i, r := range s {
		switch {
		case escaped:
			current = append(current, r)
			escaped = false
		case r == '\\':
      if nestedQuote != '\'' && quote != '\'' {
        if quote == 0 || (quote != 0 && (s[i+1] == '"' || s[i+1] == '\\' || s[i+1] == '$')) {
          escaped = true
        } else {
			    current = append(current, r)
        }
      } else {
        current = append(current, r)
      }
		case quote != 0:
			if r == quote {
				quote = 0
			} else {
        if r == '"' || r == '\'' {
          if nestedQuote == r {
            nestedQuote = 0
          } else {
            nestedQuote = r
          }
        }
				current = append(current, r)
			}
		case r == '"' || r == '\'':
			quote = r
		case unicode.IsSpace(r):
			if len(current) > 0 {
				result = append(result, string(current))
				current = nil
			}
		default:
			current = append(current, r)
		}
	}

	if len(current) > 0 {
		result = append(result, string(current))
	}

	return result
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

    sanitizedUserInput := splitString(strings.TrimSuffix(userInput, "\n"))
    
    command := sanitizedUserInput[0]
    args := sanitizedUserInput[1:]
    builtInCommands := []string { "exit", "echo", "type", "pwd", "cd" }

    switch command {
      case "exit":
        os.Exit(0)
      case "echo":
        output := strings.Join(args, " ")
        fmt.Println(output)
      case "type":
        typeCommand := args[0]

        if slices.Contains(builtInCommands, typeCommand) {
          fmt.Println(typeCommand + " is a shell builtin")
        } else if len(pathEnv) > 0 {
          execPath, pathFound := findExecutablePath(typeCommand)

          if pathFound {
            fmt.Println(typeCommand + " is " + execPath)
          } else {
            commandNotFound(typeCommand)
          }
        } else {
          commandNotFound(typeCommand)
        }
      case "pwd":
        currentWorkingDirectory, _ := os.Getwd()
        fmt.Println(currentWorkingDirectory)
      case "cd":
        path := args[0]
        if args[0] == "~" {
          path = os.Getenv("HOME")
        }

        err := os.Chdir(path)

        if err != nil {
          fmt.Println("cd: " + path + ": No such file or directory")
        }
      case "cat":
        for _, filePath := range args {
          fileContent, err := os.ReadFile(filePath)
          if err != nil {
            fmt.Print("Something went wrong")
          }
          fmt.Print(string(fileContent))
        }

      default:
        if execPath, pathFound := findExecutablePath(command); pathFound {
          cmd := exec.Command(execPath, args...)
          cmdOutput, _ := cmd.Output()
          fmt.Print(string(cmdOutput))
        } else {
          commandNotFound(command)
        }
    }
  }
}
