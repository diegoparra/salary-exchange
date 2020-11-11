package generate

import (
	"fmt"
	"time"
)
// Hours define the worked hours per day
const Hours float64 = 8

// ValuePerHour define amount earned per hour
const ValuePerHour float64 = 15.0

// Generate make the calculation and return the total salary in Euro / REAL
func Generate(m map[string]int, n map[string]int, e float64) {
	month := time.Now().Month().String()
	for k, v := range m {
		if k == month {
			fmt.Println("Month:", k)
			fmt.Println("Total business days:", v)

			totalHours := float64(v) * Hours
			fmt.Println("Total worked hours:", totalHours)

			totalSalaryEuro := totalHours * ValuePerHour
			fmt.Println("Total value in EURO:", totalSalaryEuro)

			totalSalaryReal := totalSalaryEuro * e
			fmt.Printf("Total value in REAL: %.2f\n", totalSalaryReal)

			fmt.Println("Client name: Client ABC")
			fmt.Println("Client Address: address")
			fmt.Println("Activity code: 108098")
			fmt.Println("Description. Month:", n[month], "/2020, INVOICE No00XX. total Value: €", totalSalaryEuro, "today one Euro value in real: R$", e)

		}
	}
}
