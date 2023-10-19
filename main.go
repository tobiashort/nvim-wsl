package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func toWSLPath(winpath string) (string, error) {
  cmd := exec.Command("wsl.exe", "-e", "wslpath", winpath)
  out, err := cmd.CombinedOutput()
  outString := string(out)
  if err != nil {
    fmt.Fprintf(os.Stderr, outString)
    return "", err
  }
  wslpath := strings.TrimSpace(outString)
  return wslpath, err
}

func runNvim(wslpath string) error {
  var cmd *exec.Cmd
  if wslpath == "" {
    cmd = exec.Command("wsl.exe", "-e", "nvim")
  } else {
    cmd = exec.Command("wsl.exe", "-e", "nvim", wslpath)
  }
  cmd.Stdin = os.Stdin
  cmd.Stdout = os.Stdout
  cmd.Stderr = os.Stderr
  return cmd.Run()
}

func main() {
  if len(os.Args) > 2 {
    fmt.Fprintf(os.Stderr, "Too many arguments: %v\n", os.Args)
    os.Exit(1)
  } 

  var err error
  var winpath string
  var wslpath string

  if len(os.Args) == 2 {
    winpath = os.Args[1]
  }

  if winpath != "" {
    wslpath, err = toWSLPath(winpath)
    if err != nil {
      fmt.Fprintf(os.Stderr, "toWSLPath failed: %s", err)
      os.Exit(1)
    }
  }

  err = runNvim(wslpath)
  if err != nil {
      fmt.Fprintf(os.Stderr, "runNvim failed: %s", err)
      os.Exit(1)
  }
}
