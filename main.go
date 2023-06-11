package main

import (
	"fmt"
)

type StackMachine struct {
	stack []int
}

func (m *StackMachine) Push(value int) {
	m.stack = append(m.stack, value)
}

func (m *StackMachine) Pop() (int, error) {
	stackLen := len(m.stack)
	if stackLen > 0 {
		value := m.stack[stackLen-1]
		m.stack = m.stack[:stackLen-1]
		return value, nil
	}
	return 0, fmt.Errorf("Stack is empty")
}

func (m *StackMachine) Execute(bytecode []int) error {
	for i := 0; i < len(bytecode); i++ {
		instruction := bytecode[i]
		switch instruction {
		case 0: // Push instruction
			i++
			if i >= len(bytecode) {
				return fmt.Errorf("Missing operand for push instruction")
			}
			operand := bytecode[i]
			m.Push(operand)
		case 1: // Add instruction
			err := m.Add()
			if err != nil {
				return err
			}
		case 2: // Subtract instruction
			err := m.Subtract()
			if err != nil {
				return err
			}
		case 3: // Multiply instruction
			err := m.Multiply()
			if err != nil {
				return err
			}
		case 4: // Divide instruction
			err := m.Divide()
			if err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unknown instruction: %d", instruction)
		}
	}
	return nil
}

func (m *StackMachine) Add() error {
	operand2, err := m.Pop()
	if err != nil {
		return err
	}
	operand1, err := m.Pop()
	if err != nil {
		return err
	}
	result := operand1 + operand2
	m.Push(result)
	return nil
}

func (m *StackMachine) Subtract() error {
	operand2, err := m.Pop()
	if err != nil {
		return err
	}
	operand1, err := m.Pop()
	if err != nil {
		return err
	}
	result := operand1 - operand2
	m.Push(result)
	return nil
}

func (m *StackMachine) Multiply() error {
	operand2, err := m.Pop()
	if err != nil {
		return err
	}
	operand1, err := m.Pop()
	if err != nil {
		return err
	}
	result := operand1 * operand2
	m.Push(result)
	return nil
}

func (m *StackMachine) Divide() error {
	operand2, err := m.Pop()
	if err != nil {
		return err
	}
	operand1, err := m.Pop()
	if err != nil {
		return err
	}
	result := operand1 / operand2
	m.Push(result)
	return nil
}

func main() {
	machine := &StackMachine{}
	bytecode := []int{0, 5, 0, 3, 1, 0, 10, 0, 4, 2, 0, 2, 0, 7, 3, 0, 15, 0, 3}

	err := machine.Execute(bytecode)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		result, _ := machine.Pop()
		fmt.Println("Result of bytcode: with opcode instruction : ", result) // Output: 5
	}
}
