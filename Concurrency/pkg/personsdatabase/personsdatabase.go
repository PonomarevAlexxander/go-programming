package personsdatabase

import (
	"go/course/concurrency/pkg/person"
	"sync"
)

type DataBase struct {
	persons         map[int]person.Person
	personByPhone   map[string]person.Person
	personByAddress map[string]person.Person
	personByName    map[string]person.Person
	lock            sync.RWMutex
}

func Init() *DataBase {
	return &DataBase{
		persons:         make(map[int]person.Person),
		personByPhone:   make(map[string]person.Person),
		personByAddress: make(map[string]person.Person),
		personByName:    make(map[string]person.Person),
		lock:            sync.RWMutex{},
	}
}

func (this *DataBase) Put(person person.Person) {
	this.lock.Lock()
	defer this.lock.Unlock()
	this.persons[person.Id] = person
	this.personByAddress[person.Address] = person
	this.personByName[person.Name] = person
	this.personByPhone[person.PhoneNumber] = person
}

func (this *DataBase) Remove(id int) (oldValue person.Person) {
	this.lock.Lock()
	defer this.lock.Unlock()

	oldValue, ok := this.persons[id]
	if !ok {
		return
	}
	delete(this.persons, id)
	delete(this.personByAddress, oldValue.Address)
	delete(this.personByName, oldValue.Name)
	delete(this.personByPhone, oldValue.PhoneNumber)
	return
}

func (this *DataBase) GetByName(name string) person.Person {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.personByName[name]
}

func (this *DataBase) GetByPhone(phone string) person.Person {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.personByPhone[phone]
}

func (this *DataBase) GetByAddress(address string) person.Person {
	this.lock.RLock()
	defer this.lock.RUnlock()
	return this.personByAddress[address]
}
