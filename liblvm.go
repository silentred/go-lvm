package lvm

//#cgo LDFLAGS: -llvm2app
//#include "macro_wrapper.h"
import "C"
import (
	"fmt"
	"unsafe"
)

var libh *C.struct_lvm

func init() {
	libh = C.lvm_init(nil)
}

// ######################## LVM methods #######################

// GetVersion returns library version
func GetVersion() *C.char {
	return C.lvm_library_get_version()
}

// GC cleans up libh
func GC() {
	C.lvm_quit(libh)
	libh = nil
}

//VgOpen opens volume group
func VgOpen(vgname string, mode string) C.vg_t {
	if mode == "" {
		mode = "r"
	}
	vg := C.lvm_vg_open(libh, C.CString(vgname), C.CString(mode), 0)
	return vg
}

//VgCreate creates VG
func VgCreate(vgname string) C.vg_t {
	return C.lvm_vg_create(libh, C.CString(vgname))
}

// ConfigFind checks if config could be found or not.
func ConfigFind(config string) (bool, error) {
	rval := C.lvm_config_find_bool(libh, C.CString(config), -10)
	if rval == -10 {
		return false, fmt.Errorf("config path not found")
	}
	if C.int(rval) == 0 {
		return false, nil
	}
	return true, nil
}

// ConfigReload reloads config.
func ConfigReload(config string) error {
	if C.lvm_config_reload(libh) == -1 {
		return getLastError()
	}
	return nil
}

// ConfigOverride overrides config.
func ConfigOverride(config string) error {
	if C.lvm_config_override(libh, C.CString(config)) == -1 {
		return getLastError()
	}
	return nil
}

// Scan scans libh
func Scan() {
	C.lvm_scan(libh)
}

// ListVgNames returns LVM name list
func ListVgNames() []string {
	vgnames := C.lvm_list_vg_names(libh)
	if vgnames != nil {
		// TODO
		fmt.Printf("nil\n")
	}
	cargs := C.makeCharArray(C.int(0))
	n := C.wrapper_dm_list_iterate_items(vgnames, cargs)
	gs := goStrings(n, cargs)
	return gs
}

// ListVgUUIDs returns LVM uuid list
func ListVgUUIDs() []string {
	uuids := C.lvm_list_vg_uuids(libh)
	if uuids != nil {
		// TODO
		fmt.Printf("nil\n")
	}
	cargs := C.makeCharArray(C.int(0))
	n := C.wrapper_dm_list_iterate_items(uuids, cargs)
	gs := goStrings(n, cargs)
	return gs
}

// PercentToFloat converts percent to float.
func PercentToFloat(percent C.percent_t) C.float {
	// TODO C.percent_t should be golang type.
	return C.lvm_percent_to_float(percent)
}

// VgNameValidate validates if the vg name is valid.
func VgNameValidate(name string) error {
	ret := C.lvm_vg_name_validate(libh, C.CString(name))
	if ret < 0 {
		return getLastError()
	}
	return nil
}

// VgNameFromPvID returns VG name from PVID.
func VgNameFromPvID(pvid string) *C.char {
	ret := C.lvm_vgname_from_pvid(libh, C.CString(pvid))
	// TODO
	msg := C.lvm_errmsg(libh)
	fmt.Printf("msg : %#v\n", C.GoString(msg))
	return ret
}

// VgNameFromPvDevice returns VG name from PV device.
func VgNameFromPvDevice(pvdevice string) string {
	ret := C.lvm_vgname_from_device(libh, C.CString(pvdevice))
	if ret == nil {
		// TODO
		//		return getLastError()
	}
	return C.GoString(ret)
}

// ######################## vg methods #######################

// VgObject is an object of VG.
type VgObject struct {
	Vgt C.vg_t
}

// GetName gets name of VG.
func (v *VgObject) GetName() string {
	return C.GoString(C.lvm_vg_get_name(v.Vgt))
}

// GetUuid gets UUID of VG.
func (v *VgObject) GetUuid() *C.char {
	return C.lvm_vg_get_uuid(v.Vgt)
}

// Close closes VG object.
func (v *VgObject) Close() error {
	if C.lvm_vg_close(v.Vgt) == -1 {
		return getLastError()
	}
	return nil
}

// Remove removes VG.
func (v *VgObject) Remove() error {
	if C.lvm_vg_remove(v.Vgt) == -1 {
		return getLastError()
	}
	if C.lvm_vg_write(v.Vgt) == -1 {
		return getLastError()
	}
	return v.Close()
}

