package pkg

import (
	"fmt"
	"math/rand/v2"
	"time"
)

func IncreaseNestedWithPointer(a *int) *int {
	*a = *a + 1

	return a
}

func IncreaseByOneWithPointer(a *int) {
	*a = *a + 1

}

func IncreaseByOnwWithoutPointer(a int) {
	a += 1
}

func PrintFiveMessages(message string) {
	for i := 0; i < 5; i++ {
		fmt.Println(message)
		val := rand.IntN(200)
		time.Sleep(time.Duration(val) * time.Millisecond)
	}
}

func Second() {

	PrintFiveMessages("message A")

	go PrintFiveMessages("message B")

}
