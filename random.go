package main

import (
	"fmt"
	"time"
	"math/rand"
	"net/url"
	"strings"

	"github.com/brianvoe/gofakeit"
)

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
	/**errorLogs := []string{
		"[Mon Oct 05 10:51:04.713723 2020] [mpm_prefork:notice] [pid 942] AH00170: caught SIGWINCH, shutting down gracefully",
		"[Mon Oct 05 10:51:05.760946 2020] [suexec:notice] [pid 25862] AH01232: suEXEC mechanism enabled (wrapper: /usr/sbin/suexec)",
		"[Mon Oct 05 10:51:05.770094 2020] [lbmethod_heartbeat:notice] [pid 25862] AH02282: No slotmem from mod_heartmonitor",
		"[Mon Oct 05 10:51:05.772601 2020] [mpm_prefork:notice] [pid 25862] AH00163: Apache/2.4.6 (CentOS) configured -- resuming normal operations",
		"[Mon Oct 05 10:51:05.772620 2020] [core:notice] [pid 25862] AH00094: Command line: '/usr/sbin/httpd -D FOREGROUND'",
		"[Mon Oct 05 10:51:14.147557 2020] [mpm_prefork:notice] [pid 25862] AH00170: caught SIGWINCH, shutting down gracefully",
		"[Mon Oct 05 10:51:53.059422 2020] [suexec:notice] [pid 25891] AH01232: suEXEC mechanism enabled (wrapper: /usr/sbin/suexec)",
		"[Mon Oct 05 10:51:53.068955 2020] [lbmethod_heartbeat:notice] [pid 25891] AH02282: No slotmem from mod_heartmonitor",
		"[Mon Oct 05 10:51:53.072773 2020] [mpm_prefork:notice] [pid 25891] AH00163: Apache/2.4.6 (CentOS) configured -- resuming normal operations",
		"[Mon Oct 05 10:51:53.072792 2020] [core:notice] [pid 25891] AH00094: Command line: '/usr/sbin/httpd -D FOREGROUND'",
	}
	**/
	errorLogs := []string{
		"AH00170: caught SIGWINCH, shutting down gracefully",
		"AH01232: suEXEC mechanism enabled (wrapper: /usr/sbin/suexec)",
		"AH02282: No slotmem from mod_heartmonitor",
		"AH00094: Command line: '/usr/sbin/httpd -D FOREGROUND'",
		"AH00163: Apache/2.4.6 (CentOS) configured -- resuming normal operations",
	}
	rand.Seed(time.Now().Unix())
	n := rand.Int() % len(errorLogs)
	return fmt.Sprint(errorLogs[n]), "notice"
}
// NOTE: end_bluemedora
