package main

import (
	"errors"
	"fmt"
	"log"
)

type Cache[T any] struct {
	Data map[string]*T
}

func CreateCache[T any](initialCapacity int) (*Cache[T], error) {
	if initialCapacity < 1 {
		return nil, errors.New("error: initial capacity of cache must be greater than 0")
	}

	return &Cache[T]{
		Data: make(map[string]*T, initialCapacity),
	}, nil
}

func (cache *Cache[T]) GetLength() int {
	return len(cache.Data)
}

func (cache *Cache[T]) Lookup(key string) (*T, bool) {
	value, ok := cache.Data[key]
	if ok {
		return value, true
	}

	return new(T), false
}

func (cache *Cache[T]) Upsert(key string, value *T) {
	cache.Data[key] = value
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
		Email: "john@gmail.com",
		Password: "fakepassword",
	}
	secondUser := UserLogin{
		Username: "janedoe123",
		Email: "janedoe@gmail.com",
		Password: "fakepassword",
	}
	thirdUser := UserLogin{
		Username: "smith321",
		Email: "smith@gmail.com",
		Password: "fakepassword",
	}

	userLoginCache.Upsert(firstUser.Username, &firstUser)
	fmt.Printf("Current cache length: %d\n", userLoginCache.GetLength())

	userLoginCache.Upsert(secondUser.Username, &secondUser)
	fmt.Printf("Current cache length: %d\n", userLoginCache.GetLength())

	userLoginCache.Upsert(thirdUser.Username, &thirdUser)
	fmt.Printf("Current cache length: %d\n", userLoginCache.GetLength())

	user, _ := userLoginCache.Lookup(firstUser.Username)
	fmt.Printf("%#v\n", *user)

	user, _ = userLoginCache.Lookup(secondUser.Username)
	fmt.Printf("%#v\n", *user)

	user, _ = userLoginCache.Lookup(thirdUser.Username)
	fmt.Printf("%#v\n", *user)
}
