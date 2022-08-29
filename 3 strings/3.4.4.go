package main

import (
	"fmt"
)

func main() {

	fmt.Print("Добро пожаловать в симулятор развдения бамбука.", "\n", "\n",
		"Для корректной работы программы, пожалуйста, учтоните значения для след. объёктов:", "\n")

	var saplingLength int
	fmt.Print("Укажите рост саженца в сантиметрах.", "\n")
	fmt.Scan(&saplingLength)

	var BambooBuff int
	fmt.Print("Укажите рост бамбука за день в сантиметрах.", "\n")
	fmt.Scan(&BambooBuff)

	var BambooDebuff int
	fmt.Print("Укажите, сколько сантиметров бамбука съедают гусеницы", "\n", "за 1 (арабскую) ноооочь (волшебный востоооок).", "\n")
	fmt.Scan(&BambooDebuff)

	var bigBambooLength int
	fmt.Print("Укажите, какого размера наш желанный Bamboo гигачад в сантиметрах.", "\n")
	fmt.Scan(&bigBambooLength)

	dailyGrowth := BambooBuff - BambooDebuff

	daysToGrow := (bigBambooLength - saplingLength) / dailyGrowth
	fmt.Print("\n", "\n", "С учётом ежежневного роста бамбука на ", dailyGrowth, " сантиметров и изначальной длинны,", "\n", "ему потребуется ", daysToGrow, " дней на рост до ", bigBambooLength, " сантиметров.")
}
