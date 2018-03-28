/*
Copyright IBM Corp. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package state

import (
	"crypto/rand"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/ledgerone/fabric-ledgerone/gossip/util"
	proto "github.com/ledgerone/fabric-ledgerone/protos/gossip"
	"github.com/stretchr/testify/assert"
)

func init() {
	util.SetupTestLogging()
}

func randomPayloadWithSeqNum(seqNum uint64) (*proto.Payload, error) {
	data := make([]byte, 64)
	_, err := rand.Read(data)
	if err != nil {
		return nil, err
	}
	return &proto.Payload{
		SeqNum: seqNum,
		Data:   data,
	}, nil
}

func TestNewPayloadsBuffer(t *testing.T) {
	payloadsBuffer := NewPayloadsBuffer(10)
	assert.Equal(t, payloadsBuffer.Next(), uint64(10))
}

func TestPayloadsBufferImpl_Push(t *testing.T) {
	buffer := NewPayloadsBuffer(5)

	payload, err := randomPayloadWithSeqNum(4)

	if err != nil {
		t.Fatal("Wasn't able to generate random payload for test")
	}

	t.Log("Pushing new payload into buffer")
	buffer.Push(payload)

	// Payloads with sequence number less than buffer top
	// index should not be accepted
	t.Log("Getting next block sequence number")
	assert.Equal(t, buffer.Next(), uint64(5))
	t.Log("Check block buffer size")
	assert.Equal(t, buffer.Size(), 0)

	// Adding new payload with seq. number equal to top
	// payload should not be added
	payload, err = randomPayloadWithSeqNum(5)
	if err != nil {
		t.Fatal("Wasn't able to generate random payload for test")
	}

	t.Log("Pushing new payload into buffer")
	buffer.Push(payload)
	t.Log("Getting next block sequence number")
	assert.Equal(t, buffer.Next(), uint64(5))
	t.Log("Check block buffer size")
	assert.Equal(t, buffer.Size(), 1)
}

func TestPayloadsBufferImpl_Ready(t *testing.T) {
	fin := make(chan struct{})
	buffer := NewPayloadsBuffer(1)
	assert.Equal(t, buffer.Next(), uint64(1))

	go func() {
		<-buffer.Ready()
		fin <- struct{}{}
	}()

	time.AfterFunc(100*time.Millisecond, func() {
		payload, err := randomPayloadWithSeqNum(1)

		if err != nil {
			t.Fatal("Wasn't able to generate random payload for test")
		}
		buffer.Push(payload)
	})

	select {
	case <-fin:
		payload := buffer.Pop()
		assert.Equal(t, payload.SeqNum, uint64(1))
	case <-time.After(500 * time.Millisecond):
		t.Fail()
	}
}

// Test to push several concurrent blocks into the buffer
// with same sequence number, only one expected to succeed
func TestPayloadsBufferImpl_ConcurrentPush(t *testing.T) {

	// Test setup, next block num to expect and
	// how many concurrent pushes to simulate
	nextSeqNum := uint64(7)
	concurrency := 10

	buffer := NewPayloadsBuffer(nextSeqNum)
	assert.Equal(t, buffer.Next(), uint64(nextSeqNum))

	startWG := sync.WaitGroup{}
	startWG.Add(1)

	finishWG := sync.WaitGroup{}
	finishWG.Add(concurrency)

	payload, err := randomPayloadWithSeqNum(nextSeqNum)
	assert.NoError(t, err)

	ready := int32(0)
	readyWG := sync.WaitGroup{}
	readyWG.Add(1)
	go func() {
		// Wait for next expected block to arrive
		<-buffer.Ready()
		atomic.AddInt32(&ready, 1)
		readyWG.Done()
	}()

	for i := 0; i < concurrency; i++ {
		go func() {
			buffer.Push(payload)
			startWG.Wait()
			finishWG.Done()
		}()
	}
	startWG.Done()
	finishWG.Wait()

	readyWG.Wait()
	assert.Equal(t, int32(1), atomic.LoadInt32(&ready))
	// Buffer size has to be only one
	assert.Equal(t, 1, buffer.Size())
}
