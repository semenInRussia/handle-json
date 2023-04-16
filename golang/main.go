package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
)

type Data struct {
	Data []Item `json:"data"`
}

type Item interface{}

type User struct {
	Firstname string    `json:"firstname"`
	Lastname  string    `json:"lastname"`
	Id        int       `json:"id"`
	Payments  []Payment `json:"payments"`
	Addresses []Address `json:"addresses"`
}

type Payment struct {
	Id     int `json:"id"`
	Amount int `json:"amount"`
}

type Address struct {
	Id      int    `json:"id"`
	Address string `json:"address"`
}

func main() {
	var input Data
	err := json.NewDecoder(os.Stdin).Decode(&input)

	if err != nil {
		log.Fatal(err)
	}

	userMap := make(map[int]*User, 0)

	// save users in userMap with their IDs
	for _, item := range input.Data {
		if user, err := ItemToUser(item); err == nil {
			userMap[user.Id] = &user
		}
	}

	// save payments and addresses to their users
	for _, item := range input.Data {
		if payment, err := ItemToPayment(item); err == nil {
			user := userMap[payment.Id]
			user.Payments = append(user.Payments, payment)
		} else if address, err := ItemToAddress(item); err == nil {
			user := userMap[address.Id]
			user.Addresses = append(user.Addresses, address)
		}
	}

	userSlice := make([]User, 0, len(userMap))

	for _, user := range userMap {
		userSlice = append(userSlice, *user)
	}

	bytes, _ := json.MarshalIndent(userSlice, "", " ")
	fmt.Fprintln(os.Stdout, string(bytes))
}

func ItemToUser(item Item) (user User, err error) {
	s, err := json.Marshal(item)

	if err != nil {
		_ = fmt.Errorf("An Item of the `data` is not correct: %s",
			err)
		return
	}

	err = json.Unmarshal(s, &user)

	if user.Firstname == "" {
		err = errors.New("invalid JSON for `User`")
	}

	return
}

func ItemToPayment(item Item) (payment Payment, err error) {
	s, err := json.Marshal(item)

	if err != nil {
		_ = fmt.Errorf("An Item of the `data` is not correct: %s",
			err)
		return
	}

	err = json.Unmarshal(s, &payment)

	if payment.Amount == 0 {
		err = errors.New("invalid JSON for `Payment`")
	}

	return
}

func ItemToAddress(item Item) (address Address, err error) {
	s, err := json.Marshal(item)

	if err != nil {
		_ = fmt.Errorf("An Item of the `data` is not correct: %s",
			err)
		return
	}

	err = json.Unmarshal(s, &address)

	if address.Address == "" {
		err = errors.New("invalid JSON for `Address`")
	}

	return
}
