package cliff_test

import (
	"os"
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/cliff"
)

func TestFlags(t *testing.T) {
	is := is.New(t)

	type Config struct {
		host  string
		port  int
		https bool
	}
	config := Config{}
	flags := cliff.Flags{
		"host":  cliff.F(&config.host, 0, "127.0.0.1", "host to serve on"),
		"port":  cliff.F(&config.port, 'p', 8080, "port to listen to"),
		"https": cliff.F(&config.https, 0, true, "force https"),
	}
	args := []string{"example", "--host", "localhost"}
	err := flags.Parse(os.Stderr, args)
	is.NoErr(err)
	expected := Config{host: "localhost", port: 8080, https: true}
	is.Equal(config, expected)
}

func TestParse(t *testing.T) {
	is := is.New(t)

	type Config struct {
		host  string
		port  int
		https bool
	}

	args := []string{"example", "--host", "localhost"}
	config, err := cliff.Parse(os.Stderr, args, func(c *Config) cliff.Flags {
		return cliff.Flags{
			"host":  cliff.F(&c.host, 0, "127.0.0.1", "host to serve on"),
			"port":  cliff.F(&c.port, 'p', 8080, "port to listen to"),
			"https": cliff.F(&c.https, 0, true, "force https"),
		}
	})
	is.NoErr(err)
	expected := Config{host: "localhost", port: 8080, https: true}
	is.Equal(config, expected)
}
