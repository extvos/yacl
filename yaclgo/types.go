package yaclgo

// HandleFunc the function handling call backs.
// handle function will be called as (directive string, params ...interface{}),
// return none nil error will cause parsing stop if parser was configured as Strict=true.
type HandleFunc func(string, ...interface{}) error

// Handler the handler interface which implemented Handle function.
type Handler interface {
	Handle(string, ...interface{}) error
}

// Config for parser.
type Config struct {
	Strict bool
}

// Parser the exported struct for yaclgo parser.
type Parser struct {
	cfg *Config
}
