package udev

/*
#cgo LDFLAGS: -ludev
#include <libudev.h>
#include <linux/types.h>
#include <stdlib.h>
*/
import "C"
import "unsafe"

type Error int

func (e Error) Error() string {
	return ""
}

type DeviceType uint8

type DevNum C.dev_t

type Udev struct {
	ptr *C.struct_udev
}

type Device struct {
	ptr *C.struct_udev_device
}

type Monitor struct {
	ptr *C.struct_udev_monitor
}

func (u *Udev) Ref() {
	C.udev_ref(u.ptr)
}

func (u *Udev) Unref() {
	C.udev_unref(u.ptr)
}

func NewUdev() *Udev {
	u := C.udev_new()
	if u != nil {
		return &Udev{u}
	}
	return nil
}

func (u *Udev) GetUserdata() unsafe.Pointer {
	return C.udev_get_userdata(u.ptr)
}

func (u *Udev) SetUserdata(userdata unsafe.Pointer) {
	C.udev_set_userdata(u.ptr, userdata)
}

func (d *Device) Udev() *Udev {
	u := C.udev_device_get_udev(d.ptr)
	if u != nil {
		return &Udev{u}
	}
	return nil
}

func (u *Udev) DeviceFromSysPath(syspath string) *Device {
	d := C.udev_device_new_from_syspath(u.ptr, C.CString(syspath))
	if d != nil {
		return &Device{d}
	}
	return nil
}

func DeviceFromDevNum(u *Udev, t DeviceType, num DevNum) *Device {
	d := C.udev_device_new_from_devnum(u.ptr, C.char(t), C.dev_t(num))
	if d != nil {
		return &Device{d}
	}
	return nil
}

func (u *Udev) NewDeviceFromSubsystemSysName(subsystem string, sysname string) *Device {
	d := C.udev_device_new_from_subsystem_sysname(u.ptr, C.CString(subsystem), C.CString(sysname))
	if d != nil {
		return &Device{d}
	}
	return nil
}

func (d *Device) Parent() *Device {
	p := C.udev_device_get_parent(d.ptr)
	if p != nil {
		return &Device{p}
	}
	return nil
}

func (d *Device) ParentWithSubsystemDevType(subsystem string, devType string) *Device {
	p := C.udev_device_get_parent_with_subsystem_devtype(d.ptr, C.CString(subsystem), C.CString(devType))
	if p != nil {
		return &Device{p}
	}
	return nil
}

func (d *Device) DevPath() string {
	return C.GoString(C.udev_device_get_devpath(d.ptr))
}

func (d *Device) Subsystem() string {
	return C.GoString(C.udev_device_get_subsystem(d.ptr))
}

func (d *Device) DevType() string {
	return C.GoString(C.udev_device_get_devtype(d.ptr))
}

func (d *Device) SysPath() string {
	return C.GoString(C.udev_device_get_syspath(d.ptr))
}

func (d *Device) SysName() string {
	return C.GoString(C.udev_device_get_sysname(d.ptr))
}

func (d *Device) SysNum() string {
	return C.GoString(C.udev_device_get_sysnum(d.ptr))
}

func (d *Device) DevNode() string {
	return C.GoString(C.udev_device_get_devnode(d.ptr))
}

func (d *Device) PropertyValue(key string) string {
	return C.GoString(C.udev_device_get_property_value(d.ptr, C.CString(key)))
}

func (d *Device) Driver() string {
	return C.GoString(C.udev_device_get_driver(d.ptr))
}

func (d *Device) DevNum() DevNum {
	return DevNum(C.udev_device_get_devnum(d.ptr))
}

func (d *Device) Action() string {
	return C.GoString(C.udev_device_get_action(d.ptr))
}

func (d *Device) SysAttrValue(sysattr string) string {
	return C.GoString(C.udev_device_get_sysattr_value(d.ptr, C.CString(sysattr)))
}

func (d *Device) SeqNum() uint64 {
	return uint64(C.udev_device_get_seqnum(d.ptr))
}

func (m *Monitor) Ref() {
	C.udev_monitor_ref(m.ptr)
}

func (m *Monitor) Unref() {
	C.udev_monitor_unref(m.ptr)
}

func (m *Monitor) Udev() *Udev {
	u := C.udev_monitor_get_udev(m.ptr)
	if u != nil {
		return &Udev{u}
	}
	return nil
}

func NewMonitorFromNetlink(u *Udev, name string) *Monitor {
	return &Monitor{C.udev_monitor_new_from_netlink(u.ptr, C.CString(name))}
}

func (m *Monitor) EnableReceiving() error {
	err := C.udev_monitor_enable_receiving(m.ptr)
	if err == 0 {
		return nil
	}
	return Error(err)
}

func (m *Monitor) Fd() error {
	err := C.udev_monitor_get_fd(m.ptr)
	if err == 0 {
		return nil
	}
	return Error(err)
}

func (m *Monitor) ReceiveDevice() *Device {
	d := C.udev_monitor_receive_device(m.ptr)
	if d != nil {
		return &Device{d}
	}
	return nil
}

