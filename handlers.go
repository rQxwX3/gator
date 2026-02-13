package main

func (c *commands) register(name string, f func(*state, command) error) error {
	if _, ok := c.cmdMap[name]; !ok {
		c.cmdMap[name] = f
	}

	return nil
}
