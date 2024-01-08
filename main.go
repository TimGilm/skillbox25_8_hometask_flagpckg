/*
Написать программу для нахождения подстроки в кириллической подстроке.
Программа должна запускаться с помощью команды:
go run main.go --str "строка для поиска" --substr "поиска"
Для реализации такой работы с флагами воспользуйтесь пакетом flags, а
для поиска подстроки в строке вам понадобятся руны.
*/
package main

import (
	"flag"
	"fmt"
)

var strMain string
var strSub string

func main() {

	flag.StringVar(&strMain, "str", "", "set main string")
	flag.StringVar(&strSub, "substr", "", "set substring")
	flag.Parse()

	if substringSearch(strMain, strSub) == true {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}

func substringSearch(str string, subStr string) bool {
	mainS := []rune(str)
	subS := []rune(subStr)

cycle:
	for i, letterMainString := range mainS {
		if letterMainString == subS[0] {
			if i+len(subS) > len(mainS) {
				break
			} else {
				for j := 0; j < len(subS); j++ {
					if mainS[i+j] != subS[j] {
						break cycle
					}
				}
				return true
			}
		}
	}
	return false
}
