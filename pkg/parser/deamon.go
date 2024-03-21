package parser

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
	"time"
)

func hexToInt(s string) int {
	parsed, _ := strconv.ParseInt(strings.Replace(s, "0x", "", -1), 16, 32)
	return int(parsed)
}

func intToHex(num int) string {
	return fmt.Sprintf("0x%x", num)
}

type daemon struct {
	lock        sync.RWMutex
	latestBlock int
	client      *client
	subscribers map[string]bool
}

func newDaemon(url string) *daemon {
	return &daemon{sync.RWMutex{}, -1, newClient(url), make(map[string]bool)}
}

func (daemon *daemon) subscribe(address string) bool {
	daemon.lock.Lock()
	defer daemon.lock.Unlock()
	if daemon.subscribers[address] {
		return false
	}
	daemon.subscribers[address] = true
	return true
}

func (daemon *daemon) lastParsedBlock() int {
	daemon.lock.RLock()
	defer daemon.lock.RUnlock()
	return daemon.latestBlock
}

func (daemon *daemon) run() {
	for {
		daemon.tick()
		time.Sleep(1 * time.Second)
	}
}

func (daemon *daemon) tick() {
	blockNumberResp, err := daemon.client.getRecentBlockNumber()
	if err != nil {
		panic(err)
	}
	blockNumber := hexToInt(blockNumberResp.Result)

	daemon.lock.Lock()
	defer daemon.lock.Unlock()

	if daemon.latestBlock == -1 {
		daemon.parseBlockByBlockNum(blockNumber)
	} else {
		for blockNum := daemon.latestBlock; blockNum <= blockNumber; blockNum++ {
			daemon.parseBlockByBlockNum(blockNum)
		}
	}

	daemon.latestBlock = blockNumber
}

func (daemon *daemon) parseBlockByBlockNum(block int) {
	blockByNumberResp, err := daemon.client.getBlockByNumber(intToHex(block))
	if err != nil {
		panic(err)
	}

	for _, t := range blockByNumberResp.Result.Transactions {
		if daemon.subscribers[t.To] {
			insert(t.To, t)
		}

		if daemon.subscribers[t.From] {
			insert(t.From, t)
		}
	}
}
