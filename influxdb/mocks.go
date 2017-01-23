// +build small

/*
http://www.apache.org/licenses/LICENSE-2.0.txt
Copyright 2016 Intel Corporation
Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at
    http://www.apache.org/licenses/LICENSE-2.0
Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package influxdb

var mockStatResults = `{
    "results": [
        {
            "series": [
                {
                    "name": "runtime",
                    "columns": [
                        "Alloc",
                        "Frees",
                        "HeapAlloc",
                        "HeapIdle",
                        "HeapInUse",
                        "HeapObjects",
                        "HeapReleased",
                        "HeapSys",
                        "Lookups",
                        "Mallocs",
                        "NumGC",
                        "NumGoroutine",
                        "PauseTotalNs",
                        "Sys",
                        "TotalAlloc"
                    ],
                    "values": [
                        [
                            18243880,
                            175514,
                            18243880,
                            5308416,
                            19169280,
                            60646,
                            0,
                            24477696,
                            216,
                            236160,
                            8,
                            34,
                            3886029,
                            28465400,
                            43560664
                        ]
                    ]
                },
                {
                    "name": "queryExecutor",
                    "columns": [
                        "queriesActive",
                        "queriesExecuted",
                        "queriesFinished",
                        "queryDurationNs"
                    ],
                    "values": [
                        [
                            1,
                            18,
                            17,
                            10024900
                        ]
                    ]
                },
                {
                    "name": "shard",
                    "tags": {
                        "database": "snap",
                        "engine": "tsm1",
                        "id": "1",
                        "path": "/var/lib/influxdb/data/snap/autogen/1",
                        "retentionPolicy": "autogen",
                        "walPath": "/var/lib/influxdb/wal/snap/autogen/1"
                    },
                    "columns": [
                        "diskBytes",
                        "fieldsCreate",
                        "seriesCreate",
                        "writeBytes",
                        "writePointsDropped",
                        "writePointsErr",
                        "writePointsOk",
                        "writeReq",
                        "writeReqErr",
                        "writeReqOk"
                    ],
                    "values": [
                        [
                            162900,
                            12,
                            7,
                            0,
                            0,
                            0,
                            3397,
                            405,
                            0,
                            405
                        ]
                    ]
                },
                {
                    "name": "tsm1_engine",
                    "tags": {
                        "database": "snap",
                        "engine": "tsm1",
                        "id": "1",
                        "path": "/var/lib/influxdb/data/snap/autogen/1",
                        "retentionPolicy": "autogen",
                        "walPath": "/var/lib/influxdb/wal/snap/autogen/1"
                    },
                    "columns": [
                        "cacheCompactionDuration",
                        "cacheCompactionErr",
                        "cacheCompactions",
                        "cacheCompactionsActive",
                        "tsmFullCompactionDuration",
                        "tsmFullCompactionErr",
                        "tsmFullCompactions",
                        "tsmFullCompactionsActive",
                        "tsmLevel1CompactionDuration",
                        "tsmLevel1CompactionErr",
                        "tsmLevel1Compactions",
                        "tsmLevel1CompactionsActive",
                        "tsmLevel2CompactionDuration",
                        "tsmLevel2CompactionErr",
                        "tsmLevel2Compactions",
                        "tsmLevel2CompactionsActive",
                        "tsmLevel3CompactionDuration",
                        "tsmLevel3CompactionErr",
                        "tsmLevel3Compactions",
                        "tsmLevel3CompactionsActive",
                        "tsmOptimizeCompactionDuration",
                        "tsmOptimizeCompactionErr",
                        "tsmOptimizeCompactions",
                        "tsmOptimizeCompactionsActive"
                    ],
                    "values": [
                        [
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0
                        ]
                    ]
                },
                {
                    "name": "tsm1_cache",
                    "tags": {
                        "database": "snap",
                        "engine": "tsm1",
                        "id": "1",
                        "path": "/var/lib/influxdb/data/snap/autogen/1",
                        "retentionPolicy": "autogen",
                        "walPath": "/var/lib/influxdb/wal/snap/autogen/1"
                    },
                    "columns": [
                        "WALCompactionTimeMs",
                        "cacheAgeMs",
                        "cachedBytes",
                        "diskBytes",
                        "memBytes",
                        "snapshotCount",
                        "writeDropped",
                        "writeErr",
                        "writeOk"
                    ],
                    "values": [
                        [
                            0,
                            404171,
                            0,
                            0,
                            76927,
                            0,
                            0,
                            0,
                            405
                        ]
                    ]
                },
                {
                    "name": "tsm1_filestore",
                    "tags": {
                        "database": "snap",
                        "engine": "tsm1",
                        "id": "1",
                        "path": "/var/lib/influxdb/data/snap/autogen/1",
                        "retentionPolicy": "autogen",
                        "walPath": "/var/lib/influxdb/wal/snap/autogen/1"
                    },
                    "columns": [
                        "diskBytes",
                        "numFiles"
                    ],
                    "values": [
                        [
                            0,
                            0
                        ]
                    ]
                },
                {
                    "name": "tsm1_wal",
                    "tags": {
                        "database": "snap",
                        "engine": "tsm1",
                        "id": "1",
                        "path": "/var/lib/influxdb/data/snap/autogen/1",
                        "retentionPolicy": "autogen",
                        "walPath": "/var/lib/influxdb/wal/snap/autogen/1"
                    },
                    "columns": [
                        "currentSegmentDiskBytes",
                        "oldSegmentsDiskBytes",
                        "writeErr",
                        "writeOk"
                    ],
                    "values": [
                        [
                            169206,
                            0,
                            0,
                            405
                        ]
                    ]
                },
                {
                    "name": "shard",
                    "tags": {
                        "database": "_internal",
                        "engine": "tsm1",
                        "id": "2",
                        "path": "/var/lib/influxdb/data/_internal/monitor/2",
                        "retentionPolicy": "monitor",
                        "walPath": "/var/lib/influxdb/wal/_internal/monitor/2"
                    },
                    "columns": [
                        "diskBytes",
                        "fieldsCreate",
                        "seriesCreate",
                        "writeBytes",
                        "writePointsDropped",
                        "writePointsErr",
                        "writePointsOk",
                        "writeReq",
                        "writeReqErr",
                        "writeReqOk"
                    ],
                    "values": [
                        [
                            126889,
                            101,
                            18,
                            0,
                            0,
                            0,
                            714,
                            40,
                            0,
                            40
                        ]
                    ]
                },
                {
                    "name": "tsm1_engine",
                    "tags": {
                        "database": "_internal",
                        "engine": "tsm1",
                        "id": "2",
                        "path": "/var/lib/influxdb/data/_internal/monitor/2",
                        "retentionPolicy": "monitor",
                        "walPath": "/var/lib/influxdb/wal/_internal/monitor/2"
                    },
                    "columns": [
                        "cacheCompactionDuration",
                        "cacheCompactionErr",
                        "cacheCompactions",
                        "cacheCompactionsActive",
                        "tsmFullCompactionDuration",
                        "tsmFullCompactionErr",
                        "tsmFullCompactions",
                        "tsmFullCompactionsActive",
                        "tsmLevel1CompactionDuration",
                        "tsmLevel1CompactionErr",
                        "tsmLevel1Compactions",
                        "tsmLevel1CompactionsActive",
                        "tsmLevel2CompactionDuration",
                        "tsmLevel2CompactionErr",
                        "tsmLevel2Compactions",
                        "tsmLevel2CompactionsActive",
                        "tsmLevel3CompactionDuration",
                        "tsmLevel3CompactionErr",
                        "tsmLevel3Compactions",
                        "tsmLevel3CompactionsActive",
                        "tsmOptimizeCompactionDuration",
                        "tsmOptimizeCompactionErr",
                        "tsmOptimizeCompactions",
                        "tsmOptimizeCompactionsActive"
                    ],
                    "values": [
                        [
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0,
                            0
                        ]
                    ]
                },
                {
                    "name": "tsm1_cache",
                    "tags": {
                        "database": "_internal",
                        "engine": "tsm1",
                        "id": "2",
                        "path": "/var/lib/influxdb/data/_internal/monitor/2",
                        "retentionPolicy": "monitor",
                        "walPath": "/var/lib/influxdb/wal/_internal/monitor/2"
                    },
                    "columns": [
                        "WALCompactionTimeMs",
                        "cacheAgeMs",
                        "cachedBytes",
                        "diskBytes",
                        "memBytes",
                        "snapshotCount",
                        "writeDropped",
                        "writeErr",
                        "writeOk"
                    ],
                    "values": [
                        [
                            0,
                            395270,
                            0,
                            0,
                            96464,
                            0,
                            0,
                            0,
                            40
                        ]
                    ]
                },
                {
                    "name": "tsm1_filestore",
                    "tags": {
                        "database": "_internal",
                        "engine": "tsm1",
                        "id": "2",
                        "path": "/var/lib/influxdb/data/_internal/monitor/2",
                        "retentionPolicy": "monitor",
                        "walPath": "/var/lib/influxdb/wal/_internal/monitor/2"
                    },
                    "columns": [
                        "diskBytes",
                        "numFiles"
                    ],
                    "values": [
                        [
                            0,
                            0
                        ]
                    ]
                },
                {
                    "name": "tsm1_wal",
                    "tags": {
                        "database": "_internal",
                        "engine": "tsm1",
                        "id": "2",
                        "path": "/var/lib/influxdb/data/_internal/monitor/2",
                        "retentionPolicy": "monitor",
                        "walPath": "/var/lib/influxdb/wal/_internal/monitor/2"
                    },
                    "columns": [
                        "currentSegmentDiskBytes",
                        "oldSegmentsDiskBytes",
                        "writeErr",
                        "writeOk"
                    ],
                    "values": [
                        [
                            126889,
                            0,
                            0,
                            40
                        ]
                    ]
                },
                {
                    "name": "database",
                    "tags": {
                        "database": "snap"
                    },
                    "columns": [
                        "numMeasurements",
                        "numSeries"
                    ],
                    "values": [
                        [
                            5,
                            7
                        ]
                    ]
                },
                {
                    "name": "database",
                    "tags": {
                        "database": "_internal"
                    },
                    "columns": [
                        "numMeasurements",
                        "numSeries"
                    ],
                    "values": [
                        [
                            12,
                            18
                        ]
                    ]
                },
                {
                    "name": "write",
                    "columns": [
                        "pointReq",
                        "pointReqLocal",
                        "req",
                        "subWriteDrop",
                        "subWriteOk",
                        "writeDrop",
                        "writeError",
                        "writeOk",
                        "writeTimeout"
                    ],
                    "values": [
                        [
                            4111,
                            4111,
                            445,
                            0,
                            445,
                            0,
                            0,
                            445,
                            0
                        ]
                    ]
                },
                {
                    "name": "subscriber",
                    "columns": [
                        "pointsWritten",
                        "writeFailures"
                    ],
                    "values": [
                        [
                            0,
                            0
                        ]
                    ]
                },
                {
                    "name": "cq",
                    "columns": [
                        "queryFail",
                        "queryOk"
                    ],
                    "values": [
                        [
                            0,
                            0
                        ]
                    ]
                },
                {
                    "name": "httpd",
                    "tags": {
                        "bind": ":8086"
                    },
                    "columns": [
                        "authFail",
                        "clientError",
                        "pingReq",
                        "pointsWrittenDropped",
                        "pointsWrittenFail",
                        "pointsWrittenOK",
                        "queryReq",
                        "queryReqDurationNs",
                        "queryRespBytes",
                        "req",
                        "reqActive",
                        "reqDurationNs",
                        "serverError",
                        "statusReq",
                        "writeReq",
                        "writeReqActive",
                        "writeReqBytes",
                        "writeReqDurationNs"
                    ],
                    "values": [
                        [
                            0,
                            1,
                            0,
                            0,
                            0,
                            3397,
                            19,
                            31657607,
                            267596,
                            424,
                            1,
                            433538725,
                            0,
                            0,
                            405,
                            0,
                            324645,
                            354251646
                        ]
                    ]
                }
            ]
        }
    ]
}`

var mockDiagnosticResults = `{
    "results": [
        {
            "series": [
                {
                    "name": "build",
                    "columns": [
                        "Branch",
                        "Build Time",
                        "Commit",
                        "Version"
                    ],
                    "values": [
                        [
                            "master",
                            "",
                            "e47cf1f2e83a02443d7115c54f838be8ee959644",
                            "1.1.1"
                        ]
                    ]
                },
                {
                    "name": "network",
                    "columns": [
                        "hostname"
                    ],
                    "values": [
                        [
                            "7d64bd9def1c"
                        ]
                    ]
                },
                {
                    "name": "runtime",
                    "columns": [
                        "GOARCH",
                        "GOMAXPROCS",
                        "GOOS",
                        "version"
                    ],
                    "values": [
                        [
                            "amd64",
                            2,
                            "linux",
                            "go1.7.4"
                        ]
                    ]
                },
                {
                    "name": "system",
                    "columns": [
                        "PID",
                        "currentTime",
                        "started",
                        "uptime"
                    ],
                    "values": [
                        [
                            1,
                            "2017-01-21T01:03:18.387766728Z",
                            "2017-01-21T00:32:06.589263607Z",
                            "31m11.798503239s"
                        ]
                    ]
                }
            ]
        }
    ]
}
`
