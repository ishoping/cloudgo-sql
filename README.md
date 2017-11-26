# 压力测试


## 查询压力测试

    ljx@ljx-X550JX:~$ ab -n 1000 -c 100 http://localhost:8080/service/userinfo?userid=1
    This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
	Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
	Licensed to The Apache Software Foundation, http://www.apache.org/

	Benchmarking localhost (be patient)
	Completed 100 requests
	Completed 200 requests
	Completed 300 requests
	Completed 400 requests
	Completed 500 requests
	Completed 600 requests
	Completed 700 requests
	Completed 800 requests
	Completed 900 requests
	Completed 1000 requests
	Finished 1000 requests


	Server Software:        
	Server Hostname:        localhost
	Server Port:            8080

	Document Path:          /service/userinfo?userid=1
	Document Length:        95 bytes

	Concurrency Level:      100
	Time taken for tests:   0.574 seconds
	Complete requests:      1000
	Failed requests:        0
	Non-2xx responses:      1000
	Total transferred:      227000 bytes
	HTML transferred:       95000 bytes
	Requests per second:    1741.66 [#/sec] (mean)
	Time per request:       57.417 [ms] (mean)
	Time per request:       0.574 [ms] (mean, across all concurrent requests)
	Transfer rate:          386.09 [Kbytes/sec] received

	Connection Times (ms)
				  min  mean[+/-sd] median   max
	Connect:        0    2   3.5      1      21
	Processing:     1   54  60.5     39     277
	Waiting:        0   53  60.3     37     272
	Total:          1   56  60.7     41     279

	Percentage of the requests served within a certain time (ms)
	  50%     41
	  66%     50
	  75%     58
	  80%     65
	  90%    103
	  95%    218
	  98%    256
	  99%    272
	 100%    279 (longest request)

## 插入压力测试

	$ ab -n 1000  -c 100 -p 'post.txt' -T 'application/x-www-form-urlencoded' http://localhost:8080/service/userinfo
	This is ApacheBench, Version 2.3 <$Revision: 1706008 $>
	Copyright 1996 Adam Twiss, Zeus Technology Ltd, http://www.zeustech.net/
	Licensed to The Apache Software Foundation, http://www.apache.org/

	Benchmarking localhost (be patient)
	Completed 100 requests
	Completed 200 requests
	Completed 300 requests
	Completed 400 requests
	Completed 500 requests
	Completed 600 requests
	Completed 700 requests
	Completed 800 requests
	Completed 900 requests
	Completed 1000 requests
	Finished 1000 requests


	Server Software:        
	Server Hostname:        localhost
	Server Port:            8080

	Document Path:          /service/userinfo
	Document Length:        110 bytes

	Concurrency Level:      100
	Time taken for tests:   0.813 seconds
	Complete requests:      1000
	Failed requests:        999
	   (Connect: 0, Receive: 0, Length: 999, Exceptions: 0)
	Total transferred:      237780 bytes
	Total body sent:        194000
	HTML transferred:       113780 bytes
	Requests per second:    1230.44 [#/sec] (mean)
	Time per request:       81.272 [ms] (mean)
	Time per request:       0.813 [ms] (mean, across all concurrent requests)
	Transfer rate:          285.72 [Kbytes/sec] received
							233.11 kb/s sent
							518.83 kb/s total

	Connection Times (ms)
				  min  mean[+/-sd] median   max
	Connect:        0    1   1.2      0       7
	Processing:    41   78  17.0     77     131
	Waiting:       41   77  17.0     77     131
	Total:         41   79  17.2     78     132

	Percentage of the requests served within a certain time (ms)
	  50%     78
	  66%     85
	  75%     89
	  80%     92
	  90%    101
	  95%    107
	  98%    122
	  99%    126
	 100%    132 (longest request)
