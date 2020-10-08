package main

import (
	"fmt"
	"time"
	"math/rand"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit"
)

func RandomIP() string {
	ip := []string{
		"197.45.118.47",
		"10.128.82.233",
		"103.116.130.97",
		"103.63.241.210",
		"108.175.77.46",
		"109.225.174.38",
		"112.105.252.221",
		"113.135.247.216",
		"114.88.88.237",
		"124.22.78.149",
		"126.240.114.10",
		"127.135.118.63",
		"128.106.39.157",
		"129.53.82.110",
		"129.72.202.180",
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(ip)
	return fmt.Sprint(ip[n])
}

func CustomRandomURI() string {
	uri := []string{
		"/v1/login",
		"/v1/resources",
		"/v1/agents",
		"/v1/docs",
		"/v1/upload",
		"/status",
		"/about-us",
		"/help",
		"/home",
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(uri)
	return fmt.Sprint(uri[n])
}

// RandResourceURI generates a random resource URI
func RandResourceURI() string {
	var uri string
	num := gofakeit.Number(1, 4)
	for i := 0; i < num; i++ {
		uri += "/" + url.QueryEscape(gofakeit.BS())
	}
	uri = strings.ToLower(uri)
	return uri
}

// RandAuthUserID generates a random auth user id
func RandAuthUserID() string {
	candidates := []string{"-", strings.ToLower(gofakeit.Username())}
	return candidates[rand.Intn(2)]
}

// RandHTTPVersion returns a random http version
func RandHTTPVersion() string {
	versions := []string{"HTTP/1.0", "HTTP/1.1", "HTTP/2.0"}
	return versions[rand.Intn(3)]
}

// NOTE: bluemedora
func RandomApacheErrorLog() (string, string) {
	errorLogs := make(map[string]string)
	errorLogs["AH00170: caught SIGWINCH, shutting down gracefully"] = "notice"
	errorLogs["AH01232: suEXEC mechanism enabled (wrapper: /usr/sbin/suexec)"] = "notice"
	errorLogs["AH02282: No slotmem from mod_heartmonitor"] = "notice"
	errorLogs["AH00094: Command line: '/usr/sbin/httpd -D FOREGROUND'"] = "notice"
	errorLogs["AH00163: Apache/2.4.6 (CentOS) configured -- resuming normal operations"] = "notice"
	errorLogs["AH01276: Cannot serve directory /var/www/html/: No matching DirectoryIndex (index.html) found, and server-generated directory index forbidden by Options directive"] = "error"
	errorLogs["AH01626: authorization result of Require all granted: granted"] = "debug"
	errorLogs["AH01626: authorization result of <RequireAny>: granted"] = "debug"
	errorLogs["AH00175: File does not exist: /var/www/html/robots.txt"] = "error"
	errorLogs["AH00855: Connection refused: proxy: HTTP: attempt to connect to 127.0.0.1:8484 (localhost) failed"] = "error"
	errorLogs["AH01272: client denied by server configuration: /export/home/live/ap/htdocs/test"] = "error"

	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(errorLogs)

	i := 0
	for k,v := range errorLogs {
		if i == n {
			return k,v
		}
		i += 1
	}
	return "", ""
}
// NOTE: end_bluemedora
