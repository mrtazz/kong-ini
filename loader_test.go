package kongini

import (
	"github.com/alecthomas/kong"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var testdata = `
debug = true
string = "a string"
int = 1
slice = bla,foo bar,test
intslice = 1,2,3
floatslice = 1.2,2.3,3.4
map = """a=1;b=2;c=3"""

[command]
isFoo = false
value = bar
`

var cli struct {
	Config     kong.ConfigFlag `help:"Load configuration."`
	Debug      bool            `kong:"name='debug'"`
	String     string          `kong:"name='string'"`
	Int        int             `kong:"name='int'"`
	Slice      []string        `kong:"name='slice'"`
	Intslice   []int           `kong:"name='intslice'"`
	Floatslice []float64       `kong:"name='floatslice'"`
	Map        map[string]int  `kong:"name='map'"`
	Command    struct {
		IsFoo bool   `kong:"name='isFoo'"`
		Value string `kong:"name='value'"`
	} `kong:"cmd"`
	Default struct {
	} `kong:"cmd"`
}

func init() {
	debug = false
}

func parseCli(cmd []string) error {
	r := strings.NewReader(testdata)
	resolver, err := Loader(r)
	if err != nil {
		return err
	}
	parser, err := kong.New(&cli, kong.Resolvers(resolver))
	if err != nil {
		return err
	}
	_, err = parser.Parse(cmd)
	return err
}

func TestResolvingBoolValues(t *testing.T) {
	err := parseCli([]string{"default"})
	require.NoError(t, err)
	assert.Equal(t, true, cli.Debug)
}

func TestResolvingStringValues(t *testing.T) {
	err := parseCli([]string{"default"})
	require.NoError(t, err)
	assert.Equal(t, "a string", cli.String)
}

func TestResolvingSubcommandBoolValues(t *testing.T) {
	err := parseCli([]string{"command"})
	require.NoError(t, err)
	assert.Equal(t, false, cli.Command.IsFoo)
}

func TestResolvingSubcommandStringValues(t *testing.T) {
	err := parseCli([]string{"command"})
	require.NoError(t, err)
	assert.Equal(t, "bar", cli.Command.Value)
}

func TestResolvingIntValues(t *testing.T) {
	err := parseCli([]string{"default"})
	require.NoError(t, err)
	assert.Equal(t, 1, cli.Int)
}

func TestResolvingStringSliceValues(t *testing.T) {
	err := parseCli([]string{"default"})
	require.NoError(t, err)
	assert.Equal(t, []string{"bla", "foo bar", "test"}, cli.Slice)
}

func TestResolvingIntSliceValues(t *testing.T) {
	err := parseCli([]string{"default"})
	require.NoError(t, err)
	assert.Equal(t, []int{1, 2, 3}, cli.Intslice)
}

func TestResolvingFloatSliceValues(t *testing.T) {
	err := parseCli([]string{"default"})
	require.NoError(t, err)
	assert.Equal(t, []float64{1.2, 2.3, 3.4}, cli.Floatslice)
}

func TestResolvingMapValues(t *testing.T) {
	err := parseCli([]string{"default"})
	require.NoError(t, err)
	assert.Equal(t, map[string]int{"a": 1, "b": 2, "c": 3}, cli.Map)
}
