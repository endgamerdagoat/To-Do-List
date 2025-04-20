package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"
	"github.com/aquasecurity/table"
	"slices"
)

type Todo struct {
	Name        string
	Done         bool
	TimeCreated  time.Time
	TimeFinished *time.Time
}

type Todos []Todo 

func (todos *Todos) add(name string) {
	todo:=Todo {
		Name: name,
		Done: false,
		TimeFinished: nil,
		TimeCreated: time.Now(),
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*todos) {
		err := errors.New("invalid index")
		fmt.Println(err)
		return err
	}
	return nil
}

func (todos *Todos) delete(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	*todos = slices.Delete(t, index, index+1)
	return nil
}

func (todos *Todos) toggle(index int) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	isDone := t[index].Done

	if !isDone {
		finishTime := time.Now()
		t[index].TimeFinished = &finishTime
	}

	t[index].Done = !isDone

	return nil
}

func (todos *Todos) edit(index int, name string) error {
	t := *todos

	if err := t.validateIndex(index); err != nil {
		return err
	}

	t[index].Name = name

	return nil
}

func (todos *Todos) print() {
	table := table.New(os.Stdout)
	table.SetRowLines(false)
	table.SetHeaders("#", "Name", "Done?", "Time Created", "Time Finished")
	for index, t := range *todos {
		done := "❌"
		timeFinished := ""

		if t.Done {
			done = "✅"

			if t.TimeFinished != nil {
				timeFinished = t.TimeFinished.Format(time.RFC1123)
			}
		}

		table.AddRow(strconv.Itoa(index), t.Name, done, t.TimeCreated.Format(time.RFC1123), timeFinished)
	}

	table.Render()

}