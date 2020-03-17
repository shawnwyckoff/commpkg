[![Minimum Go version](https://img.shields.io/badge/go-1.13.0+-9cf.svg)](#go-version-requirements)
[![Go Report Card](https://goreportcard.com/badge/github.com/shawnwyckoff/gpkg)](https://goreportcard.com/report/github.com/shawnwyckoff/gpkg)
[![PRs Welcome](https://img.shields.io/badge/PRs-welcome-brightgreen.svg)](https://github.com/shawnwyckoff/gpkg/pulls)
[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)

**gpkg** is a general purpose Go utility with 100+ packages.

*NOTE: it is still in active development and may not be stable at all times.*

![gpkg logo from golang.org](https://github.com/shawnwyckoff/gpkg/raw/master/gophermart.png)

# Usage

**Install**

```
go get -u github.com/shawnwyckoff/gpkg
```

**gpkg** has a lot of dependencies, if go compiler tell they are required, please install them.

**Example**

```
package main

import (
	"fmt"
	"github.com/shawnwyckoff/gpkg/dsa/gvolume"
)

func main() {
	vol, err := gvolume.ParseString("10 MB")
	fmt.Println(vol.String(), err)
}
```

# Contributors

**gpkg** was developed by [shawnwyckoff](https://github.com/shawnwyckoff) and it requires **a lot of dependencies**.

# In Production

# Packages

114 packages in total

**net**

addr  ddns  dialers  ftp  htmls  httpserver  httpz  icmp  kcps  listeners  mail  mkcp  mq  mtu  mux  p2pdns  probe  proxy  quic  smux  sniffer  socks5  ssh  tun  tuntap  upnp  utp

**sys**

cache  chans  charset  clock  cmd  concurrent_counter  counter  country  cpulimit  cron  deep_copy  desktop  file_format  firewall  fs  hdd  ios  keyboard  machine_id  mem  proc  routine  signals  syncs  sysinfo  users

**apputil**

au  bindata  deploy  dump  errorz  jsonconfig  logger  logz  panic  profile  progress  test

**database**

connect_string  driver  mongo  redis  sqldb

**dsa**

apriori  binaries  bits  bloom  bytez  combinations  crypto  decimals  encrypt  forecast  geo  geometry  gob  hash  interfaces  jsons  list  maps  nonce  num  permutations  poly  queue  randoms  ranges  score  set  sha  speed  state3  stringz  structs  taskqueue  ternary  volume

**encoding**

barcode  chart  color  csv  excel  ffmpeg  format  multimedia  zip