package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/olekukonko/tablewriter"
)

type Task struct {
	Title       string    `json:"title"`
	Completed   bool      `json:"completed"`
	CreatedAt   time.Time `json:"created_at"`
	CompletedAt time.Time `json:"completed_at"`
}

type TodoList []Task

func (t *TodoList) Add(title string) {
	newTask := Task{Title: title, Completed: false, CreatedAt: time.Now(), CompletedAt: time.Time{}}
	*t = append(*t, newTask)
}

func (t *TodoList) ValidateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		err := errors.New("index out of bounds")
		fmt.Println(err)
		return err
	}
	return nil
}

func (t *TodoList) Delete(index int) error {
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
	return nil
}

func (t *TodoList) Toggle(index int) error {
	if err := t.ValidateIndex(index); err != nil {
		return err
	}
	(*t)[index].Completed = !(*t)[index].Completed
	(*t)[index].CompletedAt = time.Now()
	return nil
}

func (t *TodoList) Load(filename string) error {
	file, err := os.ReadFile(filename)
	if err != nil {
		return err
	}
	json.Unmarshal(file, t)
	return nil
}

func (t *TodoList) Save(filename string) error {
	file, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	json.NewEncoder(file).Encode(t)
	return nil
}

func (t *TodoList) Print() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"No", "Title", "Completed", "Created At", "Completed At"})
	for i, task := range *t {
		table.Append([]string{fmt.Sprintf("%d", i+1), task.Title, fmt.Sprintf("%t", task.Completed), task.CreatedAt.Format("2006-01-02 15:04:05"), task.CompletedAt.Format("2006-01-02 15:04:05")})
	}
	table.Render()

}
