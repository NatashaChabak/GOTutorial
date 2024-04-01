package main

import "fmt"

// 3.1. Необходимо создать глобальную константу и вывести её значение в консоль.
// 3.2. Необходимо создать локальную константу и вывести её значение в консоль.
// 3.3. Необходимо создать глобальную константу. Далее затенить глобальную константу локальной и вывести результат в консоль
// 3.4. Необходимо создать группу из пяти глобальных констант и вывести их значения в консоль.
// 3.5. Необходимо создать группу из пяти локальных констант и вывести их значения в консоль.
// 3.6. Необходимо создать типизированную локальную константу n типа int со значением 5.
// 		Также необходимо создать локальную переменную m, значение которой должно определяться выражением 3.4 + n.
// 		Также нужно вывести значение переменной в консоль (в консоли должно отобразиться значение 8.4).
// 3.7. Использую генератор iota необходимо создать 5 констант. Их значения должны представлять собой 2 : 1, 2, 4, 8, 16.

const globalConst = "I am a global constant"
const testConst = 5

const (
	one   = 0
	two   = 1
	three = 2
	four  = 3
	five  = 4
)

func main() {
	const localConst = "I am a local constant"
	const testConst = 7

	const (
		oneLocal   = 0
		twoLocal   = 1
		threeLocal = 2
		fourLocal  = 3
		fiveLocal  = 4
	)

	const n int = 5       // Типизированная локальная константа типа int со значением 5
	m := 3.4 + float64(n) // Локальная переменная m, значение которой определяется выражением 3.4 + n

	const (
		myConst1 = 1 << iota
		myConst2
		myConst3
		myConst4
		myConst5
	)

	fmt.Println("Global: " + globalConst)                             //3.1
	fmt.Println("Local: " + localConst)                               //3.2
	fmt.Println(testConst)                                            //3.3
	fmt.Println(one, two, three, four, five)                          //3.4
	fmt.Println(oneLocal, twoLocal, threeLocal, fourLocal, fiveLocal) //3.5
	fmt.Println("Значение переменной m:", m)                          //3.6
	fmt.Println(myConst1, myConst2, myConst3, myConst4, myConst5)     //3.7
}