func (m *Monitor) AddFilter(subsystem string, devtype string) error {
	var err C.int
	if len(devtype) == 0 {
		err = C.udev_monitor_filter_add_match_subsystem_devtype(m.ptr, C.CString(subsystem), nil)
	} else {
		err = C.udev_monitor_filter_add_match_subsystem_devtype(m.ptr, C.CString(subsystem), C.CString(devtype))
	}
	if err == 0 {
		return nil
	}
	return Error(err)
}

func (m *Monitor) UpdateFilter() error {
	err := C.udev_monitor_filter_update(m.ptr)
	if err == 0 {
		return nil
	}
	return Error(err)
}

func (m *Monitor) RemoveFilter() error {
	err := C.udev_monitor_filter_remove(m.ptr)
	if err == 0 {
		return nil
	}
	return Error(err)
}

type Queue struct {
	ptr *C.struct_udev_queue
}

func (q *Queue) Ref() {
	C.udev_queue_ref(q.ptr)
}

func (q *Queue) Unref() {
	C.udev_queue_unref(q.ptr)
}
func (q *Queue) Udev() *Udev {
	u := C.udev_queue_get_udev(q.ptr)
	if u != nil {
		return &Udev{u}
	}
	return nil
}

func (u *Udev) NewQueue() *Queue {
	q := C.udev_queue_new(u.ptr)
	if q != nil {
		return &Queue{q}
	}
	return nil
}

func (q *Queue) IsActive() bool {
	b := C.udev_queue_get_udev_is_active(q.ptr)
	return b == 0
}

func (q *Queue) IsEmpty() bool {
	b := C.udev_queue_get_queue_is_empty(q.ptr)
	return b == 0
}

type Enumerate struct {
	ptr *C.struct_udev_enumerate
}

func (e *Enumerate) Ref() {
	C.udev_enumerate_ref(e.ptr)
}

func (e *Enumerate) Unref() {
	C.udev_enumerate_unref(e.ptr)
}

func (e *Enumerate) Udev() *Udev {
	u := C.udev_enumerate_get_udev(e.ptr)
	if u != nil {
		return &Udev{u}
	}
	return nil
}

func (u *Udev) NewEnumerate() *Enumerate {
	e := C.udev_enumerate_new(u.ptr)
	if e != nil {
		return &Enumerate{e}
	}
	return nil
}

func (e *Enumerate) AddMatchSubsystem(subsystem string) error {
	err := C.udev_enumerate_add_match_subsystem(e.ptr, C.CString(subsystem))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e *Enumerate) AddNoMatchSubsystem(subsystem string) error {
	err := C.udev_enumerate_add_nomatch_subsystem(e.ptr, C.CString(subsystem))
	if err != 0 {
		return Error(err)
	}
	return nil
}
func (e *Enumerate) AddMatchSysAttr(sysattr string, value string) error {
	err := C.udev_enumerate_add_match_sysattr(e.ptr, C.CString(sysattr), C.CString(value))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e *Enumerate) AddNoMatchSysAttr(sysattr string, value string) error {
	err := C.udev_enumerate_add_match_sysattr(e.ptr, C.CString(sysattr), C.CString(value))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e *Enumerate) AddMatchProperty(property string, value string) error {
	err := C.udev_enumerate_add_match_property(e.ptr, C.CString(property), C.CString(value))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e *Enumerate) AddMatchSysName(sysname string) error {
	err := C.udev_enumerate_add_match_sysname(e.ptr, C.CString(sysname))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e *Enumerate) AddSysPath(sysPath string) error {
	err := C.udev_enumerate_add_match_sysname(e.ptr, C.CString(sysPath))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e *Enumerate) ScanDevices() error {
	err := C.udev_enumerate_scan_devices(e.ptr)
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e *Enumerate) ScanSubsystems() error {
	err := C.udev_enumerate_scan_subsystems(e.ptr)
	if err != 0 {
		return Error(err)
	}
	return nil
}

type ListEntry struct {
	ptr *C.struct_udev_list_entry
}

func (l *ListEntry) Next() *ListEntry {
	e := C.udev_list_entry_get_next(l.ptr)
	if e != nil {
		return &ListEntry{e}
	}
	return nil
}

func (l *ListEntry) ByName(name string) *ListEntry {
	e := C.udev_list_entry_get_by_name(l.ptr, C.CString(name))
	if e != nil {
		return &ListEntry{e}
	}
	return nil
}

func (l *ListEntry) Name() string {
	return C.GoString(C.udev_list_entry_get_name(l.ptr))
}

func (l *ListEntry) Value() string {
	return C.GoString(C.udev_list_entry_get_value(l.ptr))
}

func (e *Enumerate) First() *ListEntry {
	l := C.udev_enumerate_get_list_entry(e.ptr)
	if l != nil {
		return &ListEntry{l}
	}
	return nil
}
func (d *Device) FirstDevLinks() *ListEntry {
	l := C.udev_device_get_devlinks_list_entry(d.ptr)
	if l != nil {
		return &ListEntry{l}
	}
	return nil
}

func (d *Device) FirstProperties() *ListEntry {
	l := C.udev_device_get_properties_list_entry(d.ptr)
	if l != nil {
		return &ListEntry{l}
	}
	return nil
}
