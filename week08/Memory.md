机器配置：8 Core， 32 GB

```shell
$ redis-benchmark -t get,set -n 100000 -q -d 10

SET: 111482.72 requests per second, p50=0.223 msec
GET: 111856.8 requests per second, p50=0.223 msec

$ redis-benchmark -t get,set -n 100000 -q -d 20
SET: 110987.79 requests per second, p50=0.231 msec
GET: 110619.47 requests per second, p50=0.231 msec

$ redis-benchmark -t get,set -n 100000 -q -d 50
SET: 110011.00 requests per second, p50=0.231 msec
GET: 112994.35 requests per second, p50=0.223 msec

$ redis-benchmark -t get,set -n 100000 -q -d 100
SET: 109289.62 requests per second, p50=0.231 msec
GET: 110987.79 requests per second, p50=0.223 msec

$ redis-benchmark -t get,set -n 100000 -q -d 200
SET: 108108.11 requests per second, p50=0.231 msec
GET: 110497.24 requests per second, p50=0.231 msec

$ redis-benchmark -t get,set -n 100000 -q -d 1024
SET: 107991.36 requests per second, p50=0.231 msec
GET: 109170.30 requests per second, p50=0.231 msec

$ redis-benchmark -t get,set -n 100000 -q -d 5120
SET: 104275.29 requests per second, p50=0.247 msec
GET: 105374.08 requests per second, p50=0.239 msec
```

写入50w数据

```markdown
 /usr/local/Cellar/redis/6.2.6/bin  redis-benchmark -t set -r 100000 -n 500000 -d 10
====== SET ======
  500000 requests completed in 4.59 seconds
  50 parallel clients
  10 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.135 milliseconds (cumulative count 1)
50.000% <= 0.263 milliseconds (cumulative count 259529)
75.000% <= 0.343 milliseconds (cumulative count 381383)
87.500% <= 0.415 milliseconds (cumulative count 441868)
93.750% <= 0.479 milliseconds (cumulative count 469401)
96.875% <= 0.543 milliseconds (cumulative count 485132)
98.438% <= 0.599 milliseconds (cumulative count 492748)
99.219% <= 0.663 milliseconds (cumulative count 496376)
99.609% <= 0.735 milliseconds (cumulative count 498157)
99.805% <= 0.807 milliseconds (cumulative count 499064)
99.902% <= 0.903 milliseconds (cumulative count 499541)
99.951% <= 1.007 milliseconds (cumulative count 499757)
99.976% <= 1.279 milliseconds (cumulative count 499878)
99.988% <= 2.255 milliseconds (cumulative count 499939)
99.994% <= 3.519 milliseconds (cumulative count 499970)
99.997% <= 4.439 milliseconds (cumulative count 499985)
99.998% <= 4.559 milliseconds (cumulative count 499993)
99.999% <= 4.735 milliseconds (cumulative count 499997)
100.000% <= 4.799 milliseconds (cumulative count 499999)
100.000% <= 4.839 milliseconds (cumulative count 500000)
100.000% <= 4.839 milliseconds (cumulative count 500000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
4.557% <= 0.207 milliseconds (cumulative count 22785)
66.593% <= 0.303 milliseconds (cumulative count 332966)
87.393% <= 0.407 milliseconds (cumulative count 436965)
95.254% <= 0.503 milliseconds (cumulative count 476271)
98.691% <= 0.607 milliseconds (cumulative count 493453)
99.501% <= 0.703 milliseconds (cumulative count 497506)
99.813% <= 0.807 milliseconds (cumulative count 499064)
99.908% <= 0.903 milliseconds (cumulative count 499541)
99.951% <= 1.007 milliseconds (cumulative count 499757)
99.966% <= 1.103 milliseconds (cumulative count 499829)
99.974% <= 1.207 milliseconds (cumulative count 499868)
99.976% <= 1.303 milliseconds (cumulative count 499880)
99.977% <= 1.407 milliseconds (cumulative count 499887)
99.979% <= 1.503 milliseconds (cumulative count 499894)
99.980% <= 1.607 milliseconds (cumulative count 499900)
99.982% <= 2.007 milliseconds (cumulative count 499911)
99.984% <= 2.103 milliseconds (cumulative count 499922)
99.993% <= 3.103 milliseconds (cumulative count 499963)
99.994% <= 4.103 milliseconds (cumulative count 499972)
100.000% <= 5.103 milliseconds (cumulative count 500000)

Summary:
  throughput summary: 108956.20 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.297     0.128     0.263     0.503     0.639     4.839
```

