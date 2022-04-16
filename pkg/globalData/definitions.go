/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/4/16 17:14
 * @Desc:
 */

package globalData

var (
	DpdkDrivers   = [...]string{"igb_uio", "vfio-pci", "uio_pci_generic"}
	LoadedModules = make([]string, 0)
)
