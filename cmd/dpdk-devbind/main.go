/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/4/16 16:29
 * @Desc:
 */

package main

import (
	"github.com/AghostPrj/dpdk-devbind/pkg/cli"
	"github.com/AghostPrj/dpdk-devbind/pkg/globalData"
	"github.com/AghostPrj/dpdk-devbind/pkg/utils/kernelModuleCheckUtils"
	"github.com/AghostPrj/dpdk-devbind/pkg/utils/systemCheckUtils"
	"os"
)

func main() {
	checkLspciResult := systemCheckUtils.CheckLsPci()
	if !checkLspciResult {
		panic("'lspci' not found - please install 'pciutils'")
	}

	for _, name := range globalData.DpdkDrivers {
		if kernelModuleCheckUtils.CheckModuleLoaded(name) {
			globalData.LoadedModules = append(globalData.LoadedModules, name)
		}
	}

	err := cli.BuildCLiApp().Run(os.Args)
	if err != nil {
		panic(err)
	}
}
