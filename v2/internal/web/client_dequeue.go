package web

//Dequeue removes one element from the list
func (c *client) Dequeue() (bool, []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if !c.HasElements() {
		return false, nil
	}

	result := c.queueData[0]
	c.queueData = c.queueData[1:]
	return true, result
}
