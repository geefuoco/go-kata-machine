package main

import (
  "os"
  "io"
  "fmt"
  "runtime"
  "path/filepath"
)

const DIR = "../katas"

func crashOnError(err error) {
  if err != nil {
    panic(err)
  }
}

func generateKatas() {
  _, caller, _, _ := runtime.Caller(0)
  ROOT := filepath.Dir(caller)
  KATA := filepath.Join(ROOT, DIR)
  TESTS := filepath.Join(KATA, "tests")
  println("Generating katas")
  err := os.RemoveAll(KATA)
  crashOnError(err)
  err = os.Mkdir(KATA, 0774)
  crashOnError(err)
  err = os.Mkdir(filepath.Join(KATA, "tests"), 0774)
  crashOnError(err)
  err = copyDir(filepath.Join(ROOT, "templates"), KATA)
  crashOnError(err)
  err = copyDir(filepath.Join(ROOT, "tests_templates"), TESTS)
  crashOnError(err)
  fmt.Println("Created katas directory")
  crashOnError(err)
  fmt.Println("Katas created.")
  fmt.Println(`
    cd ./katas to start

    run tests with 'go test ./katas/tests/<filename>'

    run all tests with 'go test ./katas/tests'
   `)
}

func copyDir(src, dst string) error {
  var err error

  if _, err = os.Stat(src); os.IsNotExist(err) {
    return err
  } 

  entries, err := os.ReadDir(src)
  for _, file := range entries {
    err = copyFile(filepath.Join(src, file.Name()),filepath.Join(dst, file.Name()), file.Type().Perm())
    if err != nil {
      return err
    }
  }
  return nil
}

func copyFile(src, dest string, mode os.FileMode) error {
  var err error
  var start *os.File
  var end *os.File

  if start, err = os.Open(src); err != nil {
    fmt.Println("Error when opening file")
    return err
  }
  defer start.Close()

  if end, err = os.Create(dest); err != nil {
    fmt.Println("Error when creating file")
    return err
  }
  defer end.Close()

  if _, err = io.Copy(end, start); err != nil {
    fmt.Println("Error when copying file")
    return err
  }
  return os.Chmod(dest, 0664)
}

func main() {
  generateKatas()
}
