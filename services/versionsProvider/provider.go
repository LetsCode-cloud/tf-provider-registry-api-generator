package versionsProvider

import (
	"encoding/json"
	"fmt"
	"github.com/LetsCode-cloud/tf-provider-registry-api-generator/types"
	"io"
	"log"
	"net/http"
	"time"
)

const apiProvidersFullPath = "%s/v1/providers"
const apiVersionsFullPath = apiProvidersFullPath + "/%s/%s/versions"

func GetInitialVersions(version string, cfg types.Config) *types.ProviderVersions {
	url := "https://" + fmt.Sprintf(apiVersionsFullPath, cfg.RegistryUrl, cfg.OrganizationName, cfg.PluginName)
	versions := types.ProviderVersions{}.Construct()

	httpClient := http.Client{
		Timeout: time.Second * 10, // Timeout after 10 seconds
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}

	req.Header.Set("User-Agent", "tf-provider-registry-api-generator")

	res, getErr := httpClient.Do(req)
	if getErr != nil {
		//act as no file exists
		return versions
	}

	if res.Body != nil {
		defer res.Body.Close()
	}

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		log.Fatal(readErr)
	}

	log.Println(body)

	jsonErr := json.Unmarshal(body, &versions)
	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	return versions
}
