/*
This license only applies to code part of the following packages:
https://github.com/joncalhoun/drip/tree/c0eb27d26abd52c7617ccf9d3cce73ed17a39f56

Modifications will remain under project license!


The MIT License (MIT)

Copyright (c) 2015 Jon Calhoun

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package flow

import (
	"time"
)

func (b *Bucket) Start() bool {
	if b.started {
		return false
	}

	go b.controller()

	return true
}

func (b *Bucket) controller() {
	ticker := time.NewTicker(b.DripInterval)
	b.started = true
	b.kill = make(chan bool, 1)

	for {
		select {
		case <-ticker.C:
			b.m.Lock()
			b.consumed -= b.PerDrip
			if b.consumed < 0 {
				b.consumed = 0
			}
			b.m.Unlock()
		case <-b.kill:
			return
		}
	}
}

func (b *Bucket) Stop() bool {
	if !b.started {
		return false
	}

	b.kill <- true

	return true
}

func (b *Bucket) Consume(amt int) bool {
	b.m.Lock()
	defer b.m.Unlock()
	if b.Capacity-b.consumed < amt {
		return false
	}

	b.consumed += amt
	return true
}

func (b *Bucket) ConsumeWithTimeout(amt int) bool {
	for {
		if b.Consume(amt) {
			break
		}
		time.Sleep(2 * time.Second)
	}
	return true
}
