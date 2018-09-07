// Use and distribution licensed under the Apache license version 2.
//
// See the COPYING file in the root project directory for full text.
//

package ghw

func topologyFillInfo(info *TopologyInfo) error {
	info.Nodes = nil
	if len(info.Nodes) == 1 {
		info.Architecture = SMP
	} else {
		info.Architecture = NUMA
	}
	return nil
}
