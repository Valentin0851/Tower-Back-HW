package main

import (
	"bufio"
	"flag"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

// type Options struct {
//   inputChan  io.Reader
//   outputChan io.Writer
// }

// func newOptinos1(s io.Reader) *Options {
//   return &Options{
//     inputChan: s,
//   }
// }

// func newOptinos2(s1 io.Reader, s2 io.Writer) *Options {
//   return &Options{
//     inputChan:  s1,
//     outputChan: s2,
//   }
// }

func UniqStrings(dataSlice, other []string, output io.Writer) {
	prevStr := ":wq!"
	cnt := 0
	for i, val := range other {
		if val != prevStr && prevStr != ":wq!" {
			output.Write([]byte(dataSlice[i-cnt] + "\n"))
			cnt = 0
		}
		prevStr = val
		cnt++
	}
	if len(dataSlice) > 2 && other[len(dataSlice)-1] != other[len(dataSlice)-2] {
		output.Write([]byte(dataSlice[len(dataSlice)-1]))
	} else {
		output.Write([]byte(dataSlice[len(dataSlice)-cnt]))
	}
}

func UniqStringsWithc(dataSlice, other []string, output io.Writer) {
	cnt := 0
	prevStr := ":wq!"
	for i, val := range other {
		if val != prevStr && prevStr != ":wq!" {
			output.Write([]byte(strconv.Itoa(cnt) + " " + dataSlice[i-cnt] + "\n"))
			cnt = 0
		}
		prevStr = val
		cnt++
	}
	if other[len(other)-1] != other[len(other)-2] {
		output.Write([]byte("1" + " " + dataSlice[len(dataSlice)-1] + "\n"))
	} else {
		output.Write([]byte(strconv.Itoa(cnt) + " " + dataSlice[len(dataSlice)-cnt] + "\n"))
	}
}

func UniqStringsWithd(dataSlice, other []string, output io.Writer) {
	cnt := 0
	prevStr := ":wq!"
	for i, val := range other {
		if val != prevStr && prevStr != ":wq!" {
			if cnt > 1 {
				output.Write([]byte(dataSlice[i-cnt] + "\n"))
			}
			cnt = 0
		}
		prevStr = val
		cnt++
	}
	if cnt > 1 && other[len(dataSlice)-cnt] == prevStr {
		output.Write([]byte(prevStr + "\n"))
	}
}

func UniqStringsWithu(dataSlice, other []string, output io.Writer) {
	cnt := 0
	prevStr := ":wq!"
	for i, val := range other {
		if val != prevStr && prevStr != ":wq!" {
			if cnt == 1 {
				output.Write([]byte(dataSlice[i-cnt] + "\n"))
			}
			cnt = 0
		}
		prevStr = val
		cnt++
	}
	if other[len(dataSlice)-1] != other[len(dataSlice)-2] {
		output.Write([]byte(dataSlice[len(dataSlice)-1] + "\n"))
	}
}

func UniqStringsWithi(dataSlice []string) []string {
	var ans []string
	for _, val := range dataSlice {
		ans = append(ans, strings.ToLower(val))
	}
	return ans
}

func UniqStringsWithNumFields(other []string, numFields int) []string {
	var ans []string
	for _, val := range other {
		lineSplit := strings.Fields(val)
		if len(lineSplit) > numFields {
			ans = append(ans, strings.Join(lineSplit[numFields:], " "))
		} else {
			ans = append(ans, "")
		}
	}
	return ans
}

func UniqStringsWithNumChars(other []string, numChars int) []string {
	var ans []string
	for _, val := range other {
		if len(val) > numChars {
			line := ""
			for i := range val {
				if i >= numChars {
					line += string(val[i])
				}
			}
			ans = append(ans, line)
		} else {
			ans = append(ans, "")
		}
	}
	return ans
}

func readDataFromChan(input io.Reader) []string {
	var ans []string
	in := bufio.NewScanner(input)
	for in.Scan() {
		ans = append(ans, in.Text())
	}
	return ans
}

// func writeDataToChan(dataSlice []string, output io.Writer){
//   for _,val := range dataSlice{
//     output.Write([]byte(val+"\n"))
//   }
// }

func main() {
	cFlag := flag.Bool("c", false, "count of every string") // print string with number of repetition
	dFlag := flag.Bool("d", false, "only repeated strings") // print repeated strings
	uFlag := flag.Bool("u", false, "uniq strings")          // print uniq strings
	iFlag := flag.Bool("i", false, "Ignore case")
	numFieldsFlag := flag.Int("f", 0, "number of fiels which i need to miss")
	numCharsFlag := flag.Int("s", 0, "number of chars without spaces which i need miss")

	flag.Parse()
	var other []string

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
		dataSlice := readDataFromChan(inFile)
		if *iFlag {
			other = UniqStringsWithi(other)
		}
		if *numFieldsFlag != 0 {
			other = UniqStringsWithNumFields(other, *numFieldsFlag)
		}
		if *numCharsFlag != 0 {
			other = UniqStringsWithNumChars(other, *numCharsFlag)
		}
		if *cFlag || *dFlag || *uFlag {
			if *cFlag {
				UniqStringsWithc(dataSlice, other, outFile)
				return
			}
			if *dFlag {
				UniqStringsWithd(dataSlice, other, outFile)
				return
			}
			if *uFlag {
				UniqStringsWithu(dataSlice, other, outFile)
				return
			}
		} else {
			UniqStrings(dataSlice, other, outFile)
			return
		}
	}

	if cntPosArgs == 1 {
		inFile, errin := os.Open(flag.Arg(0))
		if errin != nil {
			log.Fatal("Error opening input file")
		}
		defer inFile.Close()
		dataSlice := readDataFromChan(inFile)
		other = dataSlice
		if *iFlag {
			other = UniqStringsWithi(other)
		}
		if *numFieldsFlag != 0 {
			other = UniqStringsWithNumFields(other, *numFieldsFlag)
		}
		if *numCharsFlag != 0 {
			other = UniqStringsWithNumChars(other, *numCharsFlag)
		}
		if *cFlag || *dFlag || *uFlag {
			if *cFlag {
				UniqStringsWithc(dataSlice, other, os.Stdout)
				return
			}
			if *dFlag {
				UniqStringsWithd(dataSlice, other, os.Stdout)
				return
			}
			if *uFlag {
				UniqStringsWithu(dataSlice, other, os.Stdout)
				return
			}
		} else {
			UniqStrings(dataSlice, other, os.Stdout)
			return
		}
	}

	dataSlice := readDataFromChan(os.Stdin)
	other = dataSlice
	if *iFlag {
		other = UniqStringsWithi(other)
	}
	if *numFieldsFlag != 0 {
		other = UniqStringsWithNumFields(other, *numFieldsFlag)
	}
	if *numCharsFlag != 0 {
		other = UniqStringsWithNumChars(other, *numCharsFlag)
	}
	if *cFlag || *dFlag || *uFlag {
		if *cFlag {
			UniqStringsWithc(dataSlice, other, os.Stdout)
			return
		}
		if *dFlag {
			UniqStringsWithd(dataSlice, other, os.Stdout)
			return
		}
		if *uFlag {
			UniqStringsWithu(dataSlice, other, os.Stdout)
			return
		}
	} else {
		UniqStrings(dataSlice, other, os.Stdout)
		return
	}
}
