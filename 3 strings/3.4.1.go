package main

import (
	"fmt"
)

func main() {
	saplingLength := 100
	BambooBuff := 50
	BambooDebuff := 20
	fmt.Print("Что мы знаем о разведении бамбука в Мытищах?", "\n", "\n",
		"Длинна саженца при посадке составляет ", saplingLength, " сантиметров,", "\n",
		"За день бамбук отрастает на ", BambooBuff, " сантиметров,", "\n",
		"А гусеницы за ночь съедуют ", BambooDebuff, " сантиметров бамбука.", "\n", "\n",
		"Какой же удивительный мир Мытищинской природы!", "\n", "\n",
		"Тем не менее, давайте подсчитаем, на сколько вымахает наш Babboo за 2,5 дня.", "\n")
	totalBuff := 2*BambooBuff + BambooBuff/2
	totalDebuff := 2 * BambooDebuff
	fmt.Print("За 2.5 дня наш бамбук вырастет на ", totalBuff, " сантиметра.", "\n",
		"Будет пожран гусеницами на ", totalDebuff, " сантиметров.", "\n",
		"И с учётом его изначального размера его длинна будет составлять ", totalBuff-totalDebuff+saplingLength, " сантиметров.")
}
