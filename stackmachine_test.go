package main

import (
	"testing"
)

func TestStackMachine_Execute(t *testing.T) {
	machine := &StackMachine{}
	bytecode := []int{0, 5, 0, 3, 1, 0, 10, 0, 4, 2, 0, 2, 0, 7, 3, 0, 15, 0, 3}

	err := machine.Execute(bytecode)
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	result, _ := machine.Pop()
	expectedResult := 15 // Updated expected result
	if result != expectedResult {
		t.Errorf("Expected result: %d, got: %d", expectedResult, result)
	}
}

func TestStackMachine_Execute_MissingOperand(t *testing.T) {
	machine := &StackMachine{}
	bytecode := []int{0, 5, 1} // Missing operand for push instruction

	err := machine.Execute(bytecode)
	expectedError := "Missing operand for push instruction"
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != expectedError {
		t.Errorf("Expected error: %s, got: %v", expectedError, err)
	}
}


func TestStackMachine_Execute_UnknownInstruction(t *testing.T) {
	machine := &StackMachine{}
	bytecode := []int{5} // Unknown instruction

	err := machine.Execute(bytecode)
	expectedError := "Unknown instruction: 5"
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != expectedError {
		t.Errorf("Expected error: %s, got: %v", expectedError, err)
	}
}

func TestStackMachine_Execute_EmptyStack(t *testing.T) {
	machine := &StackMachine{}
	bytecode := []int{1} // Add instruction with empty stack

	err := machine.Execute(bytecode)
	expectedError := "Stack is empty"
	if err == nil {
		t.Error("Expected an error, but got nil")
	} else if err.Error() != expectedError {
		t.Errorf("Expected error: %s, got: %v", expectedError, err)
	}
}

func TestStackMachine_Add(t *testing.T) {
	machine := &StackMachine{}
	machine.Push(5)
	machine.Push(3)

	err := machine.Add()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	result, _ := machine.Pop()
	expectedResult := 8
	if result != expectedResult {
		t.Errorf("Expected result: %d, got: %d", expectedResult, result)
	}
}

func TestStackMachine_Subtract(t *testing.T) {
	machine := &StackMachine{}
	machine.Push(5)
	machine.Push(3)

	err := machine.Subtract()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	result, _ := machine.Pop()
	expectedResult := 2
	if result != expectedResult {
		t.Errorf("Expected result: %d, got: %d", expectedResult, result)
	}
}

func TestStackMachine_Multiply(t *testing.T) {
	machine := &StackMachine{}
	machine.Push(5)
	machine.Push(3)

	err := machine.Multiply()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	result, _ := machine.Pop()
	expectedResult := 15
	if result != expectedResult {
		t.Errorf("Expected result: %d, got: %d", expectedResult, result)
	}
}

func TestStackMachine_Divide(t *testing.T) {
	machine := &StackMachine{}
	machine.Push(10)
	machine.Push(2)

	err := machine.Divide()
	if err != nil {
		t.Errorf("Expected no error, got: %v", err)
	}

	result, _ := machine.Pop()
	expectedResult := 5
	if result != expectedResult {
		t.Errorf("Expected result: %d, got: %d", expectedResult, result)
	}
}
