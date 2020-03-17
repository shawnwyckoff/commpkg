package addr

import (
	"github.com/pkg/errors"
	"github.com/shawnwyckoff/gpkg/dsa/stringz"
	"github.com/shawnwyckoff/gpkg/net/htmls"
	"github.com/shawnwyckoff/gpkg/net/httpz"
	"github.com/shawnwyckoff/gpkg/net/probe/xonline"
	"net"
	"strings"
	"time"
)

// get my wan IPs by 3rd party service
func GetWanIpOL(proxy string) (net.IP, error) {
	var ipstr string
	var exist bool
	var firstCheckWanOnline = true

	// best choices: plain text of ip string
	endpoints := []string{
		"http://whatismyip.akamai.com",
		"http://ident.me",
		"http://myip.dnsomatic.com",
		"http://icanhazip.com",
		"http://ifconfig.co/ip"}
	eps := stringz.Shuffle(endpoints)
	for _, url := range eps {
		resp, err := httpz.Get(url, proxy, time.Second*3, true)
		if err != nil {
			if firstCheckWanOnline {
				if !xonline.IsWanOnline(proxy) {
					return nil, errors.New("Can't get WAN ip because of internet offline ")
				}
				firstCheckWanOnline = false
			}
			continue
		}
		ipstr, _ = httpz.ReadBodyString(resp)
		resp.Body.Close()
		ipstr = strings.Trim(ipstr, "\r") // icanhazip.com 的返回结果会带换行符
		ipstr = strings.Trim(ipstr, "\n")
		t := CheckIPString(ipstr)
		if t == IPv4_WAN {
			return ParseIP(ipstr)
		}
	}

	// backup choices
	htmlString, err := httpz.GetString("http://bot.whatismyipaddress.com", proxy, time.Second*5)
	if err != nil {
		return nil, err
	}
	doc, err := htmls.NewDocFromHtmlSrc(&htmlString)
	if err == nil {
		ipstr = doc.Text()
		t := CheckIPString(ipstr)
		if t == IPv4_WAN {
			return ParseIP(ipstr)
		}
	}
	htmlString, err = httpz.GetString("http://network-tools.com", proxy, time.Second*5)
	if err != nil {
		return nil, err
	}
	doc, err = htmls.NewDocFromHtmlSrc(&htmlString)
	if err == nil {
		ipstr, exist = doc.Find("#field").First().Attr("value")
		if exist {
			t := CheckIPString(ipstr)
			if t == IPv4_WAN {
				return ParseIP(ipstr)
			}
		}
	}

	if !xonline.IsWanOnline(proxy) {
		return nil, errors.New("Can't get WAN ip because of internet offline ")
	} else {
		return nil, errors.New("Can't get WAN ip, unknown error")
	}
}
