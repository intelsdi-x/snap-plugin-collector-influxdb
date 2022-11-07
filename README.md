DISCONTINUATION OF PROJECT. 

This project will no longer be maintained by Intel.

This project has been identified as having known security escapes.

Intel has ceased development and contributions including, but not limited to, maintenance, bug fixes, new releases, or updates, to this project.  

Intel no longer accepts patches to this project.
# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**

# Snap collector plugin - influxdb

This plugin collects statistical and diagnostic information about each InfluxDB node. This information can be very useful to assist with troubleshooting and performance analysis of the database itself.

This plugin has ability to gather InfluxDB internal system monitoring information in response to the following commands: `SHOW STATS` and `SHOW DIAGNOSTICS`.

The plugin is used in the [Snap framework](http://github.com/intelsdi-x/snap).				
1. [Getting Started](#getting-started)
   * [System Requirements](#system-requirements)
   * [Operating systems](#operating-systems)
   * [Installation](#installation)
   * [Configuration and Usage](#configuration-and-usage)
2. [Documentation](#documentation)
   * [Global Config](#global-config)
   * [Collected Metrics](#collected-metrics)
   * [Examples](#examples)
   * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### System Requirements

- Linux 64bit system
- [InfluxDB](https://github.com/influxdata/influxdb) (version 0.9 or higher)

### Operating systems

All OSs currently supported by Snap:
* Linux/amd64

### Installation
#### Download the plugin binary:

You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-collector-influxdb/releases) page. Download the plugin from the latest release and load it into `snapteld` (`/opt/snap/plugins` is the default location for Snap packages).

#### To build the plugin binary:

Fork https://github.com/intelsdi-x/snap-plugin-collector-influxdb  
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-influxdb.git
```

Build the Snap influxdb plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

### Configuration and Usage

* Set up the [Snap framework](https://github.com/intelsdi-x/snap#getting-started)
* Load the plugin and create a task, see example in [Examples](#examples).

## Documentation

To learn more about influxDB System Monitoring, visit:

* [InfluxDB Server Monitoring Documentation](https://docs.influxdata.com/influxdb/v0.9/administration/statistics/)
* [InfluxDB Server Monitoring README](https://github.com/influxdata/influxdb/blob/master/monitor/README.md)
* blog post ["How to use the show stats command to monitor InfluxDB"](https://influxdata.com/blog/how-to-use-the-show-stats-command-and-the-_internal-database-to-monitor-influxdb/)

### Global Config

Global configuration file is described in Snap's documentation. For this plugin section "influxdb" in "collector" specifing the following options needs to be added (see an [exemplary config file](examples/configs/snap-config-sample.json)):

Name 	  	| Data Type | Description
------------|-----------|-----------------------
"host" 		| string 	| hostname of InfluxDB http API
"port" 		| int	 	| port of InfluxDB http API (by default 8086)
"user" 		| string 	| user name
"password" 	| string 	| user password

### Collected Metrics

List of collected metrics is described in [METRICS.md](METRICS.md).

### Examples

Example of running Snap influxdb collector and writing data to file.
Download an [example of Snap global config](examples/configs/snap-config-sample.json) file.
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-influxdb/master/examples/configs/snap-config-sample.json
```
Ensure [Snap daemon is running](https://github.com/intelsdi-x/snap#running-snap) with provided configuration file:
* command line: `snapteld -l 1 -t 0 --config snap-config-sample.json&`

Download and load Snap plugins:
```
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-collector-influxdb/latest/linux/x86_64/snap-plugin-collector-influxdb
$ wget http://snap.ci.snap-telemetry.io/plugins/snap-plugin-publisher-file/latest/linux/x86_64/snap-plugin-publisher-file
$ chmod 755 snap-plugin-*
$ snaptel plugin load snap-plugin-collector-influxdb
$ snaptel plugin load snap-plugin-publisher-file
```

See all available metrics:

```
$ snaptel metric list

NAMESPACE                                                VERSIONS
/intel/influxdb/diagn/build/Branch                       2
/intel/influxdb/diagn/build/Commit                       2
/intel/influxdb/diagn/build/Version                      2
/intel/influxdb/diagn/network/hostname                   2
/intel/influxdb/diagn/runtime/GOARCH                     2
/intel/influxdb/diagn/runtime/GOMAXPROCS                 2
/intel/influxdb/diagn/runtime/GOOS                       2
/intel/influxdb/diagn/runtime/version                    2
/intel/influxdb/diagn/system/PID                         2
/intel/influxdb/diagn/system/currentTime                 2
/intel/influxdb/diagn/system/started                     2
/intel/influxdb/diagn/system/uptime                      2
/intel/influxdb/stat/engine/1/blks_write                 2
/intel/influxdb/stat/engine/1/blks_write_bytes           2
/intel/influxdb/stat/engine/1/blks_write_bytes_c         2
/intel/influxdb/stat/engine/1/points_write               2
/intel/influxdb/stat/engine/1/points_write_dedupe        2
/intel/influxdb/stat/engine/2/blks_write                 2
/intel/influxdb/stat/engine/2/blks_write_bytes           2
/intel/influxdb/stat/engine/2/blks_write_bytes_c         2
/intel/influxdb/stat/engine/2/points_write               2
/intel/influxdb/stat/engine/2/points_write_dedupe        2
/intel/influxdb/stat/httpd/auth_fail                     2
/intel/influxdb/stat/httpd/ping_req                      2
/intel/influxdb/stat/httpd/points_written_ok             2
/intel/influxdb/stat/httpd/query_req                     2
/intel/influxdb/stat/httpd/query_resp_bytes              2
/intel/influxdb/stat/httpd/req                           2
/intel/influxdb/stat/httpd/write_req                     2
/intel/influxdb/stat/httpd/write_req_bytes               2
/intel/influxdb/stat/runtime/Alloc                       2
/intel/influxdb/stat/runtime/Frees                       2
/intel/influxdb/stat/runtime/HeapAlloc                   2
/intel/influxdb/stat/runtime/HeapIdle                    2
/intel/influxdb/stat/runtime/HeapInUse                   2
/intel/influxdb/stat/runtime/HeapObjects                 2
/intel/influxdb/stat/runtime/HeapReleased                2
/intel/influxdb/stat/runtime/HeapSys                     2
/intel/influxdb/stat/runtime/Lookups                     2
/intel/influxdb/stat/runtime/Mallocs                     2
/intel/influxdb/stat/runtime/NumGC                       2
/intel/influxdb/stat/runtime/NumGoroutine                2
/intel/influxdb/stat/runtime/PauseTotalNs                2
/intel/influxdb/stat/runtime/Sys                         2
/intel/influxdb/stat/runtime/TotalAlloc                  2
/intel/influxdb/stat/shard/1/fields_create               2
/intel/influxdb/stat/shard/1/series_create               2
/intel/influxdb/stat/shard/1/write_points_ok             2
/intel/influxdb/stat/shard/1/write_req                   2
/intel/influxdb/stat/shard/2/fields_create               2
/intel/influxdb/stat/shard/2/series_create               2
/intel/influxdb/stat/shard/2/write_points_ok             2
/intel/influxdb/stat/shard/2/write_req                   2
/intel/influxdb/stat/write/point_req                     2
/intel/influxdb/stat/write/point_req_local               2
/intel/influxdb/stat/write/req                           2
/intel/influxdb/stat/write/write_ok                      2
```

Download an [exemplary task manifest](examples/tasks/influxdb-file.json) and load it:
```
$ curl -sfLO https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-influxdb/master/examples/tasks/influxdb-file.json
$ snaptel task create -t influxdb-file.json
Using task manifest to create task
Task created
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
Name: Task-02dd7ff4-8106-47e9-8b86-70067cd0a850
State: Running
```

See realtime output from `snaptel task watch <task_id>` (CTRL+C to exit)
```
$ snaptel task watch 02dd7ff4-8106-47e9-8b86-70067cd0a850
Watching Task (02dd7ff4-8106-47e9-8b86-70067cd0a850):
NAMESPACE                                        DATA                    TIMESTAMP                                       SOURCE
/intel/influxdb/diagn/system/PID                 16191                   2016-02-26 09:25:47.353886681 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/httpd/auth_fail             63                      2016-02-26 09:25:47.354068464 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/httpd/ping_req              14                      2016-02-26 09:25:47.354073486 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/httpd/points_written_ok     6.69802551e+08          2016-02-26 09:25:47.354137322 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/httpd/query_req             2310                    2016-02-26 09:25:47.354132981 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/httpd/query_resp_bytes      2.0678075e+07           2016-02-26 09:25:47.354113672 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/httpd/req                   6.824003e+06            2016-02-26 09:25:47.354028193 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/httpd/write_req             6.821614e+06            2016-02-26 09:25:47.354095848 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/httpd/write_req_bytes       7.1455502846e+10        2016-02-26 09:25:47.354105454 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/runtime/Alloc               1.07946848e+08          2016-02-26 09:25:47.354209458 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/runtime/Frees               1.5858776371e+10        2016-02-26 09:25:47.354412092 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/write/point_req             6.82720753e+08          2016-02-26 09:25:47.354579209 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/write/point_req_local       6.82720753e+08          2016-02-26 09:25:47.354619612 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/write/req                   7.000631e+06            2016-02-26 09:25:47.354624219 +0000 UTC         node-25.domain.tld
/intel/influxdb/stat/write/write_ok              7.000823e+06            2016-02-26 09:25:47.354609029 +0000 UTC         node-25.domain.tld
```

This data is published to a file `/tmp/published_influxdb_internal_monitoring` per task specification

Stop task:
```
$ snaptel task stop 02dd7ff4-8106-47e9-8b86-70067cd0a850
Task stopped:
ID: 02dd7ff4-8106-47e9-8b86-70067cd0a850
```

### Roadmap

There isn't a current roadmap for this plugin, but it is in active development. As we launch this plugin, we do not have any outstanding requirements for the next release.

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-influxdb/issues) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-influxdb/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support) or visit [Slack](http://slack.snap-telemetry.io).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
Snap, along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
* Author: 	[Izabella Raulin](https://github.com/IzabellaRaulin)
* Author: 	[Marcin Krolik](https://github.com/marcin-krolik)
