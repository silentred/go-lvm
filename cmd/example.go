package main

import (
	"fmt"

	"github.com/silentred/go-lvm"
)

func main() {
	// List volume group
	vglist := lvm.ListVgNames()
	availableVG := ""
	// Create a VG object
	vgo := &lvm.VgObject{}
	for i := 0; i < len(vglist); i++ {
		vgo.Vgt = lvm.VgOpen(vglist[i], "r")
		if vgo.GetFreeSize() > 0 {
			availableVG = vglist[i]
			vgo.Close()
			break
		}
		vgo.Close()
	}
	if availableVG == "" {
		fmt.Printf("no VG that has free space found\n")
		return
	}

	// Open VG in write mode
	vgo.Vgt = lvm.VgOpen(availableVG, "w")
	defer vgo.Close()

	// Output some data of the VG
	fmt.Printf("size: %d GiB\n", uint64(vgo.GetSize())/1024/1024/1024)
	fmt.Printf("pvlist: %v\n", vgo.ListPVs())
	fmt.Printf("Free size: %d MiB\n", uint64(vgo.GetFreeSize())/1024/1024)

	// Create a LV object
	l := &lvm.LvObject{}

	// Create a LV
	l, err := vgo.CreateLvLinear("go-lvm-example-test-lv", int64(vgo.GetFreeSize())/1024)
	if err != nil {
		fmt.Printf("error: %v")
		return
	}

	// Output uuid of LV
	fmt.Printf("Created\n\tuuid: %s\n\tname: %s\n\tattr: %s\n\torigin: %s\n\tsize: %d",
		l.GetUuid(), l.GetName(), l.GetAttr(), l.GetOrigin(), uint64(l.GetSize()))

	// Output uuid of LV
	l.Remove()
}
