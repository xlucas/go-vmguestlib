# go-vmguestlib

A wrapper of VMware® vSphere Guest native API for Go.

## Project status
This is a work-in-progress.

## Documentation

The GoDoc of the project is accessible [here](https://godoc.org/github.com/xlucas/go-vmguestlib/vmguestlib).

## Requirements
In order to use this library, to build it or to run the tests, you will need to :

* Install the VMware® tools.
* Add the following to `/etc/ld.so.conf.d/vmware-tools-libraries.conf` :
```
/usr/lib/vmware-tools/lib32/libvmtools.so
/usr/lib/vmware-tools/lib64/libvmtools.so
```
* Make sure the dynamic linker's cache is up to date by running `sudo ldconfig` after these steps.

## Disclaimer
This is not an official VMware® product.

## License
This library is distributed under the [GNU GPL V2.0 License](LICENSE).

## Authors
See the accompanying [AUTHORS](AUTHORS) file.
