package pattern

/*
Chain of responsibility - поведенческий паттерн проектирования, заключающийся в передаче запроса по цепочке обработчиков. 
Каждый обработчик сам решает, обработать запрос самостоятельно или передать следующему.

Основные случаи применения:
1) программа должна обрабатывать разнообразные запросы несколькими способами,
но заранее неизвестно, какие конкретно запросы будут приходить и какие обработчики для них понадобятся.
2) важно чтобы обработчики выполнялись в строгом порядке.
3) Когда набор объектов, способных обработать запрос, должен задаваться динамически.

Паттерн уменьшает зависимость между клиентом и обработчиками, реализует принципы единственной обязанности
и открытости/закрытости, облегчает обязанностей между объектами
*/

type RequestFilter interface {
	execute(request string)
	setNext(handler RequestFilter)
}

type AuthFilter struct {
	next RequestFilter
}

func (f *AuthFilter) execute(request string) {
	f.next.execute(request)
	return
}

func (f *AuthFilter) setNext(handler RequestFilter) {
	f.next = handler
}

type ErrorFilter struct {
	next RequestFilter
}

func (f *ErrorFilter) execute(request string) {
	f.next.execute(request)
	return
}

func (f *ErrorFilter) setNext(handler RequestFilter) {
	f.next = handler
}

type ValidationFiler struct {
	next RequestFilter
}

func (f *ValidationFiler) execute(request string) {
	println("Validating request headers, params and body:", request)
	return
}

func (f *ValidationFiler) setNext(handler RequestFilter) {
	f.next = handler
}
