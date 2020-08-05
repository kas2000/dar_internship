package main

import (
	"fmt"
	"math/rand"
	"time"
	"strconv"
)

func main() {
	//аллоцировали память для структуры ConnectionPool
	connectionPool := ConnectionPool{}

	//проинициализировали его пятью соединениями
	connectionPool.init(5)

	userId1 := "qwe123"
	userId2 := "qwe124"
	userId3 := "qwe125"

	//3 наших юзера с разными userId забрали по соединению
	connectionPool.getConnection(userId1)
	connectionPool.getConnection(userId2)
	connectionPool.getConnection(userId3)

	//print должен показывать 2 свободных и 3 занятых соединения
	fmt.Println(connectionPool)

	connectionPool.returnConnection(userId1)

	//print должен показывать 3 свободных и 2 занятых соединения
	fmt.Println(connectionPool)
}

type ConnectionPool struct {
	//connection -> user_id
	connections map[Connection]string
	//Initial number of connections
	initSize int
	//Number of connections currently in use
	checkedOut int
	// Connectionx1 -> "qwe123"
	// Connectionx2 -> "qwe124"
	// Connectionx3 -> "qwe125"
}

type Connection struct {
	Id string
	//"50" - 50 секунд
	Timeout string
}

func (c *ConnectionPool) getConnection(userId string) Connection{
	for key, value := range c.connections {
		if (value == "N/A") {
			c.connections[key] = userId
			c.checkedOut++
			return key
		}
	}

	return Connection{"size exceeded", "50"}
}

func (c *ConnectionPool) returnConnection(userId string) {
	for key, value := range c.connections {
		if (value == userId) {
			c.connections[key] = "N/A"
			c.checkedOut--
		}
	}
}

func (c *ConnectionPool) init(size int) {
	c.initSize = size
	c.connections = make(map[Connection]string)
	for i := 0; i < size; i++ {
		rand.Seed(time.Now().UnixNano())
		conn := Connection{strconv.Itoa(rand.Intn(10000)), "50"}
		c.connections[conn] = "N/A"
	}
}

func (c ConnectionPool) String() string {
	return fmt.Sprintf("%v свободных, %v занятых", c.initSize - c.checkedOut, c.checkedOut)
}
