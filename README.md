# kong-ini

A [kong](https://github.com/alecthomas/kong) configuration resolver for `.ini`
files.

## Usage

```
var cli struct {
    Config kong.ConfigFlag `help:"Load configuration."`
}
parser, err := kong.New(&cli, kong.Configuration(kongini.Loader, "/etc/myapp/config.ini", "~/.myapp.ini))
```

### Config Format
Given the following app definition:

```
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
```

This would be the corresponding `.ini` file:

```
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
```

## Installation

```
go get github.com/mrtazz/kong-ini
```

## Inspiration
- [kong yaml resolver](https://github.com/alecthomas/kong-yaml)
- [globalconf](https://github.com/rakyll/globalconf)
