// Copyright © 2022-2023 Obol Labs Inc. Licensed under the terms of a Business Source License 1.1

// Code generated by "stringer -type=OrderStart -trimprefix=Start"; DO NOT EDIT.

package lifecycle

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[StartTracker-0]
	_ = x[StartAggSigDB-1]
	_ = x[StartRelay-2]
	_ = x[StartMonitoringAPI-3]
	_ = x[StartValidatorAPI-4]
	_ = x[StartP2PPing-5]
	_ = x[StartP2PRouters-6]
	_ = x[StartP2PConsensus-7]
	_ = x[StartSimulator-8]
	_ = x[StartScheduler-9]
	_ = x[StartP2PEventCollector-10]
	_ = x[StartPeerInfo-11]
	_ = x[StartParSigDB-12]
}

const _OrderStart_name = "TrackerAggSigDBRelayMonitoringAPIValidatorAPIP2PPingP2PRoutersP2PConsensusSimulatorSchedulerP2PEventCollectorPeerInfoParSigDB"

var _OrderStart_index = [...]uint8{0, 7, 15, 20, 33, 45, 52, 62, 74, 83, 92, 109, 117, 125}

func (i OrderStart) String() string {
	if i < 0 || i >= OrderStart(len(_OrderStart_index)-1) {
		return "OrderStart(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _OrderStart_name[_OrderStart_index[i]:_OrderStart_index[i+1]]
}
