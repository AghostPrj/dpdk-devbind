/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/4/16 16:29
 * @Desc:
 */

package main

import (
	"github.com/AghostPrj/dpdk-devbind/pkg/utils/systemCheckUtils"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	checkLspciResult := systemCheckUtils.CheckLsPci()
	if !checkLspciResult {
		panic("'lspci' not found - please install 'pciutils'")
	}

	app := &cli.App{
		Name:      "dpdk-devbind",
		Usage:     "Utility to bind and unbind devices from Linux kernel",
		UsageText: "dpdk-devbind [options] DEVICE1 DEVICE2 ....",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "status",
				Aliases: []string{"s"},
				Usage: "Print the current status of all known network, crypto, event\n        and mempool devices.\n" +
					"        For each device, it displays the PCI domain, bus, slot and function,\n" +
					"        along with a text description of the device. Depending upon whether the\n" +
					"        device is being used by a kernel driver, the igb_uio driver, or no\n" +
					"        driver, other relevant information will be displayed:\n" +
					"        * the Linux interface name e.g. if=eth0\n" +
					"        * the driver being used e.g. drv=igb_uio\n" +
					"        * any suitable drivers not currently using that device\n" +
					"            e.g. unused=igb_uio\n" +
					"        NOTE: if this flag is passed along with a bind/unbind option, the\n" +
					"        status display will always occur after the other operations have taken\n" +
					"        place.",
			},
			&cli.StringFlag{
				Name:    "bind",
				Aliases: []string{"b"},
				Usage:   "Select the driver to use or \"none\" to unbind the device",
			},
			&cli.StringFlag{
				Name:    "unbind",
				Aliases: []string{"u"},
				Usage:   "Unbind a device (Equivalent to \"-b none\")",
			}, &cli.BoolFlag{
				Name: "force",
				Usage: "By default, network devices which are used by Linux - as indicated by\n" +
					"        having routes in the routing table - cannot be modified. Using the\n" +
					"        --force flag overrides this behavior, allowing active links to be\n" +
					"        forcibly unbound.\n" +
					"        WARNING: This can lead to loss of network connection and should be used\n" +
					"        with caution.",
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		panic(err)
	}
}
