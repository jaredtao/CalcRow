package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

type CalcInfo struct {
	Suffix    string
	FileCount int
	Line      int
}

var suffix = [...]string{".h", ".c", ".hpp", ".cpp", ".go", ".qml"}
var fileCount int

var calcMap map[string]*CalcInfo = make(map[string]*CalcInfo)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], " Path")
		os.Exit(1)
	}
	for _, s := range suffix {
		calcMap[s] = &CalcInfo{Suffix: s}
	}

	ch := make(chan bool)
	go calcRow(os.Args[1], ch)
	<-ch
	output()
}
func output() {
	sumLine := 0
	sumCount := 0
	for _, s := range suffix {
		fmt.Println(
			calcMap[s].Suffix,
			" file count ", calcMap[s].FileCount,
			" line ", calcMap[s].Line)
		sumLine += calcMap[s].Line
		sumCount += calcMap[s].FileCount
	}
	fmt.Println("sum line ", sumLine)
	fmt.Println("sum file count ", sumCount)
}
func calcRow(path string, ch chan bool) {
	var group sync.WaitGroup
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return nil
		}
		if info.IsDir() {
			return nil
		}
		group.Add(1)
		go func() {
			for _, s := range suffix {
				if strings.HasSuffix(info.Name(), s) {
					calcMap[s].FileCount++
					calcMap[s].Line += calcLine(path)

					break
				}
			}
			group.Done()
		}()
		return nil
	})
	group.Wait()
	if err != nil {
		fmt.Println(err)
	}
	ch <- true
}
func calcLine(path string) int {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return bytes.Count(data, []byte("\n"))
}
