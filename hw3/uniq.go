package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
)

// type Options struct {
// 	inputChan  io.Reader
// 	outputChan io.Writer
// }

// func newOptinos1(s io.Reader) *Options {
// 	return &Options{
// 		inputChan: s,
// 	}
// }

// func newOptinos2(s1 io.Reader, s2 io.Writer) *Options {
// 	return &Options{
// 		inputChan:  s1,
// 		outputChan: s2,
// 	}
// }

func UniqStrings(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	prevStr := ""
	for in.Scan() {
		txt := in.Text()
		if txt != prevStr {
			output.Write([]byte(txt + "\n"))
		}
		prevStr = txt
	}
	return nil
}

func cFlag(input io.Reader, output io.Writer) {
	in := bufio.NewScanner(input)
	prevStr := ""
	cnt := 1
	for in.Scan() {
		txt := in.Text()
		if txt != prevStr {
			output.Write([]byte(strconv.Itoa(cnt) + " " + txt + "\n"))
			cnt = 0
		}
		cnt++
		prevStr = txt
	}
}

func main() {
	// var cFlag *int
	// var dFlag, uFlag, f_num_fieldsFlag, s_num_charsFlag, iFlag string
	// flag.IntVar(cFlag, "c", 0, "count of every string")
	// flag.Func("c", "Count of every string in input", func(s string) error {

	// })

	flag.Parse()

	cntPosArgs := len(flag.Args())
	if cntPosArgs == 2 {
		inFile, errin := os.Open(flag.Arg(0))
		outFile, errout := os.OpenFile(flag.Arg(1), os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if errin != nil {
			log.Fatal("Error opening input file")
		}
		if errout != nil {
			log.Fatal("Error opening output file")
		}
		defer inFile.Close()
		defer outFile.Close()
		UniqStrings(inFile, outFile)
	}

	if cntPosArgs == 1 {
		inFile, errin := os.Open(flag.Arg(0))
		if errin != nil {
			log.Fatal("Error opening input file")
		}
		defer inFile.Close()
		UniqStrings(inFile, os.Stdout)
	}

	UniqStrings(os.Stdin, os.Stdout)

}
