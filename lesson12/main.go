package main

import "fmt"

type Task struct {
	Identifier int
	Priority   int
}

type Scheduler struct {
	list []Task
}

func NewScheduler() Scheduler {
	return Scheduler{
		list: make([]Task, 0),
	}
}

func (s *Scheduler) AddTask(t Task) {
	s.list = append(s.list, t)
	i := len(s.list) - 1
	parent := (i - 1) / 2

	for i > 0 && s.list[parent].Priority < s.list[i].Priority {
		s.list[i], s.list[parent] = s.list[parent], s.list[i]

		i = parent
		parent = (i - 1) / 2
	}
}

func (s *Scheduler) ChangeTaskPriority(taskID int, newPriority int) {
	for i := 0; i < len(s.list); i++ {
		if s.list[i].Identifier == taskID {
			s.list[i].Priority = newPriority
		}
	}
	s.heapify(0)
}

func (s *Scheduler) GetTask() Task {
	if len(s.list) <= 0 {
		return Task{}
	}

	task := s.list[0]
	s.list[0] = s.list[len(s.list)-1]
	s.list = s.list[:len(s.list)-1]
	s.heapify(0)
	return task
}

func (s *Scheduler) heapify(i int) {
	var leftChild int
	var rightChild int
	var largestChild int

	for {
		leftChild = 2*i + 1
		rightChild = 2*i + 2
		largestChild = i

		if leftChild < len(s.list) && s.list[leftChild].Priority > s.list[largestChild].Priority {
			largestChild = leftChild
		}

		if rightChild < len(s.list) && s.list[rightChild].Priority > s.list[largestChild].Priority {
			largestChild = rightChild
		}

		if largestChild == i {
			break
		}

		s.list[i], s.list[largestChild] = s.list[largestChild], s.list[i]
	}
}

func main() {

	task1 := Task{Identifier: 1, Priority: 10}
	task2 := Task{Identifier: 2, Priority: 20}
	task3 := Task{Identifier: 3, Priority: 30}
	task4 := Task{Identifier: 4, Priority: 40}
	task5 := Task{Identifier: 5, Priority: 50}

	scheduler := Scheduler{}
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
