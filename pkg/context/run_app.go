package context

func (c *Context) RunApp() {
	authorDict := c.ParseRepository()
	c.Formatter.FormatAndPrint(authorDict)
}