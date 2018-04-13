/*
 * Copyright © 2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * @author		Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @copyright 	2015-2018 Aeneas Rekkas <aeneas+oss@aeneas.io>
 * @license 	Apache-2.0
 */

package metrics

import (
	"runtime"
	"sync"
)

type MemoryStatistics struct {
	sync.Mutex
	Alloc        uint64 `json:"alloc"`
	TotalAlloc   uint64 `json:"totalAlloc"`
	Sys          uint64 `json:"sys"`
	Lookups      uint64 `json:"lookups"`
	Mallocs      uint64 `json:"mallocs"`
	Frees        uint64 `json:"frees"`
	HeapAlloc    uint64 `json:"heapAlloc"`
	HeapSys      uint64 `json:"heapSys"`
	HeapIdle     uint64 `json:"heapIdle"`
	HeapInuse    uint64 `json:"heapInuse"`
	HeapReleased uint64 `json:"heapReleased"`
	HeapObjects  uint64 `json:"heapObjects"`
	NumGC        uint32 `json:"numGC"`
}

func (ms *MemoryStatistics) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"alloc":          ms.Alloc,
		"totalAlloc":     ms.TotalAlloc,
		"sys":            ms.Sys,
		"lookups":        ms.Lookups,
		"mallocs":        ms.Mallocs,
		"frees":          ms.Frees,
		"heapAlloc":      ms.HeapAlloc,
		"heapSys":        ms.HeapSys,
		"heapIdle":       ms.HeapIdle,
		"heapInuse":      ms.HeapInuse,
		"heapReleased":   ms.HeapReleased,
		"heapObjects":    ms.HeapObjects,
		"numGC":          ms.NumGC,
		"nonInteraction": 1,
	}
}

func (ms *MemoryStatistics) Update() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)

	ms.Lock()
	defer ms.Unlock()
	ms.Alloc = m.Alloc
	ms.TotalAlloc = m.TotalAlloc
	ms.Sys = m.Sys
	ms.Lookups = m.Lookups
	ms.Mallocs = m.Mallocs
	ms.Frees = m.Frees
	ms.HeapAlloc = m.HeapAlloc
	ms.HeapSys = m.HeapSys
	ms.HeapIdle = m.HeapIdle
	ms.HeapInuse = m.HeapInuse
	ms.HeapReleased = m.HeapReleased
	ms.HeapObjects = m.HeapObjects
	ms.NumGC = m.NumGC
}
