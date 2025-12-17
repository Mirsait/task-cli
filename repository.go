package main

import (
	"log"
	"slices"
	"time"
)

type Repository struct {
	tasks    []Task
	filename string
}

func (repo *Repository) load() {
	repo.tasks = loadData(repo.filename)
}

func (repo *Repository) getAll() []Task {
	return repo.tasks
}

func (repo *Repository) filter(predicate func(t Task) bool) []Task {
	var result []Task
	for _, t := range repo.tasks {
		if predicate(t) {
			result = append(result, t)
		}
	}
	return result
}

func (repo *Repository) add(text string) {
	id := 1
	if len(repo.tasks) > 0 {
		last := repo.tasks[len(repo.tasks)-1]
		id = last.Id + 1
	}
	repo.tasks = append(repo.tasks, newTask(id, text))
	repo.save()
}

func (repo *Repository) delete(id int) {
	if !repo.isExists(id) {
		log.Printf("The task with ID:%d not found.\n", id)
		return
	}

	tasks := slices.DeleteFunc(repo.tasks, func(t Task) bool {
		return t.Id == id
	})
	repo.tasks = tasks
	repo.save()
	log.Printf("The task with ID:%d has been deleted.\n", id)
}

func (repo *Repository) update(id int, newText string) {
	if !repo.isExists(id) {
		log.Printf("The task with ID:%d has not found.\n", id)
		return
	}
	result := make([]Task, len(repo.tasks))
	copy(result, repo.tasks)
	for j, v := range repo.tasks {
		if v.Id == id {
			result[j].Description = newText
			result[j].UpdatedAt = time.Now()
			break
		}
	}
	repo.tasks = result
	repo.save()
	log.Printf("The task with ID:%d has been updated.\n", id)
}

func (repo *Repository) save() {
	saveData(repo.filename, repo.tasks)
}

func (repo *Repository) mark(id int, status Status) {
	if !repo.isExists(id) {
		log.Printf("The task with ID:%d has not found.\n", id)
		return
	}
	result := make([]Task, len(repo.tasks))
	copy(result, repo.tasks)
	for j, v := range repo.tasks {
		if v.Id == id {
			result[j].Status = status
			result[j].UpdatedAt = time.Now()
			break
		}
	}
	repo.tasks = result
	repo.save()
	log.Printf("The task with ID:%d has been marked as `%s`.\n", id, status)
}

func (repo *Repository) isExists(id int) bool {
	return slices.ContainsFunc(
		repo.tasks,
		func(t Task) bool { return t.Id == id })
}
