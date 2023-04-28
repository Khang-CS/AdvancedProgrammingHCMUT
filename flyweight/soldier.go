package flyweight

import "fmt"

/*Soldier is struct*/

type Soldier struct {
	name string
}

func newSoldier(name string) *Soldier {
	return &Soldier{
		name: name,
	}
}

/*Promote is function*/
func (s *Soldier) Promote(context *Context) {
	fmt.Printf("%s %s promoted %d\n", s.name, context.id, context.star)
}
