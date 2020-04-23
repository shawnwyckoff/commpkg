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
	"github.com/shawnwyckoff/gpkg/container/gvolume"
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

119 packages in total

**apputil**

gau  gbindata  gdeploy  gdump  gerror  ginstance  gjsonconfig  glog  glogger  glogs  gpanic  gparam  gprofile  gprogress  gtest

**encoding**

gbarcode  gchart  gcolor  gcsv  gexcel  gffmpeg  gmultimedia  gzip

**sys**

gcache  gcharset  gcmd  gconcurrentcounter  gcounter  gcountry  gcpulimit  gcron  gdesktop  gfirewall  gfs  ghdd  gio  gkeyboard  gmachineid  gmem  gproc  groutine  gsignal  gsync  gsysinfo  gtime  gusers

**container**

gapriori  gbinary  gbit  gbloom  gbyte  gcombination  gdecimal  gforecast  ggeo  ggeometry  ginterface  gjson  glist  gmap  gnonce  gnum  gob  gpermutation  gpoly  gqueue  grandom  grange  gscore  gset  gspeed  gstate3  gstring  gstruct  gtaskqueue  gternary  gtimeseries  gvolume  gfileformat

**crypto**

g2fa  gencrypt  ghash  gsha

**database**

gconnectstring  gdriver  gmongo  gredis  gsqldb

**net**

gaddr  gdialer  gheadless  ghtml  ghttp  ghttpserver  ghttputils  gicmp  gkcp  glistener  gmail  gmkcp  gmq  gmtu  gmux  gp2pdns  gprobe  gproxy  gquic  gsmux  gsniffer  gsocks5  gssh  gtun  gtuntap  gupnp  gutp  gweb  gzk

**safe**

gchan  gdeepcopy  gwg