package yametrics

// MemStorage struct
type MemStorage struct {
	Gauge    map[string]float64
	Counter  map[string]int64
	FilePath string
}

func NewMemStorage(filepath string) *MemStorage {
	return &MemStorage{
		Gauge:    make(map[string]float64),
		Counter:  make(map[string]int64),
		FilePath: filepath,
	}
}
