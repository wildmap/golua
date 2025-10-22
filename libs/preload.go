package libs

import (
	lua "github.com/yuin/gopher-lua"

	"github.com/wildmap/golua/libs/base64"
	"github.com/wildmap/golua/libs/bit"
	"github.com/wildmap/golua/libs/cmd"
	"github.com/wildmap/golua/libs/crypto"
	"github.com/wildmap/golua/libs/filepath"
	"github.com/wildmap/golua/libs/hex"
	"github.com/wildmap/golua/libs/humanize"
	"github.com/wildmap/golua/libs/inspect"
	"github.com/wildmap/golua/libs/ioutil"
	"github.com/wildmap/golua/libs/json"
	"github.com/wildmap/golua/libs/log"
	"github.com/wildmap/golua/libs/regexp"
	"github.com/wildmap/golua/libs/runtime"
	"github.com/wildmap/golua/libs/shellescape"
	"github.com/wildmap/golua/libs/ssh"
	"github.com/wildmap/golua/libs/strings"
	"github.com/wildmap/golua/libs/tac"
	"github.com/wildmap/golua/libs/time"
	"github.com/wildmap/golua/libs/xmlpath"
	"github.com/wildmap/golua/libs/yaml"
)

// Preload preload all gopher lua packages
func Preload(L *lua.LState) {
	base64.Preload(L)
	bit.Preload(L)
	cmd.Preload(L)
	crypto.Preload(L)
	filepath.Preload(L)
	hex.Preload(L)
	humanize.Preload(L)
	inspect.Preload(L)
	ioutil.Preload(L)
	json.Preload(L)
	log.Preload(L)
	regexp.Preload(L)
	runtime.Preload(L)
	shellescape.Preload(L)
	ssh.Preload(L)
	strings.Preload(L)
	tac.Preload(L)
	time.Preload(L)
	xmlpath.Preload(L)
	yaml.Preload(L)
}
