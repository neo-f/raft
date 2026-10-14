//go:build protoopaque

// Copyright 2026 The etcd Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Opaque-API counterpart of raft_test.go's TestProtoMemorySizes: the opaque
// layout replaces per-field pointers with values plus a presence bitmap, so
// the expected sizes differ from the default (hybrid/open) build. Only
// 64-bit values are asserted; the 32-bit layout has not been measured.

package raftpb

import (
	"math/bits"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestProtoMemorySizes(t *testing.T) {
	if bits.UintSize != 64 {
		t.Skip("sizes only asserted on 64-bit")
	}

	var sm SnapshotMetadata
	assert.Equal(t, uintptr(72), unsafe.Sizeof(sm), "SnapshotMetadata size check")

	var s Snapshot
	assert.Equal(t, uintptr(80), unsafe.Sizeof(s), "Snapshot size check")

	var m Message
	assert.Equal(t, uintptr(176), unsafe.Sizeof(m), "Message size check")

	var hs HardState
	assert.Equal(t, uintptr(72), unsafe.Sizeof(hs), "HardState size check")

	var cs ConfState
	assert.Equal(t, uintptr(144), unsafe.Sizeof(cs), "ConfState size check")

	var cc ConfChange
	assert.Equal(t, uintptr(96), unsafe.Sizeof(cc), "ConfChange size check")

	var ccs ConfChangeSingle
	assert.Equal(t, uintptr(64), unsafe.Sizeof(ccs), "ConfChangeSingle size check")

	var ccv2 ConfChangeV2
	assert.Equal(t, uintptr(88), unsafe.Sizeof(ccv2), "ConfChangeV2 size check")
}
