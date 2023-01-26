package storage

type Storage interface {
	Get(m interface{}) (interface{}, int)
	Put(m interface{}) int
}
