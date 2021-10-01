package main

import "fmt"

const (
	isAdmin          = 1 << iota // 1 --> 1
	isHeadquaters                // 2 --> 10
	canSeeFinancials             // 4 --> 100

	canSeeAfrica       // 8 --> 1000
	canSeeAsia         // 16 --> 10000
	canSeeEurope       // 32 --> 100000
	canSeeNorthAmerica // 64 --> 1000000
	canSeeSouthAmerica // 128 --> 10000000
)

func main() {
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope

	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v", isAdmin&roles == isAdmin)
	fmt.Printf("Is HQ? %v", isHeadquaters&roles == isHeadquaters)
}
