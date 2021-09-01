package storage

type Storager interface {
	Exists(file string) bool
}
