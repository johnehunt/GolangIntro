jeh@Johns-iMac modules % go mod init 10-packages-and-modules/modules
go: creating new go.mod: module 10-packages-and-modules/modules

The go.mod file only appears in the root of the module. 

Packages in subdirectories have import paths consisting of the module path 
plus the path to the subdirectory.

The command go list -m all lists the current module and all its dependencies:

With Go modules, versions are referenced with semantic version tags. 

A semantic version has three parts: major, minor, and patch. 

For example, for v0.1.2, the major version is 0, the minor version is 1, and 
the patch version is 2. Let's walk through a couple minor version upgrades. 



go list -m all
10-packages-and-modules/modules

Then add in util dpeendecnies on 
    "rsc.io/quote"
    quoteV3 "rsc.io/quote/v3"

jeh@Johns-iMac modules % go list -m all
go: finding golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
go: finding rsc.io/quote v1.5.2
go: finding rsc.io/quote/v3 v3.1.0
go: finding rsc.io/sampler v1.3.0
10-packages-and-modules/modules
golang.org/x/text v0.0.0-20170915032832-14c0d48ead0c
rsc.io/quote v1.5.2
rsc.io/quote/v3 v3.1.0
rsc.io/sampler v1.3.0
jeh@Johns-iMac modules % 

go.mod file now contains

require (
	rsc.io/quote v1.5.2
	rsc.io/quote/v3 v3.1.0
)

