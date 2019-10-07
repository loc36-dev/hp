package lib

func NewConf () (*Conf, error) {}

type Conf map[string]string

func (c *Conf) Get (name string) (string) {}
