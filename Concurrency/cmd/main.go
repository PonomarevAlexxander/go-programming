package main

import (
	"fmt"
	"go/course/concurrency/pkg/person"
	"go/course/concurrency/pkg/personsdatabase"
	"sync"
)

func storePerson(db *personsdatabase.DataBase, upstream <-chan person.Person, downstream chan<- string, wg *sync.WaitGroup) {
	defer wg.Done()
	for data := range upstream {
		db.Put(data)
		downstream <- data.Name
	}
}

func prccessNames(db *personsdatabase.DataBase, upstream <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for name := range upstream {
		fmt.Printf("Person %s is stored with id %d\n", name, db.GetByName(name).Id)
	}
}

func main() {
	var producer sync.WaitGroup
	var consumer sync.WaitGroup
	persons := personsdatabase.Init()
	upstream := make(chan person.Person, 3)
	downstream := make(chan string)

	for index := 0; index < 3; index++ {
		producer.Add(1)
		consumer.Add(1)
		go storePerson(persons, upstream, downstream, &producer)
		go prccessNames(persons, downstream, &consumer)
	}

	list := []person.Person{
		{1, "Alex", "+7575", "London"},
		{255, "Simon", "+8181", "Berlin"},
		{123, "Bill", "+9999", "St. Petersburg"},
		{11, "Tim", "+1234", "Moscow"},
		{11, "Tim", "+7768", "Moscow"},
	}
	for _, person := range list {
		upstream <- person
	}

	close(upstream)
	producer.Wait()
	close(downstream)
	consumer.Wait()

	fmt.Println("All data proccessed")
	tim := persons.GetByName("Tim")
	fmt.Printf("%s last number is %s\n", tim.Name, tim.PhoneNumber)
}
