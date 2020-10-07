package main

import (
	"os"
	"strconv"
	"fmt"
	"strings"
	"time"

	"github.com/brianvoe/gofakeit"
)

const (
	// ApacheCommonLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	ApacheCommonLog = "%s - %s [%s] \"%s %s %s\" %d %d"
	// ApacheCombinedLog : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes} "{referrer}" "{agent}"
	ApacheCombinedLog = "%s - %s [%s] \"%s %s %s\" %d %d \"%s\" \"%s\""
	// ApacheErrorLog : [{timestamp}]                      [{module}:{severity}]       [pid {pid}:tid {thread-id}] [client %{client}:{port}] %{message}
	//                  [Mon Oct 05 10:51:53.068955 2020] [lbmethod_heartbeat:notice] [pid 25891] AH02282: No slotmem from mod_heartmonitor
	//ApacheErrorLog = "[%s] [%s:%s] [pid %d:tid %d] [client %s:%d] %s"
	ApacheErrorLog = "[%s] [%s:%s] [pid %d:tid %d] %s"
	// RFC3164Log : <priority>{timestamp} {hostname} {application}[{pid}]: {message}
	RFC3164Log = "<%d>%s %s %s[%d]: %s"
	// RFC5424Log : <priority>{version} {iso-timestamp} {hostname} {application} {pid} {message-id} {structured-data} {message}
	RFC5424Log = "<%d>%d %s %s %s %d ID%d %s %s"
	// CommonLogFormat : {host} {user-identifier} {auth-user-id} [{datetime}] "{method} {request} {protocol}" {response-code} {bytes}
	CommonLogFormat = "%s - %s [%s] \"%s %s %s\" %d %d"
	// JSONLogFormat : {"host": "{host}", "user-identifier": "{user-identifier}", "datetime": "{datetime}", "method": "{method}", "request": "{request}", "protocol": "{protocol}", "status", {status}, "bytes": {bytes}, "referer": "{referer}"}
	JSONLogFormat = `{"host":"%s", "user-identifier":"%s", "datetime":"%s", "method": "%s", "request": "%s", "protocol":"%s", "status":%d, "bytes":%d, "referer": "%s"}`
)

// NOTE: bluemedora
func validHTTPStatus(s int) bool {
	allowedCodes := []int{200,201,202,203,204,205,206,300,301,303,304,401,402,403,404,405,500,501,502,503,504,505}
	statusLimit := os.Getenv("STATUS_LIMIT")
	if statusLimit != "" {
		allowedCodes = []int{}
		codes := strings.Split(statusLimit, ",")
		for _, code := range codes {
			x, err := strconv.Atoi(code)
			if err != nil {
				fmt.Println("status code in STATUS_LIMIT is bad: " + code)
				fmt.Println(err.Error())
				os.Exit(1)
			}
			allowedCodes = append(allowedCodes, x)
		}
	}

	for _, code := range allowedCodes {
		if s == code {
			return true
		}
	}
	return false
}

func validURI(uri string) bool {
	badURI := []string{"sexy"}
	for _, i := range badURI {
		if strings.Contains(uri, i) {
			return false
		}
	}
	return true
}

// NOTE: end_bluemedora

// NewApacheCommonLog creates a log string with apache common log format
func NewApacheCommonLog(t time.Time) string {
	var s int
	for {
		s = gofakeit.StatusCode()
		if validHTTPStatus(s) {
			break
		}
	}

	var uri string
	for {
		uri = RandResourceURI()
		if validURI(uri) {
			break
		}
	}

	var x string
	for {
		x = fmt.Sprintf(
			ApacheCommonLog,
			gofakeit.IPv4Address(),
			RandAuthUserID(),
			t.Format(Apache),
			gofakeit.HTTPMethod(),
			uri,
			RandHTTPVersion(),
			s,
			gofakeit.Number(0, 30000),
		)
		if validURI(x) {
			break
		}
	}
	return x
}

// NewApacheCombinedLog creates a log string with apache combined log format
func NewApacheCombinedLog(t time.Time) string {
	var s int
	for {
		s = gofakeit.StatusCode()
		if validHTTPStatus(s) {
			break
		}
	}

	var x string
	for {
		x = fmt.Sprintf(
			ApacheCombinedLog,
			gofakeit.IPv4Address(),
			RandAuthUserID(),
			t.Format(Apache),
			gofakeit.HTTPMethod(),
			RandResourceURI(),
			RandHTTPVersion(),
			s,
			gofakeit.Number(30, 100000),
			gofakeit.URL(),
			gofakeit.UserAgent(),
		)
		if validURI(x) {
			break
		}
	}
	return x
}

// NewApacheErrorLog creates a log string with apache error log format
func NewApacheErrorLog(t time.Time) string {
	/*x := fmt.Sprintf(
		ApacheErrorLog,
		t.Format(ApacheError),
		gofakeit.Word(),
		gofakeit.LogLevel("apache"),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 10000),
		gofakeit.IPv4Address(),
		gofakeit.Number(1, 65535),
		gofakeit.HackerPhrase(),
	)
	return x*/

	message, sev := RandomApacheErrorLog()

	return fmt.Sprintf(
		ApacheErrorLog,
		t.Format(ApacheError),
		//gofakeit.Word(),
		gofakeit.Word(),
		//gofakeit.LogLevel("apache"),
		sev,
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 10000),
		//gofakeit.IPv4Address(),
		//gofakeit.Number(1, 65535),
		message,
	)
}

// NewRFC3164Log creates a log string with syslog (RFC3164) format
func NewRFC3164Log(t time.Time) string {
	return fmt.Sprintf(
		RFC3164Log,
		gofakeit.Number(0, 191),
		t.Format(RFC3164),
		strings.ToLower(gofakeit.Username()),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.HackerPhrase(),
	)
}

// NewRFC5424Log creates a log string with syslog (RFC5424) format
func NewRFC5424Log(t time.Time) string {
	return fmt.Sprintf(
		RFC5424Log,
		gofakeit.Number(0, 191),
		gofakeit.Number(1, 3),
		t.Format(RFC5424),
		gofakeit.DomainName(),
		gofakeit.Word(),
		gofakeit.Number(1, 10000),
		gofakeit.Number(1, 1000),
		"-", // TODO: structured data
		gofakeit.HackerPhrase(),
	)
}

// NewCommonLogFormat creates a log string with common log format
func NewCommonLogFormat(t time.Time) string {
	return fmt.Sprintf(
		CommonLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
	)
}

// NewJSONLogFormat creates a log string with json log format
func NewJSONLogFormat(t time.Time) string {
	return fmt.Sprintf(
		JSONLogFormat,
		gofakeit.IPv4Address(),
		RandAuthUserID(),
		t.Format(CommonLog),
		gofakeit.HTTPMethod(),
		RandResourceURI(),
		RandHTTPVersion(),
		gofakeit.StatusCode(),
		gofakeit.Number(0, 30000),
		gofakeit.URL(),
	)
}
