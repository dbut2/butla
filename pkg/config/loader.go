package config

type Loader struct {
	Env    *Env    `yaml:"Env"`
	Secret *Secret `yaml:"gsmResourceID"`
}

var _ loader = new(Loader)

func (l Loader) load(c any) error {
	if l.Env != nil {
		return l.Env.load(c)
	}
	if l.Secret != nil {
		return l.Secret.load(c)
	}
	return nil
}

type loader interface {
	load(c any) error
}

func Load(l loader) error {
	if l == nil {
		return nil
	}
	return l.load(l)
}
