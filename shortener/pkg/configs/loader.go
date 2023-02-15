package configs

import (
	"gopkg.in/yaml.v3"
)

var skipLoading bool = false

func SkipLoading() {
	skipLoading = true
}

type Loader[T any] struct {
	Config T       `yaml:",inline"`
	Env    *Env    `yaml:"env"`
	Secret *Secret `yaml:"secret"`
}

type tempLoader[T any] Loader[T]

func (l *Loader[T]) UnmarshalYAML(value *yaml.Node) error {
	var tmp tempLoader[T]
	err := value.Decode(&tmp)
	if err != nil {
		return err
	}
	*l = Loader[T](tmp)

	if skipLoading {
		return nil
	}

	var loaders []loader
	if l.Env != nil {
		loaders = append(loaders, l.Env)
	}
	if l.Secret != nil {
		loaders = append(loaders, l.Secret)
	}

	return l.load(loaders...)
}

type loader interface {
	load() ([]byte, error)
}

func (l *Loader[T]) load(loaders ...loader) error {
	for _, lr := range loaders {
		bytes, err := lr.load()
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(bytes, &l.Config)
		if err != nil {
			return err
		}
	}
	return nil
}
