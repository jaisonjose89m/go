package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/jaisonjose89m/go/concurrency/models"
)

var cache = map[int]models.Book{}
var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

func main() {
	wg := &sync.WaitGroup{}
	mutex := &sync.RWMutex{}
	for i := 0; i < 10; i++ {
		id := rnd.Intn(10) + 1
		wg.Add(2)
		go queryCacheAndPrint(id, wg, mutex)
		go queryDatabaseAndPrint(id, wg, mutex)
	}
	wg.Wait()
}

func queryDatabaseAndPrint(id int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	bookFromDb, gotIt := queryDatabase(id, mutex)
	if gotIt {
		fmt.Printf("From DB\n%v", bookFromDb)
	}
	wg.Done()
}

func queryCacheAndPrint(id int, wg *sync.WaitGroup, mutex *sync.RWMutex) {
	book, ok := queryCache(id, mutex)
	if ok {
		fmt.Printf("From cache\n%v", book)
	}
	wg.Done()
}

func queryCache(id int, mutex *sync.RWMutex) (models.Book, bool) {
	mutex.RLock()
	book, ok := cache[id]
	mutex.RUnlock()
	return book, ok
}

func queryDatabase(id int, mutex *sync.RWMutex) (models.Book, bool) {
	time.Sleep(100 * time.Millisecond)
	for _, b := range models.GetBooks() {
		if b.ID == id {
			mutex.Lock()
			cache[id] = b
			mutex.Unlock()
			return b, true
		}
	}
	return models.Book{}, false
}
