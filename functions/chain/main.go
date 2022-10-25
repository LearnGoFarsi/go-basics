package main

import (
	"fmt"
)

// !+ Step
type ExecutionFunc func() error

type Step struct {
	nextStep      *Step
	executionFunc ExecutionFunc
}

func NewStep(f ExecutionFunc) *Step {
	return &Step{
		executionFunc: f,
	}
}

func (s *Step) SetNextStep(nextStep *Step) {
	s.nextStep = nextStep
}

func (s *Step) Execute() {

	if s.executionFunc == nil {
		panic("No function is presented")
	}

	if err := s.executionFunc(); err != nil {
		fmt.Printf("Error: %s \n", err)
	}

	if s.nextStep != nil {
		s.nextStep.Execute()
	}
}

// !- Step

// !+ ChainBuilder
type ChainBuilder struct {
	firstStep, lastStep *Step
}

func NewChainBuilder() *ChainBuilder {
	return &ChainBuilder{}
}

func (c *ChainBuilder) SetNextStep(step *Step) *ChainBuilder {
	if c.firstStep == nil {
		c.firstStep = step
	}

	if c.lastStep != nil {
		c.lastStep.SetNextStep(step)
	}
	c.lastStep = step

	return c
}

func (c *ChainBuilder) Build() *Step {
	return c.firstStep
}

// !- ChainBuilder

func f1() error {
	fmt.Println("It is F1")
	return nil
}

func f2() error {
	fmt.Println("It is F2")
	return nil
}

func f3() error {
	return fmt.Errorf("%s", "F3 raised error")
}

func main() {

	cb := NewChainBuilder()
	firstStep := cb.SetNextStep(NewStep(f1)).
		SetNextStep(NewStep(f2)).
		SetNextStep(NewStep(f3)).
		Build()

	firstStep.Execute()
}
