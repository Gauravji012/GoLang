package storage

// interfaces is an important concept. with the help of interfaces we can make plugins easily. we can handle failures , switching easily.
// we can switch from one application to another one easily. like sqlite -> postgresql

type Storage interface {
	CreateStudent(name string, email string, age int) (int64, error)
}
