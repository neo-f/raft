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

package raftpb

import (
	"testing"
)

func TestLeaveJoint(t *testing.T) {
	tests := []struct {
		name     string
		cc       *ConfChangeV2
		expected bool
	}{
		{
			name:     "empty conf change (all zero values)",
			cc:       ConfChangeV2_builder{}.Build(),
			expected: true,
		},
		{
			name: "empty conf change with auto transition explicitly set",
			cc: ConfChangeV2_builder{
				Transition: ConfChangeTransition_ConfChangeTransitionAuto.Enum(),
			}.Build(),
			expected: true,
		},
		{
			name: "empty changes with context (context should be ignored)",
			cc: ConfChangeV2_builder{
				Transition: ConfChangeTransition_ConfChangeTransitionAuto.Enum(),
				Context:    []byte("some context"),
			}.Build(),
			expected: true,
		},

		{
			name: "empty changes slice with auto transition",
			cc: ConfChangeV2_builder{
				Transition: ConfChangeTransition_ConfChangeTransitionAuto.Enum(),
				Changes:    []*ConfChangeSingle{},
			}.Build(),
			expected: true,
		},
		{
			name: "non-empty changes",
			cc: ConfChangeV2_builder{
				Transition: ConfChangeTransition_ConfChangeTransitionAuto.Enum(),
				Changes: []*ConfChangeSingle{
					ConfChangeSingle_builder{
						Type:   ConfChangeAddNode.Enum(),
						NodeId: new(uint64),
					}.Build(),
				},
			}.Build(),
			expected: false,
		},
		{
			name: "joint implicit transition with empty changes",
			cc: ConfChangeV2_builder{
				Transition: ConfChangeTransition_ConfChangeTransitionJointImplicit.Enum(),
			}.Build(),
			expected: false,
		},
		{
			name: "joint explicit transition with empty changes",
			cc: ConfChangeV2_builder{
				Transition: ConfChangeTransition_ConfChangeTransitionJointExplicit.Enum(),
			}.Build(),
			expected: false,
		},
		{
			name: "auto transition with multiple changes",
			cc: ConfChangeV2_builder{
				Transition: ConfChangeTransition_ConfChangeTransitionAuto.Enum(),
				Changes: []*ConfChangeSingle{
					ConfChangeSingle_builder{Type: ConfChangeAddNode.Enum(), NodeId: new(uint64)}.Build(),
					ConfChangeSingle_builder{Type: ConfChangeRemoveNode.Enum(), NodeId: new(uint64)}.Build(),
				},
			}.Build(),
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.cc.LeaveJoint()
			if got != tt.expected {
				t.Errorf("LeaveJoint() = %v, want %v", got, tt.expected)
			}
		})
	}
}
