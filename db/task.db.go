package db

import (
	"github.com/simple-rest-API/model"
)

// creamos la estructura de un nuevo registro en memoria de una tarea
type TaskRecord struct {
	// le asignamos un id al registro
	RecordID int
	// y creamos un campo tareas donde se almacenan las tareas en un mapa con un id unico y de valor una instancia de una tarea es decir un objeto tarea
	Tasks map[int]model.Task
}

// creamos una funcion que instancia la estructura TaskRecord para poder hacer nuevos registros en esa estructura ya que nesecitamos inicializar el mapa que contiene las tareas
func InitTaskRecords() TaskRecord { // esta funcion retorna un puntero a TaskRecord
	recordInstance := make(map[int]model.Task)
	return TaskRecord{ // retornamos la estructura TaskRecord ya inicializada
		RecordID: 0,
		Tasks:    recordInstance,
	}
}

// func crear tarea esta funcion es un metodo de la estructura TaskRecord que recibe un puntero a una estructura task
func (TR *TaskRecord) CreateTask(task *model.Task) error { // esta puede retornar un error
	TR.RecordID++
	if task == nil {
		return model.ErrTaskCanNotBeNil
	}

	TR.Tasks[task.Id] = *task

	// por ultimo retornamos el objeto tarea que acabamos de instanciar
	return nil
}

func (TR *TaskRecord) GetAllTasks() (model.Tasks, error) {
	// Creamos un nuevo slice de tareas vacío
	var tasks model.Tasks

	// Iteramos sobre el mapa Tasks de la estructura TaskRecord
	for _, task := range TR.Tasks {
		// Agregamos cada tarea al slice tasks
		tasks = append(tasks, task)
	}

	// Retornamos el slice tasks que contiene todas las tareas almacenadas en TaskRecord
	return tasks, nil
}

// GetTaskByID obtiene una tarea por su ID.
// Recibe el ID de la tarea a buscar y devuelve la tarea correspondiente si existe.
func (TR TaskRecord) GetTaskByID(ID int) (model.Task, error) {
	// Buscamos la tarea en el mapa Tasks de la estructura TaskRecord con la llave o Id de la tarea
	task, ok := TR.Tasks[ID]
	if !ok {
		// Si la tarea no existe (el ID es 0), devolvemos un error
		return model.Task{}, model.ErrtaskDoNotExist
	}

	// Si la tarea existe, la devolvemos
	return task, nil
}

func (TR *TaskRecord) UpdateTask(ID int, title, content string) (model.Task, error) {
	task, ok := TR.Tasks[ID]
	if !ok {
		return model.Task{}, model.ErrCanotUpatetask
	}

	task.Title = title
	task.Content = content

	return task, nil
}

func (TR *TaskRecord) DeleteTask(ID int) error {
	_, ok := TR.Tasks[ID]
	if !ok {
		return model.ErrtaskDoNotExist
	}

	delete(TR.Tasks, ID)

	return nil
}
