// Copyright 2021 Chaos Mesh Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package server

import (
	"context"

	v1 "k8s.io/api/core/v1"

	"github.com/chaos-mesh/chaos-mesh/pkg/bpm"
)

// GetIpset returns result of ipset list
func (r *Resolver) GetIpset(ctx context.Context, obj *v1.Pod) (string, error) {
	cmd := "ipset list"
	return r.ExecBypass(ctx, obj, cmd, bpm.PidNS, bpm.NetNS)
}

// GetIpset returns result of tc qdisc list
func (r *Resolver) GetTcQdisc(ctx context.Context, obj *v1.Pod) (string, error) {
	cmd := "tc qdisc list"
	return r.ExecBypass(ctx, obj, cmd, bpm.PidNS, bpm.NetNS)
}

// GetIptables returns result of iptables --list
func (r *Resolver) GetIptables(ctx context.Context, obj *v1.Pod) (string, error) {
	cmd := "iptables --list"
	return r.ExecBypass(ctx, obj, cmd, bpm.PidNS, bpm.NetNS)
}