// Extend extends PV by adding vg.
func (v *VgObject) Extend(device string) error {
	if C.lvm_vg_extend(v.Vgt, C.CString(device)) == -1 {
		return getLastError()
	}
	if C.lvm_vg_write(v.Vgt) == -1 {
		return getLastError()
	}
	return nil
}

// Reduce reduces VG from PV.
func (v *VgObject) Reduce(device string) error {
	if C.lvm_vg_reduce(v.Vgt, C.CString(device)) == -1 {
		return getLastError()
	}
	if C.lvm_vg_write(v.Vgt) == -1 {
		return getLastError()
	}
	return nil
}

// AddTag adds tag to VG.
func (v *VgObject) AddTag(stag string) error {
	tag := C.CString(stag)
	if C.lvm_vg_add_tag(v.Vgt, tag) == -1 {
		return getLastError()
	}
	if C.lvm_vg_write(v.Vgt) == -1 {
		return getLastError()
	}
	return nil
}

// RemoveTag removes tag from VG.
func (v *VgObject) RemoveTag(stag string) error {
	tag := C.CString(stag)
	if C.lvm_vg_remove_tag(v.Vgt, tag) == -1 {
		return getLastError()
	}
	if C.lvm_vg_write(v.Vgt) == -1 {
		return getLastError()
	}
	return nil
}

// SetExtentSize sets extent size.
func (v *VgObject) SetExtentSize(size uint32) error {
	if C.lvm_vg_set_extent_size(v.Vgt, C.uint32_t(size)) == 1 {
		return getLastError()
	}
	return nil
}

// IsClustered checks clustered or not.
func (v *VgObject) IsClustered() bool {
	if C.lvm_vg_is_clustered(v.Vgt) == 1 {
		return true
	}
	return false
}

// IsExported checks exported or not.
func (v *VgObject) IsExported() bool {
	if C.lvm_vg_is_exported(v.Vgt) == 1 {
		return true
	}
	return false
}

// IsPartial checks partial or not.
func (v *VgObject) IsPartial() bool {
	if C.lvm_vg_is_partial(v.Vgt) == 1 {
		return true
	}
	return false
}

// GetSeqno returns sequence number of VG.
func (v *VgObject) GetSeqno() C.uint64_t {
	return C.lvm_vg_get_seqno(v.Vgt)
}

// GetSize returns size of VG
func (v *VgObject) GetSize() C.uint64_t {
	return C.lvm_vg_get_size(v.Vgt)
}

// GetFreeSize returns free size of VG
func (v *VgObject) GetFreeSize() C.uint64_t {
	return C.lvm_vg_get_free_size(v.Vgt)
}

// GetExtentSize returns extent size of VG.
func (v *VgObject) GetExtentSize() C.uint64_t {
	return C.lvm_vg_get_extent_size(v.Vgt)
}

// GetExtentCount returns extent count of VG.
func (v *VgObject) GetExtentCount() C.uint64_t {
	return C.lvm_vg_get_extent_count(v.Vgt)
}

// GetFreeExtentCount returns free extent count of VG.
func (v *VgObject) GetFreeExtentCount() C.uint64_t {
	return C.lvm_vg_get_free_extent_count(v.Vgt)
}

// GetProperty returns properties of VG.
func (v *VgObject) GetProperty(name string) (properties, error) {
	prop_value := C.lvm_vg_get_property(v.Vgt, C.CString(name))
	return get_property(&prop_value)
}

//        { "setProperty",        (PyCFunction)_liblvm_lvm_vg_set_property, METH_VARARGS },

// GePvCount returns the number of PV.
func (v *VgObject) GetPvCont() C.uint64_t {
	return C.lvm_vg_get_pv_count(v.Vgt)
}

// GetMaxPv returns maximum value of PV.
func (v *VgObject) GetMaxPv() C.uint64_t {
	return C.lvm_vg_get_max_pv(v.Vgt)
}

// GetMaxPv returns maximum value of LV.
func (v *VgObject) GetMaxLV() C.uint64_t {
	return C.lvm_vg_get_max_lv(v.Vgt)
}

// ListLVs lists of lvs from VG
func (v *VgObject) ListLVs() []string {
	lvl := C.lvm_vg_list_lvs(v.Vgt)
	if lvl == nil {
		return []string{""}
	}
	cargs := C.makeCharArray(C.int(0))
	n := C.wrapper_dm_list_iterate_items(lvl, cargs)
	gs := goStrings(n, cargs)
	fmt.Printf("(test)lvList: %#v\n", gs)
	return gs
}

// ListPVs lists of pvs from VG
func (v *VgObject) ListPVs() []string {
	pvs := C.lvm_vg_list_pvs(v.Vgt)
	if pvs == nil {
		return []string{""}
	}
	cargs := C.makeCharArray(C.int(0))
	n := C.wrapper_dm_list_iterate_items(pvs, cargs)
	gs := goStrings(n, cargs)
	fmt.Printf("(test)pvsList: %#v\n", gs)
	return gs
}

