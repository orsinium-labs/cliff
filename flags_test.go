package cliff_test

import (
	"testing"

	"github.com/matryer/is"
	"github.com/orsinium-labs/cliff"
)

func Test(t *testing.T) {
	is := is.New(t)

	type Config struct {
		host  string
		port  int
		https bool
	}
	config := Config{}
	flags := cliff.Flags{
		"host":  cliff.Fn(&config.host, 'h', "127.0.0.1", "host to serve on"),
		"port":  cliff.Fn(&config.port, 'p', 8080, "port to listen to"),
		"https": cliff.Fn(&config.https, 0, true, "force https"),
	}
	args := []string{"example", "--host", "localhost"}
	err := flags.Parse(args)
	is.NoErr(err)
	expected := Config{host: "localhost", port: 8080, https: true}
	is.Equal(config, expected)
}
