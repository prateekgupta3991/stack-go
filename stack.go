package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("This is a stack")
	s := NewStack()
	if t, err := s.Pop(); err != nil {
		fmt.Errorf("Error - $s", err)
	} else {
		s = t;
		if val, err := s.Pop(); err != nil {
			fmt.Errorf("Error while popping value - %s", err)
		} else {
			fmt.Printf("Poppd value - %s", val)
		}
	}
	s = s.Push(356)
	s = s.Push("Multi type values")
	s = s.Push(84668434)
	if val, err := s.Top(); err != nil {
		fmt.Errorf("Error while top - %s", val)
	} else {
		fmt.Printf("Top value - ", val)
	}
	if t, err := s.Pop(); err != nil {
		fmt.Errorf("Error - $s", err)
	} else {
		s = t;
		if val, err := s.Top(); err != nil {
			fmt.Errorf("Error while popping value - %s", err)
		} else {
			fmt.Printf("Poppd value - %s", val)
		}
	}
	if t, err := s.Pop(); err != nil {
		fmt.Errorf("Error - $s", err)
	} else {
		s = t;
		if val, err := s.Top(); err != nil {
			fmt.Errorf("Error while popping value - %s", err)
		} else {
			fmt.Printf("Poppd value - %s", val)
		}
	}
}

type Stack interface {
	Top() (interface{}, error)
	Push(val interface{}) Stack
	Pop() (Stack, error)
}

type StackImpl struct {
	elem interface{}
	next Stack
}

func NewStack() Stack {
	return &StackImpl {
		elem: nil,
		next: nil,
	}
}

func (s *StackImpl) Top() (interface{}, error) {
	if s.elem == nil {
		return nil, errors.New("Top error due to empty stack")
	} else {
		return s.elem, nil
	}
}

func (s *StackImpl) Push(val interface{}) Stack {
	if s.elem == nil {
		s.elem = val
	} else {
		tmp := StackImpl {
			elem: val,
			next: s,
		}
		s = &tmp
	}
	return s
}

func (s StackImpl) Pop() (Stack, error) {
	if s.elem == nil {
		return nil, errors.New("Pop error due to empty stack")
	} else {
		return s.next, nil
	}
}