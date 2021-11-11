package hyper_interactive

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

var (
	TermWidth = 80
)

func printSelect(options []string, startFrom int) {
	optionDisplay := make([]string, len(options))
	for idx, str := range options {
		optionDisplay[idx] = fmt.Sprintf("(%d) %s", idx+startFrom, str)
	}
	longest := 0
	for _, str := range optionDisplay {
		if len(str)+2 > longest {
			longest = len(str) + 2
		}
	}
	columns := TermWidth / longest
	rows := len(options) / columns
	if len(options)%columns >= 0 {
		rows++
	}
	cols := make([][]string, columns)
	for idx, _ := range cols {
		cols[idx] = make([]string, rows)
	}
	offset := 0
	for cidx := 0; cidx < len(cols); cidx++ {
		for ridx := 0; ridx < len(cols[cidx]); ridx++ {
			if offset < len(optionDisplay) {
				cols[cidx][ridx] = optionDisplay[offset]
			}
			offset++
		}
	}

	for ridx := 0; ridx < rows; ridx++ {
		for cidx := 0; cidx < columns; cidx++ {
			fmt.Print(padString(cols[cidx][ridx], longest))
		}
		fmt.Println()
	}
}

func Select(question string, options []string, startFrom, defaultOption int, confirm bool) int {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("%s:\n", question)
		printSelect(options, startFrom)
		fmt.Printf("Choose from number above [default : (%d) %s] ? ", defaultOption, options[defaultOption-startFrom])
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			return defaultOption
		}
		if chosen, err := strconv.Atoi(text); err != nil {
			fmt.Println("Invalid input")
			continue
		} else {
			if chosen >= startFrom && chosen < len(options)+startFrom {
				if confirm {
					if Confirm(fmt.Sprintf("(%d) %s - Are you sure ? ", chosen, options[chosen-startFrom]), true) {
						return chosen
					}
				} else {
					return chosen
				}
				continue
			} else {
				fmt.Printf("Number %d is not valid option\n", chosen)
				continue
			}
		}
	}
}

func AskNumber(question string, from, to, def int, confirm bool) int {
	scanner := bufio.NewScanner(os.Stdin)
	var min, max int
	if from < to {
		min = from
		max = to
	} else {
		min = to
		max = from
	}
	for {
		fmt.Printf("%s (%d - %d) [default %d] : ", question, min, max, def)
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			if confirm {
				if Confirm(fmt.Sprintf("%d, are you sure?", def), true) {
					return def
				} else {
					continue
				}
			} else {
				return def
			}
		}
		num, err := strconv.Atoi(text)
		if err != nil {
			fmt.Printf("%s is not a valid number\n", text)
			continue
		}
		if num < min || num > max {
			fmt.Printf("%d is not in range between %d and %d\n", num, min, max)
			continue
		}
		if Confirm(fmt.Sprintf("%d, are you sure?", num), true) {
			return num
		} else {
			continue
		}
	}
}

func AskTime(question string, def time.Time, confirm bool) time.Time {
	format := "2006-01-02 15:04:05 -0700"
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s [default \"%s\"] : ", question, def.Format(format))
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			if confirm {
				if Confirm(fmt.Sprintf("\"%s\", are you sure?", def.Format(format)), true) {
					return def
				} else {
					continue
				}
			} else {
				return def
			}
		}
		theTime, err := time.Parse(format, text)
		if err != nil {
			fmt.Printf("\"%s\" is not a valid time format\n")
			continue
		}
		if Confirm(fmt.Sprintf("\"%s\", are you sure?", text), true) {
			return theTime
		} else {
			continue
		}
	}
}

func Ask(question, defaultAnswer string, confirm bool) string {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf("%s [default \"%s\"] : ", question, defaultAnswer)
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			if confirm {
				if Confirm(fmt.Sprintf("\"%s\", are you sure?", defaultAnswer), true) {
					return defaultAnswer
				} else {
					continue
				}
			} else {
				return defaultAnswer
			}
		}
		if Confirm(fmt.Sprintf("\"%s\", are you sure?", text), true) {
			return text
		} else {
			continue
		}
	}
}

func Confirm(question string, def bool) bool {
	scanner := bufio.NewScanner(os.Stdin)
	var defText string
	if def {
		defText = "Y"
	} else {
		defText = "N"
	}
	for {
		fmt.Printf("%s (y/n/Y/N) [default : %s] ? ", question, defText)
		scanner.Scan()
		text := scanner.Text()
		if len(text) == 0 {
			return def
		}
		if strings.ToUpper(text) == "Y" {
			return true
		}
		if strings.ToUpper(text) == "N" {
			return false
		}
		fmt.Printf("Invalid input")
	}
}

func padString(str string, length int) string {
	buff := &bytes.Buffer{}
	buff.WriteString(str)
	for buff.Len() < length {
		buff.WriteString(" ")
	}
	return buff.String()
}
