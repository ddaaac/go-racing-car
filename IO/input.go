package IO

import (
	"../Exception"
	"bufio"
	"os"
	"strings"
)

const newLineByte = '\n'
const newLineString = "\n"

var reader *bufio.Reader

func init() {
	reader = bufio.NewReader(os.Stdin)
}

func GetCarsName() []string {
	question := "차이름을 입력해주세요"
	carsName := getInput(question)
	carsNameArray, err := Exception.ValidateCarsName(carsName)
	if err != nil {
		PrintError(err)
		return GetCarsName()
	}
	return carsNameArray
}

func GetTryNumber() int {
	question := "시도할 횟수를 입력해주세요"
	tryNumberString := getInput(question)
	tryNumber, err := Exception.ValidateTryNumber(tryNumberString)
	if err != nil {
		PrintError(err)
		return GetTryNumber()
	}
	return tryNumber
}

func getInput(question string) string {
	PrintLine(question)
	input, _ := reader.ReadString(newLineByte)
	return strings.TrimSuffix(input, newLineString)
}
