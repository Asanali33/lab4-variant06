package main

import (
	"fmt"

	"github.com/fatih/color"
	"github.com/google/uuid"

	"lab4-variant06/pkg/waterbill"
)

func main() {

	owner := "Иванов И.И."
	prev := 120.5
	curr := 135.8
	tariff := 45.30
	penalty := 10.0

	// Генерация ID отчёта через внешний пакет
	reportID := uuid.New()

	// 1. Расход
	cubic, err := waterbill.WaterUsage(prev, curr)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Расход воды: %.3f м³\n", cubic)

	// 2. Стоимость
	cost, err := waterbill.WaterCost(cubic, tariff)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Стоимость без штрафа: %.2f руб.\n", cost)

	// 3. Применение штрафа
	err = waterbill.ApplyPenalty(&cost, penalty)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Printf("Стоимость со штрафом: %.2f руб.\n", cost)

	// 4. Формирование отчёта
	report, err := waterbill.FormatWaterReport(owner, cubic, cost)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}

	color.Cyan("Report ID: %s", reportID.String())
	color.Green(report)
}
