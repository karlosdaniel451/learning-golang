package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"sync"
)

type Product struct {
	Code               string `json:"code"`
	Name               string `json:"name"`
	PriceInDollarCents int    `json:"price_in_dollar_cents"`
}

func (product Product) String() string {
	return fmt.Sprintf(
		"Name: %s, Price: $%.2f",
		product.Name, float64(product.PriceInDollarCents/100))
}

const (
	host string = "localhost"
	port int    = 8000
)

var defaultProducts = []*Product{
	&Product{
		Code:               "1",
		Name:               "Shoes",
		PriceInDollarCents: 50_00,
	},
	&Product{
		Code:               "2",
		Name:               "Socks",
		PriceInDollarCents: 5_00,
	},
}

func insertDefaultProducts() {
	for _, product := range defaultProducts {
		productsCacheDB.Upsert(product.Code, product)
	}
}

var productsCacheDB *Cache[Product]

func main() {
	var err error
	productsCacheDB, err = CreateCache[Product](100)
	if err != nil {
		log.Fatal(err)
	}

	insertDefaultProducts()

	mux := http.NewServeMux()

	mux.HandleFunc("/products", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getAllProductsHandler(w, r)
		case http.MethodPost:
			insertProductHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			log.Printf("%s %s %s %s", r.Method, r.URL, r.Proto, "405 Method Not Allowed")
		}
	})

	mux.HandleFunc("/products/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			getProductByCodeHandler(w, r)
		default:
			http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
			log.Printf("%s %s %s %s", r.Method, r.URL, r.Proto, "405 Method Not Allowed")
		}
	})

	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", host, port),
		Handler: mux, // *ServeMux implements http.Handler, thus, http.Server as well.
	}

	// // Assert at compile time that a value of `*SimpleBuffer`` satisfies `io.ReadWriter`.
	var _ http.Handler = new(http.Server).Handler
	var _ http.Handler = http.NewServeMux()

	log.Printf("Trying to listen requests at %s:%d...", host, port)
	log.Fatal(server.ListenAndServe())
}

func getAllProductsHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Fprintf(w, "%s", products) http.ResponseWriter embeds io.Writer
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(productsCacheDB.GetAllItems())
	log.Printf("%s %s %s %s", r.Method, r.URL, r.Proto, "200 OK")
}

func getProductByCodeHandler(w http.ResponseWriter, r *http.Request) {
	productCode := r.URL.Path[len("/products/"):]
	product, ok := productsCacheDB.Lookup(productCode)
	if !ok {
		http.Error(w, "Not found product with code "+productCode, http.StatusNotFound)
		log.Printf("%s %s %s %s", r.Method, r.URL, r.Proto, "404 Not Found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
	log.Printf("%s %s %s %s", r.Method, r.URL, r.Proto, "200 OK")
}

func insertProductHandler(w http.ResponseWriter, r *http.Request) {
	var product Product
	bodyJSONDecoder := json.NewDecoder(r.Body)

	err := bodyJSONDecoder.Decode(&product)
	// Check for errors in the body decoding.
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log.Printf("%s %s %s %s", r.Method, r.URL, r.Proto, "400 Bad Request")
		return
	}

	// Check if product already exists
	_, productAlreadyExists := productsCacheDB.Lookup(product.Code)
	if productAlreadyExists {
		http.Error(w, "Product with code "+product.Code+" already exists", http.StatusConflict)
		log.Printf("%s %s %s %s", r.Method, r.URL, r.Proto, "409 Conflict")
		return
	}

	productsCacheDB.Upsert(product.Code, &product)

	// json.NewEncoder(w).Encode(product)

	w.Header().Add("Location", "/products/"+product.Code)
	w.WriteHeader(http.StatusCreated)
	log.Printf("%s %s %s %s", r.Method, r.URL, r.Proto, "201 Created")
}

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

func (cache *Cache[T]) GetAllItems() []T {
	cache.Lock()
	defer cache.Unlock()

	items := make([]T, 0, cache.GetLength())

	for _, item := range cache.items {
		items = append(items, *item)
	}

	return items
}

func (cache *Cache[T]) Lookup(key string) (*T, bool) {
	cache.Lock()
	value, ok := cache.items[key]
	cache.Unlock()
	return value, ok
}

func (cache *Cache[T]) Upsert(key string, value *T) {
	cache.items[key] = value
}
