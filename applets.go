package main

// Applet imports
import (
	"github.com/tiborvass/gobox/applets/cat"
	"github.com/tiborvass/gobox/applets/chroot"
	"github.com/tiborvass/gobox/applets/echo"
	"github.com/tiborvass/gobox/applets/grep"
	"github.com/tiborvass/gobox/applets/gzip"
	"github.com/tiborvass/gobox/applets/head"
	"github.com/tiborvass/gobox/applets/httpd"
	"github.com/tiborvass/gobox/applets/kill"
	"github.com/tiborvass/gobox/applets/ls"
	"github.com/tiborvass/gobox/applets/mkdir"
	"github.com/tiborvass/gobox/applets/mknod"
	"github.com/tiborvass/gobox/applets/mount"
	"github.com/tiborvass/gobox/applets/ps"
	"github.com/tiborvass/gobox/applets/rm"
	"github.com/tiborvass/gobox/applets/shell"
	"github.com/tiborvass/gobox/applets/telnetd"
	"github.com/tiborvass/gobox/applets/umount"
	"github.com/tiborvass/gobox/applets/wget"
)

// This map contains the mappings from callname
// to applet function.
var Applets map[string]Applet = map[string]Applet{
	"echo":    echo.Echo,
	"shell":   shell.Shell,
	"telnetd": telnetd.Telnetd,
	"ls":      ls.Ls,
	"rm":      rm.Rm,
	"httpd":   httpd.Httpd,
	"wget":    wget.Wget,
	"kill":    kill.Kill,
	"cat":     cat.Cat,
	"mknod":   mknod.Mknod,
	"mount":   mount.Mount,
	"umount":  umount.Umount,
	"chroot":  chroot.Chroot,
	"ps":      ps.Ps,
	"mkdir":   mkdir.Mkdir,
	"head":    head.Head,
	"grep":    grep.Grep,
	"gzip":    gzip.Gzip,
	"gunzip":  gzip.Gunzip,
	"zcat":    gzip.Zcat,
}

// Signature of applet functions.
// call is like os.Argv, and therefore contains the
// name of the applet itself in call[0].
// If the returned error is not nil, it is printed
// to stdout.
type Applet func(call []string) error
