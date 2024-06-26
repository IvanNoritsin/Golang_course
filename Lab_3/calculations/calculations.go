package calculations

import "fmt"

func Calculate(n int64, useLogging bool) int64 {
	if useLogging {
		fmt.Println("Start calculations...")
		fmt.Printf("Calculate %d!\n", n)
	}

	var result int64 = 1
	for i := int64(1); i <= n; i++ {
		result *= i
	}

	if useLogging {
		fmt.Println("Calculations complete!")
	}

	return result
}
