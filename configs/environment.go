package configs

type Environment string

const (
	Development Environment = "development"
	Production  Environment = "production"
)

func (e Environment) String() string {
	return string(e)
}
