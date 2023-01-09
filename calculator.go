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
	var x, y, rim1, rim2 int

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
	x, rim1 = check(s1[0], m)

	y, rim2 = check(s1[2], m)

	err_rim := ckeck_rim(rim1, rim2)
	if err_rim != nil {
		fmt.Println(err_rim)
		return
	}

	if s1[1] == "+" {
		if rim1 == 0 {
			fmt.Println(x + y)
		} else {
			fmt.Println(int_to_rim(x + y))
		}

	} else if s1[1] == "-" {
		if rim1 == 0 {
			fmt.Println(x - y)
		} else {
			err := ckeck_rim2(x, y)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Println(int_to_rim(x - y))
		}

	} else if s1[1] == "*" {
		if rim1 == 0 {
			fmt.Println(x * y)
		} else {
			fmt.Println(int_to_rim(x * y))
		}

	} else if s1[1] == "/" {
		if rim1 == 0 {
			fmt.Println(x / y)
		} else {
			fmt.Println(int_to_rim(x / y))
		}
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

func int_to_rim(a int) string {
	m := [][]string{
		{"", "M", "MM", "MMM"},
		{"", "C", "CC", "CCC", "CD", "D", "DC", "DCC", "DCCC", "CM"},
		{"", "X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"},
		{"", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}}
	n := []int{1000, 100, 10, 1}
	result := ""
	for k, v := range n {
		result += m[k][a/v]
		a %= v
	}
	return result
}

func check(s1 string, m map[string]int) (int, int) {
	var a, b int
	if d, err := strconv.Atoi(s1); err == nil {
		err_new := ckeck_ch(d)
		if err_new != nil {
			fmt.Println(err_new)
			return 0, 0
		}
		a = d
		b = 0

	} else {
		val, ok := m[s1]
		if ok {
			a = val
			b = 1
		}
	}
	return a, b
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

func ckeck_ch(n int) error {
	if n < 1 || n > 10 {
		return errors.New("Ошибка, значения должны быть от 1 до 10.")
	}
	return nil
}
