package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/lnashier/viper"
)

var (
	sourceDir string
	dstDir    string
)

func init() {
	flag.StringVar(&sourceDir, "s", "./", "template directory path")
	flag.StringVar(&dstDir, "d", "./k8s-install", "the target directory path in which the generated files are put")
}

func mergeConfig(configType, srcPath, srcFile, overridePath, overrideFile, target string) error {
	v1 := viper.New()
	v1.SetConfigType(configType) // e.g. "yml"
	v1.AddConfigPath(srcPath)    // file directory
	v1.SetConfigName(srcFile)    // filename(without postfix)
	err := v1.ReadInConfig()
	if err != nil {
		return err
	}

	v2 := viper.New()
	v2.SetConfigType(configType)
	v2.AddConfigPath(overridePath)
	v2.SetConfigName(overrideFile)
	err = v2.ReadInConfig()
	if err != nil {
		return err
	}

	overrideKeys := v2.AllKeys()

	// override special keys
	prefixKey := srcFile + "." + configType + "." // e.g "a.yml."
	for _, key := range overrideKeys {
		if !strings.HasPrefix(key, prefixKey) {
			continue
		}

		stripKey := strings.TrimPrefix(key, prefixKey)
		val := v2.Get(key)
		v1.Set(stripKey, val)
	}

	// write the final result after overriding
	return v1.WriteConfigAs(target)
}

func main() {
	var err error
	flag.Parse()

	// create target directory if not exist
	err = os.MkdirAll(dstDir+"/conf", 0775)
	if err != nil {
		fmt.Printf("create %s error: %s\n", dstDir+"/conf", err)
		return
	}

	err = os.MkdirAll(dstDir+"/manifests", 0775)
	if err != nil {
		fmt.Printf("create %s error: %s\n", dstDir+"/manifests", err)
		return
	}

	// override manifests files with same config item in custom/manifests.yml,
	// store the final result to the target directory
	err = mergeManifestsFiles()
	if err != nil {
		fmt.Printf("override and generate manifests files error: %s\n", err)
		return
	}
	fmt.Printf("override and generate manifests files ok\n")
}

var (
	manifestFiles = []string{
		"a.yml",
		"b.yml",
	}
)

func mergeManifestsFiles() error {
	for _, file := range manifestFiles {
		// check whether the file exist
		srcFile := sourceDir + "/templates/manifests/" + file
		_, err := os.Stat(srcFile)
		if os.IsNotExist(err) {
			fmt.Printf("%s not exist, ignore it\n", srcFile)
			continue
		}

		err = mergeConfig("yml", sourceDir+"/templates/manifests", strings.TrimSuffix(file, ".yml"),
			sourceDir+"/custom", "manifests", dstDir+"/manifests/"+file)
		if err != nil {
			fmt.Println("mergeConfig error: ", err)
			return err
		}
		fmt.Printf("mergeConfig %s ok\n", file)

	}
	return nil
}
