// Copyright © 2022 Obol Labs Inc.
//
// This program is free software: you can redistribute it and/or modify it
// under the terms of the GNU General Public License as published by the Free
// Software Foundation, either version 3 of the License, or (at your option)
// any later version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT
// ANY WARRANTY; without even the implied warranty of  MERCHANTABILITY or
// FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for
// more details.
//
// You should have received a copy of the GNU General Public License along with
// this program.  If not, see <http://www.gnu.org/licenses/>.

// Code generated by "stringer -type=uponRule -trimprefix=upon"; DO NOT EDIT.

package qbft

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[uponNothing-0]
	_ = x[uponJustifiedPrePrepare-1]
	_ = x[uponQuorumPrepares-2]
	_ = x[uponQuorumCommits-3]
	_ = x[uponUnjustQuorumRoundChanges-4]
	_ = x[uponFPlus1RoundChanges-5]
	_ = x[uponQuorumRoundChanges-6]
	_ = x[uponJustifiedDecided-7]
}

const _uponRule_name = "NothingJustifiedPrePrepareQuorumPreparesQuorumCommitsUnjustQuorumRoundChangesFPlus1RoundChangesQuorumRoundChangesJustifiedDecided"

var _uponRule_index = [...]uint8{0, 7, 26, 40, 53, 77, 95, 113, 129}

func (i uponRule) String() string {
	if i < 0 || i >= uponRule(len(_uponRule_index)-1) {
		return "uponRule(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _uponRule_name[_uponRule_index[i]:_uponRule_index[i+1]]
}
