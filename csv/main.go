package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("txns.csv")
	if err != nil {
		panic(err)
	}

	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	meses := map[int]string{
		1:  "Enero",
		2:  "Febrero",
		3:  "Marzo",
		4:  "Abril",
		5:  "Mayo",
		6:  "Junio",
		7:  "Julio",
		8:  "Agosto",
		9:  "Septiembre",
		10: "Octubre",
		11: "Noviembre",
		12: "Diciembre",
	}

	total, debits, credits := 0.0, 0.0, 0.0
	count := make(map[string]int)
	for i, row := range records {
		if i == 0 {
			continue
		}
		amount, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Printf("Error parsing amount in row %d: %v\n", i, err)
			continue
		}

		partes := strings.Split(row[1], "/")

		if len(partes) >= 1 {
			mesInt, err := strconv.Atoi(partes[0])

			if err != nil {
				fmt.Printf("Error parsing month in row %d: %v\n", i, err)
				continue
			}

			monthName, ok := meses[mesInt]
			if !ok {
				fmt.Printf("Error: Mes no encontrado en la fila %d\n", i)
				continue
			}

			count[monthName]++

		} else {
			fmt.Println("Error: No se encontro la fecha")
			continue
		}

		if amount < 0 {
			debits = debits + amount
		} else {
			credits = credits + amount
		}
		total = total + amount
	}
	fmt.Println("Total:", total)

	for mes, total := range count {
		fmt.Printf("Number of transactions in %s: %d\n", mes, total)
	}

	fmt.Println("Average debit amount:", debits/float64(len(count)))
	fmt.Println("Average credit amount:", credits/float64(len(count)))
}
