package flyweight

/*Context is struct*/
type Context struct {
	id   string
	star int
}

/*NewContext is function*/
func NewContext(id string, star int) *Context {
	return &Context{
		id:   id,
		star: star,
	}
}
