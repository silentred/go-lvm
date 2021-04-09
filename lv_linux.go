package lvm

//#cgo LDFLAGS: -llvm2app
//#include "cfiles/macro_wrapper.c"
import "C"

// ######################################## LV methods ###################################

// LvObject represents LV.
type LvObject struct {
	Lvt      C.lv_t
	parentVG *VgObject
}

// GetAttr gets LV attr
func (l *LvObject) GetAttr() string {
	return C.GoString(C.lvm_lv_get_attr(l.Lvt))
}

// GetName gets LV name
func (l *LvObject) GetName() string {
	return C.GoString(C.lvm_lv_get_name(l.Lvt))
}

// GetOrigin gets LV origin
func (l *LvObject) GetOrigin() string {
	return C.GoString(C.lvm_lv_get_origin(l.Lvt))
}

// GetUuid gets LV uuid
func (l *LvObject) GetUuid() string {
	return C.GoString(C.lvm_lv_get_uuid(l.Lvt))
}

// Activate activates LV.
func (l *LvObject) Activate() error {
	if C.lvm_lv_activate(l.Lvt) == -1 {
		return getLastError()
	}
	return nil
}

// Deactivate deactivates LV.
func (l *LvObject) Deactivate() error {
	if C.lvm_lv_deactivate(l.Lvt) == -1 {
		return getLastError()
	}
	return nil
}

// Remove removes LV.
func (l *LvObject) Remove() error {
	if C.lvm_vg_remove_lv(l.Lvt) == -1 {
		return getLastError()
	}
	return nil
}

// GetProperty returns properties of LV.
func (l *LvObject) GetProperty(name string) (properties, error) {
	prop_value := C.lvm_lv_get_property(l.Lvt, C.CString(name))
	return get_property(&prop_value)
}

// GetSize returns size of LV.
func (l *LvObject) GetSize() C.uint64_t {
	return C.lvm_lv_get_size(l.Lvt)
}

// IsActive checks active LV or not.
func (l *LvObject) IsActive() bool {
	if C.lvm_lv_is_active(l.Lvt) == 1 {
		return true
	}
	return false
}

// IsSuspended checks suspended LV or not.
func (l *LvObject) IsSuspended() bool {
	if C.lvm_lv_is_suspended(l.Lvt) == 1 {
		return true
	}
	return false
}

// AddTag adds tag to LV.
func (l *LvObject) AddTag(stag string) error {
	tag := C.CString(stag)
	if C.lvm_lv_add_tag(l.Lvt, tag) == -1 {
		return getLastError()
	}
	if C.lvm_vg_write(l.parentVG.Vgt) == -1 {
		return getLastError()
	}
	return nil
}

// RemoveTag removes tag from LV.
func (l *LvObject) RemoveTag(stag string) error {
	tag := C.CString(stag)
	if C.lvm_lv_remove_tag(l.Lvt, tag) == -1 {
		return getLastError()
	}
	if C.lvm_vg_write(l.parentVG.Vgt) == -1 {
		return getLastError()
	}
	return nil
}

// GetTags returns tag list of LV.
func (l *LvObject) GetTags() []string {
	tagsl := C.lvm_lv_get_tags(l.Lvt)
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

// Rename rename the name of LV.
func (l *LvObject) Rename(name string) error {
	if C.lvm_lv_rename(l.Lvt, C.CString(name)) == -1 {
		return getLastError()
	}
	return nil
}

// Resize resizes the size of LV.
func (l *LvObject) Resize(size uint64) error {
	if C.lvm_lv_resize(l.Lvt, C.uint64_t(size)) == -1 {
		return getLastError()
	}
	return nil
}

//        { "listLVsegs",         (PyCFunction)_liblvm_lvm_lv_list_lvsegs, METH_NOARGS },

// Snapshot creates a LV snapshot.
func (l *LvObject) Snapshot(snapname string, size uint64) (*LvObject, error) {

	lvp := C.lvm_lv_params_create_snapshot(l.Lvt, C.CString(snapname), C.uint64_t(size))
	if lvp == nil {
		return nil, getLastError()
	}
	lv := C.lvm_lv_create(lvp)
	if lv == nil {
		return nil, getLastError()
	}
	return createGoLv(l.parentVG, lv), nil
}
