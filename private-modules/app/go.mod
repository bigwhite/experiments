module app

go 1.19

require (
	github.com/bigwhite/privatemodule3 v0.0.0-20230227061700-3762215e798f
	mycompany.com/go/privatemodule1 v1.0.0
	mycompany.com/go/privatemodule2 v1.0.0
)

replace (
	mycompany.com/go/privatemodule1 v1.0.0 => 10.10.30.30/ard/incubators/privatemodule1.git v0.0.0-20230227061032-c4a6ea813d1a
	mycompany.com/go/privatemodule2 v1.0.0 => github.com/bigwhite/privatemodule2 v0.0.0-20230227061454-a2de3aaa7b27
)
