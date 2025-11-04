package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculateAndPrint(total, step, daysOff, currentPace, prevCalendarDays float64) float64 {
	workingDays := math.Ceil(total / step)

	var totalCalendarDays float64
	var daysSavedStr string

	workingDaysPerWeek := 7.0 - daysOff

	if workingDaysPerWeek <= 0 {
		totalCalendarDays = math.Inf(1)
	} else if daysOff == 0 {
		totalCalendarDays = workingDays
	} else {
		fullWeeks := math.Floor(workingDays / workingDaysPerWeek)
		remainingWorkDays := math.Mod(workingDays, workingDaysPerWeek)

		if remainingWorkDays == 0 && workingDays > 0 {
			totalCalendarDays = (workingDays / workingDaysPerWeek) * 7
		} else {
			totalCalendarDays = (fullWeeks * 7) + remainingWorkDays
		}
	}

	if math.IsInf(prevCalendarDays, 1) {
		daysSavedStr = "---"
	} else {
		saved := prevCalendarDays - totalCalendarDays
		if math.IsInf(saved, 1) {
			daysSavedStr = "ОГРОМНАЯ"
		} else {
			daysSavedStr = fmt.Sprintf("%.0f д.", saved)
		}
	}

	var daysStr string
	if math.IsInf(totalCalendarDays, 1) {
		daysStr = "бесконечно"
	} else {
		daysStr = fmt.Sprintf("%.0f", totalCalendarDays)
	}

	marker := ""
	if step == currentPace {
		marker = " <-- ВАШ ТЕМП"
	}

	fmt.Printf("  Шаг (в день): %-10.0f | Календ. дней: %-12s | Экономия: %-10s %s\n", step, daysStr, daysSavedStr, marker)

	return totalCalendarDays
}

func readFloat(prompt string) (float64, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}
		trimmedInput := strings.TrimSpace(input)
		value, err := strconv.ParseFloat(trimmedInput, 64)
		if err != nil {
			fmt.Println("Ошибка: Введите число.")
			continue
		}
		if value < 0 {
			fmt.Println("Ошибка: Введено некорректное положительное число.")
			continue
		}
		return value, nil
	}
}

func readInt(prompt string, min, max int) (int, error) {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print(prompt)
		input, err := reader.ReadString('\n')
		if err != nil {
			return 0, err
		}
		trimmedInput := strings.TrimSpace(input)
		value, err := strconv.Atoi(trimmedInput)
		if err != nil {
			fmt.Println("Ошибка: Введите число.")
			continue
		}
		if value < min || value > max {
			fmt.Printf("Ошибка: Введите число от %d до %d.\n", min, max)
			continue
		}
		return value, nil
	}
}

func main() {
	fmt.Println("--- Калькулятор прогресса и 'точки баланса' ---")

	total, err := readFloat("Введите общий объем (например, я хочу читать 500 страниц в день): ")
	if err != nil {
		return
	}

	currentPace, err := readFloat("Введите ваш текущий темп 'шага в день' (0, если нет): ")
	if err != nil {
		return
	}

	daysOff, err := readInt("Введите кол-во ВЫХОДНЫХ дней в неделю (0-6): ", 0, 6)
	if err != nil {
		return
	}

	fmt.Println("\n==================================================================")
	fmt.Printf("Расчет для объема: %.0f. Выходных в неделю: %d.\n", total, daysOff)
	fmt.Println("==================================================================\n")

	prevCalendarDays := math.Inf(1)

	fmt.Println("--- Маленькие шаги (с 1 до 10) ---")
	limit := math.Min(10, total)
	for step := 1.0; step <= limit; step++ {
		prevCalendarDays = calculateAndPrint(total, step, float64(daysOff), currentPace, prevCalendarDays)
	}

	if total > 10 {
		fmt.Println("\n--- Средние шаги (с 12 до 30) ---")
		for step := 12.0; step <= 30 && step <= total; step += 2 {
			prevCalendarDays = calculateAndPrint(total, step, float64(daysOff), currentPace, prevCalendarDays)
		}
	}

	if total > 30 {
		fmt.Println("\n--- Средне-большие шаги (с 35 до 50) ---")
		for step := 35.0; step <= 50 && step <= total; step += 5 {
			prevCalendarDays = calculateAndPrint(total, step, float64(daysOff), currentPace, prevCalendarDays)
		}
	}

	if total > 50 {
		fmt.Println("\n--- Большие шаги (с 60 до 100) ---")
		for step := 60.0; step <= 100 && step <= total; step += 10 {
			prevCalendarDays = calculateAndPrint(total, step, float64(daysOff), currentPace, prevCalendarDays)
		}
	}

	fmt.Println("\n==================================================================")
	fmt.Println("Изучите колонку 'Экономия'.")
	fmt.Println("Где экономия становится 1 или 0 дней - ваш 'баланс' достигнут.")
}
