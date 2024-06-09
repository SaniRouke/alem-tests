package main

import (
	"fmt"
	"os"
)

func Example1() {
	main()
	fmt.Fprint(os.Stdin, "2 3\n111\n111")
	// Output:
	// _______________________
	// |       |       |       |
	// |       |       |       |
	// |_______|_______|_______|
	// |       |       |       |
	// |       |       |       |
	// |_______|_______|_______|
}

func Example2() {
	main()
	fmt.Fprint(os.Stdin, "1 1\n2")
	// Output:
	// _______
	// |       |
	// |   >   |
	// |_______|
}

func Example3() {
	main()
	fmt.Fprint(os.Stdin, "1 1\n2")
	// Output:
	// _______
	// |XXXXXXX|
	// |XXXXXXX|
	// |XXXXXXX|
}

func Example4() {
	main()
	fmt.Fprint(os.Stdin, `8 8
	00000000
	01103010
	01101110
	01010100
	01111110
	00010010
	02111010
	00000000`)
	// Output:
	// _______________________________________________________________
	// |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
	// |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
	// |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
	// |XXXXXXX|       |       |XXXXXXX|       |XXXXXXX|       |XXXXXXX|
	// |XXXXXXX|       |       |XXXXXXX|   @   |XXXXXXX|       |XXXXXXX|
	// |XXXXXXX|_______|_______|XXXXXXX|_______|XXXXXXX|_______|XXXXXXX|
	// |XXXXXXX|       |       |XXXXXXX|       |       |       |XXXXXXX|
	// |XXXXXXX|       |       |XXXXXXX|       |       |       |XXXXXXX|
	// |XXXXXXX|_______|_______|XXXXXXX|_______|_______|_______|XXXXXXX|
	// |XXXXXXX|       |XXXXXXX|       |XXXXXXX|       |XXXXXXX|XXXXXXX|
	// |XXXXXXX|       |XXXXXXX|       |XXXXXXX|       |XXXXXXX|XXXXXXX|
	// |XXXXXXX|_______|XXXXXXX|_______|XXXXXXX|_______|XXXXXXX|XXXXXXX|
	// |XXXXXXX|       |       |       |       |       |       |XXXXXXX|
	// |XXXXXXX|       |       |       |       |       |       |XXXXXXX|
	// |XXXXXXX|_______|_______|_______|_______|_______|_______|XXXXXXX|
	// |XXXXXXX|XXXXXXX|XXXXXXX|       |XXXXXXX|XXXXXXX|       |XXXXXXX|
	// |XXXXXXX|XXXXXXX|XXXXXXX|       |XXXXXXX|XXXXXXX|       |XXXXXXX|
	// |XXXXXXX|XXXXXXX|XXXXXXX|_______|XXXXXXX|XXXXXXX|_______|XXXXXXX|
	// |XXXXXXX|       |       |       |       |XXXXXXX|       |XXXXXXX|
	// |XXXXXXX|   >   |       |       |       |XXXXXXX|       |XXXXXXX|
	// |XXXXXXX|_______|_______|_______|_______|XXXXXXX|       |XXXXXXX|
	// |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
	// |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
	// |XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|XXXXXXX|
}

func ExampleBonus1() {
	main()
	fmt.Fprint(os.Stdin, `1 4
	0123
	KP$`)
	// Output:
	// _______________________________
	// |KKKKKKK|       |       |       |
	// |KKKKKKK|       |   P   |   $   |
	// |KKKKKKK|_______|_______|_______|
}
