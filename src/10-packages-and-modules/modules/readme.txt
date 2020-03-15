jeh@Johns-iMac modules % go mod init 10-packages-and-modules/modules
go: creating new go.mod: module 10-packages-and-modules/modules

The go.mod file only appears in the root of the module. 

Packages in subdirectories have import paths consisting of the module path 
plus the path to the subdirectory.

