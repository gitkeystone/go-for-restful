package docs

import "github.com/go-openapi/spec"

func EnrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Description:    "UserService",
			Title:          "Resource for managing Users",
			TermsOfService: "",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "john",
					URL:   "john@doe.rp",
					Email: "https://johndoe.org",
				},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "https://mit.org",
				},
			},
			Version: "1.0.0",
		},
	}

	swo.Tags = []spec.Tag{
		{
			TagProps: spec.TagProps{
				Description: "Managing users",
				Name:        "users",
			},
		},
	}
}
