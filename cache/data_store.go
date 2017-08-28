package cache

type DataStore interface {
	GetCount() (uint, error)
	Set(key string, object interface{}) error
	Find(key string, object interface{}) error
	RemoveOld() error
	Close()
}
