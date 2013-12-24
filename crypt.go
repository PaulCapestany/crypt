package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"text/tabwriter"
)

// Diffie-Hellman Key Exchange flags
var A = flag.String("A", "3", "Alice's secret")
var B = flag.String("B", "6", "Bob's secret")
var Y = flag.String("Y", "2", "base")
var P = flag.String("P", "11", "mod P")

// RSA flags
var p = flag.String("p", "17", "p")
var q = flag.String("q", "11", "q")
var e = flag.String("e", "7", "e")
var M = flag.String("M", "88", "M")

var w = new(tabwriter.Writer)

// ANSI color control escape sequences.
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

func modSeries(base int, exponent int, mod int) {
	// Format right-aligned in space-separated columns of minimal width 5
	// and at least three blanks of padding (so wider column entries do not
	// touch each other).
	w.Init(os.Stdout, 5, 0, 3, ' ', tabwriter.AlignRight)
	fmt.Fprintf(w, underscore+bright+"Equation\tResult\tMod %v\t"+reset+"\n", mod)

	value := base
	for i := 1; i <= exponent; i++ {
		modVal := value % mod

		if i%2 == 0 {
			fmt.Fprintf(w, underscore+bright+"%v^%v\t%v\t%v\t"+reset+"\n", base, i, value, modVal)
		} else {
			fmt.Fprintf(w, underscore+dim+"%v^%v\t%v\t%v\t"+reset+"\n", base, i, value, modVal)
		}
		value *= base
	}
	w.Flush()
}

func diffieHellmanKey() {
	// Diffie-Hellman flags
	Y := fromBase10(*Y)
	P := fromBase10(*P)
	A := fromBase10(*A)
	B := fromBase10(*B)

	fmt.Println(fgCyan + bright + underscore + "Paul's 1st Crypto" + reset)
	fmt.Println(fgCyan + "  ↳ Y^x(mod P)" + reset)
	fmt.Printf(fgGreen+"(Public)\n  Y = %v\n  P = %v\n", Y, P)
	fmt.Printf(fgRed+"(Private)\n  A = %v\n  B = %v\n", A, B)

	// Y must be smaller than P
	fmt.Printf(fgYellow + "(Results)\n")
	alpha := new(big.Int).Exp(Y, A, P)
	fmt.Printf("  α = %v\n", alpha)

	beta := new(big.Int).Exp(Y, B, P)
	fmt.Printf("  β = %v\n", beta)

	aliceKey := new(big.Int).Exp(beta, A, P)
	fmt.Printf(bright+"KEY = %v\n", aliceKey)
}

func fromBase10(base10 string) *big.Int {
	i := new(big.Int)
	i.SetString(base10, 10)
	return i
}

func tryRSA() {
	// RSA flags
	p := fromBase10(*p)
	q := fromBase10(*q)
	e := fromBase10(*e)
	M := fromBase10(*M)

	// Alice must pick two prime numbers (these are SECRET)
	fmt.Printf("p = %v\n", p)
	fmt.Printf("q = %v\n", q)

	// Alice must pick a number for 'e'
	// 'e' should be 1 < e < Φ
	fmt.Printf("e = %v\n", e)

	// 'e', along with 'N', are Alice's public key
	N := new(big.Int).Mul(p, q)
	fmt.Printf("N = %v\n", N)

	// Calculate the Φ with: (p-1)*(q-1)
	var bigOne = big.NewInt(1)
	Φ := new(big.Int).Mul(new(big.Int).Sub(p, bigOne), new(big.Int).Sub(q, bigOne))
	fmt.Printf("Φ = %v\n", Φ)

	// To encrypt a message, the message must first be converted into a number, 'M'
	// 'M' *must* be smaller than 'N'
	// Text is changed into ASCII binary digits as 'M' which then gives ciphertext 'C'
	fmt.Printf("M = %v\n", M)
	C := new(big.Int).Exp(M, e, N)
	fmt.Printf("C = %v\n", C)

	// Calculate decryption key, 'd'
	d := new(big.Int).ModInverse(e, Φ)
	fmt.Printf("d = %v\n", d)

	// To decrypt message use: C^d(mod N)
	W := new(big.Int).Exp(C, d, N)
	fmt.Printf("W = %v\n", W)
}

func main() {
	flag.Parse()

	// diffieHellmanKey()
	tryRSA()
}