写入50w 20数据

```markdown
 /usr/local/Cellar/redis/6.2.6/bin  redis-benchmark -t set -r 100000 -n 500000 -d 20
====== SET ======
  500000 requests completed in 4.51 seconds
  50 parallel clients
  20 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.119 milliseconds (cumulative count 2)
50.000% <= 0.231 milliseconds (cumulative count 279731)
75.000% <= 0.255 milliseconds (cumulative count 385054)
87.500% <= 0.287 milliseconds (cumulative count 439456)
93.750% <= 0.335 milliseconds (cumulative count 470361)
96.875% <= 0.399 milliseconds (cumulative count 485325)
98.438% <= 0.455 milliseconds (cumulative count 492421)
99.219% <= 0.503 milliseconds (cumulative count 496195)
99.609% <= 0.551 milliseconds (cumulative count 498177)
99.805% <= 0.599 milliseconds (cumulative count 499063)
99.902% <= 0.663 milliseconds (cumulative count 499544)
99.951% <= 0.751 milliseconds (cumulative count 499768)
99.976% <= 0.863 milliseconds (cumulative count 499883)
99.988% <= 1.039 milliseconds (cumulative count 499939)
99.994% <= 1.703 milliseconds (cumulative count 499970)
99.997% <= 1.895 milliseconds (cumulative count 499985)
99.998% <= 2.015 milliseconds (cumulative count 499993)
99.999% <= 2.055 milliseconds (cumulative count 499997)
100.000% <= 2.095 milliseconds (cumulative count 499999)
100.000% <= 2.103 milliseconds (cumulative count 500000)
100.000% <= 2.103 milliseconds (cumulative count 500000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
3.720% <= 0.207 milliseconds (cumulative count 18600)
90.857% <= 0.303 milliseconds (cumulative count 454284)
97.302% <= 0.407 milliseconds (cumulative count 486511)
99.239% <= 0.503 milliseconds (cumulative count 496195)
99.830% <= 0.607 milliseconds (cumulative count 499152)
99.935% <= 0.703 milliseconds (cumulative count 499677)
99.968% <= 0.807 milliseconds (cumulative count 499839)
99.983% <= 0.903 milliseconds (cumulative count 499914)
99.987% <= 1.007 milliseconds (cumulative count 499935)
99.988% <= 1.103 milliseconds (cumulative count 499942)
99.989% <= 1.207 milliseconds (cumulative count 499945)
99.990% <= 1.303 milliseconds (cumulative count 499950)
99.991% <= 1.607 milliseconds (cumulative count 499956)
99.994% <= 1.703 milliseconds (cumulative count 499970)
99.996% <= 1.807 milliseconds (cumulative count 499979)
99.997% <= 1.903 milliseconds (cumulative count 499986)
99.998% <= 2.007 milliseconds (cumulative count 499992)
100.000% <= 2.103 milliseconds (cumulative count 500000)

Summary:
  throughput summary: 110741.98 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.247     0.112     0.231     0.351     0.487     2.103
```

写入50w  50数据

