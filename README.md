# go-vmguestlib

A wrapper of VMware® vSphere Guest API for Go.

## Project status
This is a work-in-progress.

## Disclaimer
This is not an official VMware® product.

## Documentation

The GoDoc of the project is accessible [here](https://godoc.org/github.com/xlucas/go-vmguestlib/vmguestlib).

## Requirements
In order to use this library, to build it or to run the tests, you will need to :

* Install the VMware® tools.
* Run the following commands :
```bash
# Update dynamic linker search paths
sudo sh -c 'echo "/usr/lib/vmware-tools/lib32/libvmtools.so" >> /etc/ld.so.conf.d/vmware-tools-libraries.conf'
sudo sh -c 'echo "/usr/lib/vmware-tools/lib64/libvmtools.so" >> /etc/ld.so.conf.d/vmware-tools-libraries.conf'

# Update the dynamic linker cache
sudo ldconfig
```

## Author(s)
See the accompanying [AUTHORS](AUTHORS) file.

## License
This library is distributed under the [GNU GPL V2.0 License](LICENSE).
