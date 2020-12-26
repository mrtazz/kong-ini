package kongini

import (
	"github.com/alecthomas/kong"
	"github.com/pkg/errors"
	"gopkg.in/ini.v1"
	"io"
	"io/ioutil"
	"log"
)

var debug = false

func logDebug(format string, data ...interface{}) {
	if debug {
		log.Printf(format, data...)
	}
}

// Resolver is the receiver for the interface methods
type Resolver struct {
	config *ini.File
}

// Validate implements the kong Resolver interface
func (r *Resolver) Validate(app *kong.Application) error {
	return nil
}

// Resolve implements the kong Resolver interface
func (r *Resolver) Resolve(context *kong.Context, parent *kong.Path, flag *kong.Flag) (interface{}, error) {
	logDebug("Getting key for '%s'.\n", flag.Name)
	logDebug("Parent is '%s'.\n", parent.Node().Name)

	var section string
	if parent.Node().Type != kong.ApplicationNode && parent.Node().Name != "" {
		section = parent.Node().Name
	}

	// return if the key isn't configured
	if !r.config.Section(section).HasKey(flag.Name) {
		return nil, nil
	}

	value := r.config.Section(section).Key(flag.Name).Value()
	logDebug("Found data '%v' for '%s'.\n", value, flag.Name)

	return value, nil
}

// Loader is a Kong configuration loader for HCL.
func Loader(r io.Reader) (kong.Resolver, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	cfg, err := ini.Load(data)
	if err != nil {
		return nil, errors.Wrap(err, "unable to load ini file")
	}

	return &Resolver{config: cfg}, nil
}