```markdown
 /usr/local/Cellar/redis/6.2.6/bin  redis-benchmark -t set -r 100000 -n 500000 -d 50
====== SET ======
  500000 requests completed in 4.53 seconds
  50 parallel clients
  50 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.111 milliseconds (cumulative count 1)
50.000% <= 0.231 milliseconds (cumulative count 252435)
75.000% <= 0.263 milliseconds (cumulative count 390158)
87.500% <= 0.295 milliseconds (cumulative count 442515)
93.750% <= 0.335 milliseconds (cumulative count 469246)
96.875% <= 0.399 milliseconds (cumulative count 484986)
98.438% <= 0.463 milliseconds (cumulative count 492445)
99.219% <= 0.519 milliseconds (cumulative count 496505)
99.609% <= 0.559 milliseconds (cumulative count 498089)
99.805% <= 0.607 milliseconds (cumulative count 499086)
99.902% <= 0.655 milliseconds (cumulative count 499535)
99.951% <= 0.727 milliseconds (cumulative count 499765)
99.976% <= 0.847 milliseconds (cumulative count 499880)
99.988% <= 1.295 milliseconds (cumulative count 499940)
99.994% <= 1.439 milliseconds (cumulative count 499970)
99.997% <= 1.551 milliseconds (cumulative count 499985)
99.998% <= 1.639 milliseconds (cumulative count 499993)
99.999% <= 1.687 milliseconds (cumulative count 499997)
100.000% <= 1.711 milliseconds (cumulative count 499999)
100.000% <= 1.719 milliseconds (cumulative count 500000)
100.000% <= 1.719 milliseconds (cumulative count 500000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
3.382% <= 0.207 milliseconds (cumulative count 16910)
90.074% <= 0.303 milliseconds (cumulative count 450372)
97.231% <= 0.407 milliseconds (cumulative count 486156)
99.112% <= 0.503 milliseconds (cumulative count 495558)
99.817% <= 0.607 milliseconds (cumulative count 499086)
99.943% <= 0.703 milliseconds (cumulative count 499714)
99.971% <= 0.807 milliseconds (cumulative count 499857)
99.979% <= 0.903 milliseconds (cumulative count 499894)
99.981% <= 1.007 milliseconds (cumulative count 499903)
99.983% <= 1.103 milliseconds (cumulative count 499916)
99.986% <= 1.207 milliseconds (cumulative count 499929)
99.988% <= 1.303 milliseconds (cumulative count 499940)
99.993% <= 1.407 milliseconds (cumulative count 499964)
99.996% <= 1.503 milliseconds (cumulative count 499978)
99.998% <= 1.607 milliseconds (cumulative count 499990)
100.000% <= 1.703 milliseconds (cumulative count 499998)
100.000% <= 1.807 milliseconds (cumulative count 500000)

Summary:
  throughput summary: 110399.65 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.250     0.104     0.231     0.359     0.503     1.719
```

写入50w 100数据

```markdown
 /usr/local/Cellar/redis/6.2.6/bin  redis-benchmark -t set -r 100000 -n 500000 -d 100
====== SET ======
  500000 requests completed in 4.65 seconds
  50 parallel clients
  100 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.119 milliseconds (cumulative count 1)
50.000% <= 0.239 milliseconds (cumulative count 314917)
75.000% <= 0.263 milliseconds (cumulative count 389617)
87.500% <= 0.295 milliseconds (cumulative count 441574)
93.750% <= 0.343 milliseconds (cumulative count 470991)
96.875% <= 0.407 milliseconds (cumulative count 485303)
98.438% <= 0.471 milliseconds (cumulative count 492819)
99.219% <= 0.519 milliseconds (cumulative count 496158)
99.609% <= 0.567 milliseconds (cumulative count 498076)
99.805% <= 0.615 milliseconds (cumulative count 499025)
99.902% <= 0.679 milliseconds (cumulative count 499528)
99.951% <= 0.759 milliseconds (cumulative count 499772)
99.976% <= 0.839 milliseconds (cumulative count 499883)
99.988% <= 0.935 milliseconds (cumulative count 499940)
99.994% <= 1.823 milliseconds (cumulative count 499970)
99.997% <= 2.127 milliseconds (cumulative count 499985)
99.998% <= 2.255 milliseconds (cumulative count 499993)
99.999% <= 2.303 milliseconds (cumulative count 499997)
100.000% <= 2.319 milliseconds (cumulative count 499999)
100.000% <= 2.335 milliseconds (cumulative count 500000)
100.000% <= 2.335 milliseconds (cumulative count 500000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
1.961% <= 0.207 milliseconds (cumulative count 9805)
89.931% <= 0.303 milliseconds (cumulative count 449653)
97.061% <= 0.407 milliseconds (cumulative count 485303)
99.054% <= 0.503 milliseconds (cumulative count 495268)
99.783% <= 0.607 milliseconds (cumulative count 498916)
99.924% <= 0.703 milliseconds (cumulative count 499619)
99.972% <= 0.807 milliseconds (cumulative count 499859)
99.986% <= 0.903 milliseconds (cumulative count 499929)
99.990% <= 1.007 milliseconds (cumulative count 499949)
99.990% <= 1.103 milliseconds (cumulative count 499950)
99.991% <= 1.503 milliseconds (cumulative count 499956)
99.992% <= 1.607 milliseconds (cumulative count 499959)
99.993% <= 1.703 milliseconds (cumulative count 499963)
99.994% <= 1.807 milliseconds (cumulative count 499968)
99.995% <= 1.903 milliseconds (cumulative count 499973)
99.996% <= 2.007 milliseconds (cumulative count 499979)
99.997% <= 2.103 milliseconds (cumulative count 499984)
100.000% <= 3.103 milliseconds (cumulative count 500000)

Summary:
  throughput summary: 107573.16 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.252     0.112     0.239     0.359     0.503     2.335
```

