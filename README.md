# paramigo
Send te command from local to remote server through ssh.

# Usage

### Usage of StdinCommand:
```
  -host string
    	ip address
  -port string
    	port (default 22)
  -u string
    	user
  -p string
     	password
  -key string
        SSH private key
  -cmd string
        command that you want to send
```

### Example

```
package main

import (
	"github.com/ryojsb/paramigo"
)

func main() {
	paramigo.StdinCommand()
}
```
```
$ go run exapmle.go -host <IP> -port <port> -u <user> -p <password> -cmd <command>
```


### Usage of InnerCommand:
```
  host string
    	ip address
  port string
    	port (default 22)
  u string
    	user
  p string
     	password
  cmd string
        command that you want to send
```

### Example

```
package main

import (
	"github.com/ryojsb/paramigo"
)

func main() {
    var host string = "192.198.1.1"
    var port string = "22"
    var u string = "user"
    var p string = "password"
    var cmd string = "ls"
	
    paramigo.InnerCommand(host, port, u, p, cmd)
}
```
```
$ go run exapmle.go
```

# Installation

```
$ go get github.com/ryojsb/paramigo
```