package config

type DBType string

const (
	MemoryDBType DBType = "memory"
)

func (e DBType) String() string {
	return string(e)
}
