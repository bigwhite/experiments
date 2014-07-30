/*
 * Copyright 2012 iWoBi Technologies. All rights reserved.
 *
 * memcached client api
 *
 *
 * memcached protocol:
 *  store cmd:  <command name> <key> <flags> <exptime> <bytes> [noreply]\r\n<data block>\r\n
 *  store response:
 *          - "STORED\r\n", to indicate success.
 *          - "NOT_STORED\r\n" to indicate the data was not stored, but not
 *              because of an error. This normally means that the
 *              condition for an "add" or a "replace" command wasn't met.
 *          - "EXISTS\r\n" to indicate that the item you are trying to store with
 *              a "cas" command has been modified since you last fetched it.
 *          - "NOT_FOUND\r\n" to indicate that the item you are trying to store
 *              with a "cas" command did not exist.
 *
 *  incr: incr <key> <value> [noreply]\r\n
 *  incr response:
 *          - "NOT_FOUND\r\n" to indicate the item with this value was not found
 *          - <value>\r\n , where <value> is the new value of the item's data,
 *             after the increment/decrement operation was carried out.
 *
 */

package libmemcached

import (
	"errors"
	"fmt"
	"net"
	"strconv"
	"strings"
)

type MemcachedConn struct {
	conn net.Conn
}

var ErrNotFound = errors.New("key in memcached can not be found")
var ErrExists = errors.New("the key exists")
var ErrNotStored = errors.New("the key has not been stored")
var ErrUnknown = errors.New("Unknown error code")

func Connect(host string, port string) (c *MemcachedConn, err error) {
	service := host + ":" + port
	tcpAddr, err := net.ResolveTCPAddr("tcp", service)
	if err != nil {
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		return nil, err
	}

	return &MemcachedConn{conn: conn}, nil
}

func Close(c *MemcachedConn) error {
	return c.conn.Close()
}

func sendRequest(c *MemcachedConn, req string) error {
	_, err := c.conn.Write([]byte(req))
	return err
}

func recvResponse(c *MemcachedConn) ([]byte, int, error) {
	var resp [512]byte
	n, err := c.conn.Read(resp[:])
	if err != nil {
		return nil, 0, err
	}

	return resp[:], n, nil
}

func Incr(c *MemcachedConn, key string, offset string, value *string) error {
	/*
	 * incr: incr <key> <value> [noreply]\r\n
	 */
	request := "incr " + key + " " + offset + "\r\n"
	//fmt.Printf("Send Request - [%s]\n", request)
	err := sendRequest(c, request)
	if err != nil {
		return err
	}

	/*
	 * "NOT_FOUND\r\n" to indicate the item with this value was not found
	 * <value>\r\n , where <value> is the new value of the item's data,
	 */
	resp, n, err := recvResponse(c)
	if err != nil {
		return err
	}

	i := strings.Index(string(resp[:n]), "\r\n")
	//fmt.Printf("Response from Memcachecd Server is = [%s]\n", string(resp[:i]))
	if string(resp[:i]) == "NOT_FOUND" {
		return ErrNotFound
	}

	*value = string(resp[:i])
	return nil
}

func Add(c *MemcachedConn, key string, value string, exptime string) error {
	/* add <key> <flags> <exptime> <bytes> [noreply]\r\n<data block>\r\n */

	bytes := len(value)
	request := "add " + key + " 0 " + exptime + " " + strconv.Itoa(bytes) + "\r\n" + value + "\r\n"
	//fmt.Printf("Send Request - [%s]\n", request)
	err := sendRequest(c, request)
	if err != nil {
		return err
	}

	resp, n, err := recvResponse(c)
	if err != nil {
		return err
	}

	/*
	 * "STORED\r\n"
	 * "NOT_STORED\r\n"
	 * "EXISTS\r\n"
	 * "NOT_FOUND\r\n"
	 */
	i := strings.Index(string(resp[:n]), "\r\n")
	//fmt.Printf("Response from Memcachecd Server is = [%s]\n", string(resp[:i]))
	switch string(resp[:i]) {
	case "NOT_FOUND":
		return ErrNotFound
	case "STORED":
		return nil
	case "EXISTS":
		return ErrExists
	case "NOT_STORED":
		return ErrNotStored
	default:
		fmt.Printf("Errcode from Memcachecd Server is = [%s]\n", string(resp[:i]))
		return ErrUnknown
	}

	return nil
}
