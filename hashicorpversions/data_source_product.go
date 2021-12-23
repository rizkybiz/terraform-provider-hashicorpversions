package hashicorpversions

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"time"

	"github.com/Masterminds/semver"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

// Version is a struct for unmarshaling
type Version struct {
	Name     string                 `json:"name"`
	Versions map[string]VersionInfo `json:"versions"`
}

type VersionInfo struct {
	Name              string      `json:"name"`
	Version           string      `json:"version"`
	SHASums           string      `json:"shasums"`
	SHASumsSignature  string      `json:"shasums_signature"`
	SHASumsSignatures []string    `json:"shasums_signatures"`
	Builds            []BuildInfo `json:"builds"`
}

type BuildInfo struct {
	Name     string `json:"name"`
	Version  string `json:"version"`
	OS       string `json:"os"`
	Arch     string `json:"arch"`
	Filename string `json:"filename"`
	URL      string `json:"url"`
}

func dataSourceProduct() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceProductRead,
		Schema: map[string]*schema.Schema{
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"shasums": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shasums_signature": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"shasums_signatures": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"builds": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"version": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"os": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"arch": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"filename": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"url": {
							Type:     schema.TypeString,
							Computed: true,
						},
					},
				},
			},
		},
	}
}

func dataSourceProductRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// setup the HTTP client
	client := &http.Client{Timeout: 10 * time.Second}

	// warnings or errors can be collected in a slice type
	var diags diag.Diagnostics

	// assemble the HTTP request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://releases.hashicorp.com/%s/index.json", d.Get("name")), nil)
	if err != nil {
		return diag.FromErr(err)
	}

	// execute the HTTP request
	r, err := client.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	// process the response body
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return diag.FromErr(err)
	}

	// unmarshal JSON body into VersionInfo
	vi := Version{}
	err = json.Unmarshal(body, &vi)
	if err != nil {
		return diag.FromErr(err)
	}

	// compile a slice of versions, check if they are simple SemVer,
	// make a slice of semver objects, sort the semvers, then return
	// the most recent
	validVers := regexp.MustCompile(`^(0|[1-9]\d*)\.(0|[1-9]\d*)\.(0|[1-9]\d*)$`)
	var versions []string
	for vers := range vi.Versions {
		if validVers.MatchString(vers) {
			versions = append(versions, vers)
		}
	}
	vs := make([]*semver.Version, len(versions))
	for i, r := range versions {
		v, err := semver.NewVersion(r)
		if err != nil {
			return diag.FromErr(err)
		}
		vs[i] = v
	}
	sort.Sort(semver.Collection(vs))
	recent := vs[len(vs)-1].String()
	info := vi.Versions[vs[len(vs)-1].String()]

	// set the info within the data source
	err = setDataSourceInfo(d, recent, info)
	if err != nil {
		diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	// default return
	return diags
}

func setDataSourceInfo(d *schema.ResourceData, version string, info VersionInfo) error {
	err := d.Set("version", version)
	if err != nil {
		return err
	}
	err = d.Set("name", info.Name)
	if err != nil {
		return err
	}
	err = d.Set("shasums", info.SHASums)
	if err != nil {
		return err
	}
	err = d.Set("shasums_signature", info.SHASumsSignature)
	if err != nil {
		return err
	}
	err = d.Set("shasums_signatures", info.SHASumsSignatures)
	if err != nil {
		return err
	}

	var builds []interface{}
	for i := len(info.Builds) - 1; i >= 0; i-- {
		builds = append(builds, parseVersionBuilds(info.Builds[i]))
	}

	err = d.Set("builds", builds)
	if err != nil {
		return err
	}
	return nil
}

func parseVersionBuilds(build BuildInfo) map[string]interface{} {
	ret := map[string]interface{}{
		"arch":     build.Arch,
		"filename": build.Filename,
		"name":     build.Name,
		"os":       build.OS,
		"url":      build.URL,
		"version":  build.Version,
	}

	return ret
}
