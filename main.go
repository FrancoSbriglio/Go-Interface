package main

import (
	"fmt"
	"interfaces/users"
)

func main() {

	var frank users.Cashier = users.NewEmploye("Frank")

	var roberto users.Admin = users.NewEmploye("Roberto")

	total := frank.CalcTotal(90, 65, 92.6)

	fmt.Println(total)

	roberto.DesactivateEmployee(frank)

	fmt.Println(frank.CalcTotal(90, 65, 92.6))
}
