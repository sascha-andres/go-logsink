package web

//HasElements returns true if there are elements to read from queue
func (c *client) HasElements() bool {
	return len(c.queueData) > 0
}

