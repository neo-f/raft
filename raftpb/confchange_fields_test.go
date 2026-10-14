//go:build !protoopaque

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

// The Go-struct layout of generated messages depends on the protobuf API
// level (open/hybrid vs opaque), so the reflect-based field count guard only
// works with the default (non-opaque) build.

package raftpb

import (
	"reflect"
	"testing"
)

// TestConfChangeV2FieldCount verifies that ConfChangeV2 has exactly the
// expected fields. This ensures that if new fields are added to ConfChangeV2,
// this test will fail.
func TestConfChangeV2FieldCount(t *testing.T) {
	ccv2Type := reflect.TypeOf((*ConfChangeV2)(nil)).Elem()
	// 3 user-visible fields (Transition, Changes, Context) + 3 internal protobuf
	// fields (state, unknownFields, sizeCache) added by protoc-gen-go v2.
	expectedFieldCount := 6
	if ccv2Type.NumField() != expectedFieldCount {
		t.Fatalf("ConfChangeV2 has %d fields, expected %d. If new fields were added, please update this test.",
			ccv2Type.NumField(), expectedFieldCount)
	}
}
