Smcroute
========

Smcroute is client library for [smcroute](http://troglobit.com/smcroute.html) written in [Go](http://golang.org/) programming language.

Library uses unix domain socket for communication.

[Documentation](https://godoc.org/github.com/BradburyLab/go/smcroute)

Quick-start
-----------

```bash
$ sudo chmod 777 /var/run/smcroute
```

```go
package main

import (
  "fmt"

  "github.com/BradburyLab/go/smcroute"
)

// sudo smcroute -j eth0.1 239.255.1.1
func main() {
  client := smcroute.NewClient()
  cmdJoin := smcroute.NewCMD(smcroute.CMD_JOIN, "eth0.1", "239.255.1.1")
  resp, e := client.Exec(cmdJoin)
  fmt.Println(resp, e)
}
```

Join and leave
--------------

```bash
$ sudo chmod 777 /var/run/smcroute
$ ip maddr show eth0.1
2:      eth0.1
        link  33:33:00:00:00:01
        link  01:00:5e:00:00:01
        link  33:33:ff:93:e9:07
        link  33:33:00:00:02:02
        inet  224.0.0.1
        inet6 ff02::202
        inet6 ff02::1:ff93:e907
        inet6 ff02::1
        inet6 ff01::1
```

```go
package main

import (
  "fmt"

  "github.com/BradburyLab/go/smcroute"
)

// sudo smcroute -j eth0.1 239.255.1.1
func main() {
  client := smcroute.NewClient()
  cmdJoin := smcroute.NewCMD(smcroute.CMD_JOIN, "eth0.1", "239.255.1.1")
  resp, e := client.Exec(cmdJoin)
  fmt.Println(resp, e)
}
```

```bash
$ ip maddr show eth0.1
2:      eth0.1
        link  33:33:00:00:00:01
        link  01:00:5e:00:00:01
        link  33:33:ff:93:e9:07
        link  33:33:00:00:02:02
        inet  224.0.0.1
        inet  239.255.1.1        # <--- +1
        inet6 ff02::202
        inet6 ff02::1:ff93:e907
        inet6 ff02::1
        inet6 ff01::1
```

```go
package main

import (
  "fmt"

  "github.com/BradburyLab/go/smcroute"
)

// sudo smcroute -l eth0.1 239.255.1.1
func main() {
  client := smcroute.NewClient()
  cmdLeave := smcroute.NewCMD(smcroute.CMD_LEAVE, "eth0.1", "239.255.1.1")
  resp, e := client.Exec(cmdLeave)
  fmt.Println(resp, e)
}
```

```bash
$ ip maddr show eth0.1
2:      eth0.1
        link  33:33:00:00:00:01
        link  01:00:5e:00:00:01
        link  33:33:ff:93:e9:07
        link  33:33:00:00:02:02
        inet  224.0.0.1
        inet6 ff02::202
        inet6 ff02::1:ff93:e907
        inet6 ff02::1
        inet6 ff01::1
```