写入50w 200数据

```markdown
 /usr/local/Cellar/redis/6.2.6/bin  redis-benchmark -t set -r 100000 -n 500000 -d 200
====== SET ======
  500000 requests completed in 4.59 seconds
  50 parallel clients
  200 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.151 milliseconds (cumulative count 2)
50.000% <= 0.231 milliseconds (cumulative count 264545)
75.000% <= 0.255 milliseconds (cumulative count 379229)
87.500% <= 0.295 milliseconds (cumulative count 445134)
93.750% <= 0.335 milliseconds (cumulative count 469291)
96.875% <= 0.407 milliseconds (cumulative count 485168)
98.438% <= 0.471 milliseconds (cumulative count 492545)
99.219% <= 0.527 milliseconds (cumulative count 496491)
99.609% <= 0.567 milliseconds (cumulative count 498093)
99.805% <= 0.615 milliseconds (cumulative count 499092)
99.902% <= 0.663 milliseconds (cumulative count 499514)
99.951% <= 0.719 milliseconds (cumulative count 499757)
99.976% <= 0.863 milliseconds (cumulative count 499887)
99.988% <= 2.079 milliseconds (cumulative count 499940)
99.994% <= 2.351 milliseconds (cumulative count 499971)
99.997% <= 2.511 milliseconds (cumulative count 499985)
99.998% <= 2.727 milliseconds (cumulative count 499993)
99.999% <= 2.871 milliseconds (cumulative count 499997)
100.000% <= 2.943 milliseconds (cumulative count 499999)
100.000% <= 2.975 milliseconds (cumulative count 500000)
100.000% <= 2.975 milliseconds (cumulative count 500000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
2.001% <= 0.207 milliseconds (cumulative count 10006)
90.499% <= 0.303 milliseconds (cumulative count 452497)
97.034% <= 0.407 milliseconds (cumulative count 485168)
99.010% <= 0.503 milliseconds (cumulative count 495049)
99.796% <= 0.607 milliseconds (cumulative count 498979)
99.943% <= 0.703 milliseconds (cumulative count 499713)
99.971% <= 0.807 milliseconds (cumulative count 499854)
99.980% <= 0.903 milliseconds (cumulative count 499898)
99.980% <= 1.007 milliseconds (cumulative count 499900)
99.981% <= 1.903 milliseconds (cumulative count 499905)
99.985% <= 2.007 milliseconds (cumulative count 499926)
99.988% <= 2.103 milliseconds (cumulative count 499941)
100.000% <= 3.103 milliseconds (cumulative count 500000)

Summary:
  throughput summary: 108837.62 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.250     0.144     0.231     0.359     0.503     2.975
```

写入50w 1024数据

