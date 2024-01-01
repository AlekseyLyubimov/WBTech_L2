package patterns

import "fmt"

/*
	Паттерн visitor позволяет добавлять в программу новые операции,
	не изменяя классы объектов, над которыми эти операции могут выполняться.

	Паттерн упрощает добавление операций, работающих со сложными структурами объектов, 
	объединяет родственные операции в одном классе и позволяет накапливать состояние при обходе структуры элементов.

	Применение паттерна требует стабильной иерархии элементов
*/

type Shape interface {
    getType() string
    accept(Visitor)
}

type Square struct {
    side int
}

func (s *Square) accept(v Visitor) {
    v.visitForSquare(s)
}

func (s *Square) getType() string {
    return "Square"
}

type Circle struct {
    radius int
}

func (c *Circle) accept(v Visitor) {
    v.visitForCircle(c)
}

func (c *Circle) getType() string {
    return "Circle"
}

type Visitor interface {
    visitForSquare(*Square)
    visitForCircle(*Circle)
}

type AreaCalculator struct {
    area int
}

func (a *AreaCalculator) visitForSquare(s *Square) {
    fmt.Println("Calculating area for square")
}

func (a *AreaCalculator) visitForCircle(s *Circle) {
    fmt.Println("Calculating area for circle")
}