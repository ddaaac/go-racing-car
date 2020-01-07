package IO

import (
	"../Exception"
	"bufio"
	"fmt"
	"os"
	"strings"
)

const newLineByte = '\n'
const newLineString = "\n"
const separator = ","

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func GetCarsName() []string {
	fmt.Println("차이름을 입력해주세요")
	carsName, _ := reader.ReadString(newLineByte)
	carsName = strings.TrimSuffix(carsName, newLineString)
	err := Exception.ValidateCarsName(carsName)
	if err != nil {
		PrintError(err)
		return GetCarsName()
	}
	return strings.Split(carsName, separator)
}

