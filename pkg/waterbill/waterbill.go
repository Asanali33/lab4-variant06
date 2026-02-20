// Package waterbill предоставляет функции для расчёта
// потребления воды, стоимости и формирования отчёта.
package waterbill

import (
	"fmt"
)

// WaterUsage вычисляет расход воды по предыдущему и текущему показанию счётчика.
// Возвращает ошибку, если значения отрицательные или текущий показатель меньше предыдущего.
func WaterUsage(prev, curr float64) (float64, error) {
	if prev < 0 || curr < 0 {
		return 0, fmt.Errorf("показания не могут быть отрицательными")
	}
	if curr < prev {
		return 0, fmt.Errorf("текущее показание меньше предыдущего")
	}
	return curr - prev, nil
}

// WaterCost вычисляет стоимость воды.
// Возвращает ошибку, если расход или тариф отрицательные.
func WaterCost(cubic, tariff float64) (float64, error) {
	if cubic < 0 {
		return 0, fmt.Errorf("расход не может быть отрицательным")
	}
	if tariff <= 0 {
		return 0, fmt.Errorf("тариф должен быть больше нуля")
	}
	return cubic * tariff, nil
}

// ApplyPenalty увеличивает стоимость на указанный процент штрафа.
// Использует указатель для изменения исходного значения.
// Возвращает ошибку при некорректном проценте.
func ApplyPenalty(cost *float64, penaltyPercent float64) error {
	if cost == nil {
		return fmt.Errorf("nil указатель на стоимость")
	}
	if penaltyPercent < 0 {
		return fmt.Errorf("процент штрафа не может быть отрицательным")
	}
	*cost += *cost * penaltyPercent / 100
	return nil
}

// FormatWaterReport формирует строку отчёта по воде.
// Возвращает ошибку, если входные данные некорректны.
func FormatWaterReport(owner string, cubic, cost float64) (string, error) {
	if owner == "" {
		return "", fmt.Errorf("имя владельца не может быть пустым")
	}
	if cubic < 0 || cost < 0 {
		return "", fmt.Errorf("некорректные данные для отчёта")
	}

	report := fmt.Sprintf(
		"Отчёт по воде\nВладелец: %s\nРасход: %.3f м³\nИтоговая стоимость: %.2f руб.\n",
		owner, cubic, cost,
	)

	return report, nil
}
