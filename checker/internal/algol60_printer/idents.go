package algol60_printer

import (
	"regexp"
	"strings"
)

var isPredefinedName = map[string]bool{
	"exit": true, "outchar": true, "outinteger": true, "outstring": true, "integer2string": true,
	"openRW": true, "openRO": true, "openWOConfidential": true,
	"readchar": true, "readstring": true, "writechar": true, "writestring": true,
	"$top":   true,
	"offset": true, // fails in assembler
}

// names with seem like something about pwning
var pwnyNames = enrichNames([]string{
	"base", "baseOffset", "baseAddress",
	"binary", "binaryOffset", "binaryAddress",
	"libc", "libcOffset", "libcAddress",
	"address",
	"rop", "ropChain", "gadget", "gadgetAddress",
	"syscall", "syscallGadget",
	"oneGadget",
	"ret2libc", "ret2csu", "ret2dlresolve",
	"leak", "leakedAddress", "leakedOffset",
	"shellcode", "shellcodeOffset", "shellcodeAddress",
	"flag", "flagId", "flagPwd", "flagPassword",
	"overflow", "bufferOverflow",
})

func enrichNames(strs []string) []string {
	result := make([]string, 0, len(strs))
	for _, str := range strs {
		result = append(result, str)
		upper := regexp.MustCompile("[A-Z]").FindStringIndex(str)
		if upper == nil {
			continue
		}
		// replaces camelCase with space: camel case
		newStr := str[:upper[0]] + " " + strings.ToLower(str[upper[0]:upper[1]]) + str[upper[1]:]
		result = append(result, newStr)
	}
	return result
}
