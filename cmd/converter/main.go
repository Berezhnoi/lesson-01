// Команда converter зчитує суму та курс з аргументів командного рядка
// і виводить результат конвертації, застосовуючи converter.ConvertCurrency.
package main

import (
	"fmt"
	"os"
	"strconv"

	"example.com/lesson01/converter"
)

func parseArgs(args []string) (float64, float64, error) {
	if len(args) != 2 {
		return 0, 0, fmt.Errorf("usage: converter <amount> <rate>\nexample: converter 100 0.91")
	}

	amount, err := strconv.ParseFloat(args[0], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid amount %q: expected a number, for example: converter 100 0.91", args[0])
	}

	rate, err := strconv.ParseFloat(args[1], 64)
	if err != nil {
		return 0, 0, fmt.Errorf("invalid rate %q: expected a number, for example: converter 100 0.91", args[1])
	}

	return amount, rate, nil
}

func main() {
	amount, rate, err := parseArgs(os.Args[1:])
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	result, err := converter.ConvertCurrency(amount, rate)
	if err != nil {
		fmt.Fprintf(os.Stderr, "conversion error: %v\nexample: converter 100 0.91\n", err)
		os.Exit(1)
	}

	fmt.Printf("%.2f = %.2f\n", amount, result)
}
