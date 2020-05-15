package stringutil

import (
	"errors"
	"fmt"
	"sync"
	"time"
)

// create of solution by snowflake
type IdCreator struct {
	startTime             int64
	workerIdBits          uint
	datacenterIdBits      uint
	maxWorkerId           int64
	maxDatacenterId       int64
	sequenceBits          uint
	workerIdLeftShift     uint
	datacenterIdLeftShift uint
	timestampLeftShift    uint
	sequenceMask          int64
	workerId              int64
	datacenterId          int64
	sequence              int64
	lastTimestamp         int64
	signMask              int64
	mLock                 *sync.Mutex
}

// singleton instance
var mIdCreator = &IdCreator{
	startTime:             1577808000000,
	workerIdBits:          5,
	datacenterIdBits:      5,
	maxWorkerId:           (-1) ^ ((-1) << 5),
	maxDatacenterId:       (-1) ^ ((-1) << 5),
	sequenceBits:          12,
	workerIdLeftShift:     12,
	datacenterIdLeftShift: 17,
	timestampLeftShift:    22,
	sequenceMask:          (-1) ^ ((-1) << 12),
	sequence:              0,
	lastTimestamp:         -1,
	signMask:              ^(-1) + 1,
	mLock:                 &sync.Mutex{},
	workerId:              0,
	datacenterId:          0,
}

// set creator config
// workerId
// datacenterId
func SetCreator(workerId, datacenterId int64) {
	mIdCreator.mLock.Lock()
	defer mIdCreator.mLock.Unlock()

	mIdCreator.workerId = workerId
	mIdCreator.datacenterId = datacenterId
}

// create new Id with error
func TryCreateNewId() (int64, error) {
	return mIdCreator.newId()
}

// create new Id
func CreateNewId() int64 {

	r, err := mIdCreator.newId()

	if err != nil {
		fmt.Printf("stringutil.CreateNewId error:%s", err)
	}

	return r
}

// create Id
func (c *IdCreator) newId() (int64, error) {

	c.mLock.Lock()

	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	if timestamp < c.lastTimestamp {

		c.mLock.Unlock()

		return -1, errors.New(fmt.Sprintf("create id error, for %d milliseconds", timestamp))
	}

	if timestamp == c.lastTimestamp {
		c.sequence = (c.sequence + 1) & c.sequenceMask

		// if current time Millisecond , sequence is zero
		// set timestamp = current Millisecond
		if c.sequence == 0 {
			timestamp = c.tilNextMillis()
			c.sequence = 0
		}
	} else {
		c.sequence = 0
	}

	c.lastTimestamp = timestamp

	c.mLock.Unlock()

	// 0 - 0000000000 0000000000 0000000000 0000000000 0 - 00000 - 00000 - 000000000000
	// first is zero
	// second part is timestamp
	// third part is data center group Id
	// forth part is worker id or module id
	// fifth part is count ,total count in the same Millisecond
	r := ((timestamp - c.startTime) << c.timestampLeftShift) |
		(c.datacenterId << c.datacenterIdLeftShift) |
		(c.workerId << c.workerIdLeftShift) |
		c.sequence

	if r < 0 {
		r = -r
	}

	return r, nil
}

// tilNextMillis
func (c *IdCreator) tilNextMillis() int64 {
	t := time.Now().UnixNano()

	if t <= c.lastTimestamp {
		t = time.Now().UnixNano() / int64(time.Millisecond)
	}

	return t
}
