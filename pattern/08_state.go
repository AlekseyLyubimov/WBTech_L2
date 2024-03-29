package pattern

import (
	"fmt"
	"reflect"
)

/*
State - поведенческий паттерн проектирования, который позволяет объектам менять поведение в зависимости от своего
состояния. Извне создаётся впечатление, что изменился класс объекта.

Очень важным нюансом, отличающим этот паттерн от Стратегии, является то, что и контекст, и сами конкретные
состояния могут знать друг о друге и инициировать переходы от одного состояния к другому.

Основные сценарии применения - объект, поведение которого зависит от внутреннего состояния,
причём типов состояний много, и их код часто меняется, или сознательное использование
табличной машины состояний, построенной на условных операторах,
вынуждающее мириться с дублированием кода для похожих состояний и переходов.

Паттрн избавляет от множества больших условных операторов машины состояний, концентрирует 
в одном месте код, связанный с определённым состоянием и упрощает код контекста.
*/

type State interface {
	onHome()
	onPowerOff()
}

type Phone struct {
	off          State
	locked       State
	ready        State
	currentState State
}

func (p *Phone) NewPhone() {
	p.locked = &LockedState{phone: p}
	p.off = &OffState{phone: p}
	p.ready = &ReadyState{phone: p}
	p.currentState = p.off
}

func (p *Phone) setState(state State) {
	p.currentState = state
}

type LockedState struct {
	phone *Phone
}

func (p *LockedState) onHome() {
	fmt.Println("from LOCKED to READY")
	p.phone.setState(p.phone.ready)
}

func (p *LockedState) onPowerOff() {
	fmt.Println("from LOCKED to OFF")
	p.phone.setState(p.phone.off)
}

type ReadyState struct {
	phone *Phone
}

func (o *ReadyState) onHome() {
	fmt.Println("from READY to READY")
	o.phone.setState(o.phone.ready)
}

func (o *ReadyState) onPowerOff() {
	fmt.Println("from READY to OFF")
	o.phone.setState(o.phone.off)
}

type OffState struct {
	phone *Phone
}

func (s *OffState) onHome() {
	fmt.Println("from OFF to LOCKED")
	s.phone.setState(s.phone.locked)
}
func (s *OffState) onPowerOff() {
	fmt.Println("from OFF to LOCKED")
	s.phone.setState(s.phone.locked)
}

func printState(p *Phone) {
	str := fmt.Sprintf("State: %s", reflect.TypeOf(p.currentState).String())
	fmt.Println(str)
}
