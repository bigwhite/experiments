package main
import (
	"context"
	"fmt"
	"os"
	"github.com/tetratelabs/wazero"
	"github.com/tetratelabs/wazero/api"
	"github.com/tetratelabs/wazero/imports/wasi_snapshot_preview1"
)

func main() {
	ctx := context.Background()
	r := wazero.NewRuntime(ctx)
	defer r.Close(ctx)
	buf, err := os.ReadFile(os.Args[1])
	if err != nil {
		panic(err)
	}
	config := wazero.NewModuleConfig().
		WithStdout(os.Stdout).WithStderr(os.Stderr).
		WithStartFunctions() // don't call _start
	wasi_snapshot_preview1.MustInstantiate(ctx, r)
	m, err := r.InstantiateWithConfig(ctx, buf, config)
	if err != nil {
		panic(err)
	}

	// get export functions from the module
	F := func(a int64, b int64) int64 {
		exp := m.ExportedFunction("Add")
		r, err := exp.Call(ctx, api.EncodeI64(a), api.EncodeI64(b))
		if err != nil {
			panic(err)
		}
	        rr := int64(r[0])
                fmt.Printf("host: Add %d + %d = %d\n", a,b,rr)
                return rr
	}

	// Library mode.
	entry := m.ExportedFunction("_initialize")
	fmt.Println("Library mode: initialize")
	_, err = entry.Call(ctx)
	if err != nil {
		panic(err)
	}
	fmt.Println("\nLibrary mode: call export functions")
	println(F(5,6))
}

