package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"text/tabwriter"
	// "io/ioutil"
	// "encoding/json"
	// "github.com/____/____"
)

var A = flag.String("A", "3", "Alice's secret")
var B = flag.String("B", "6", "Bob's secret")
var Y = flag.String("Y", "2", "base")
var e = flag.String("e", "3", "exponent")
var P = flag.String("P", "11", "mod P")

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

func aliceAndBob(aliceSecret *big.Int, bobSecret *big.Int, Y *big.Int, P *big.Int) {
	fmt.Println(fgCyan + bright + underscore + "Paul's 1st Crypto" + reset)
	fmt.Println(fgCyan + "  ↳ Y^x(mod P)" + reset)
	fmt.Printf(fgGreen+"(Public)\n  Y = %v\n  P = %v\n", Y, P)
	fmt.Printf(fgRed+"(Private)\n  A = %v\n  B = %v\n", aliceSecret, bobSecret)

	// 1234567890123456789 → limit of 19 digits for normal int64
	// 100100100100100100100
	// 101101101101101101101
	//
	// Y must be smaller than P

	fmt.Printf(fgYellow + "(Results)\n")
	alpha := new(big.Int).Exp(Y, aliceSecret, P)
	fmt.Printf("  α = %v\n", alpha)

	beta := new(big.Int).Exp(Y, bobSecret, P)
	fmt.Printf("  β = %v\n", beta)

	aliceKey := new(big.Int).Exp(beta, aliceSecret, P)
	fmt.Printf(bright+"KEY = %v\n", aliceKey)

	// bobKey := new(big.Int).Exp(alpha, bobSecret, P)
	// fmt.Printf("B Key: %v\n", bobKey)
}

func fromBase10(base10 string) *big.Int {
	i := new(big.Int)
	i.SetString(base10, 10)
	return i
}

func tryRSA() {
	b_A := fromBase10("3490529510847650949147849619903898133417764638493387843990820577")
	b_B := fromBase10("32769132993266709549961988190834461413177642967992942539798288533")
	b_C := new(big.Int).Mul(b_A, b_B)
	fmt.Printf("b_A = %v\nb_B = %v\nb_C = %v\n", b_A, b_B, b_C)
}

func main() {
	// fmt.Println("crypt says hello!\n")
	flag.Parse()
	zA := fromBase10(*A)
	zB := fromBase10(*B)
	zY := fromBase10(*Y)
	zP := fromBase10(*P)
	// modSeries(*b, *e, *m)
	aliceAndBob(zA, zB, zY, zP)
	// tryRSA()
}
