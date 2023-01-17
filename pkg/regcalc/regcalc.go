package regcalc

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"module34/module346/pkg/calc"
	"os"
	"regexp"
	"strconv"
)

func RegFile(inputFile, outputFile string) error {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		return err
	}
	re := regexp.MustCompile(`([0-9]+)([+-/*]{1})([0-9]+)=\?(\r\n|$)`)

	submatches := re.FindAllStringSubmatch(string(data), -1)

	f, err := os.Create(outputFile)
	if err != nil {
		return err
	}
	defer f.Close()
	w := bufio.NewWriter(f)
	c := calc.NewCalculator()
	var a, b float64
	for _, s := range submatches {
		if a, err = strconv.ParseFloat(s[1], 64); err != nil {
			return err
		}
		if b, err = strconv.ParseFloat(s[3], 64); err != nil {
			return err
		}
		res := c.Calculate(a, b, s[2])
		_, err = w.WriteString(fmt.Sprintf("%s%s%s=%.2f\n", s[1], s[2], s[3], res))
		if err != nil {
			return err
		}
	}
	w.Flush()
	return nil
}
