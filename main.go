package go_aussie_payslip

import (
	"fmt"
	"encoding/csv"
	"os"
	"bufio"
	"io"
)

func main()  {
    	f, _ := os.Open('aussie_payroll.csv')
	r := csv.NewReader(bufio.NewReader(f))
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}

	}
}