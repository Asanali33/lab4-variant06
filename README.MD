# Lab4 - WaterBill (Вариант 6)

## Описание
Проект реализует пакет `waterbill` для расчёта расхода воды, стоимости и формирования отчёта.  
Пакет содержит функции:

1. `WaterUsage(prev, curr float64) (float64, error)` — вычисляет расход воды.
2. `WaterCost(cubic, tariff float64) (float64, error)` — вычисляет стоимость.
3. `ApplyPenalty(cost *float64, penaltyPercent float64) error` — применяет штраф к стоимости.
4. `FormatWaterReport(owner string, cubic, cost float64) (string, error)` — формирует строку отчёта.

Все функции документированы и используют PascalCase.

---

## Структура проекта
