package main

import (
	"errors"
	"github.com/rs/xid"
	"sync"
)

var (
	list []Todo
	mtx  sync.RWMutex
)

type Todo struct {
	ID       string `json:"id"`
	Message  string `json:"message"`
	Complete bool   `json:"complete"`
}

func Get() []Todo {
	return list
}

func GetItem(id string) (*Todo, error) {
	for i, t := range list {
		if t.ID == id {
			return &list[i], nil
		}
	}
	return nil, errors.New("not found")
}

func Add(message string) string {
	t := Todo{ID: xid.New().String(), Message: message, Complete: false}
	mtx.Lock()
	list = append(list, t)
	mtx.Unlock()
	return t.ID
}

func Delete(id string) error {
	mtx.Lock()
	defer mtx.Unlock()
	for i, t := range list {
		if t.ID == id {
			list = append(list[:i], list[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func Complete(id string) error {
	mtx.Lock()
	defer mtx.Unlock()
	for i, t := range list {
		if t.ID == id {
			list[i].Complete = true
			return nil
		}
	}
	return errors.New("not found")
}
