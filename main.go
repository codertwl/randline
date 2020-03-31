package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"time"
)

func main() {
	inFile := ""
	outFile := ""
	inMax := 10
	outMax := 100

	flag.StringVar(&inFile, "i", "", "input file full path")
	flag.StringVar(&outFile, "o", "./tmp/out.txt", "output file full path")
	flag.IntVar(&inMax, "a", 10, "read max count")
	flag.IntVar(&outMax, "b", 100, "random write max count")
	flag.Parse()

	index := strings.LastIndex(outFile, "/")
	if index >= 0 {
		err := os.MkdirAll(outFile[:index], os.ModePerm)
		if err != nil {
			fmt.Println("mkdir outpath err:", err)
			return
		}
	}

	randLine(inFile, outFile, inMax, outMax)
}

func randLine(inFile, outFile string, inMax, outMax int) {

	inFp, err := os.OpenFile(inFile, os.O_RDONLY, 0666)
	if err != nil {
		panic("open inFile err:" + err.Error())
	}
	defer inFp.Close()

	rBuf := bufio.NewReader(inFp)
	lines := make([]string, 0, inMax)
	for {
		line, err := rBuf.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("read inFile over...")
				break
			} else {
				panic("read inFile line err:" + err.Error())
			}
		}

		lines = append(lines, line)
		if len(lines) >= inMax {
			break
		}
	}

	outFp, err := os.Create(outFile)
	if err != nil {
		panic("create outFile err:" + err.Error())
	}
	defer outFp.Close()

	wBuf := bufio.NewWriter(outFp)
	rand.Seed(time.Now().Unix())
	for outMax > 0 {
		index := rand.Intn(len(lines))
		wBuf.WriteString(lines[index])
		outMax--
	}
	wBuf.Flush()
}
