package def

import (
	"context"

	"github.com/magodo/terrafix-sdk/tfxsdk"
)

var NullResource = tfxsdk.DefinitionConfigUpgraders{
	0: tfxsdk.DefinitionConfigUpgrader{
		DefinitionConfigUpgrader: func(ctx context.Context, req tfxsdk.UpgradeDefinitionConfigRequest, resp *tfxsdk.UpgradeDefinitionConfigResponse) {
			wbody := req.WriteBody
			wbody.SetAttributeRaw("the_triggers", wbody.RemoveAttribute("triggers").Expr().BuildTokens(nil))
		},
	},
}
