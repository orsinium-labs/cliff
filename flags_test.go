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
	flags := cliff.Flags("example",
		cliff.Fn(&config.host, "host", 'h', "127.0.0.1", "host to serve on"),
		cliff.Fn(&config.port, "port", 'p', 8080, "port to listen to"),
		cliff.Fn(&config.https, "https", 0, true, "force https"),
	)
	err := flags.Parse([]string{"--host", "localhost"})
	is.NoErr(err)
	expected := Config{host: "localhost", port: 8080, https: true}
	is.Equal(config, expected)
}
