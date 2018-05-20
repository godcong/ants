// Copyright (c) 2018 Andy Pan
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell copies
// of the Software, and to permit persons to whom the Software is furnished to do so,
// subject to the following conditions:
//
//The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
//THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED,
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A
// PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT
// HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
// OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
// SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.

package ants

import "github.com/iris-contrib/errors"

const (
	DEFAULT_POOL_SIZE           = 50000
	DEFAULT_CLEAN_INTERVAL_TIME = 30
)

var defaultPool, _ = NewPool(DEFAULT_POOL_SIZE)

func Push(task f) error {
	return defaultPool.Push(task)
}

func Running() int {
	return defaultPool.Running()
}

func Cap() int {
	return defaultPool.Cap()
}

func Free() int {
	return defaultPool.Free()
}

func Release() {

}

var (
	PoolSizeInvalidError = errors.New("invalid size for pool")
	PoolClosedError      = errors.New("this pool has been closed")
)
