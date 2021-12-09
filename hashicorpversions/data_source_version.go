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

// VersionInfo is a struct for unmarshaling
type VersionInfo struct {
	Name     string                     `json:"name"`
	Versions map[string]json.RawMessage `json:"versions"`
}

func dataSourceVersion() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceVersionRead,
		Schema: map[string]*schema.Schema{
			"version": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"product": {
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}

func dataSourceVersionRead(ctx context.Context, d *schema.ResourceData, m interface{}) diag.Diagnostics {
	// setup the HTTP client
	client := &http.Client{Timeout: 10 * time.Second}

	// warnings or errors can be collected in a slice type
	var diags diag.Diagnostics

	// assemble the HTTP request
	req, err := http.NewRequest("GET", fmt.Sprintf("https://releases.hashicorp.com/%s/index.json", d.Get("product")), nil)
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
	vi := VersionInfo{}
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

	// set the version within the data source
	err = d.Set("version", recent)
	if err != nil {
		return diag.FromErr(err)
	}

	// always run
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))

	// default return
	return diags
}
