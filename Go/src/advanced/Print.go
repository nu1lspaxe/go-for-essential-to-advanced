package advanced

import "fmt"

func RunPrint() {
	test := true
	message := fmt.Sprintf(`
HELLO WORLD
I AM YOUR ADMIN
YOU MUST LISTEN TO ME :p
Are you sure to join me? %t`, test)

	fmt.Println(message)
}
