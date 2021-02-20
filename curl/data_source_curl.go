package curl

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-plugin-sdk/v2/diag"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

func dataSource() *schema.Resource {
	return &schema.Resource{
		ReadContext: dataSourceRead,
		Schema: map[string]*schema.Schema{
			"uri": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "URI of resource you'd like to retrieve via HTTP(s).",
			},
			"http_method": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
				Description: "HTTP method like GET, POST, PUT, DELETE, PATCH.",
			},
			"response": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Description: "Valued returned by the HTTP request.",
			},
		},
	}
}

func dataSourceRead(ctx context.Context, d *schema.ResourceData, meta interface{}) diag.Diagnostics {
	//this will retrieve the MyHttpClient that we instantiated in the 'providerConfigure' method (in provider)
	myClient := meta.(*MyHttpClient)

	// Warning or errors can be collected in a slice type
	var diags diag.Diagnostics
	uri := d.Get("uri").(string)
	httpMethod := d.Get("http_method").(string)

	req, err := http.NewRequest(httpMethod, uri, nil)
	if err != nil {
		return diag.FromErr(err)
	}

	if myClient.oauthConfig != nil {
		spToken, err := myClient.oauthConfig.ServicePrincipalToken()
		if err != nil {
			fmt.Errorf("failed %s\n", err.Error())
		}

		err = spToken.EnsureFresh()
		if err != nil {
			return diag.FromErr(err)
		}
		if err != nil {
			return diag.FromErr(err)
		}
		req.Header.Set("Authorization", "Bearer " + spToken.OAuthToken())
	}

	log.Printf("[INFO] ******* About to send request")
	r, err := myClient.httpClient.Do(req)
	if err != nil {
		return diag.FromErr(err)
	}
	defer r.Body.Close()

	responseData, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return diag.FromErr(err)
	}
	responseString := string(responseData)

	//set the actual value we're going to return into, associated with the 'response' key name.
	if err := d.Set("response", responseString); err != nil {
		return diag.FromErr(err)
	}

	// force that it always sets for the newest json object by changing the id of the object
	d.SetId(strconv.FormatInt(time.Now().Unix(), 10))
	return diags
}