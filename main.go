package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
  cmd := exec.Command("wsl.exe", "-e", "wslpath", os.Args[1])
  out, err := cmd.CombinedOutput()
  outString := string(out)
  if err != nil {
    fmt.Fprintf(os.Stderr, outString)
    panic(err)
  }
  outString = strings.TrimSpace(outString)
  cmd = exec.Command("wsl.exe", "-e", "nvim", string(outString))
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  cmd.Run()
}
