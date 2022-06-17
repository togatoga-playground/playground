package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"sync"
	"testing"
	"time"
)

func initSimple() {
	daemonStarted := startNetworkDaemon()
	daemonStarted.Wait()
}

func initWithWarm() {
	daemonStarted := startNetworkDaemonWithWarm()
	daemonStarted.Wait()
}

func connectToService() interface{} {
	time.Sleep(time.Second)
	return struct{}{}
}

func startNetworkDaemon() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		server, err := net.Listen("tcp", "localhost:8080")
		if err != nil {
			log.Fatalf("connot listen: %v", err)
		}

		defer server.Close()
		wg.Done()
		for {
			conn, err := server.Accept()

			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			connectToService()
			fmt.Fprintf(conn, "")
			conn.Close()
		}
	}()
	return &wg
}

func warmServiceConnCache() *sync.Pool {
	p := &sync.Pool{
		New: connectToService,
	}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func startNetworkDaemonWithWarm() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		connPool := warmServiceConnCache()
		server, err := net.Listen("tcp", "localhost:8080")

		if err != nil {
			log.Fatalf("connot listen: %v", err)
		}

		defer server.Close()
		wg.Done()

		for {
			conn, err := server.Accept()
			if err != nil {
				log.Printf("cannot accept connection: %v", err)
				continue
			}
			svcConn := connPool.Get()
			fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			conn.Close()
		}

	}()
	return &wg
}

var onceBenchmarkNetworkRequest sync.Once

func BenchmarkNetworkRequest(b *testing.B) {

	onceBenchmarkNetworkRequest.Do(func() { initSimple() })

	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannnot dial host: %v", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}

var onceBenchmarkNetworkRequestWithWarm sync.Once

func BenchmarkNetworkRequestWithWarm(b *testing.B) {
	onceBenchmarkNetworkRequestWithWarm.Do(func() { initWithWarm() })
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", "localhost:8080")
		if err != nil {
			b.Fatalf("cannnot dial host: %v", err)
		}

		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("cannot read: %v", err)
		}
		conn.Close()
	}
}
