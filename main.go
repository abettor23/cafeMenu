package main

import (
	"fmt"
)

func main() {
	fmt.Println("Введите, пожалуйста, день недели от 1 (понедельник) до 7 (воскресенье):")
	var dayWeek int
	fmt.Scan(&dayWeek)

	if dayWeek > 7 {
		fmt.Println("Вы ввели какую то ерунду. Попробуйте попасть по цифрам правильно.")
	} else if dayWeek > 5 {
		fmt.Println("Добро пожаловать в наше кафе!")
		fmt.Println("На сегодня у нас в меню:")
		fmt.Println("Шашлык (свинина, курица)")
		fmt.Println("Картофель пюре")
		fmt.Println("Овощной салат")
		fmt.Println("Суп вермишелевый")
		fmt.Println("Компот, лимонад")
		fmt.Println("Кофе, чай")
	} else if dayWeek >= 1 {
		fmt.Println("Добро пожаловать в наше кафе!")
		fmt.Println("На сегодня у нас в меню:")
		fmt.Println("Шашлык (свинина, курица)")
		fmt.Println("Картофель пюре")
		fmt.Println("Овощной салат")
		fmt.Println("Суп вермишелевый")
		fmt.Println("Компот, лимонад")
		fmt.Println("Кофе, чай")
		fmt.Println("Бизнес ланч: Ролл Филадельфия, Удон по домашнему, Саке")
	} else {
		fmt.Println("Вы ввели какую то ерунду. Попробуйте попасть по цифрам правильно.")
	}
}
