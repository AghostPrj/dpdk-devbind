/**
 * @Author: aghost<ggg17226@gmail.com>
 * @Date: 2022/4/16 17:22
 * @Desc:
 */

package kernelModuleCheckUtils

import (
	"golang.org/x/sys/unix"
	"os"
	"strings"
)

func CheckModuleLoaded(moduleName string) bool {
	sysfsPath := "/sys/module/"
	fileList, err := os.ReadDir(sysfsPath)
	if err != nil {
		panic("read /sys/module/ error " + err.Error())
	}

	for _, f := range fileList {
		if !f.IsDir() {
			continue
		}
		if f.Name() == moduleName {
			return true
		}
	}

	utsName := unix.Utsname{}
	unix.Uname(&utsName)
	var tmpReleaseByte []byte

	pos := 0
	for i := range utsName.Release {
		if utsName.Release[i] == 0 {
			pos = i
			break
		}
	}

	if pos == 0 {
		return false
	}
	tmpReleaseByte = make([]byte, pos)
	for i := 0; i < pos; i++ {
		tmpReleaseByte[i] = utsName.Release[i]
	}

	releaseString := string(tmpReleaseByte)

	kernelBuildInModulesListPath := "/lib/modules/" + releaseString + "/modules.builtin"

	fileContent, err := os.ReadFile(kernelBuildInModulesListPath)
	if err != nil {
		panic("read " + kernelBuildInModulesListPath + " error: " + err.Error())
	}

	lines := strings.Split(string(fileContent), "\n")

	for _, line := range lines {
		pathArr := strings.Split(line, "/")
		if len(pathArr) < 1 {
			continue
		}
		nameArr := strings.Split(pathArr[len(pathArr)-1], ".")
		if len(nameArr) < 1 {
			continue
		}
		if nameArr[0] == moduleName {
			return true
		}
	}

	return false
}
