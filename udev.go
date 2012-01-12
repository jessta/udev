package udev

/* 
  #cgo LDFLAGS: -ludev
  #include <libudev.h>
  #include <linux/types.h>
  #include <stdlib.h>
*/
import "C"

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

func NewUdev() Udev {
	var u = C.udev_new()
	return Udev{u}
}

func (u Udev) Ref() {
	C.udev_ref(u.ptr)
}

func (u Udev) Unref() {
	C.udev_unref(u.ptr)
}

/*
func (u Udev) SetLogger() {
	C.udev_set_log_fn(u.ptr)
}

func (u Udev) GetLogPriority() {
	C.udev_get_log_priority(u.ptr)
}

func (u Udev) SetLogPriority() {
	C.udev_set_log_priority(u.ptr)
}
*/

func (u Udev) SysPath() string {
	return C.GoString(C.udev_get_sys_path(u.ptr))
}

func (u Udev) DevPath() string {
	return C.GoString(C.udev_get_dev_path(u.ptr))
}

func (d Device) Udev() Udev {
	return Udev{C.udev_device_get_udev(d.ptr)}
}

func (u Udev) DeviceFromSysPath(syspath string) Device {
	return Device{C.udev_device_new_from_syspath(u.ptr, C.CString(syspath))}
}
func DeviceFromDevNum(u Udev, t DeviceType, num DevNum) Device {
	return Device{C.udev_device_new_from_devnum(u.ptr, C.char(t), C.dev_t(num))}
}

func (u Udev) NewDeviceFromSubsystemSysName(subsystem string, sysname string) Device {
	return Device{C.udev_device_new_from_subsystem_sysname(u.ptr, C.CString(subsystem), C.CString(sysname))}
}

func (d Device) Parent() Device {
	return Device{C.udev_device_get_parent(d.ptr)}
}
func (d Device) IsNil() bool {
	return d.ptr == nil

}
func (d Device) ParentWithSubsystemDevType(subsystem string, devType string) Device {
	return Device{C.udev_device_get_parent_with_subsystem_devtype(d.ptr, C.CString(subsystem), C.CString(devType))}
}

func (d Device) DevPath() string {
	return C.GoString(C.udev_device_get_devpath(d.ptr))

}

func (d Device) Subsystem() string {
	return C.GoString(C.udev_device_get_subsystem(d.ptr))
}

func (d Device) DevType() string {
	return C.GoString(C.udev_device_get_devtype(d.ptr))
}

func (d Device) SysPath() string {
	return C.GoString(C.udev_device_get_syspath(d.ptr))
}

func (d Device) SysName() string {
	return C.GoString(C.udev_device_get_sysname(d.ptr))
}

func (d Device) SysNum() string {
	return C.GoString(C.udev_device_get_sysnum(d.ptr))
}

func (d Device) DevNode() string {
	return C.GoString(C.udev_device_get_devnode(d.ptr))
}

func (d Device) PropertyValue(key string) string {
	return C.GoString(C.udev_device_get_property_value(d.ptr, C.CString(key)))
}

func (d Device) Driver() string {
	return C.GoString(C.udev_device_get_driver(d.ptr))
}

func (d Device) DevNum() DevNum {
	return DevNum(C.udev_device_get_devnum(d.ptr))
}

func (d Device) Action() string {
	return C.GoString(C.udev_device_get_action(d.ptr))
}

func (d Device) SysAttrValue(sysattr string) string {
	return C.GoString(C.udev_device_get_sysattr_value(d.ptr, C.CString(sysattr)))
}

func (d Device) SeqNum() uint64 {
	return uint64(C.udev_device_get_seqnum(d.ptr))
}

func (m Monitor) Ref() {
	C.udev_monitor_ref(m.ptr)
}

func (m Monitor) Unref() {
	C.udev_monitor_unref(m.ptr)
}

func (m Monitor) Udev() Udev {
	return Udev{C.udev_monitor_get_udev(m.ptr)}
}

func NewMonitorFromNetlink(u Udev, name string) Monitor {
	return Monitor{C.udev_monitor_new_from_netlink(u.ptr, C.CString(name))}
}

func NewMonitorFromSocket(u Udev, socketPath string) Monitor {
	return Monitor{C.udev_monitor_new_from_socket(u.ptr, C.CString(socketPath))}
}
func (m Monitor) EnableReceiving() error {
	err := C.udev_monitor_enable_receiving(m.ptr)
	if err == 0 {
		return nil
	}
	return Error(err)

}

func (m Monitor) Fd() error {
	err := C.udev_monitor_get_fd(m.ptr)
	if err == 0 {
		return nil
	}
	return Error(err)
}

func (m Monitor) ReceiveDevice() Device {
	return Device{C.udev_monitor_receive_device(m.ptr)}
}

func (m Monitor) AddFilter(subsystem string, devtype string) error {
	err := C.udev_monitor_filter_add_match_subsystem_devtype(m.ptr, C.CString(subsystem), C.CString(devtype))
	if err == 0 {
		return nil
	}
	return Error(err)

}

func (m Monitor) UpdateFilter() error {
	err := C.udev_monitor_filter_update(m.ptr)
	if err == 0 {
		return nil
	}
	return Error(err)

}

