package core

import (
	"hsubdomain/gologger"
)

const Version = "0.7"
const banner = `
 _  __   _____       _         _                       _
| |/ /  / ____|     | |       | |                     (_)
| ' /  | (___  _   _| |__   __| | ___  _ __ ___   __ _ _ _ __
|  <    \___ \| | | | '_ \ / _| |/ _ \| '_   _ \ / _  | | '_ \
| . \   ____) | |_| | |_) | (_| | (_) | | | | | | (_| | | | | |
|_|\_\ |_____/ \__,_|_.__/ \__,_|\___/|_| |_| |_|\__,_|_|_| |_|

`

func ShowBanner() {
	return
	gologger.Printf(banner)
	gologger.Infof("Current Version: %s\n", Version)
}
