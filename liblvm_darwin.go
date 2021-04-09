package lvm

type VgObject struct {
}

type LvObject struct {
}

// ListVgNames returns LVM name list
func ListVgNames() []string {
	return nil
}

func VgOpen(vgname string, mode string) *VgObject {
	vg := &VgObject{}
	return vg
}

// CreateLvLinear creates LV Object. s is the size of logical volume in bytes, size should be round up by 4MiB
func (v *VgObject) CreateLvLinear(n string, s int64) (*LvObject, error) {
	return nil, nil
}

// GetName gets name of VG.
func (v *VgObject) GetName() string {
	return ""
}

// GetUuid gets UUID of VG.
func (v *VgObject) GetUuid() string {
	return ""
}

// Close closes VG object.
func (v *VgObject) Close() error {
	return nil
}

// Remove removes VG.
func (v *VgObject) Remove() error {
	return nil
}

// Extend extends PV by adding vg.
func (v *VgObject) Extend(device string) error {
	return nil
}

// Reduce reduces VG from PV.
func (v *VgObject) Reduce(device string) error {
	return nil
}

// AddTag adds tag to VG.
func (v *VgObject) AddTag(stag string) error {
	return nil
}

// RemoveTag removes tag from VG.
func (v *VgObject) RemoveTag(stag string) error {
	return nil
}

// SetExtentSize sets extent size.
func (v *VgObject) SetExtentSize(size uint32) error {
	return nil
}

// IsClustered checks clustered or not.
func (v *VgObject) IsClustered() bool {
	return false
}

// IsExported checks exported or not.
func (v *VgObject) IsExported() bool {
	return false
}

// IsPartial checks partial or not.
func (v *VgObject) IsPartial() bool {
	return false
}

// GetSeqno returns sequence number of VG.
func (v *VgObject) GetSeqno() int {
	return 0
}

// GetSize returns size of VG
func (v *VgObject) GetSize() int {
	return 0
}

// GetFreeSize returns free size of VG
func (v *VgObject) GetFreeSize() int {
	return 0
}

// GetExtentSize returns extent size of VG.
func (v *VgObject) GetExtentSize() int {
	return 0
}

// GetExtentCount returns extent count of VG.
func (v *VgObject) GetExtentCount() int {
	return 0
}

// GetFreeExtentCount returns free extent count of VG.
func (v *VgObject) GetFreeExtentCount() int {
	return 0
}

// GePvCount returns the number of PV.
func (v *VgObject) GetPvCont() int {
	return 0
}

// GetMaxPv returns maximum value of PV.
func (v *VgObject) GetMaxPv() int {
	return 0
}

// GetMaxPv returns maximum value of LV.
func (v *VgObject) GetMaxLV() int {
	return 0
}

// ListLVs lists of lvs from VG
func (v *VgObject) ListLVs() []string {
	return nil
}

// ListPVs lists of pvs from VG
func (v *VgObject) ListPVs() []string {
	return nil
}

// LvFromName returns LV object from name of VG.
func (v *VgObject) LvFromName(sname string) (*LvObject, error) {
	return nil, nil
}

// LvFromUuid returns LV object from UUID of VG.
func (v *VgObject) LvFromUuid(suuid string) (*LvObject, error) {
	return nil, nil
}

// LvNameValidate validates if the lv name is valid inside VG.
func (v *VgObject) LvNameValidate(name string) error {
	return nil
}

//  ==== LVObject

// GetAttr gets LV attr
func (l *LvObject) GetAttr() string {
	return ""
}

// GetName gets LV name
func (l *LvObject) GetName() string {
	return ""
}

// GetOrigin gets LV origin
func (l *LvObject) GetOrigin() string {
	return ""
}

// GetUuid gets LV uuid
func (l *LvObject) GetUuid() string {
	return ""
}

// Activate activates LV.
func (l *LvObject) Activate() error {
	return nil
}

// Deactivate deactivates LV.
func (l *LvObject) Deactivate() error {
	return nil
}

// Remove removes LV.
func (l *LvObject) Remove() error {
	return nil
}

// GetSize returns size of LV.
func (l *LvObject) GetSize() int {
	return 0
}

// IsActive checks active LV or not.
func (l *LvObject) IsActive() bool {

	return false
}

// IsSuspended checks suspended LV or not.
func (l *LvObject) IsSuspended() bool {
	return false
}

// AddTag adds tag to LV.
func (l *LvObject) AddTag(stag string) error {
	return nil
}

// RemoveTag removes tag from LV.
func (l *LvObject) RemoveTag(stag string) error {
	return nil
}

// GetTags returns tag list of LV.
func (l *LvObject) GetTags() []string {
	return nil
}

// Rename rename the name of LV.
func (l *LvObject) Rename(name string) error {
	return nil
}

// Resize resizes the size of LV.
func (l *LvObject) Resize(size uint64) error {
	return nil
}

//        { "listLVsegs",         (PyCFunction)_liblvm_lvm_lv_list_lvsegs, METH_NOARGS },

// Snapshot creates a LV snapshot.
func (l *LvObject) Snapshot(snapname string, size uint64) (*LvObject, error) {
	return nil, nil
}
