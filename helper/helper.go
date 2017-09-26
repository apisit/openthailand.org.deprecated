package helper

import (
	"crypto/md5"
	"fmt"
	"io"
	"net/url"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SliceContains(slice []string, item string) bool {
	set := make(map[string]struct{}, len(slice))
	for _, s := range slice {
		set[s] = struct{}{}
	}

	_, ok := set[item]
	return ok
}

func GenerateImageFileName(value string) string {
	time := strconv.FormatInt(time.Now().Unix(), 10)
	str := fmt.Sprintf("%s%s", value, time)
	h := md5.New()
	io.WriteString(h, str)
	return fmt.Sprintf("%x", h.Sum(nil))
}

func PrettyUrl(name string) string {
	reg, err := regexp.Compile(`[^a-zA-Z0-9]`) //find everything that is not alphanumeric
	if err != nil {
		return ""
	}
	prettyurl := reg.ReplaceAllString(name, "-") //replace it with -
	prettyurl = strings.Trim(prettyurl, "-")
	if len(prettyurl) == 0 { //incase it's not english we simply escape the string
		return url.QueryEscape(name)
	}
	return strings.ToLower(prettyurl)
}
