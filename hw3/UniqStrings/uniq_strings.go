package UniqStrings //чет с импортом не получилось :((((
import (
	"bufio"
	"fmt"
	"io"
)

func UniqStrings(input io.Reader, output io.Writer) error {
	in := bufio.NewScanner(input)
	var prevStr string
	for in.Scan() {
		txt := in.Text()
		if txt != prevStr {
			fmt.Fprintln(output, txt)
		}
		prevStr = txt
	}
	return nil
}
