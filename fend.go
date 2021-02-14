package main

import (
    "fmt"
    // "io/ioutil"
    "log"
	"os"
	"path/filepath"
)

func main() {
    // files, err := ioutil.ReadDir("./")
    // if err != nil {
    //     log.Fatal(err)
    // }
 
    // for _, f := range files {
    //         fmt.Println(f.Name())
    // }

    // file, err := os.Open(".dotfile")
    // if err != nil {
    //     log.Fatal(err)
    // }
    // defer func() {
    //     if err = file.Close(); err != nil {
    //         log.Fatal(err)
    //     }
    // }()


  	//b, err := ioutil.ReadAll(file)
 	//fmt.Print(b)
    subDirToSkip := ".git"
  	err := filepath.Walk(".",func(path string, info os.FileInfo, err error) error {
  		if err != nil {
	  		return err
  		}
		if info.IsDir() && info.Name() == subDirToSkip	{
			//Skip
            return filepath.SkipDir
		} else {
            result, err := checkLineEnding(path)
            if err != nil {
                return err
            }
            fmt.Println(path, result)
		}
  		return nil
	})
if err != nil {
  log.Println(err)
}

  //readFile(".dotfile")
}

func checkLineEnding(fname string) (bool, error) {
    newLine := "\r\n"
    file, err := os.Open(fname)
    if err != nil {
        panic(err)
    }
    defer file.Close()
    buf := make([]byte, 2)
    stat, err := os.Stat(fname)
	//fmt.Print(stat.Size())
    start := stat.Size() - 2
    _, err = file.ReadAt(buf, start)
    if err == nil {
        fmt.Printf("%s\n", buf)
    }
	fmt.Print(buf)
	//s := string([]byte{buf})
	myString := string(buf)
	fmt.Print(myString)
    fmt.Print(newLine)
    b := []byte(newLine)
    fmt.Print(b)
    if myString == newLine {
        return true, nil
    }
	return false, nil
}