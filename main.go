package main

import "fmt"
import "strconv"
import "strings"

func main() {
	var pr string
	fmt.Println("Введите пример:")
	fmt.Scan(&pr)
	b := []byte(pr)
	var c int
	for i := 0; i < len(b); i++ {
		if b[i] == 43 || b[i] == 45 || b[i] == 47 || b[i] == 42 {
			break
		}
		c = i
	}
	if c == len(b)-1 {
		panic("Некорректные данные")
	}
	if (romanToArabic(string(b[:c+1])) == 0 && romanToArabic(string(b[c+2:])) == 0) ||
		(romanToArabic(string(b[:c+1])) == 0 || romanToArabic(string(b[c+2:])) == 0) {
		firstNumb, err := strconv.Atoi(string(b[:c+1]))
		if err != nil {
			panic("Некорректные данные")
		}
		secondNumb, err := strconv.Atoi(string(b[c+2:]))
		if err != nil {
			panic("Некорректные данные")
		}
		switch {
		case firstNumb > 10:
			panic("Некорректное число")
		case secondNumb > 10:
			panic("Некорректное число")
		}
		resArabic(b[c+1], firstNumb, secondNumb)
	} else {
		resRoman(b[c+1], romanToArabic(string(b[:c+1])), romanToArabic(string(b[c+2:])))
	}
}

func add(x, y int) int {
	return x + y
}
func sub(x, y int) int {
	return x - y
}
func mul(x, y int) int {
	return x * y
}
func div(x, y int) int {
	return x / y
}
func romanToArabic(x string) int {
	a := map[string]int{
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
	b := a[x]
	return b
}
func resArabic(x uint8, y, z int) {
	switch x {
	case 43:
		fmt.Println("Равно", add(y, z))
	case 45:
		fmt.Println("Равно", sub(y, z))
	case 42:
		fmt.Println("Равно", mul(y, z))
	case 47:
		fmt.Println("Равно", div(y, z))
	}
}

func resRoman(x uint8, y, z int) {
	switch x {
	case 43:
		if add(y, z) < 1 {
			panic("Результат некорректный")
		} else {
			intToRoman(add(y, z))
		}
	case 45:
		if sub(y, z) < 1 {
			panic("Результат некорректный")
		} else {
			intToRoman(sub(y, z))
		}
	case 42:
		if mul(y, z) < 1 {
			panic("Результат некорректный")
		} else {
			intToRoman(mul(y, z))
		}
	case 47:
		if div(y, z) < 1 {
			panic("Результат некорректный")
		} else {
			intToRoman(div(y, z))
		}
	}
}

func intToRoman(s int) {
	rMap := map[int]string{
		100: "C",
		90:  "XC",
		50:  "L",
		40:  "XL",
		10:  "X",
		9:   "IX",
		5:   "V",
		4:   "IV",
		1:   "I",
	}
	var mapMas []int
	for key, _ := range rMap {
		mapMas = append(mapMas, key)
	}
	for i := 0; i < len(mapMas)-1; i++ {
		for mapMas[i] < mapMas[i+1] {
			mapMas[i], mapMas[i+1] = mapMas[i+1], mapMas[i]
			i = 0
		}
	}
	roman := strings.Builder{}
	for idx, elem := range mapMas {
		for s >= mapMas[idx] {
			roman.WriteString(rMap[elem])
			s = s - mapMas[idx]
		}
	}
	fmt.Print("Равно ", roman.String())
}
