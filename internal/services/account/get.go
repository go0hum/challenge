package account

import (
	domain "challenge/internal/domain/account"
	"context"
	"encoding/csv"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

func (s *Service) Get(ctx context.Context) (domain.AccountFile, error) {
	var transactionsFile []domain.TransactionFile
	var monthsFile domain.MonthsFile = make(domain.MonthsFile)

	data, err := s.Account.ReadCSV(ctx)
	if err != nil {
		fmt.Println(err)
	}
	defer data.Body.Close()

	reader := csv.NewReader(data.Body)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV:", err)
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
		fmt.Printf("Fila %d: %v", i, row)
		if i == 0 {
			continue
		}
		value, err := strconv.ParseFloat(row[2], 64)
		if err != nil {
			fmt.Printf("Error parsing amount in row %d: %v\n", i, err)
			continue
		}

		amount := float32(value)

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

		mesInt, err := strconv.Atoi(partes[0])
		if err != nil {
			fmt.Printf("Error parsing month in row %d: %v\n", i, err)
			continue
		}

		day, err := strconv.Atoi(partes[1])
		if err != nil {
			fmt.Printf("Error parsing year in row %d: %v\n", i, err)
			continue
		}

		finalDate := time.Date(time.Now().Year(), time.Month(mesInt), day, 0, 0, 0, 0, time.UTC)

		transactionsFile = append(transactionsFile, domain.TransactionFile{
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

	totalFloat32 := float32(math.Round(float64(total*100)) / 100)
	debitFloat32 := float32(math.Round(float64(debits/float32(len(monthsFile))*100)) / 100)
	creditFloat32 := float32(math.Round(float64(credits/float32(len(monthsFile))*100)) / 100)

	return domain.AccountFile{
		Name:         "txns.csv",
		Total:        totalFloat32,
		Debit:        debitFloat32,
		Credit:       creditFloat32,
		Transactions: transactionsFile,
		Months:       monthsFile,
	}, nil
}
