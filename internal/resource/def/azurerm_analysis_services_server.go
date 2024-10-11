package def

import (
	"context"

	"github.com/magodo/terrafix-sdk/tfxsdk"
	"github.com/zclconf/go-cty/cty"
)

var AnalysisServicesServer = tfxsdk.DefinitionConfigUpgraders{
	0: tfxsdk.DefinitionConfigUpgrader{
		DefinitionConfigUpgrader: func(ctx context.Context, req tfxsdk.UpgradeDefinitionConfigRequest, resp *tfxsdk.UpgradeDefinitionConfigResponse) {
			state, sbody, wbody := req.State, req.SyntaxBody, req.WriteBody

			// enable_power_bi_service -> power_bi_service_enabled
			wbody.SetAttributeRaw("power_bi_service_enabled", wbody.RemoveAttribute("enable_power_bi_service").Expr().BuildTokens(nil))

			// querypool_connection_mode O+C -> O (with All as default)
			// In case it is absent in config and its remote value not equals to "All", we'll need to explicitly set the current value learned from state.
			if state != nil {
				if _, ok := sbody.Attributes["querypool_connection_mode"]; !ok {
					if v, ok := state["querypool_connection_mode"]; ok && v.(string) != "All" {
						wbody.SetAttributeValue("querypool_connection_mode", cty.StringVal(v.(string)))
					}
				}
			}
		},
	},
}
