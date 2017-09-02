package yaclgo

// VERSION the version of yaclgo as a constant string.
const VERSION = "0.0.1"

// New ...
// Create a new parser
func New(cfg *Config) (*Parser, error) {
	return nil, nil
}

// DefaultConfig ...
// Generate a new default config
func DefaultConfig() *Config {
	return &Config{Strict: true}
}

// Parse ...
// Start to parse the file content.
func (p *Parser) Parse(filename string) {

}

// Close ...
// Close the parser and the opened file handle(s).
func (p *Parser) Close() {

}