func (m Monitor) RemoveFilter() error {
	err := C.udev_monitor_filter_remove(m.ptr)

	if err == 0 {
		return nil
	}
	return Error(err)

}

type Queue struct {
	ptr *C.struct_udev_queue
}

func (q Queue) Ref() {
	C.udev_queue_ref(q.ptr)
}

func (q Queue) Unref() {
	C.udev_queue_unref(q.ptr)
}
func (q Queue) Udev() Udev {
	return Udev{C.udev_queue_get_udev(q.ptr)}
}

func (u Udev) NewQueue() Queue {
	return Queue{C.udev_queue_new(u.ptr)}
}
func (q Queue) IsActive() bool {
	b := C.udev_queue_get_udev_is_active(q.ptr)
	return b == 0
}

func (q Queue) IsEmpty() bool {
	b := C.udev_queue_get_queue_is_empty(q.ptr)
	return b == 0
}

func (q Queue) SeqNumIsFinished(seqNum uint64) bool {
	b := C.udev_queue_get_seqnum_is_finished(q.ptr, C.ulonglong(seqNum))
	return b == 0
}

func (q Queue) SeqNumSequenceIsFinished(start uint64, end uint64) bool {
	b := C.udev_queue_get_seqnum_sequence_is_finished(q.ptr, C.ulonglong(start), C.ulonglong(end))
	return b == 0
}

func (q Queue) KernelSeqNum() uint64 {
	return uint64(C.udev_queue_get_kernel_seqnum(q.ptr))
}

func (q Queue) UdevSeqNum() uint64 {
	return uint64(C.udev_queue_get_udev_seqnum(q.ptr))
}

type Enumerate struct {
	ptr *C.struct_udev_enumerate
}

func (e Enumerate) Ref() {
	C.udev_enumerate_ref(e.ptr)
}

func (e Enumerate) Unref() {
	C.udev_enumerate_unref(e.ptr)
}

func (e Enumerate) Udev() Udev {
	return Udev{C.udev_enumerate_get_udev(e.ptr)}
}

func (u Udev) NewEnumerate() Enumerate {
	return Enumerate{C.udev_enumerate_new(u.ptr)}
}

func (e Enumerate) AddMatchSubsystem(subsystem string) error {
	err := C.udev_enumerate_add_match_subsystem(e.ptr, C.CString(subsystem))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e Enumerate) AddNoMatchSubsystem(subsystem string) error {
	err := C.udev_enumerate_add_nomatch_subsystem(e.ptr, C.CString(subsystem))
	if err != 0 {
		return Error(err)
	}
	return nil
}
func (e Enumerate) AddMatchSysAttr(sysattr string, value string) error {
	err := C.udev_enumerate_add_match_sysattr(e.ptr, C.CString(sysattr), C.CString(value))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e Enumerate) AddNoMatchSysAttr(sysattr string, value string) error {
	err := C.udev_enumerate_add_match_sysattr(e.ptr, C.CString(sysattr), C.CString(value))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e Enumerate) AddMatchProperty(property string, value string) error {
	err := C.udev_enumerate_add_match_property(e.ptr, C.CString(property), C.CString(value))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e Enumerate) AddMatchSysName(sysname string) error {
	err := C.udev_enumerate_add_match_sysname(e.ptr, C.CString(sysname))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e Enumerate) AddSysPath(sysPath string) error {
	err := C.udev_enumerate_add_match_sysname(e.ptr, C.CString(sysPath))
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e Enumerate) ScanDevices() error {
	err := C.udev_enumerate_scan_devices(e.ptr)
	if err != 0 {
		return Error(err)
	}
	return nil
}

func (e Enumerate) ScanSubsystems() error {
	err := C.udev_enumerate_scan_subsystems(e.ptr)
	if err != 0 {
		return Error(err)
	}
	return nil
}

type ListEntry struct {
	ptr *C.struct_list_entry
}

func (l ListEntry) IsNil() bool {
	return l.ptr == nil
}

func (l ListEntry) Next() ListEntry {
	return ListEntry{C.udev_list_entry_get_next(l.ptr)}
}

func (l ListEntry) ByName(name string) ListEntry {
	return ListEntry{C.udev_list_entry_get_by_name(l.ptr, C.CString(name))}
}

func (l ListEntry) Name() string {
	return C.GoString(C.udev_list_entry_get_name(l.ptr))
}

func (l ListEntry) Value() string {
	return C.GoString(C.udev_list_entry_get_value(l.ptr))
}

func (e Enumerate) First() ListEntry {
	return ListEntry{C.udev_enumerate_get_list_entry(e.ptr)}
}
func (e Enumerate) FirstDevLinks() ListEntry {
	return ListEntry{C.udev_device_get_devlinks_list_entry(e.ptr)}
}

func (e Enumerate) FirstProperties() ListEntry {
	return ListEntry{C.udev_device_get_properties_list_entry(e.ptr)}
}

func (e Enumerate) FirstQueued() ListEntry {

	return ListEntry{C.udev_queue_get_queued_list_entry(e.ptr)}
}

func (e Enumerate) FirstFailed() ListEntry {
	return ListEntry{C.udev_queue_get_failed_list_entry(e.ptr)}
}
