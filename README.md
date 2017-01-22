# httpget
httpget for golang use goroutine to control 100 concurrency to get baidu 1000 search httpcode

## Use

```
-bash-4.1$ time go run httptest.go
....
http://www.baidu.com/s?wd=test229 200
http://www.baidu.com/s?wd=test472 200
http://www.baidu.com/s?wd=test480 200
http://www.baidu.com/s?wd=test486 200
http://www.baidu.com/s?wd=test811 200

real    0m4.235s
user    0m1.091s
sys     0m0.237s

```

##SEE ALSO
when goroutine is 1000,baidu will reject most http request like :
```
http://www.baidu.com/s?wd=search522 Get http://verify.baidu.com/vcode?http://www.baidu.com/s?wd=search522&vif=1: read tcp 220.181.111.155:80: connection reset by peer
url is http://www.baidu.com/s?wd=search522, goroutine id is 526
```

or 
```
http://www.baidu.com/s?wd=search533 Get http://www.baidu.com/s?wd=search533: read tcp 220.181.111.188:80: connection reset by peer
url is http://www.baidu.com/s?wd=search533, goroutine id is 537
```
