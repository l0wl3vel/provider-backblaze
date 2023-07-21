package b2

import "github.com/upbound/upjet/pkg/config"
import xpv1 "github.com/crossplane/crossplane-runtime/apis/common/v1"

// Configure configures individual resources by adding custom ResourceConfigurators.
func Configure(p *config.Provider) {
    p.AddResourceConfigurator("b2_bucket", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "b2"
    })
    p.AddResourceConfigurator("b2_application_key", func(r *config.Resource) {
        // We need to override the default group that upjet generated for
        // this resource, which would be "github"
        r.ShortGroup = "b2"

        r.References["bucket_id"] = config.Reference{
            Type: "Bucket",
        }

		r.Sensitive.AdditionalConnectionDetailsFn = func(attr map[string]interface{}) (map[string][]byte, error) {
			return map[string][]byte{
				xpv1.ResourceCredentialsSecretUserKey:     []byte(attr["application_key_id"].(string)),
				xpv1.ResourceCredentialsSecretPasswordKey: []byte(attr["application_key"].(string)),
			}, nil
		}
    })
}