```markdown
 /usr/local/Cellar/redis/6.2.6/bin  redis-benchmark -t set -r 100000 -n 500000 -d 1024
====== SET ======
  500000 requests completed in 4.63 seconds
  50 parallel clients
  1024 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.127 milliseconds (cumulative count 2)
50.000% <= 0.239 milliseconds (cumulative count 264854)
75.000% <= 0.271 milliseconds (cumulative count 383818)
87.500% <= 0.311 milliseconds (cumulative count 443778)
93.750% <= 0.359 milliseconds (cumulative count 469710)
96.875% <= 0.423 milliseconds (cumulative count 484490)
98.438% <= 0.487 milliseconds (cumulative count 492301)
99.219% <= 0.543 milliseconds (cumulative count 496289)
99.609% <= 0.591 milliseconds (cumulative count 498062)
99.805% <= 0.655 milliseconds (cumulative count 499092)
99.902% <= 0.719 milliseconds (cumulative count 499534)
99.951% <= 0.807 milliseconds (cumulative count 499768)
99.976% <= 0.919 milliseconds (cumulative count 499882)
99.988% <= 1.063 milliseconds (cumulative count 499940)
99.994% <= 2.335 milliseconds (cumulative count 499970)
99.997% <= 2.503 milliseconds (cumulative count 499985)
99.998% <= 2.591 milliseconds (cumulative count 499993)
99.999% <= 2.623 milliseconds (cumulative count 499997)
100.000% <= 2.639 milliseconds (cumulative count 499999)
100.000% <= 2.647 milliseconds (cumulative count 500000)
100.000% <= 2.647 milliseconds (cumulative count 500000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
1.594% <= 0.207 milliseconds (cumulative count 7972)
87.185% <= 0.303 milliseconds (cumulative count 435925)
96.366% <= 0.407 milliseconds (cumulative count 481831)
98.735% <= 0.503 milliseconds (cumulative count 493673)
99.685% <= 0.607 milliseconds (cumulative count 498425)
99.891% <= 0.703 milliseconds (cumulative count 499455)
99.954% <= 0.807 milliseconds (cumulative count 499768)
99.975% <= 0.903 milliseconds (cumulative count 499874)
99.984% <= 1.007 milliseconds (cumulative count 499921)
99.989% <= 1.103 milliseconds (cumulative count 499945)
99.990% <= 1.207 milliseconds (cumulative count 499950)
100.000% <= 3.103 milliseconds (cumulative count 500000)

Summary:
  throughput summary: 107968.04 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.259     0.120     0.239     0.383     0.527     2.647
```

写入50w 5120数据

