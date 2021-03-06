# dbtester

[![Build Status](https://img.shields.io/travis/coreos/dbtester.svg?style=flat-square)](https://travis-ci.org/coreos/dbtester) [![Godoc](http://img.shields.io/badge/go-documentation-blue.svg?style=flat-square)](https://godoc.org/github.com/coreos/dbtester)

Distributed database benchmark tester: etcd, Zookeeper, Consul, zetcd, cetcd


<br><br><hr>
##### Performance Analysis

- Latest test results can be found at https://github.com/coreos/dbtester/tree/master/test-results
- Exploring Performance of etcd, Zookeeper and Consul Consistent Key-value Datastores (February 17, 2017)
  - https://coreos.com/blog/performance-of-etcd.html


<br><br><hr>
##### Project

![dbtester system architecture](./dbtester.png)

- Database Agent
  - https://github.com/coreos/dbtester/tree/master/agent
- Database Client
  - https://github.com/coreos/dbtester/tree/master/control
- System Metrics
  - https://github.com/gyuho/linux-inspect
- Test Data Analysis
  - https://github.com/coreos/dbtester/tree/master/analyze
  - https://github.com/gyuho/dataframe
  - https://github.com/gonum/plot

For etcd, we recommend [etcd benchmark tool](https://github.com/coreos/etcd/tree/master/tools/benchmark).

All logs and results can be found at https://github.com/coreos/dbtester/tree/master/test-results or https://console.cloud.google.com/storage/browser/dbtester-results/?authuser=0&project=etcd-development.


<br><br><hr>
##### Noticeable Warnings: Zookeeper

Snapshot, when writing 1-million entries (256-byte key, 1KB value value), with 500 concurrent clients

```
# snapshot warnings
cd 2017Q1-00-etcd-zookeeper-consul/02-write-1M-keys-best-throughput
grep -r -i fsync-ing\ the zookeeper-r3.4.9-java8-* | less

2017-02-10 18:55:38,997 [myid:3] - WARN  [SyncThread:3:SyncRequestProcessor@148] - Too busy to snap, skipping
2017-02-10 18:55:38,998 [myid:3] - INFO  [SyncThread:3:FileTxnLog@203] - Creating new log file: log.1000c0c51
2017-02-10 18:55:40,855 [myid:3] - INFO  [SyncThread:3:FileTxnLog@203] - Creating new log file: log.1000cd2e6
2017-02-10 18:55:40,855 [myid:3] - INFO  [Snapshot Thread:FileTxnSnapLog@240] - Snapshotting: 0x1000cd1ca to /home/gyuho/zookeeper/zookeeper.data/version-2/snapshot.1000cd1ca
2017-02-10 18:55:46,382 [myid:3] - WARN  [SyncThread:3:FileTxnLog@338] - fsync-ing the write ahead log in SyncThread:3 took 1062ms which will adversely effect operation latency. See the ZooKeeper troubleshooting guide
2017-02-10 18:55:47,471 [myid:3] - WARN  [SyncThread:3:FileTxnLog@338] - fsync-ing the write ahead log in SyncThread:3 took 1084ms which will adversely effect operation latency. See the ZooKeeper troubleshooting guide
2017-02-10 18:55:49,425 [myid:3] - WARN  [SyncThread:3:FileTxnLog@338] - fsync-ing the write ahead log in SyncThread:3 took 1142ms which will adversely effect operation latency. See the ZooKeeper troubleshooting guide
2017-02-10 18:55:51,188 [myid:3] - WARN  [SyncThread:3:FileTxnLog@338] - fsync-ing the write ahead log in SyncThread:3 took 1201ms which will adversely effect operation latency. See the ZooKeeper troubleshooting guide
2017-02-10 18:55:52,292 [myid:3] - WARN  [SyncThread:3:FileTxnLog@338] - fsync-ing the write ahead log in SyncThread:3 took 1102ms which will adversely effect operation latency. See the ZooKeeper troubleshooting guide
```

When writing more than 2-million entries (256-byte key, 1KB value value) with 500 concurrent clients

```
# leader election
cd 2017Q1-00-etcd-zookeeper-consul/04-write-too-many-keys
grep -r -i election\ took  zookeeper-r3.4.9-java8-* | less

# leader election is taking more than 10 seconds...
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:22:16,549 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Follower@61] - FOLLOWING - LEADER ELECTION TOOK - 22978
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:23:02,279 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 10210
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:23:14,498 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 203
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:23:36,303 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 9791
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:23:52,151 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 3836
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:24:13,849 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 9686
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:24:29,694 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 3573
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:24:51,392 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 8686
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:25:07,231 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 3827
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:25:28,940 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 9697
zookeeper-r3.4.9-java8-2-database.log:2017-02-10 19:25:44,772 [myid:2] - INFO  [QuorumPeer[myid=2]/0:0:0:0:0:0:0:0:2181:Leader@361] - LEADING - LEADER ELECTION TOOK - 3820
```


<br><br><hr>
##### Noticeable Warnings: Consul

Snapshot, when writing 1-million entries (256-byte key, 1KB value value), with 500 concurrent clients

```
# snapshot warnings
cd 2017Q1-00-etcd-zookeeper-consul/02-write-1M-keys-best-throughput
grep -r -i installed\ remote consul-v0.7.4-go1.7.5-* | less

    2017/02/10 18:58:43 [INFO] snapshot: Creating new snapshot at /home/gyuho/consul.data/raft/snapshots/2-900345-1486753123478.tmp
    2017/02/10 18:58:45 [INFO] snapshot: reaping snapshot /home/gyuho/consul.data/raft/snapshots/2-849399-1486753096972
    2017/02/10 18:58:46 [INFO] raft: Copied 1223270573 bytes to local snapshot
    2017/02/10 18:58:55 [INFO] raft: Compacting logs from 868354 to 868801
    2017/02/10 18:58:56 [INFO] raft: Installed remote snapshot
    2017/02/10 18:58:57 [INFO] snapshot: Creating new snapshot at /home/gyuho/consul.data/raft/snapshots/2-911546-1486753137827.tmp
    2017/02/10 18:58:59 [INFO] consul.fsm: snapshot created in 32.255µs
    2017/02/10 18:59:01 [INFO] snapshot: reaping snapshot /home/gyuho/consul.data/raft/snapshots/2-873921-1486753116619
    2017/02/10 18:59:02 [INFO] raft: Copied 1238491373 bytes to local snapshot
    2017/02/10 18:59:11 [INFO] raft: Compacting logs from 868802 to 868801
    2017/02/10 18:59:11 [INFO] raft: Installed remote snapshot
```

Logs do not tell much but average latency spikes (e.g. from 70.27517 ms to 10407.900082 ms)



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/2017Q2-01-write-1M-cpu-client-scaling.png" alt="2017Q2-01-write-1M-cpu-client-scaling">

<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/2017Q2-02-write-1M-network-traffic-best-throughput.png" alt="2017Q2-02-write-1M-network-traffic-best-throughput">

<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/2017Q2-01-write-1M-throughput-client-scaling.png" alt="2017Q2-01-write-1M-throughput-client-scaling">

<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/2017Q2-02-write-1M-latency-best-throughput.png" alt="2017Q2-02-write-1M-latency-best-throughput">


<br><br><hr>
##### Write 1M keys, 256-byte key, 1KB value, Best Throughput (etcd 1K clients with 100 conns, Zookeeper 700, Consul 500 clients)

- Google Cloud Compute Engine
- 4 machines of 16 vCPUs + 60 GB Memory + 300 GB SSD (1 for client)
- Ubuntu 16.10 (GNU/Linux kernel 4.8.0-49-generic)
- `ulimit -n` is 120000
- etcd tip (Go 1.8.3, git SHA 47a8156851b5a59665421661edb7c813f8a7993e)
- Zookeeper r3.5.3-beta
  - Java 8
  - javac 1.8.0_131
  - Java(TM) SE Runtime Environment (build 1.8.0_131-b11)
  - Java HotSpot(TM) 64-Bit Server VM (build 25.131-b11, mixed mode)
  - `/usr/bin/java -Djute.maxbuffer=33554432 -Xms50G -Xmx50G`
- Consul v0.8.4 (Go 1.8.3)


```
+---------------------------------------+------------------+-----------------------------+-----------------------+
|                                       | etcd-tip-go1.8.3 | zookeeper-r3.5.3-beta-java8 | consul-v0.8.4-go1.8.3 |
+---------------------------------------+------------------+-----------------------------+-----------------------+
|                         TOTAL-SECONDS |      27.9797 sec |                143.8585 sec |          135.7728 sec |
|                  TOTAL-REQUEST-NUMBER |        1,000,000 |                   1,000,000 |             1,000,000 |
|                        MAX-THROUGHPUT |   38,526 req/sec |              25,103 req/sec |        15,424 req/sec |
|                        AVG-THROUGHPUT |   35,740 req/sec |               6,913 req/sec |         7,365 req/sec |
|                        MIN-THROUGHPUT |   13,418 req/sec |                   0 req/sec |           195 req/sec |
|                       FASTEST-LATENCY |        5.1907 ms |                   6.7527 ms |            17.7190 ms |
|                           AVG-LATENCY |       27.9170 ms |                  55.4371 ms |            67.8635 ms |
|                       SLOWEST-LATENCY |      129.6517 ms |                4427.4805 ms |          2665.0249 ms |
|                           Latency p10 |     12.783090 ms |                15.327740 ms |          29.877078 ms |
|                           Latency p25 |     16.081346 ms |                21.706332 ms |          33.992948 ms |
|                           Latency p50 |     22.047040 ms |                37.275107 ms |          40.148835 ms |
|                           Latency p75 |     35.297635 ms |                57.453429 ms |          54.282575 ms |
|                           Latency p90 |     53.916881 ms |                79.224931 ms |         109.468689 ms |
|                           Latency p95 |     60.144462 ms |                93.233345 ms |         235.236038 ms |
|                           Latency p99 |     73.229996 ms |               456.307896 ms |         464.681161 ms |
|                         Latency p99.9 |     94.903421 ms |              2128.132040 ms |         801.018344 ms |
|      SERVER-TOTAL-NETWORK-RX-DATA-SUM |           5.0 GB |                      5.8 GB |                5.6 GB |
|      SERVER-TOTAL-NETWORK-TX-DATA-SUM |           3.8 GB |                      4.7 GB |                4.4 GB |
|           CLIENT-TOTAL-NETWORK-RX-SUM |           277 MB |                      384 MB |                207 MB |
|           CLIENT-TOTAL-NETWORK-TX-SUM |           1.4 GB |                      1.4 GB |                1.5 GB |
|                  SERVER-MAX-CPU-USAGE |         406.67 % |                    492.00 % |              405.40 % |
|               SERVER-MAX-MEMORY-USAGE |           1.2 GB |                       17 GB |                4.9 GB |
|                  CLIENT-MAX-CPU-USAGE |         468.00 % |                    208.00 % |              189.00 % |
|               CLIENT-MAX-MEMORY-USAGE |           112 MB |                      4.2 GB |                 87 MB |
|                    CLIENT-ERROR-COUNT |                0 |                       5,451 |                     0 |
|  SERVER-AVG-READS-COMPLETED-DELTA-SUM |               78 |                         247 |                    12 |
|    SERVER-AVG-SECTORS-READS-DELTA-SUM |                0 |                           0 |                     0 |
| SERVER-AVG-WRITES-COMPLETED-DELTA-SUM |           97,145 |                     335,863 |               660,796 |
|  SERVER-AVG-SECTORS-WRITTEN-DELTA-SUM |       20,655,776 |                  48,217,560 |            71,342,952 |
|           SERVER-AVG-DISK-SPACE-USAGE |           2.6 GB |                       10 GB |                2.9 GB |
+---------------------------------------+------------------+-----------------------------+-----------------------+


zookeeper__r3_5_3_beta errors:
"zk: could not connect to a server" (count 805)
"zk: connection closed" (count 4,646)
```


<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-LATENCY-MS.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-LATENCY-MS">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-LATENCY-MS-BY-KEY.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-LATENCY-MS-BY-KEY">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-LATENCY-MS-BY-KEY-ERROR-POINTS.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-LATENCY-MS-BY-KEY-ERROR-POINTS">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-THROUGHPUT.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-THROUGHPUT">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-VOLUNTARY-CTXT-SWITCHES.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-VOLUNTARY-CTXT-SWITCHES">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-NON-VOLUNTARY-CTXT-SWITCHES.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-NON-VOLUNTARY-CTXT-SWITCHES">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-CPU.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-CPU">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/MAX-CPU.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/MAX-CPU">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-VMRSS-MB.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-VMRSS-MB">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-VMRSS-MB-BY-KEY.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-VMRSS-MB-BY-KEY">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-VMRSS-MB-BY-KEY-ERROR-POINTS.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-VMRSS-MB-BY-KEY-ERROR-POINTS">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-READS-COMPLETED-DELTA.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-READS-COMPLETED-DELTA">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-SECTORS-READ-DELTA.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-SECTORS-READ-DELTA">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-WRITES-COMPLETED-DELTA.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-WRITES-COMPLETED-DELTA">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-SECTORS-WRITTEN-DELTA.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-SECTORS-WRITTEN-DELTA">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-READ-BYTES-NUM-DELTA.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-READ-BYTES-NUM-DELTA">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-WRITE-BYTES-NUM-DELTA.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-WRITE-BYTES-NUM-DELTA">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-RECEIVE-BYTES-NUM-DELTA.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-RECEIVE-BYTES-NUM-DELTA">



<img src="https://storage.googleapis.com/dbtester-results/2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-TRANSMIT-BYTES-NUM-DELTA.svg" alt="2017Q2-02-etcd-zookeeper-consul/write-1M-keys-best-throughput/AVG-TRANSMIT-BYTES-NUM-DELTA">



