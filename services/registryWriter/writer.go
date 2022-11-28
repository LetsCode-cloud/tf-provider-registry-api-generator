package registryWriter

import (
	"fmt"
	"github.com/LetsCode-cloud/tf-provider-registry-api-generator/services/io"
	"github.com/LetsCode-cloud/tf-provider-registry-api-generator/services/versionsProvider"
	"github.com/LetsCode-cloud/tf-provider-registry-api-generator/signing_key"
	"github.com/LetsCode-cloud/tf-provider-registry-api-generator/types"
	"log"
	"os"
	"path/filepath"
)

const wellKnownDirName = ".well-known"
const providersFullPath = "%s" + string(os.PathSeparator) + "v1" + string(os.PathSeparator) + "providers"
const pluginRepoDirFmt = providersFullPath + string(os.PathSeparator) + "%s" + string(os.PathSeparator) + "%s" + string(os.PathSeparator)

func WriteRegistry(cfg types.Config, binDir string) {
	outputDir := prepareDistDir()
	ps := string(os.PathSeparator)

	wellKnownFullPath := outputDir + ps + wellKnownDirName
	versionsFullPath := fmt.Sprintf(pluginRepoDirFmt, outputDir, cfg.OrganizationName, cfg.PluginName)

	sk := signing_key.GetPublicSigningKey(cfg.PgpKeyId)
	shasums := make(map[string]string)
	files := getFiles(binDir)
	fileDescr := types.CreateFromFileList(files, cfg.RegistryUrl, sk, shasums, cfg.Protocols)

	version := ""
	newVersionDownloadFullPath := ""
	archDownloadFullPathFmt := ""
	providerVersionsContent := versionsProvider.GetInitialVersions(version, cfg)
	for _, f := range fileDescr {
		newVersionDownloadFullPath = versionsFullPath + ps + f.Version + ps + "download"
		archDownloadFullPathFmt = newVersionDownloadFullPath + ps + "%s"
		pth := archDownloadFullPathFmt + ps + f.Arch
		io.WriteToJsonFile(f, pth)
	}

	for _, v := range fileDescr.ExtractVersions() {
		providerVersionsContent.Merge(*v)
	}

	io.CreateDir(wellKnownFullPath)
	makeDiscoveryFile(wellKnownFullPath)
	io.WriteToJsonFile(providerVersionsContent, versionsFullPath+string(os.PathSeparator)+"versions")
}

func prepareDistDir() string {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	outputDir := path + string(os.PathSeparator) + "registryDist"
	if io.DirExists(outputDir) {
		io.RemoveRecursivelyDir(outputDir)
	}
	io.CreateDir(outputDir)

	return outputDir
}

func makeDiscoveryFile(wellKnownFullPath string) {
	terraformJsonContent := types.Discovery{}.Construct()
	io.WriteToJsonFile(terraformJsonContent, wellKnownFullPath+string(os.PathSeparator)+"terraform.json")
}

func getFiles(binDir string) []string {
	files, err := filepath.Glob(binDir + string(os.PathSeparator) + "/terraform-provider-*")
	if err != nil {
		log.Fatal(err)
	}
	return files
}
