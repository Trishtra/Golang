package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//Ввод ключей (числа через пробел)
	keysInput, _ := reader.ReadString('\n')
	keys := strings.Fields(strings.TrimSpace(keysInput))

	//Ввод значений (строки через пробел)
	valuesInput, _ := reader.ReadString('\n')
	values := strings.Fields(strings.TrimSpace(valuesInput))

	result := make(map[int]string)

	for i := 0; i < len(keys) && i < len(values); i++ {
		if key, err := strconv.Atoi(keys[i]); err == nil {
			result[key] = values[i]
		}
	}

	//Вводим число а
	var a int

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	a, _ = strconv.Atoi(input)

	//keysToDelete := []int{}
	/*for key, value := range result {
		if num, err := strconv.Atoi(value); err == nil {
			if num%a == 0 {
				delete(result, key)
			}
		}
	}*/

	for key := range result {
		if key%a == 0 {
			delete(result, key)
		}
	}

	/*for _, key := range keysToDelete {

		delete(result, key)
	}*/
	fmt.Println(result)
}
