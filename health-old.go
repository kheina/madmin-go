//
// Copyright (c) 2015-2022 MinIO, Inc.
//
// This file is part of MinIO Object Storage stack
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.
//

package madmin

import (
	"encoding/json"
	"time"
)

// HealthInfoV0 - MinIO cluster's health Info version 0
type HealthInfoV0 struct {
	TimeStamp time.Time     `json:"timestamp,omitempty"`
	Error     string        `json:"error,omitempty"`
	Sys       SysHealthInfo `json:"sys,omitempty"`
}

// HealthInfoV2 - MinIO cluster's health Info version 2
type HealthInfoV2 struct {
	Version string `json:"version"`
	Error   string `json:"error,omitempty"`

	TimeStamp time.Time       `json:"timestamp,omitempty"`
	Sys       SysInfo         `json:"sys,omitempty"`
	Perf      PerfInfo        `json:"perf,omitempty"`
	Minio     MinioHealthInfo `json:"minio,omitempty"`
}

func (info HealthInfoV2) String() string {
	data, err := json.Marshal(info)
	if err != nil {
		panic(err) // This never happens.
	}
	return string(data)
}

// JSON returns this structure as JSON formatted string.
func (info HealthInfoV2) JSON() string {
	data, err := json.MarshalIndent(info, " ", "    ")
	if err != nil {
		panic(err) // This never happens.
	}
	return string(data)
}

// GetError - returns error from the cluster health info v2
func (info HealthInfoV2) GetError() string {
	return info.Error
}

// GetStatus - returns status of the cluster health info v2
func (info HealthInfoV2) GetStatus() string {
	if info.Error != "" {
		return "error"
	}
	return "success"
}

// GetTimestamp - returns timestamp from the cluster health info v2
func (info HealthInfoV2) GetTimestamp() time.Time {
	return info.TimeStamp
}

// Latency contains write operation latency in seconds of a disk drive.
type Latency struct {
	Avg          float64 `json:"avg"`
	Max          float64 `json:"max"`
	Min          float64 `json:"min"`
	Percentile50 float64 `json:"percentile_50"`
	Percentile90 float64 `json:"percentile_90"`
	Percentile99 float64 `json:"percentile_99"`
}

// Throughput contains write performance in bytes per second of a disk drive.
type Throughput struct {
	Avg          uint64 `json:"avg"`
	Max          uint64 `json:"max"`
	Min          uint64 `json:"min"`
	Percentile50 uint64 `json:"percentile_50"`
	Percentile90 uint64 `json:"percentile_90"`
	Percentile99 uint64 `json:"percentile_99"`
}

// DrivePerfInfo contains disk drive's performance information.
type DrivePerfInfo struct {
	Error string `json:"error,omitempty"`

	Path       string     `json:"path"`
	Latency    Latency    `json:"latency,omitempty"`
	Throughput Throughput `json:"throughput,omitempty"`
}

// DrivePerfInfos contains all disk drive's performance information of a node.
type DrivePerfInfos struct {
	NodeCommon

	SerialPerf   []DrivePerfInfo `json:"serial_perf,omitempty"`
	ParallelPerf []DrivePerfInfo `json:"parallel_perf,omitempty"`
}

// PeerNetPerfInfo contains network performance information of a node.
type PeerNetPerfInfo struct {
	NodeCommon

	Latency    Latency    `json:"latency,omitempty"`
	Throughput Throughput `json:"throughput,omitempty"`
}

// NetPerfInfo contains network performance information of a node to other nodes.
type NetPerfInfo struct {
	NodeCommon

	RemotePeers []PeerNetPerfInfo `json:"remote_peers,omitempty"`
}

// PerfInfo - Includes Drive and Net perf info for the entire MinIO cluster
type PerfInfo struct {
	Drives      []DrivePerfInfos `json:"drives,omitempty"`
	Net         []NetPerfInfo    `json:"net,omitempty"`
	NetParallel NetPerfInfo      `json:"net_parallel,omitempty"`
}

func (info HealthInfoV0) String() string {
	data, err := json.Marshal(info)
	if err != nil {
		panic(err) // This never happens.
	}
	return string(data)
}

// JSON returns this structure as JSON formatted string.
func (info HealthInfoV0) JSON() string {
	data, err := json.MarshalIndent(info, " ", "    ")
	if err != nil {
		panic(err) // This never happens.
	}
	return string(data)
}

// SysHealthInfo - Includes hardware and system information of the MinIO cluster
type SysHealthInfo struct {
	CPUInfo    []ServerCPUInfo    `json:"cpus,omitempty"`
	DiskHwInfo []ServerDiskHwInfo `json:"drives,omitempty"`
	OsInfo     []ServerOsInfo     `json:"osinfos,omitempty"`
	MemInfo    []ServerMemInfo    `json:"meminfos,omitempty"`
	ProcInfo   []ServerProcInfo   `json:"procinfos,omitempty"`
	Error      string             `json:"error,omitempty"`
}

// ServerProcInfo - Includes host process lvl information
type ServerProcInfo struct {
	Addr      string       `json:"addr"`
	Processes []SysProcess `json:"processes,omitempty"`
	Error     string       `json:"error,omitempty"`
}

// SysProcess - Includes process lvl information about a single process
type SysProcess struct {
	Pid             int32   `json:"pid"`
	Background      bool    `json:"background,omitempty"`
	CPUPercent      float64 `json:"cpupercent,omitempty"`
	Children        []int32 `json:"children,omitempty"`
	CmdLine         string  `json:"cmd,omitempty"`
	ConnectionCount int     `json:"connection_count,omitempty"`
	CreateTime      int64   `json:"createtime,omitempty"`
	Cwd             string  `json:"cwd,omitempty"`
	Exe             string  `json:"exe,omitempty"`
	Gids            []int32 `json:"gids,omitempty"`
	IsRunning       bool    `json:"isrunning,omitempty"`
	MemPercent      float32 `json:"mempercent,omitempty"`
	Name            string  `json:"name,omitempty"`
	Nice            int32   `json:"nice,omitempty"`
	NumFds          int32   `json:"numfds,omitempty"`
	NumThreads      int32   `json:"numthreads,omitempty"`
	Parent          int32   `json:"parent,omitempty"`
	Ppid            int32   `json:"ppid,omitempty"`
	Status          string  `json:"status,omitempty"`
	Tgid            int32   `json:"tgid,omitempty"`
	Uids            []int32 `json:"uids,omitempty"`
	Username        string  `json:"username,omitempty"`
}

// GetOwner - returns owner of the process
func (sp SysProcess) GetOwner() string {
	return sp.Username
}

// ServerMemInfo - Includes host virtual and swap mem information
type ServerMemInfo struct {
	Addr  string `json:"addr"`
	Error string `json:"error,omitempty"`
}

// ServerOsInfo - Includes host os information
type ServerOsInfo struct {
	Addr  string `json:"addr"`
	Error string `json:"error,omitempty"`
}

// ServerCPUInfo - Includes cpu and timer stats of each node of the MinIO cluster
type ServerCPUInfo struct {
	Addr  string `json:"addr"`
	Error string `json:"error,omitempty"`
}

// MinioHealthInfoV0 - Includes MinIO confifuration information
type MinioHealthInfoV0 struct {
	Info   InfoMessage `json:"info,omitempty"`
	Config interface{} `json:"config,omitempty"`
	Error  string      `json:"error,omitempty"`
}

// ServerDiskHwInfo - Includes usage counters, disk counters and partitions
type ServerDiskHwInfo struct {
	Addr  string `json:"addr"`
	Error string `json:"error,omitempty"`
}
