/*
 * Copyright (C) 2016 Red Hat, Inc.
 *
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *  http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 *
 */

package analyzer

import (
	"github.com/skydive-project/skydive/config"
	"github.com/skydive-project/skydive/logging"
	"github.com/skydive-project/skydive/probe"
	"github.com/skydive-project/skydive/topology/graph"
	tprobes "github.com/skydive-project/skydive/topology/probes"
)

func NewTopologyProbeBundleFromConfig(g *graph.Graph) (*probe.ProbeBundle, error) {
	probes := make(map[string]probe.Probe)
	probes["fabric"] = tprobes.NewFabricProbe(g)
	probeBundle := probe.NewProbeBundle(probes)

	for _, probe := range config.GetConfig().GetStringSlice("analyzer.topology.probes") {
		switch probe {
		case "neutron":
			neutron, err := tprobes.NewNeutronMapperFromConfig(g, probeBundle)
			if err != nil {
				logging.GetLogger().Errorf("Failed to initialize Neutron probe: %s", err.Error())
				return nil, err
			}
			probes["neutron"] = neutron
		default:
			logging.GetLogger().Errorf("unknown probe type %s", probe)
		}
	}
	return probeBundle, nil
}
