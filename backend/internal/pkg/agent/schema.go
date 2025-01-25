package agent

import "github.com/invopop/jsonschema"

type ProjectTitleAndSlug struct {
	Title       string `json:"title" jsonschema_description:"The title of the project"`
	Slug        string `json:"slug" jsonschema_description:"The slug of the project"`
	Description string `json:"description" jsonschema_description:"The description of the project"`
}

// Generate the JSON schema at initialization time
var ProjectTitleAndSlugResponseSchema = GenerateSchema[ProjectTitleAndSlug]()

func GenerateSchema[T any]() interface{} {
	reflector := jsonschema.Reflector{
		AllowAdditionalProperties: false,
		DoNotReference:            true,
	}
	var v T
	schema := reflector.Reflect(v)
	return schema
}
