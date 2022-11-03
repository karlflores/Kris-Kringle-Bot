package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
)

type Person struct {
	Name  string `json:"name"`
	Email string `json:"email"` 
}

type SantaHat struct {
	people []Person
}

func createHat(filepath string) (SantaHat, error){
	ss := SantaHat{}

	people := []Person{}

	jsonFile, err := os.Open(filepath)
	defer jsonFile.Close()	

	if err != nil { 
		log.Fatal(fmt.Sprintf("Error reading from %s", filepath))
		return ss, err 
	}

	bytes, err := ioutil.ReadAll(jsonFile)

	if err != nil {
		log.Fatal(fmt.Sprintf("Error reading from %s", filepath))
		return ss, err
	}
	json.Unmarshal([]byte(bytes), &people)
	for _,p := range people { 	
		ss.register(p)
	}

	return ss, nil
}

func (hat *SantaHat) register(person Person) {
	if hat.people == nil {
		hat.people = make([]Person, 0)
	}
	hat.people = append(hat.people, person)
}

func (hat *SantaHat) pair() map[Person]Person {
	pairings := make(map[Person]Person) 

	// first we have to shuffle the people array 
	mapping := []int{}
	for {
		mapping = rand.Perm(len(hat.people))
		flag := false
		for i := 0 ; i < len(mapping) ; i++ {
			if mapping[i] == i {
				flag = true
				break
			}
		}
		if !flag {
			break
		}
	}

	fmt.Print(mapping)
	// now we have the pairings
	for i := 0 ; i < len(mapping) ; i++ {
		pairings[hat.people[i]] = hat.people[mapping[i]]
	}

	return pairings
}