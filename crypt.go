package main

import (
	"flag"
	"fmt"
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
	fmt.Fprintf(w, "Equation\tResult\tMod %v\t\n", mod)
	fmt.Fprintf(w, "--------\t------\t------\t\n")

	value := base
	for i := 1; i <= exponent; i++ {
		modVal := value % mod
		fmt.Fprintf(w, "%v^%v\t%v\t%v\t\n", base, i, value, modVal)
		value *= base
	}
}

func main() {
	fmt.Println("crypt says hello!\n")
	flag.Parse()

	modSeries(*b, *e, *m)
	w.Flush()
}
