package main

import "fmt"

type Task struct {
	Identifier int
	Priority   int
}

type MaximumElementDataStructure interface {
	Add(id int, priority int, object interface{})
	SetPriority(id int, priority int)
	Get() (object interface{})
}

type Scheduler struct {
	dataStruct MaximumElementDataStructure
}

func NewScheduler(dataStruct MaximumElementDataStructure) Scheduler {
	return Scheduler{
		dataStruct: dataStruct,
	}
}

func (s *Scheduler) AddTask(task Task) {
	s.dataStruct.Add(task.Identifier, task.Priority, task)
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	s.dataStruct.SetPriority(taskID, newPriority)
}

func (s *Scheduler) GetTask() Task {
	return (s.dataStruct.Get()).(Task)
}

func main() {

	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := NewScheduler(NewHeap())
	scheduler.AddTask(task1)
	scheduler.AddTask(task2)
	scheduler.AddTask(task3)
	scheduler.AddTask(task4)
	scheduler.AddTask(task5)

	fmt.Println(scheduler.GetTask())
	fmt.Println(scheduler.GetTask())
	scheduler.ChangeTaskPriority(1, 100)
	fmt.Println(scheduler.GetTask())
	fmt.Println(scheduler.GetTask())
	fmt.Println(scheduler.GetTask())

	// fmt.Println(scheduler)
}
