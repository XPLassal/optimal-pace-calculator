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
			daysSavedStr = "HUGE"
		} else {
			daysSavedStr = fmt.Sprintf("%.0f d.", saved)
		}
	}

	var daysStr string
	if math.IsInf(totalCalendarDays, 1) {
		daysStr = "infinite"
	} else {
		daysStr = fmt.Sprintf("%.0f", totalCalendarDays)
	}

	marker := ""
	if step == currentPace {
		marker = " <-- YOUR PACE"
	}

	fmt.Printf("  Step (per day): %-10.0f | Calendar Days: %-12s | Days Saved: %-10s %s\n", step, daysStr, daysSavedStr, marker)

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
			fmt.Println("Error: Please enter a valid number.")
			continue
		}
		if value < 0 {
			fmt.Println("Error: Please enter a positive number or 0.")
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
			fmt.Println("Error: Please enter a number.")
			continue
		}
		if value < min || value > max {
			fmt.Printf("Error: Please enter a number between %d and %d.\n", min, max)
			continue
		}
		return value, nil
	}
}

func main() {
	fmt.Println("--- Progress and 'Balance Point' Calculator ---")

	total, err := readFloat("Enter total volume (e.g., I want to read 500 pages in a day): ")
	if err != nil {
		return
	}

	currentPace, err := readFloat("Enter your current pace 'step per day' (0 if none): ")
	if err != nil {
		return
	}

	daysOff, err := readInt("Enter number of OFF-DAYS per week (0-6): ", 0, 6)
	if err != nil {
		return
	}

	fmt.Println("\n==================================================================")
	fmt.Printf("Calculation for volume: %.0f. Off-days per week: %d.\n", total, daysOff)
	fmt.Println("==================================================================\n")

	prevCalendarDays := math.Inf(1)

	fmt.Println("--- Small Steps (1 to 10) ---")
	limit := math.Min(10, total)
	for step := 1.0; step <= limit; step++ {
		prevCalendarDays = calculateAndPrint(total, step, float64(daysOff), currentPace, prevCalendarDays)
	}

	if total > 10 {
		fmt.Println("\n--- Medium Steps (12 to 30) ---")
		for step := 12.0; step <= 30 && step <= total; step += 2 {
			prevCalendarDays = calculateAndPrint(total, step, float64(daysOff), currentPace, prevCalendarDays)
		}
	}

	if total > 30 {
		fmt.Println("\n--- Medium-Large Steps (35 to 50) ---")
		for step := 35.0; step <= 50 && step <= total; step += 5 {
			prevCalendarDays = calculateAndPrint(total, step, float64(daysOff), currentPace, prevCalendarDays)
		}
	}

	if total > 50 {
		fmt.Println("\n--- Large Steps (60 to 100) ---")
		for step := 60.0; step <= 100 && step <= total; step += 10 {
			prevCalendarDays = calculateAndPrint(total, step, float64(daysOff), currentPace, prevCalendarDays)
		}
	}

	fmt.Println("\n==================================================================")
	fmt.Println("Study the 'Days Saved' column.")
	fmt.Println("Your 'balance point' is where the savings become 1 or 0 days.")
}
