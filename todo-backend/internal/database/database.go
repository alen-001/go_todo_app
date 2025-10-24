package database

import "todo-backend/internal/models"

var Todos = map[int]models.Todo{
	1:{
		Id:        1,
		Name:      "Go to gym",
		Status: models.Pending,
	},
	2:{
		Id: 	   2,
		Name:      "Learn go",
		Status: models.InProgress,
	},
}

func Database() {

}