```markdown
 /usr/local/Cellar/redis/6.2.6/bin  redis-benchmark -t set -r 100000 -n 500000 -d 5120
====== SET ======
  500000 requests completed in 4.81 seconds
  50 parallel clients
  5120 bytes payload
  keep alive: 1
  host configuration "save": 3600 1 300 100 60 10000
  host configuration "appendonly": no
  multi-thread: no

Latency by percentile distribution:
0.000% <= 0.135 milliseconds (cumulative count 1)
50.000% <= 0.271 milliseconds (cumulative count 258376)
75.000% <= 0.335 milliseconds (cumulative count 376451)
87.500% <= 0.407 milliseconds (cumulative count 438878)
93.750% <= 0.479 milliseconds (cumulative count 469563)
96.875% <= 0.551 milliseconds (cumulative count 484756)
98.438% <= 0.631 milliseconds (cumulative count 492384)
99.219% <= 0.727 milliseconds (cumulative count 496158)
99.609% <= 0.831 milliseconds (cumulative count 498052)
99.805% <= 0.959 milliseconds (cumulative count 499056)
99.902% <= 1.111 milliseconds (cumulative count 499523)
99.951% <= 1.327 milliseconds (cumulative count 499758)
99.976% <= 1.799 milliseconds (cumulative count 499879)
99.988% <= 2.103 milliseconds (cumulative count 499939)
99.994% <= 2.359 milliseconds (cumulative count 499970)
99.997% <= 2.495 milliseconds (cumulative count 499985)
99.998% <= 2.591 milliseconds (cumulative count 499993)
99.999% <= 2.639 milliseconds (cumulative count 499997)
100.000% <= 2.663 milliseconds (cumulative count 499999)
100.000% <= 2.687 milliseconds (cumulative count 500000)
100.000% <= 2.687 milliseconds (cumulative count 500000)

Cumulative distribution of latencies:
0.000% <= 0.103 milliseconds (cumulative count 0)
0.701% <= 0.207 milliseconds (cumulative count 3506)
66.061% <= 0.303 milliseconds (cumulative count 330307)
87.776% <= 0.407 milliseconds (cumulative count 438878)
95.141% <= 0.503 milliseconds (cumulative count 475706)
98.164% <= 0.607 milliseconds (cumulative count 490821)
99.096% <= 0.703 milliseconds (cumulative count 495478)
99.548% <= 0.807 milliseconds (cumulative count 497740)
99.749% <= 0.903 milliseconds (cumulative count 498747)
99.851% <= 1.007 milliseconds (cumulative count 499254)
99.901% <= 1.103 milliseconds (cumulative count 499507)
99.933% <= 1.207 milliseconds (cumulative count 499667)
99.949% <= 1.303 milliseconds (cumulative count 499744)
99.958% <= 1.407 milliseconds (cumulative count 499790)
99.962% <= 1.503 milliseconds (cumulative count 499808)
99.964% <= 1.607 milliseconds (cumulative count 499822)
99.968% <= 1.703 milliseconds (cumulative count 499842)
99.976% <= 1.807 milliseconds (cumulative count 499882)
99.984% <= 1.903 milliseconds (cumulative count 499919)
99.987% <= 2.007 milliseconds (cumulative count 499935)
99.988% <= 2.103 milliseconds (cumulative count 499939)
100.000% <= 3.103 milliseconds (cumulative count 500000)

Summary:
  throughput summary: 103950.10 requests per second
  latency summary (msec):
          avg       min       p50       p95       p99       max
        0.305     0.128     0.271     0.503     0.695     2.687
```



写入前

```markdown
127.0.0.1:6379> info memory
# Memory
used_memory:1111968
used_memory_human:1.06M
used_memory_rss:9445376
used_memory_rss_human:9.01M
used_memory_peak:5208256
used_memory_peak_human:4.97M
used_memory_peak_perc:21.35%
used_memory_overhead:1045760
used_memory_startup:1028320
used_memory_dataset:66208
used_memory_dataset_perc:79.15%
allocator_allocated:1046368
allocator_active:9407488
allocator_resident:9407488
total_system_memory:34359738368
total_system_memory_human:32.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:8.99
allocator_frag_bytes:8361120
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:9.03
mem_fragmentation_bytes:8399008
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入后

```markdown
redis-benchmark -t set -r 100000 -n 500000
127.0.0.1:6379> info memory
# Memory
used_memory:11701568
used_memory_human:11.16M
used_memory_rss:23007232
used_memory_rss_human:21.94M
used_memory_peak:15791184
used_memory_peak_human:15.06M
used_memory_peak_perc:74.10%
used_memory_overhead:6069656
used_memory_startup:1028320
used_memory_dataset:5631912
used_memory_dataset_perc:52.77%
allocator_allocated:11635968
allocator_active:22969344
allocator_resident:22969344
total_system_memory:34359738368
total_system_memory_human:32.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.97
allocator_frag_bytes:11333376
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.98
mem_fragmentation_bytes:11371264
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入50w 10的数据

```markdown
127.0.0.1:6379> info memory
# Memory
used_memory:11694720
used_memory_human:11.15M
used_memory_rss:24932352
used_memory_rss_human:23.78M
used_memory_peak:15791184
used_memory_peak_human:15.06M
used_memory_peak_perc:74.06%
used_memory_overhead:6066736
used_memory_startup:1028320
used_memory_dataset:5627984
used_memory_dataset_perc:52.76%
allocator_allocated:11629120
allocator_active:24894464
allocator_resident:24894464
total_system_memory:34359738368
total_system_memory_human:32.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.14
allocator_frag_bytes:13265344
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:2.14
mem_fragmentation_bytes:13303232
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入50w 20的数据

```markdown
127.0.0.1:6379> info memory
# Memory
used_memory:13287680
used_memory_human:12.67M
used_memory_rss:31334400
used_memory_rss_human:29.88M
used_memory_peak:17377424
used_memory_peak_human:16.57M
used_memory_peak_perc:76.47%
used_memory_overhead:6068096
used_memory_startup:1028320
used_memory_dataset:7219584
used_memory_dataset_perc:58.89%
allocator_allocated:13222080
allocator_active:31296512
allocator_resident:31296512
total_system_memory:34359738368
total_system_memory_human:32.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.37
allocator_frag_bytes:18074432
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:2.37
mem_fragmentation_bytes:18112320
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入50w 50的数据

