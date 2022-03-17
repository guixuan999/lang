package main

import (
	"fmt"
	"io"
	"os"

	"github.com/unknwon/goconfig"
)

const IN_FILE = "ceph.conf"
const OUT_FILE = "ceph_updated.conf"

func main() {

	// func LoadFromReader(in io.Reader) (c *ConfigFile, err error)
	// func LoadFromData(data []byte) (c *ConfigFile, err error)
	// func LoadConfigFile(fileName string, moreFiles ...string) (c *ConfigFile, err error)
	c, _ := goconfig.LoadConfigFile(IN_FILE)

	sections := c.GetSectionList()

	for _, section := range sections {
		fmt.Println(section)
		keys := c.GetKeyList(section)
		for _, key := range keys {
			value, _ := c.GetValue(section, key)
			fmt.Println("    " + key + " <--------> " + value)
		}
	}

	// add a key "log to stderr = true" in section "global"
	c.SetValue("global", "log to stderr", "true")
	c.SetKeyComments("global", "log to stderr", "# this matters: without it, we can't get message like following from terminal(console window)\n# 2022-03-15T09:50:22.239中国标准时间 1  0 ceph-dokan: Mounted cephfs directory: /. Mountpoint: y")

	// add a section "mon", then add a key "mon_allow_pool_delete = true"
	// setValue will create a new section if it dose not exist.
	//c.SetValue("mon", "mon_allow_pool_delete", "true")

	// add a section "client", then add a key "keyring = C:/ProgramData/ceph/ceph.client.wsadm.keyring"
	c.SetValue("client", "keyring", "C:/ProgramData/ceph/ceph.client.wsadm.keyring")
	c.SetKeyComments("client", "keyring", "# this section is almost no need, cause we'll put client ID and it's keyring in command line\n# But put a existing keyring file as follow(whatever the content of keyring file's is will be ok, even if a blank file), the folliwng message in terminal(console window) will be refrained.\n# 2022-03-15T09:50:17.104中国标准时间 1 -1 auth: unable to find a keyring on /etc/ceph/ceph.client.fs_11.keyring,/etc/ceph/ceph.keyring,/etc/ceph/keyring,/etc/ceph/keyring.bin,: (2) No such file or directory")

	c.SetPrettyFormat(true)
	goconfig.SaveConfigFile(c, OUT_FILE)

	fmt.Println()
	fmt.Printf("CREATED %s ===>", OUT_FILE)
	fmt.Println()
	inf, _ := os.Open(OUT_FILE)
	defer inf.Close()
	io.Copy(os.Stdout, inf)

}
