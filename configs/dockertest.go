package configs

import (
	"github.com/kelseyhightower/envconfig"
)

type DockerTest struct {
	Driver  string
	Host    string
	Port    string
	Name    string
	User    string
	Pass    string
	SslMode string
}

func DockerTestCfg() DockerTest {
	var dockerTest DockerTest
	envconfig.MustProcess("DOCKERTEST", &dockerTest)

	return dockerTest
}
