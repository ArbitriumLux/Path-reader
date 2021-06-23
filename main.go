package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"sync"
	"sync/atomic"
)

var InCh chan string = make(chan string)
var OutCh chan string = make(chan string)
var FinalSize int64

func FilepathScan() {
	path := flag.String("path", " . ", "A path")
	flag.Parse()
	err := filepath.Walk(*path,
		func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil
			}
			if err != nil {
				return err
			}
			InCh <- path
			return nil
		})
	if err != nil {
		log.Println(err)
	}
	close(InCh)
}

func Size() {
	for path := range InCh {
		result := path + "\t"
		fi, err := os.Stat(path)
		if err != nil {
			result += "ERROR"
			log.Println(err)
		} else {
			size := fi.Size()
			atomic.AddInt64(&FinalSize, size)
			result += strconv.Itoa(int(size))
		}
		OutCh <- result
	}
}

func Finalizer() {
	file, err := os.Create("result.txt")
	if err != nil {
		log.Fatal("Cannot create file", err)
	}
	defer file.Close()
	for out := range OutCh {
		_, err := file.WriteString(out + "\n")
		if err != nil {
			log.Fatal("Cannot write to file", err)
		}
	}
	_, err = file.WriteString(strconv.Itoa(int(FinalSize)))
	if err != nil {
		log.Fatal("Cannot write to file", err)
	}
}

func main() {
	var wgIn sync.WaitGroup
	for i := 0; i < 3; i++ {
		wgIn.Add(1)
		go func() {
			defer wgIn.Done()
			Size()
		}()
	}
	var wgOut sync.WaitGroup
	wgOut.Add(1)
	go func() {
		defer wgOut.Done()
		Finalizer()
	}()
	FilepathScan()
	wgIn.Wait()
	close(OutCh)
	wgOut.Wait()
}
