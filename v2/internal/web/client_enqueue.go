package web

//Enqueue adds one element to the queue
func (c *client) Enqueue(value []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	c.queueData = append(c.queueData, value)
}

