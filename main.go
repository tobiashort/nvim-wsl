package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
  var file string
  var cmd *exec.Cmd
  if len(os.Args) == 2 {
    cmd = exec.Command("wsl.exe", "-e", "wslpath", os.Args[1])
    out, err := cmd.CombinedOutput()
    outString := string(out)
    if err != nil {
      fmt.Fprintf(os.Stderr, outString)
      panic(err)
    }
    file = strings.TrimSpace(outString)
  }
  if file == "" {
    cmd = exec.Command("wsl.exe", "-e", "nvim")
  } else {
    cmd = exec.Command("wsl.exe", "-e", "nvim", string(file))
  }
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}
