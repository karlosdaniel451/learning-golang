package main

import (
	"fmt"
	"sync"
	"time"

	"github.com/ugurcsen/gods-generic/queues/arrayqueue"
)

type Ticket struct {
	number int
}

func NewTicket() *Ticket {
	return &Ticket{}
}

type TicketOffice struct {
	tickets                     arrayqueue.Queue[Ticket]
	ticketsMutex                sync.Mutex
	maxNumberOfTickets          int
	nextTicketNumberToBeCreated int
	cond                        sync.Cond
}

func newTicketOffice(maxNumberOfTickets int) *TicketOffice {
	return &TicketOffice{
		tickets:                     *arrayqueue.New[Ticket](),
		ticketsMutex:                sync.Mutex{},
		maxNumberOfTickets:          maxNumberOfTickets,
		nextTicketNumberToBeCreated: 1,
		cond:                        *sync.NewCond(&sync.Mutex{}),
	}
}

func (ticketOffice *TicketOffice) numberOfTickets() int {
	ticketOffice.ticketsMutex.Lock()
	defer ticketOffice.ticketsMutex.Unlock()

	return ticketOffice.tickets.Size()
}

func (ticketOffice *TicketOffice) popTicket() (ticket *Ticket, ok bool) {
	ticketOffice.cond.L.Lock()
	defer ticketOffice.cond.L.Unlock()

	for ticketOffice.tickets.Size() == 0 {
		fmt.Println("Empty tickets queue, waiting for one item be produced...")
		ticketOffice.cond.Wait()
		fmt.Println("Wait for one item to be produced finished...")
	}

	dequeuedTicket, wasDequeuingOk := ticketOffice.tickets.Dequeue()
	ticketOffice.cond.Signal()
	return &dequeuedTicket, wasDequeuingOk
}

func (ticketOffice *TicketOffice) addTicket() (newCreatedTicket *Ticket) {
	ticketOffice.cond.L.Lock()
	defer ticketOffice.cond.L.Unlock()

	for ticketOffice.tickets.Size() == ticketOffice.maxNumberOfTickets {
		fmt.Println("Full tickets queue, waiting for one item be consumed...")
		ticketOffice.cond.Wait()
		fmt.Println("Wait for one item to be consumed finished...")
	}

	newTicket := *NewTicket()
	newTicket.number = ticketOffice.nextTicketNumberToBeCreated
	ticketOffice.nextTicketNumberToBeCreated++

	ticketOffice.tickets.Enqueue(newTicket)
	ticketOffice.cond.Signal()
	return &newTicket
}

func main() {
	const numberOfProducers int = 30
	const numberOfConsumers int = 30
	const maxNumberOfTickets int = 20

	var wg sync.WaitGroup

	ticketOffice := *newTicketOffice(maxNumberOfTickets)

	fmt.Printf("Max number of tickets: %d\n\n", maxNumberOfTickets)

	for i := 0; i < numberOfProducers; i++ {
		// go wg.Add(1)
		wg.Add(1)
		go func() {
			defer wg.Done()
			produceTickets(&ticketOffice, time.Millisecond*500)
		}()
	}

	for i := 0; i < numberOfConsumers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			consumeTickets(&ticketOffice, time.Millisecond*200)
		}()
	}

	wg.Wait()

	fmt.Printf("Final number of tickets: %d\n", ticketOffice.tickets.Size())
}

func produceTickets(ticketOffice *TicketOffice, delay time.Duration) {
	time.Sleep(delay)
	ticketProduced := ticketOffice.addTicket()
	fmt.Printf("Ticket produced with number %d\n", ticketProduced.number)

	// The following function call may generate race conditions
	// fmt.Printf("Current number of tickets: %d\n\n", ticketOffice.numberOfTickets())
}

func consumeTickets(ticketOffice *TicketOffice, delay time.Duration) {
	time.Sleep(delay)
	ticketConsumed, ok := ticketOffice.popTicket()
	if !ok {
		panic("")
	}
	fmt.Printf("Ticket consumed with number %d\n", ticketConsumed.number)

	// The following function call may generate race conditions
	// fmt.Printf("Current number of tickets: %d\n\n", ticketOffice.numberOfTickets())
}
