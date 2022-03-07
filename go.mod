module github.com/semenovem/bot

go 1.16

require (
	github.com/caarlos0/env/v6 v6.7.1
	github.com/pkg/errors v0.9.1
	github.com/semenovem/bsf v0.0.0
	github.com/sirupsen/logrus v1.8.1
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b
)

replace github.com/semenovem/bsf => ./submod/bsf
