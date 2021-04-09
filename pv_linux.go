package lvm

//#cgo LDFLAGS: -llvm2app
//#include "cfiles/macro_wrapper.h"
import "C"
import "fmt"

//PvCreate creates PV
func PvCreate(pv_name string, size uint64, pvmetadatacopies uint64, pvmetadatasize uint64,
	data_alignment uint64, data_alignment_offset uint64, zero uint64) C.vg_t {
	pv_params := C.lvm_pv_params_create(libh, C.CString(pv_name))
	if pv_params != nil {
		// TODO
	}

	// TODO return error
	C.wrapper_set_pv_prop(pv_params, C.CString("size"), C.longlong(size))
	C.wrapper_set_pv_prop(pv_params, C.CString("size"), C.longlong(size))
	C.wrapper_set_pv_prop(pv_params, C.CString("pvmetadatacopies"), C.longlong(pvmetadatacopies))
	C.wrapper_set_pv_prop(pv_params, C.CString("pvmetadatasize"), C.longlong(pvmetadatasize))
	C.wrapper_set_pv_prop(pv_params, C.CString("data_alignment"), C.longlong(data_alignment))
	C.wrapper_set_pv_prop(pv_params, C.CString("data_alignment_offset"), C.longlong(data_alignment_offset))
	C.wrapper_set_pv_prop(pv_params, C.CString("zero"), C.longlong(zero))

	C.lvm_pv_create_adv(pv_params)
	// TODO
	return nil
}

// PvRemove removes PV.
func PvRemove(pv_name string) {
	C.lvm_pv_remove(libh, C.CString(pv_name))
}

// TODO: test
// PvFromName returns PV object from VGName
func (v *VgObject) PvFromName(sname string) (*PvObject, error) {
	name := C.CString(sname)
	return pv_from_N(v.Vgt, name, func(vg *C.struct_volume_group, id *C.char) C.pv_t {
		return C.lvm_pv_from_name(vg, name)
	})
}

// PvFromUuid returns PV object from uuid.
func (v *VgObject) PvFromUuid(sid string) (*PvObject, error) {
	id := C.CString(sid)
	return pv_from_N(v.Vgt, id, func(vg *C.struct_volume_group, id *C.char) C.pv_t {
		return C.lvm_pv_from_uuid(vg, id)
	})
}

// ######################## pv list methods #######################

// Open lists PVs and get them as string array.
func Open() []string {
	pvsList := C.lvm_list_pvs(libh)
	fmt.Printf("pvsList: %#v\n", pvsList)

	cargs := C.makeCharArray(C.int(0))
	n := C.wrapper_dm_list_iterate_items(pvsList, cargs)
	gs := goStrings(n, cargs)
	return gs
}

// Close frees list of PVs.
func Close(pvs []string) {
	// TODO
	if len(pvs) > 0 {
		//C.lvm_list_pvs_free(pvs)
	}
}

// ######################## pv methods #######################

// pvObject represents PV.
type PvObject struct {
	Pvt C.pv_t
}

// GetName returns name of the PV.
func (p *PvObject) GetName() string {
	return C.GoString(C.lvm_pv_get_name(p.Pvt))
}

// GetUuid returns UUID of the PV.
func (p *PvObject) GetUuid() string {
	return C.GoString(C.lvm_pv_get_uuid(p.Pvt))
}

// GetMdaCount returns metadata count.
func (p *PvObject) GetMdaCount() C.uint64_t {
	return C.lvm_pv_get_mda_count(p.Pvt)
}

// GetProperty returns properties of PV.
func (p *PvObject) GetProperty(name string) (properties, error) {
	prop_value := C.lvm_pv_get_property(p.Pvt, C.CString(name))
	return get_property(&prop_value)
}

// GetSize returns size of PV.
func (p *PvObject) GetSize() C.uint64_t {
	return C.lvm_pv_get_size(p.Pvt)
}

// GetDevSize returns free size of PV.
func (p *PvObject) GetDevSize() C.uint64_t {
	return C.lvm_pv_get_size(p.Pvt)
}

// GetFreeSize returns free size of PV.
func (p *PvObject) GetFreeSize() C.uint64_t {
	return C.lvm_pv_get_free(p.Pvt)
}

// Resize resizes the size of PV.
func (p *PvObject) Resize(size uint64) error {
	if C.lvm_pv_resize(p.Pvt, C.uint64_t(size)) == -1 {
		return getLastError()
	}
	return nil
}