// pv_from_N returns PV.
func lv_from_N(vg *C.struct_volume_group, id *C.char, pvg *VgObject, f func(*C.struct_volume_group, *C.char) C.lv_t) (*LvObject, error) {
	lv := f(vg, id)
	if lv == nil {
		return nil, getLastError()
	}
	return &LvObject{
		Lvt:      lv,
		parentVG: pvg,
	}, nil
}

// TODO: test
// LvFromName returns LV object from name of VG.
func (v *VgObject) LvFromName(sname string) (*LvObject, error) {
	name := C.CString(sname)
	return lv_from_N(v.Vgt, name, v, func(vg *C.struct_volume_group, id *C.char) C.lv_t {
		return C.lvm_lv_from_name(vg, name)
	})
}

// LvFromUuid returns LV object from UUID of VG.
func (v *VgObject) LvFromUuid(suuid string) (*LvObject, error) {
	uuid := C.CString(suuid)
	return lv_from_N(v.Vgt, uuid, v, func(vg *C.struct_volume_group, id *C.char) C.lv_t {
		return C.lvm_lv_from_uuid(vg, uuid)
	})
}

// LvNameValidate validates if the lv name is valid inside VG.
func (v *VgObject) LvNameValidate(name string) error {
	ret := C.lvm_lv_name_validate(v.Vgt, C.CString(name))
	if ret < 0 {
		return getLastError()
	}
	return nil
}

// pv_from_N returns PV.
func pv_from_N(vg *C.struct_volume_group, id *C.char, f func(*C.struct_volume_group, *C.char) C.pv_t) (*PvObject, error) {
	pv := f(vg, id)
	if pv == nil {
		return nil, getLastError()
	}
	return &PvObject{pv}, nil
}

// GetTags returns tag list of LV.
func (v *VgObject) GetTags() []string {
	tagsl := C.lvm_vg_get_tags(v.Vgt)
	// TODO: should check error?
	if tagsl == nil {
		return []string{""}
	}
	// TODO?:  dm_list_size(vgnames?)
	cargs := C.makeCharArray(C.int(0))
	n := C.wrapper_dm_list_iterate_items(tagsl, cargs)
	gs := goStrings(n, cargs)
	return gs
}

// createGoLv creats a LV Object
func createGoLv(v *VgObject, lv C.lv_t) *LvObject {
	return &LvObject{
		Lvt:      lv,
		parentVG: v,
	}
}

// CreateLvLinear creates LV Object. s is the size of logical volume in bytes, size should be round up by 4MiB
func (v *VgObject) CreateLvLinear(n string, s int64) (*LvObject, error) {
	size := C.uint64_t(s)
	name := C.CString(n)

	lv := C.lvm_vg_create_lv_linear(v.Vgt, name, size)
	if lv == nil {
		return nil, getLastError()
	}
	return createGoLv(v, lv), nil
}

//        { "createLvThinpool",   (PyCFunction)_liblvm_lvm_vg_create_lv_thinpool, METH_VARARGS },
//        { "createLvThin",       (PyCFunction)_liblvm_lvm_vg_create_lv_thin, METH_VARARGS },

// properties represents variety of properties.
type properties struct {
	signed_integer int
	integer        int
	str            string
}

// get_property returns properties.
func get_property(prop *C.struct_lvm_property_value) (properties, error) {
	var p properties
	if C.is_valid(unsafe.Pointer(prop)) == 0 {
		return p, getLastError()
	}

	if C.is_integer(unsafe.Pointer(prop)) != 0 {
		if C.is_signed(unsafe.Pointer(prop)) != 0 {
			// TODO
		} else {
			// TODO
		}
	} else {
		// TODO
	}
	return p, nil
}

//        { "listPVsegs",         (PyCFunction)_liblvm_lvm_pv_list_pvsegs, METH_NOARGS },

// ######################## utility methods #######################
func goStrings(argc C.int, argv **C.char) []string {
	// TODO nConstraint
	length := int(argc)
	tmpslice := (*[1 << 30]*C.char)(unsafe.Pointer(argv))[:length:length]
	gostrings := make([]string, length)
	for i, s := range tmpslice {
		gostrings[i] = C.GoString(s)
	}
	return gostrings
}

func getLastError() error {
	msg := C.GoString(C.lvm_errmsg(libh))
	if msg == "" {
		return fmt.Errorf("unknown error")
	}
	return fmt.Errorf(msg)
}
