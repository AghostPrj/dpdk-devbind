/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/4/16 16:33
 * @Desc:
 */

package systemCheckUtils

import "os/exec"

func CheckLsPci() bool {
	proc := exec.Command("which", "lspci")
	err := proc.Run()
	return err == nil
}
