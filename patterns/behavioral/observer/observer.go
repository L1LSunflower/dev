package observer

import "fmt"

type Notificator struct {
	UsersList map[string]Observer
}

type Subject interface {
	Follow(observer Observer)
	Unfollow(observer Observer)
	NotifyAll()
}

type Observer interface {
}

type Customer struct {
	id string
}

func (c *Customer) update(itemName string) {
	fmt.Printf("Sending email to customer %s for item %s\n", c.id, itemName)
}

func (c *Customer) getID() string {
	return c.id
}
