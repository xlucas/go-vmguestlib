# go-vmguestlib

[![GoDoc](https://godoc.org/github.com/xlucas/go-vmguestlib/vmguestlib?status.svg)](https://godoc.org/github.com/xlucas/go-vmguestlib/vmguestlib)
[![GitHub version](https://badge.fury.io/gh/xlucas%2Fgo-vmguestlib.svg)](http://badge.fury.io/gh/xlucas%2Fgo-vmguestlib)

VMware® vSphere Guest API for Go.

## Disclaimer
This is not an official VMware® product.

## Requirements
In order to use this library, to build it or to run the tests, you will need to :

* Install the VMware® tools.
* Make sure Guest API runtime components are enabled on vSphere. If not, set `isolation.tools.guestlibGetInfo.disable = "TRUE"` in your virtual machines configurations.
* Run the following commands as root :
  * On Debian/Ubuntu :
  ```bash
  echo "/usr/lib/vmware-tools/lib/libvmtools.so" >> /etc/ld.so.conf.d/vmware-tools-libraries.conf
  ```

  * On RedHat/CentOS :
  ```bash
  echo "/usr/lib/vmware-tools/lib/libvmtools.so" >> /etc/ld.so.conf.d/vmware-tools-guestlib.conf
  ```

  * Then, for all platforms :
  ```bash
  ldconfig
  ```

## Author(s)
See the accompanying [AUTHORS](AUTHORS) file.

## License
This library is distributed under the [GNU GPL V2.0 License](LICENSE).
