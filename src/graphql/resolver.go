package graphql

import (
	"backend/src/graphql/modules/user"
	"fmt"
	"time"

	"math/rand"
)

type Resolver struct {
	OnCreatedUserEvent      chan *user.User
	OnCreatedUserSubscriber chan *user.OnCreatedUserSubscriber
	user.UserResolver
}

func (Resolver) Hello() string {
	fmt.Println("Call hello query")
	return "hello world"
}

func NewResolver() *Resolver {
	resolver := Resolver{
		OnCreatedUserEvent:      make(chan *user.User),
		OnCreatedUserSubscriber: make(chan *user.OnCreatedUserSubscriber),
	}

	go resolver.broadcastOnCreatedUser()

	return &resolver
}

func (r *Resolver) broadcastOnCreatedUser() {
	subscribers := map[string]*user.OnCreatedUserSubscriber{}
	unsubscribe := make(chan string)

	for {
		select {
		case id := <-unsubscribe:
			delete(subscribers, id)
		case s := <-r.OnCreatedUserSubscriber:
			subscribers[randomID()] = s
		case e := <-r.OnCreatedUserEvent:
			for id, s := range subscribers {
				go func(id string, s *user.OnCreatedUserSubscriber) {
					select {
					case <-s.Stop:
						unsubscribe <- id
						return
					default:
					}

					select {
					case <-s.Stop:
						unsubscribe <- id
					case s.Events <- e:
					case <-time.After(time.Second):
					}
				}(id, s)
			}

		}

	}
}

func randomID() string {
	var letter = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")

	b := make([]rune, 16)
	for i := range b {
		b[i] = letter[rand.Intn(len(letter))]
	}
	return string(b)
}