```markdown
127.0.0.1:6379> info memory
# Memory
used_memory:16465952
used_memory_human:15.70M
used_memory_rss:41078784
used_memory_rss_human:39.18M
used_memory_peak:20555312
used_memory_peak_human:19.60M
used_memory_peak_perc:80.11%
used_memory_overhead:6067776
used_memory_startup:1028320
used_memory_dataset:10398176
used_memory_dataset_perc:67.36%
allocator_allocated:16400352
allocator_active:41040896
allocator_resident:41040896
total_system_memory:34359738368
total_system_memory_human:32.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.50
allocator_frag_bytes:24640544
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:2.50
mem_fragmentation_bytes:24678432
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入50w 100的数据

```markdown
127.0.0.1:6379> info memory
# Memory
used_memory:21233120
used_memory_human:20.25M
used_memory_rss:54353920
used_memory_rss_human:51.84M
used_memory_peak:25322416
used_memory_peak_human:24.15M
used_memory_peak_perc:83.85%
used_memory_overhead:6067536
used_memory_startup:1028320
used_memory_dataset:15165584
used_memory_dataset_perc:75.06%
allocator_allocated:21167520
allocator_active:54316032
allocator_resident:54316032
total_system_memory:34359738368
total_system_memory_human:32.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.57
allocator_frag_bytes:33148512
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:2.57
mem_fragmentation_bytes:33186400
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入50w 200的数据

```markdown
127.0.0.1:6379> info memory
# Memory
used_memory:30762944
used_memory_human:29.34M
used_memory_rss:77258752
used_memory_rss_human:73.68M
used_memory_peak:34852272
used_memory_peak_human:33.24M
used_memory_peak_perc:88.27%
used_memory_overhead:6066696
used_memory_startup:1028320
used_memory_dataset:24696248
used_memory_dataset_perc:83.06%
allocator_allocated:30697344
allocator_active:77220864
allocator_resident:77220864
total_system_memory:34359738368
total_system_memory_human:32.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:2.52
allocator_frag_bytes:46523520
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:2.52
mem_fragmentation_bytes:46561408
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入50w 1024的数据

```markdown
127.0.0.1:6379> info memory
# Memory
used_memory:162655440
used_memory_human:155.12M
used_memory_rss:268283904
used_memory_rss_human:255.86M
used_memory_peak:166744320
used_memory_peak_human:159.02M
used_memory_peak_perc:97.55%
used_memory_overhead:6066936
used_memory_startup:1028320
used_memory_dataset:156588504
used_memory_dataset_perc:96.88%
allocator_allocated:162589840
allocator_active:268246016
allocator_resident:268246016
total_system_memory:34359738368
total_system_memory_human:32.00G
used_memory_lua:37888
used_memory_lua_human:37.00K
used_memory_scripts:0
used_memory_scripts_human:0B
number_of_cached_scripts:0
maxmemory:0
maxmemory_human:0B
maxmemory_policy:noeviction
allocator_frag_ratio:1.65
allocator_frag_bytes:105656176
allocator_rss_ratio:1.00
allocator_rss_bytes:0
rss_overhead_ratio:1.00
rss_overhead_bytes:37888
mem_fragmentation_ratio:1.65
mem_fragmentation_bytes:105694064
mem_not_counted_for_evict:0
mem_replication_backlog:0
mem_clients_slaves:0
mem_clients_normal:17440
mem_aof_buffer:0
mem_allocator:libc
active_defrag_running:0
lazyfree_pending_objects:0
lazyfreed_objects:0
```

写入50w 5120的数据

```markdown

```
