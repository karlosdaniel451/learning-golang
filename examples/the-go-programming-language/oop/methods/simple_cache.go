package main

import (
	"errors"
	"fmt"
	"log"
	"sync"
)

type Cache[T any] struct {
	// `items` is an encapsulated field and cannot be used from another package.
	items map[string]*T

	// It is an embedded anonymous field, thus its public fields and methods
	// are promoted to `Cache`.
	sync.Mutex
}

func CreateCache[T any](initialCapacity int) (*Cache[T], error) {
	if initialCapacity < 1 {
		return nil, errors.New("error: initial capacity of cache must be greater than 0")
	}

	return &Cache[T]{
		items: make(map[string]*T, initialCapacity),
	}, nil
}

func (cache *Cache[T]) GetLength() int {
	return len(cache.items)
}

func (cache *Cache[T]) Lookup(key string) (*T, bool) {
	// First implementation:
	// cache.Lock()
	// value, ok := cache.items[key]
	// cache.Unlock()
	// if ok {
	// 	return value, true
	// }
	// return new(T), false

	cache.Lock()
	value, ok := cache.items[key]
	cache.Unlock()
	return value, ok
}

func (cache *Cache[T]) Upsert(key string, value *T) {
	cache.items[key] = value
}

type UserLogin struct {
	Username string
	Email    string
	Password string
}

func main() {
	userLoginCache, err := CreateCache[UserLogin](100)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Current cache length: %d\n", userLoginCache.GetLength())

	firstUser := UserLogin{
		Username: "johndoe123",
		Email:    "john@gmail.com",
		Password: "fakepassword",
	}
	secondUser := UserLogin{
		Username: "janedoe123",
		Email:    "janedoe@gmail.com",
		Password: "fakepassword",
	}
	thirdUser := UserLogin{
		Username: "smith321",
		Email:    "smith@gmail.com",
		Password: "fakepassword",
	}

	userLoginCache.Upsert(firstUser.Username, &firstUser)
	fmt.Printf("Current cache length: %d\n", userLoginCache.GetLength())

	userLoginCache.Upsert(secondUser.Username, &secondUser)
	fmt.Printf("Current cache length: %d\n", userLoginCache.GetLength())

	userLoginCache.Upsert(thirdUser.Username, &thirdUser)
	fmt.Printf("Current cache length: %d\n", userLoginCache.GetLength())

	usernameToLookup := firstUser.Username
	user, ok := userLoginCache.Lookup(usernameToLookup)
	if !ok {
		fmt.Printf("Not found user with username \"%s\".\n", usernameToLookup)
	} else {
		fmt.Printf("%#v\n", *user)
	}

	usernameToLookup = secondUser.Username
	user, ok = userLoginCache.Lookup(usernameToLookup)
	if !ok {
		fmt.Printf("Not found user with username \"%s\".\n", usernameToLookup)
	} else {
		fmt.Printf("%#v\n", *user)
	}

	usernameToLookup = thirdUser.Username
	user, ok = userLoginCache.Lookup(usernameToLookup)
	if !ok {
		fmt.Printf("Not found user with username \"%s\".\n", usernameToLookup)
	} else {
		fmt.Printf("%#v\n", *user)
	}

	usernameToLookup = "any_username"
	user, ok = userLoginCache.Lookup(usernameToLookup)
	if !ok {
		fmt.Printf("Not found user with username \"%s\".\n", usernameToLookup)
	} else {
		fmt.Printf("%#v\n", *user)
	}

}
