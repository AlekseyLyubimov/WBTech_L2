package pattern

import "fmt"

/*
Фабричный метод — это порождающий паттерн проектирования, который решает проблему
создания различных продуктов, без указания конкретных классов продуктов.
В Go невозможно реализовать классический вариант паттерна Фабричный метод,
поскольу в языке отсутствуют возможности ООП, в том числе классы и наследственность.
Несмотря на это, мы все же можем реализовать базовую версию этого паттерна — Простая фабрика.
*/

type IDrink interface {
    setName(name string)
    setFlavour(flavour string)
}

type Drink struct {
    name  string
    flavour string
}

func (d *Drink) setName(name string) {
    d.name = name
}

func (d *Drink) getName() string {
    return d.name
}

func (d *Drink) setFlavour(flavour string) {
    d.flavour = flavour
}

func (d *Drink) getPower() string {
    return d.flavour
}

type Lemonade struct {
    Drink
}

func newLemonade() IDrink {
    return &Lemonade{
        Drink: Drink{
            name:  "Lemonade",
            flavour: "Lemon",
        },
    }
}

type Cola struct {
    Drink
}

func newCola() IDrink {
    return &Cola{
        Drink: Drink{
            name:  "Cola",
            flavour: "Cola",
        },
    }
}

func getDrink(drinkType string) (IDrink, error) {
    if drinkType == "Lemonade" {
        return newLemonade(), nil
    }
    if drinkType == "Cola" {
        return newCola(), nil
    }
    return nil, fmt.Errorf("Wrong drink type passed")
}