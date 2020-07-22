package main

import (
	"fmt"
	"os"
)

type Trade struct {
	CompanyName string // stock company name
	Volume int // number of shares
	Price float64 // trade price
	Buy bool // true to buy, false to sell
}

// New is mandatory
func NewTrade(compananyName string, volume int, price float64, buy bool) (*Trade, error){
	if compananyName =="" {
		return nil, fmt.Errorf("Sorry Company Name cannot be empty")
	}

	if price <0 {
		return nil, fmt.Errorf("Sorry price cannot be less than zero")
	}

	if volume <=0 {
		return nil, fmt.Errorf("you have to trade atleast 1 stock ")
	}

	trade:= &Trade{
		CompanyName: compananyName,
		Volume:  volume,
		Price: price,
		Buy: buy,
	}

	return  trade, nil
}

func (t *Trade) Value() float64  {
	value := float64(t.Volume) * t.Price

	if t.Buy {
		fmt.Println("You are buying.  with value ", value)
	}else  {
		fmt.Println("You are selling.  with value ", value)
	}
	return  value
}


func main()  {
	comp, err := NewTrade("Swiggy", 10, 56.6, true )

	if err != nil {
		fmt.Println("Error Creating trade ", err)
		os.Exit(1)
	}
	fmt.Println(comp)
	fmt.Println(comp.CompanyName)
	fmt.Println(comp.Value())
}
