package config

import (
	"gopkg.in/yaml.v3"
)

type Loader[T any] struct {
	C      T       `yaml:",inline"`
	Env    *Env    `yaml:"env"`
	Secret *Secret `yaml:"secret"`
}

type tempLoader[T any] Loader[T]

func (l *Loader[T]) UnmarshalYAML(value *yaml.Node) error {
	loaders := []loader{l.Env, l.Secret}

	var tmp tempLoader[T]
	err := value.Decode(&tmp)
	if err != nil {
		return err
	}
	*l = Loader[T](tmp)

	return l.load(loaders...)
}

type loader interface {
	load() ([]byte, error)
}

func (l *Loader[T]) load(loaders ...loader) error {
	for _, lr := range loaders {
		if lr == nil {
			return nil
		}
		bytes, err := lr.load()
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(bytes, &l.C)
		if err != nil {
			return err
		}
	}
	return nil
}
