package input

import (
	"fmt"
	"os"
	"strconv"
)

func ReadInput(episode int, quest int, part int) string {
	filename := "../../input/everybody_codes_e" + strconv.Itoa(episode) + "_q" + fmt.Sprintf("%02d", quest) + "_p" + strconv.Itoa(part) + ".txt"
	input, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return string(input)
}
