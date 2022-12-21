package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Account struct {
	Name    string
	Balance int
}

var accounts = make(map[string]*Account)

func (a Account) getBalance() {
	fmt.Println(a.Balance)
}

func (a *Account) transferToAccount(receiver *Account, amount int) {
	a.Balance -= amount
	receiver.Balance += amount
}

func (a *Account) addIncome(n int) {
	if a.Balance > 0 {
		a.Balance += int(math.Floor(float64(a.Balance) * 0.01 * float64(n)))
	}
}

func addIncomeToAll(n int) {
	for _, account := range accounts {
		account.addIncome(n)
	}
}

func addNewAccount(name string, deposit int) *Account {
	accountToAdd := Account{
		Name:    name,
		Balance: deposit}
	accounts[name] = &accountToAdd
	return &accountToAdd
}

func deposit(name string, amount int) {
	account, ok := accounts[name]
	if ok == false {
		account = addNewAccount(name, 0)
	}
	account.Balance += amount
}

func (a *Account) withdraw(amount int) {
	a.Balance -= amount
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input := strings.Fields(scanner.Text())
		switch operation := input[0]; operation {
		case "DEPOSIT":
			name := input[1]
			amount, _ := strconv.Atoi(input[2])
			deposit(name, amount)
		case "INCOME":
			n, _ := strconv.Atoi(input[1])
			addIncomeToAll(n)
		case "BALANCE":
			name := input[1]
			account, ok := accounts[name]
			if ok {
				account.getBalance()
			} else {
				fmt.Println("ERROR")
			}
		case "TRANSFER":
			senderName := input[1]
			sender, ok := accounts[senderName]
			if ok == false {
				sender = addNewAccount(senderName, 0)
			}
			receiverName := input[2]
			receiver, ok := accounts[receiverName]
			if ok == false {
				receiver = addNewAccount(receiverName, 0)
			}
			amount, _ := strconv.Atoi(input[3])
			sender.transferToAccount(receiver, amount)
		case "WITHDRAW":
			name := input[1]
			amount, _ := strconv.Atoi(input[2])
			account, ok := accounts[name]
			if ok == false {
				account = addNewAccount(name, 0)
			}
			account.withdraw(amount)
		}
	}
}
