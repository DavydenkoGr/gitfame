package context

// method runs an application: scan files and print results
func (c *Context) RunApp() {
	authorDict := c.ParseRepository()
	c.Formatter.FormatAndPrint(authorDict)
}
