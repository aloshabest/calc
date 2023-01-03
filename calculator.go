package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var x, y, r1, r2 int

	m := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	in := bufio.NewReader(os.Stdin)
	s, err := in.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода: ", err)
	}
	s = filterNewLines(s)
	s1 := strings.Split(s, " ")

	err1 := ckeck_oper(s1)
	if err1 != nil {
		fmt.Println(err1)
		return
	}

	if d, err := strconv.Atoi(s1[0]); err == nil {
		x = d
		r1 = 0
	} else {
		val, ok := m[s1[0]]
		if ok {
			x = val
			r1 = 1
		}
	}

	if q, err := strconv.Atoi(s1[2]); err == nil {
		y = q
		r2 = 0
	} else {
		val, ok := m[s1[2]]
		if ok {
			y = val
			r2 = 1
		}
	}

	err_rim := ckeck_rim(r1, r2)
	if err_rim != nil {
		fmt.Println(err_rim)
		return
	}

	if s1[1] == "+" {
		fmt.Println(x + y)
	} else if s1[1] == "-" {
		if r1 == 1 && r2 == 1 {
			err := ckeck_rim2(x, y)

			if err != nil {
				fmt.Println(err)
				return
			}
		}
		fmt.Println(x - y)
	} else if s1[1] == "*" {
		fmt.Println(x * y)
	} else if s1[1] == "/" {
		fmt.Println(x / y)
	}

}

func filterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r
		}
	}, s)
}

func ckeck_rim(r1 int, r2 int) error {
	if (r1 == 0 && r2 != 0) || (r1 != 0 && r2 == 0) {
		return errors.New("Ошибка, так как используются одновременно разные системы счисления.")
	}
	return nil
}

func ckeck_rim2(a int, b int) error {
	if a-b <= 0 {
		return errors.New("Ошибка, так как в римской системе нет отрицательных чисел.")
	}
	return nil
}

func ckeck_oper(s1 []string) error {
	if len(s1) > 3 {
		return errors.New("Ошибка, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	}
	if len(s1) < 3 {
		return errors.New("Ошибка, так как строка не является математической операцией.")
	}
	return nil
}
