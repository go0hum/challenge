package file

import (
	"challenge/csv/account"
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadFile() account.AccountFile {
	var transactionsFile []account.TransactionFile
	var monthsFile account.MonthsFile = make(account.MonthsFile)
	file, err := os.Open("../../csv/txns.csv")
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

	var total, debits, credits float32
	total, debits, credits = 0.0, 0.0, 0.0
	count := 0
	for i, row := range records {
		if i == 0 {
			continue
		}
		value, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Printf("Error parsing amount in row %d: %v\n", i, err)
			continue
		}

		amount := float32(value)

		// obtener el mes
		partes := strings.Split(row[1], "/")

		if len(partes) >= 1 {
			mesInt, err := strconv.Atoi(partes[0])

			if err != nil {
				fmt.Printf("Error parsing month in row %d: %v\n", i, err)
				continue
			}

			monthsFile[meses[mesInt]]++
			count++

		} else {
			fmt.Println("Error: No se encontro la fecha")
			continue
		}

		partesDate := strings.Split(row[1], "/")
		if len(partes) != 2 {
			fmt.Printf("Error: Fecha inválida en la fila %d\n", i)
			continue
		}

		mesInt, err := strconv.Atoi(partesDate[0])
		if err != nil {
			fmt.Printf("Error parsing month in row %d: %v\n", i, err)
			continue
		}

		day, err := strconv.Atoi(partesDate[1])
		if err != nil {
			fmt.Printf("Error parsing year in row %d: %v\n", i, err)
			continue
		}

		finalDate := time.Date(time.Now().Year(), time.Month(mesInt), day, 0, 0, 0, 0, time.UTC)

		transactionsFile = append(transactionsFile, account.TransactionFile{
			Transaction: amount,
			Date:        finalDate,
		})

		if amount < 0 {
			debits = debits + amount
		} else {
			credits = credits + amount
		}
		total = total + amount
	}

	return account.AccountFile{
		Name:         "txns.csv",
		Total:        total,
		Debit:        debits / float32(len(monthsFile)),
		Credit:       credits / float32(len(monthsFile)),
		Transactions: transactionsFile,
		Months:       monthsFile,
	}
}
