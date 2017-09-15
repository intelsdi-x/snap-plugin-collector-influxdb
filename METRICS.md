# snap plugin collector - influxdb

## Collected Metrics
This plugin has the ability to gather:

a) all **diagnostic information** of InfluxDB system itself, represented by the metrics with prefix `/intel/influxdb/diagn/`

b) all **statistical information** of InfluxDB system itself, represented by the metrics with prefix `/intel/influxdb/stat/`
                                                                                                
Metric Name | Data Type | Description
------------ | ---------|-------------
/intel/influxdb/diagn/build/Branch | string | the branch name
/intel/influxdb/diagn/build/Commit | string | the commit
/intel/influxdb/diagn/build/Version | string | the InfluxDB Version
/intel/influxdb/diagn/network/hostname | string | the InfluxDB host name
/intel/influxdb/diagn/runtime/GOARCH | string |the target architecture
/intel/influxdb/diagn/runtime/GOMAXPROCS | int | the maximum number of CPUs that can be executing simultaneously and returns the previous setting 
/intel/influxdb/diagn/runtime/GOOS | string |the target operating system
/intel/influxdb/diagn/runtime/version | string |the Go tree's version
/intel/influxdb/diagn/system/PID | int | the process id
/intel/influxdb/diagn/system/currentTime | string | the current local time
/intel/influxdb/diagn/system/started | string | the started time
/intel/influxdb/diagn/system/uptime | string | the uptime
| |
/intel/influxdb/stat/engine/\<engine_id>/blks_write | int | the number of written blocks into database engine
/intel/influxdb/stat/engine/\<engine_id>/blks_write_bytes | int | sum of blocks, in bytes, written into database engine
/intel/influxdb/stat/engine/\<engine_id>/blks_write_bytes_c | int | sum of compressed blocks, in bytes, written into database engine
/intel/influxdb/stat/engine/\<engine_id>/points_write | int | the number of written points into database engine
/intel/influxdb/stat/engine/\<engine_id>/points_write_dedupe | int | the number of written deduplicated points into database engine
/intel/influxdb/stat/httpd/auth_fail | int | the number of authentication failures
/intel/influxdb/stat/httpd/ping_req | int | the number of ping requests served
/intel/influxdb/stat/httpd/points_written_ok | int | the number of points that failed to be written
/intel/influxdb/stat/httpd/query_req | int | the number of query requests served
/intel/influxdb/stat/httpd/query_resp_bytes | int | the sum of all bytes returned in query responses
/intel/influxdb/stat/httpd/req | int | the number of HTTP requests served
/intel/influxdb/stat/httpd/write_req | int | the number of query requests served
/intel/influxdb/stat/httpd/write_req_bytes | int | the sum of all bytes in write requests
/intel/influxdb/stat/runtime/Alloc | int | bytes of the memory allocated and not yet freed
/intel/influxdb/stat/runtime/Frees | int | the number of frees
/intel/influxdb/stat/runtime/HeapAlloc | int | bytes of the heap space allocated and not yet freed
/intel/influxdb/stat/runtime/HeapIdle | int | bytes of the heap space in idle spans
/intel/influxdb/stat/runtime/HeapInUse | int | bytes of the heap space in non-idle span
/intel/influxdb/stat/runtime/HeapObjects | int | total number of allocated objects on the heap
/intel/influxdb/stat/runtime/HeapReleased | int | bytes of the heap space released to the OS
/intel/influxdb/stat/runtime/HeapSys | int | bytes of the heap space obtained from system
/intel/influxdb/stat/runtime/Lookups | int | the number of pointer lookups
/intel/influxdb/stat/runtime/Mallocs | int | the number of mallocs
/intel/influxdb/stat/runtime/NumGC | int | the garbage collector statistics, next collection will happen when HeapAlloc ≥ this amount
/intel/influxdb/stat/runtime/NumGoroutine | int | the number of goroutines that currently exist.
/intel/influxdb/stat/runtime/PauseTotalNs | int | the garbage collector statistic, pause total
/intel/influxdb/stat/runtime/Sys | int | bytes of the memory obtained from system
/intel/influxdb/stat/runtime/TotalAlloc | int | bytes of the memory allocated (even if freed)
/intel/influxdb/stat/shard/\<shard_id>/fields_create | int | the number of created fields for shard space
/intel/influxdb/stat/shard/\<shard_id>/series_create | int | the number of created series  for shard space
/intel/influxdb/stat/shard/\<shard_id>/write_points_ok | int | the number of points written successfully for shard space
/intel/influxdb/stat/shard/\<shard_id>/write_req | int | the number of write request for shard space
/intel/influxdb/stat/write/point_req | int | the number of written points
/intel/influxdb/stat/write/point_req_local | int | the number of locally written points
/intel/influxdb/stat/write/req | int | the number of write requests
/intel/influxdb/stat/write/write_ok | int | the number of successful write request

The list of available metrics might be vary depending on the influxdb version or the system configuration.

Diagnostics information are gathered only once at the beginning of collecting process, because they are constant during running the influxdb process.

In task manifest there are declaration of metrics names which will be collected and value of an interval (see [exemplary task manifest](examples/tasks/influxdb-file.json)). By default metrics are gathered once per second.
