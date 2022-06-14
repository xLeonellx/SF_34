package calculate

import (
	"bufio"
	"os"
	"regexp"
	"strconv"
)

// Функция считывает из файла список математических выражений (функция умеет считать только "+" и "-"),
//вычисляет результат и записывает в другой файл.
func Calculate(inputFile, outputFile string) error {
	file, err := os.Open(inputFile)
	if err != nil {
		return err
	}

	_ = os.Remove(outputFile)
	output, err := os.OpenFile(outputFile, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0777)
	if err != nil {
		return err
	}

	defer file.Close()
	reader := bufio.NewScanner(file)
	writer := bufio.NewWriter(output)
	for reader.Scan() {
		//проверка строки на то, что она является математическим выражением (+-X1+X2+...+Xn=?)
		m, err := regexp.MatchString(`^([\-+]\d+|\d+)+=[?]$`, reader.Text())
		if err != nil {
			return err
		}
		if !m {
			continue
		}

		//извлечение и вычисление чисел из математического выражения
		pattern := regexp.MustCompile(`([\-+]\d+|\d+)`)
		numbers := pattern.FindAllString(reader.Text(), -1)
		var result int
		for _, num := range numbers {
			n, err := strconv.Atoi(num)
			if err != nil {
				return err
			}
			result += n
		}

		//замена "?" на результат вычисления, запись в буфер
		replace := regexp.MustCompile(`[?]`)
		_, err = writer.WriteString(replace.ReplaceAllLiteralString(reader.Text(), strconv.Itoa(result)))
		if err != nil {
			return err
		}
		writer.WriteString("\n")

	}
	err = writer.Flush()
	if err != nil {
		return err
	}
	return nil
}
