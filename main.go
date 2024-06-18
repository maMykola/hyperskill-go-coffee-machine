package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

type MenuItem int

type CoffeeRecipe struct {
	Water int
	Milk  int
	Beans int
	Cost  int
}

type CoffeeMachineState struct {
	Water int
	Milk  int
	Beans int
	Cups  int
}

type CoffeeMachine struct {
	Menu    []MenuItem
	Recipes map[MenuItem]CoffeeRecipe
	State   CoffeeMachineState
	Money   int
}

const (
	espresso MenuItem = iota
	ristretto
	latte
	cappuccino
)

var names = map[MenuItem]string{
	espresso:   "Espresso",
	ristretto:  "Ristretto",
	latte:      "Latte",
	cappuccino: "Cappuccino",
}

func userInput() string {
	var action string

	fmt.Printf("> ")
	fmt.Scanln(&action)

	return action
}

func chooseAction() string {
	fmt.Println("Write action (buy, fill, take, remaining, exit):")

	return userInput()
}

func buy(cm *CoffeeMachine) {
	coffee, ok := cm.chooseCoffee()
	if !ok {
		return
	}

	if err := cm.make(coffee); err != nil {
		fmt.Printf("Sorry, not enough %v!\n", err)
	} else {
		fmt.Println("I have enough resources, making you a coffee!")
	}

	fmt.Println()
}

func fill(cm *CoffeeMachine) {
	var water, milk, beans, cups int

	fmt.Println()
	fmt.Println("Write how many ml of water you want to add:")
	water, _ = strconv.Atoi(userInput())
	fmt.Println("Write how many ml of milk you want to add:")
	milk, _ = strconv.Atoi(userInput())
	fmt.Println("Write how many grams of coffee beans you want to add:")
	beans, _ = strconv.Atoi(userInput())
	fmt.Println("Write how many disposable cups you want to add:")
	cups, _ = strconv.Atoi(userInput())
	fmt.Println()

	cm.State.Water += water
	cm.State.Milk += milk
	cm.State.Beans += beans
	cm.State.Cups += cups
}

func take(cm *CoffeeMachine) {
	fmt.Printf("\nI gave you $%d\n\n", cm.takeMoney())
}

func remaining(cm *CoffeeMachine) {
	fmt.Println()
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", cm.State.Water)
	fmt.Printf("%d ml of milk\n", cm.State.Milk)
	fmt.Printf("%d g of coffee beans\n", cm.State.Beans)
	fmt.Printf("%d disposable cups\n", cm.State.Cups)
	fmt.Printf("$%d of money\n", cm.Money)
	fmt.Println()
}

func (cm *CoffeeMachine) takeMoney() int {
	money := cm.Money
	cm.Money = 0
	return money
}
func (cm *CoffeeMachine) chooseCoffee() (CoffeeRecipe, bool) {
	fmt.Printf("\nWhat do you want to buy? (back - to main menu)\n")
	for i, item := range cm.Menu {
		fmt.Printf("  %d. %s\n", i+1, names[item])
	}

	action := userInput()

	if action == "back" {
		return CoffeeRecipe{}, false
	}

	i, err := strconv.Atoi(action)
	if err != nil {
		panic(err)
	}

	kind := cm.Menu[i-1]

	return cm.Recipes[kind], true
}

func (cm *CoffeeMachine) make(coffee CoffeeRecipe) error {
	if cm.State.Water < coffee.Water {
		return errors.New("water")
	} else if cm.State.Milk < coffee.Milk {
		return errors.New("milk")
	} else if cm.State.Beans < coffee.Beans {
		return errors.New("beans")
	} else if cm.State.Cups == 0 {
		return errors.New("disposable cups")
	}

	cm.State.Water -= coffee.Water
	cm.State.Milk -= coffee.Milk
	cm.State.Beans -= coffee.Beans
	cm.State.Cups--

	cm.Money += coffee.Cost

	return nil
}

func main() {
	var coffeeMachine = CoffeeMachine{
		Menu: []MenuItem{espresso, ristretto, latte, cappuccino},
		Recipes: map[MenuItem]CoffeeRecipe{
			espresso:   CoffeeRecipe{250, 0, 16, 4},
			ristretto:  CoffeeRecipe{250, 0, 30, 7},
			latte:      CoffeeRecipe{350, 75, 20, 7},
			cappuccino: CoffeeRecipe{200, 100, 12, 6},
		},
		State: CoffeeMachineState{400, 540, 120, 9},
		Money: 540,
	}

	for {
		switch chooseAction() {
		case "exit":
			os.Exit(0)
		case "buy":
			buy(&coffeeMachine)
		case "fill":
			fill(&coffeeMachine)
		case "take":
			take(&coffeeMachine)
		case "remaining":
			remaining(&coffeeMachine)
		}
	}
}
