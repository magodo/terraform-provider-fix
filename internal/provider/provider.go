package provider

import (
	"context"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/function"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/magodo/terrafix-sdk/tfxsdk"
	resourcedef "github.com/magodo/terraform-provider-fix/internal/resource/def"
	resourceref "github.com/magodo/terraform-provider-fix/internal/resource/ref"
)

var _ provider.ProviderWithFunctions = (*fixProvider)(nil)

type fixProvider struct{}

func New() func() provider.Provider {
	return func() provider.Provider {
		return &fixProvider{}
	}
}

func (p *fixProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
}

func (p *fixProvider) Metadata(ctx context.Context, req provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "fix"
}

func (p *fixProvider) DataSources(ctx context.Context) []func() datasource.DataSource {
	return []func() datasource.DataSource{}
}

func (p *fixProvider) Resources(ctx context.Context) []func() resource.Resource {
	return []func() resource.Resource{}
}

func (p *fixProvider) Schema(ctx context.Context, req provider.SchemaRequest, resp *provider.SchemaResponse) {
}

func (p *fixProvider) Functions(context.Context) []func() function.Function {
	definitionFixers := tfxsdk.DefinitionFixers{
		tfxsdk.BlockTypeResource: {
			"null_resource": resourcedef.NullResource,
		},
		tfxsdk.BlockTypeDataSource: {},
	}
	referenceFixers := tfxsdk.ReferenceFixers{
		tfxsdk.BlockTypeResource: {
			"null_resource": resourceref.NullResource,
		},
		tfxsdk.BlockTypeDataSource: {},
	}
	return []func() function.Function{
		func() function.Function {
			return tfxsdk.NewFixConfigDefinitionFunction(definitionFixers)
		},
		func() function.Function {
			return tfxsdk.NewFixConfigReferenceFunction(referenceFixers)
		},
	}
}
