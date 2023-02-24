package main

import (
	"fmt"
)

// this interface defines the behaviour for Fruits: and type that implements the Sum()
// method satisfies the fruits interface.
type Fruit interface {
	Sum() float32
}
///////////////////////////////////////////////////////////////////////////////////////////

type Mangoes struct {
	Amount float32
	Price  float32
}

func (g Mangoes) Sum() float32 {
	return g.Amount * g.Price
}

/////////////////////////////////////////////////////////////////////////////////////////////////

func FruitTotalPrice(f Fruit) { // defined a FruitTotalPrice function that is also accepting a parameter
	fmt.Printf("The total price of your mangoes bought are: $ %.2f\n", f.Sum())
}

func main() {
	sum := Mangoes{Amount: 5, Price: 3.50}//create a mangoes instance 
	FruitTotalPrice(sum)//we then pass it to the FruitTotalPrice function
}

