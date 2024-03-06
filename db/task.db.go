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
func InitTaskRecords() *TaskRecord { // esta funcion retorna un puntero a TaskRecord
	recordInstance := make(map[int]model.Task)
	return &TaskRecord{ // retornamos la estructura TaskRecord ya inicializada
		RecordID: 0,
		Tasks:    recordInstance,
	}
}

// func crear tarea esta funcion es un metodo de la estructura TaskRecord el cual recibe un id un titulo y el contenido de la tarea
func (TR *TaskRecord) CreateTask(id int, title, content string) model.Task { // esta retorna un objeto tarea
	// aumentamos el valor de RecordID cada vez que hacemos un registro porque ID 1 para representar el registro numero 1 y asi consecutivamente
	TR.RecordID++
	// creamos un objeto tarea y llenamos los campos de la struct tarea con los parametros que la funcion recibe
	newtask := model.Task{
		Id:      id,
		Title:   title,
		Content: content,
	}
	// creamos una nueva llave en el mapa de tareas de TR con el valor del campo Id de el objeto tarea que estamos instanciando el cual dicha clave contiene el valor del objeto tarea que instanciamos en la funcion
	TR.Tasks[newtask.Id] = newtask

	// por ultimo retornamos el objeto tarea que acabamos de instanciar
	return newtask
}

func (TR TaskRecord) GetAllTasks() []model.Task {
	// Creamos un nuevo slice de tareas vacío
	tasks := make([]model.Task, 0, len(TR.Tasks))

	// Iteramos sobre el mapa Tasks de la estructura TaskRecord
	for _, task := range TR.Tasks {
		// Agregamos cada tarea al slice tasks
		tasks = append(tasks, task)
	}

	// Retornamos el slice tasks que contiene todas las tareas almacenadas en TaskRecord
	return tasks
}

// GetTaskByID obtiene una tarea por su ID.
// Recibe el ID de la tarea a buscar y devuelve la tarea correspondiente si existe.
func (TR TaskRecord) GetTaskByID(id int) (model.Task, error) {
	// Buscamos la tarea en el mapa Tasks de la estructura TaskRecord con la llave o Id de la tarea
	task, ok := TR.Tasks[id]
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
