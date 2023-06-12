package strings

import (
	"bytes"
	"cn.yasyx.container-license/pkg/utils/logger"
	"fmt"
	"net"
	"regexp"
	"sort"
	"strings"
	"unicode"
)

// In returns if the key is in the slice.
func In(key string, slice []string) bool {
	for _, s := range slice {
		if key == s {
			return true
		}
	}
	return false
}

func InList(key string, slice []string) bool {
	return In(key, slice)
}

func NotInIPList(key string, slice []string) bool {
	for _, s := range slice {
		if s == "" {
			continue
		}
		if key == strings.Split(s, ":")[0] {
			return false
		}
	}
	return true
}

func ReduceIPList(src, dst []string) []string {
	var ipList []string
	for _, ip := range src {
		if In(ip, dst) {
			ipList = append(ipList, ip)
		}
	}
	return ipList
}

func AppendIPList(src, dst []string) []string {
	for _, ip := range dst {
		if !In(ip, src) {
			src = append(src, ip)
		}
	}
	return src
}

func IPListRemove(ss []string, s string) (result []string) {
	for _, v := range ss {
		if v != s {
			result = append(result, v)
		}
	}
	return
}
func SortIPList(iplist []string) {
	realIPs := make([]net.IP, 0, len(iplist))
	for _, ip := range iplist {
		realIPs = append(realIPs, net.ParseIP(ip))
	}

	sort.Slice(realIPs, func(i, j int) bool {
		return bytes.Compare(realIPs[i], realIPs[j]) < 0
	})

	for i := range realIPs {
		iplist[i] = realIPs[i].String()
	}
}

func Reverse(s []string) []string {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}

func ContainList(list []string, toComplete string) (containerList []string) {
	for i := range list {
		if strings.Contains(list[i], toComplete) {
			containerList = append(containerList, list[i])
		}
	}
	return
}

func IsEmptyLine(str string) bool {
	re := regexp.MustCompile(`^\s*$`)

	return re.MatchString(str)
}

func TrimWS(str string) string {
	return strings.Trim(str, "\n\t")
}

func TrimSpaceWS(str string) string {
	return strings.TrimRight(str, " \n\t")
}

func RemoveSliceEmpty(list []string) (fList []string) {
	for i := range list {
		if strings.TrimSpace(list[i]) != "" {
			fList = append(fList, list[i])
		}
	}
	return
}

func SplitRemoveEmpty(s, sep string) []string {
	data := strings.Split(s, sep)
	return RemoveSliceEmpty(data)
}

// RemoveDuplicate removes duplicate entry in the list.
func RemoveDuplicate(list []string) []string {
	var result []string
	flagMap := map[string]struct{}{}
	for _, v := range list {
		if _, ok := flagMap[v]; !ok {
			flagMap[v] = struct{}{}
			result = append(result, v)
		}
	}
	return result
}

// RemoveStrSlice remove dst element from src slice
func RemoveStrSlice(src, dst []string) []string {
	var ipList []string
	for _, ip := range src {
		if !In(ip, dst) {
			ipList = append(ipList, ip)
		}
	}
	return ipList
}

func SliceRemoveStr(ss []string, s string) (result []string) {
	for _, v := range ss {
		if v != s {
			result = append(result, v)
		}
	}
	return
}

func FormatSize(size int64) (Size string) {
	if size < 1024 {
		Size = fmt.Sprintf("%.2fB", float64(size)/float64(1))
	} else if size < (1024 * 1024) {
		Size = fmt.Sprintf("%.2fKB", float64(size)/float64(1024))
	} else if size < (1024 * 1024 * 1024) {
		Size = fmt.Sprintf("%.2fMB", float64(size)/float64(1024*1024))
	} else {
		Size = fmt.Sprintf("%.2fGB", float64(size)/float64(1024*1024*1024))
	}
	return
}

func IsLetterOrNumber(k string) bool {
	for _, r := range k {
		if r == '_' {
			continue
		}
		if !unicode.IsLetter(r) && !unicode.IsNumber(r) {
			return false
		}
	}
	return true
}

func RenderShellFromEnv(shell string, envs map[string]string) string {
	var env string
	for k, v := range envs {
		env = fmt.Sprintf("%s%s=\"%s\" ", env, k, v)
	}
	if env == "" {
		return shell
	}
	return fmt.Sprintf("export %s; %s", env, shell)
}

func RenderTextFromEnv(text string, envs map[string]string) string {
	replaces := make(map[string]string, 0)
	for k, v := range envs {
		replaces[fmt.Sprintf("$(%s)", k)] = v
		replaces[fmt.Sprintf("${%s}", k)] = v
		replaces[fmt.Sprintf("$%s", k)] = v
	}
	logger.Debug("renderTextFromEnv: replaces: %+v ; text: %s", replaces, text)
	for o, n := range replaces {
		text = strings.ReplaceAll(text, o, n)
	}
	return text
}

func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if c := s[len(s)-1]; s[0] == c && (c == '"' || c == '\'') {
			return s[1 : len(s)-1]
		}
	}
	return s
}
