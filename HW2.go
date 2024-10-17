package main

import (
 "fmt"
 "slices"
)

// stack
type Stack struct {
 values []int
}

func Newstack() *Stack {
 return &Stack{
  values: []int{},
 }
}

func (s *Stack) IsExist() bool {
 return len(s.values) != 0
}

func (s *Stack) Last() int {
 if s.IsExist() != true {
  return 0
 } else {
  return s.values[len(s.values)-1]
 }
}

func (s *Stack) Add(n int) *Stack {
 s.values = append(s.values, n)
 return s
}

func (s *Stack) Pop() int {
 if len(s.values) == 0 {
  return 0
 }
 element := s.values[len(s.values)-1]
 s.values = s.values[:len(s.values)-1]
 return element
}

// stack
// deck если я правильно понял двухсторонняя очередь
type Deck struct {
 values []int
}

func NewDeck() *Deck {
 return &Deck{
  values: []int{},
 }
}

func (s *Deck) IsExist() bool {
 return len(s.values) != 0
}

func (s *Deck) Front() int {
 if s.IsExist() != true {
  return 0
 } else {
  return s.values[0]
 }
}

func (s *Deck) Back() int {
 if s.IsExist() != true {
  return 0
 } else {
  return s.values[len(s.values)-1]
 }
}

func (s *Deck) AddBack(n int) *Deck {
 s.values = append(s.values, n)
 return s
}
func (s *Deck) AddFront(n int) *Deck {
 s.values = slices.Insert(s.values, 0, n)
 return s
}
func (s *Deck) PopBack() int {
 if len(s.values) == 0 {
  return 0
 }
 element := s.values[len(s.values)-1]
 s.values = s.values[:len(s.values)-1]
 return element
}
func (s *Deck) PopFront() int {
 if len(s.values) == 0 {
  return 0
 }
 element := s.values[0]
 if len(s.values) > 1 {
  s.values = s.values[1:]
 } else {
  s.values = []int{}
 }
 return element
}

// deck
func main() {
 //Stack
 a := []int{1, 2, 3, 4}
 b := Newstack()
 for i := 0; i < len(a); i++ { //выведем слайс
  fmt.Println(a[i])
 }
 for i := 0; i < len(a); i++ { // добавим элементы в стек
  b.Add(a[i])
 }
 fmt.Println(b.Pop(), b.IsExist()) //выведем последний элемент, удалим его из стека и проверим наличие в нем элементов
 //Deck
 c := []int{1, 2, 3, 4}
 d := NewDeck()
 for i := 0; i < len(c); i++ { // добавим элементы в queue
  d.AddFront(c[i])
 }
 for i := 0; i < len(c); i++ { // добавим элементы в queue
  d.AddBack(c[i])
 }
 fmt.Println(d.PopFront(), d.PopBack(), d.IsExist())
}