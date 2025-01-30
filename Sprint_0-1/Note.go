package main

import (
	"fmt"
	"time"
)

type Task struct {
	summary     string
	description string
	deadline    time.Time
	priority    int
}

func (t Task) IsOverdue() bool {
	return t.deadline.After(time.Now())
}

func (t Task) IsTopPriority() bool {
	if t.priority > 3 {
		return true
	}
	return false
}

type Note struct {
	title string
	text  string
}

type ToDoList struct {
	name  string
	tasks []Task
	notes []Note
}

func (t ToDoList) TasksCount() int {
	return len(t.tasks)
}

func (n ToDoList) NotesCount() int {
	return len(n.notes)
}

func (td ToDoList) CountTopPrioritiesTasks() int {
	var count int
	for i := range td.tasks {
		if Task.IsTopPriority(td.tasks[i]) {
			count++
		}
	}
	return count
}

func (tdo ToDoList) CountOverdueTasks() int {
	var counter int
	for i := range tdo.tasks {
		if Task.IsOverdue(tdo.tasks[i]) {
			counter++
		}
	}
	return counter
}

func main() {
	todo := ToDoList{name: "Gosha ToDo list", tasks: []Task{Task{summary: "Make Yandex Lyceum Task 9", deadline: time.Now().Add(-time.Hour), description: "Make Module 0, Task 9", priority: 5}, Task{summary: "Make Yandex Lyceum Task 10", deadline: time.Now().Add(time.Hour), description: "Make Module 0, Task 10", priority: 3}}, notes: []Note{Note{title: "ToDo list task", text: "ToDo list task in Yandex Lyceum is very interesting"}}}
	fmt.Println(todo.TasksCount())
	fmt.Println(todo.NotesCount())
	fmt.Println(todo.CountTopPrioritiesTasks())
	fmt.Print(todo.CountOverdueTasks())
}
