package main

import (
	"fmt"
	
	"math/rand"
	"time"
)

func main() {
	// // For loops 
	// // Number of rows for the pyramid
    // rows := 8
    // for i := 1; i <= rows; i++ {

    //     // Print spaces
    //  for j := 1; j <= rows - i ; j++ {
	// 	 fmt.Print(" ")
	//  }

    //     // Print stars
    // for k := 1; k <= 2*i - 1 ; k++ {
	// 	 fmt.Print("*")
	// }

    //     // Move to next line
    //   fmt.Println()
    // }
	// // While loops 
	// i := 1 
	// for i <= 5 {
	// 	fmt.Println(i)
	// 	i++
	// }
	source := rand.NewSource(time.Now().UnixMicro())
	random := rand.New(source)
	target := random.Intn(100) + 1
	fmt.Println("Welcome user to the guessing game")
		fmt.Println("I will generate a random number between 1 and 100 ")
	fmt.Println("Your job is to keep guessing until you found the generated random number")
	var guess int
for {
	fmt.Println("enter your guess :")
fmt.Scanln(&guess)
if  guess == target {
		fmt.Println("Congratulations ! You have won and found the number")

	break
} else if guess < target {
		fmt.Println("too low ! keep guessing ")

} else if guess > target {
		fmt.Println("too high ! keep guessing ")

}
}
}
