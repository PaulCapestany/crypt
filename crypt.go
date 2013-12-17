package main

import (
	"flag"
	"fmt"
	// "github.com/daviddengcn/go-colortext"
	"os"
	"text/tabwriter"
	// "io/ioutil"
	// "encoding/json"
	// "github.com/____/____"
)

var b = flag.Int("b", 2, "base")
var e = flag.Int("e", 3, "exponent")
var m = flag.Int("m", 7, "mod")

var w = new(tabwriter.Writer)

// ANSI color control escape sequences.
// Shamelessly copied from https://github.com/sqp/godock/blob/master/libs/log/colors.go
var (
	reset      = "\x1b[0m"
	bright     = "\x1b[1m"
	dim        = "\x1b[2m"
	underscore = "\x1b[4m"
	blink      = "\x1b[5m"
	reverse    = "\x1b[7m"
	hidden     = "\x1b[8m"
	fgBlack    = "\x1b[30m"
	fgRed      = "\x1b[31m"
	fgGreen    = "\x1b[32m"
	fgYellow   = "\x1b[33m"
	fgBlue     = "\x1b[34m"
	fgMagenta  = "\x1b[35m"
	fgCyan     = "\x1b[36m"
	fgWhite    = "\x1b[37m"
	bgBlack    = "\x1b[40m"
	bgRed      = "\x1b[41m"
	bgGreen    = "\x1b[42m"
	bgYellow   = "\x1b[43m"
	bgBlue     = "\x1b[44m"
	bgMagenta  = "\x1b[45m"
	bgCyan     = "\x1b[46m"
	bgWhite    = "\x1b[47m"
)

// var secrets = map[string]string{}

// func ReadSecrets() {
// 	someBytes, _ := ioutil.ReadFile("_IGNORE/_secrets.json")
// 	json.Unmarshal(someBytes, &secrets)
// }

func modSeries(base int, exponent int, mod int) {
	// Format right-aligned in space-separated columns of minimal width 5
	// and at least three blanks of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 0, 3, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, underscore+bright+"Equation\tResult\tMod %v\t"+reset+"\n", mod)
	// fmt.Fprintf(w, "--------\t------\t------\t\n")

	value := base
	for i := 1; i <= exponent; i++ {
		modVal := value % mod

		if i%2 == 0 {
			fmt.Fprintf(w, underscore+bright+"%v^%v\t%v\t%v\t"+reset+"\n", base, i, value, modVal)
		} else {
			fmt.Fprintf(w, underscore+dim+"%v^%v\t%v\t%v\t"+reset+"\n", base, i, value, modVal)
		}
		value *= base
		// ct.ResetColor()
	}
	w.Flush()
}

func main() {
	fmt.Println("crypt says hello!\n")
	flag.Parse()

	modSeries(*b, *e, *m)
}
