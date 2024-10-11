package def

import (
	"context"

	"github.com/hashicorp/hcl/v2/hclwrite"
	"github.com/magodo/terrafix-sdk/tfxsdk"
)

var AutomationSoftwareUpdateConfiguration = tfxsdk.DefinitionConfigUpgraders{
	0: tfxsdk.DefinitionConfigUpgrader{
		DefinitionConfigUpgrader: func(ctx context.Context, req tfxsdk.UpgradeDefinitionConfigRequest, resp *tfxsdk.UpgradeDefinitionConfigResponse) {
			wbody := req.WriteBody

			fixInBlock := func(blk *hclwrite.Block) {
				body := blk.Body()
				if v, ok := body.Attributes()["classification_included"]; ok {
					body.RemoveAttribute("classification_included")

					tks := hclwrite.TokensForTuple([]hclwrite.Tokens{v.Expr().BuildTokens(nil)})
					body.SetAttributeRaw("classifications_included", tks)
				}
			}

			// linux.classification_included -> linux.classifications_included
			if blk := wbody.FirstMatchingBlock("linux", nil); blk != nil {
				fixInBlock(blk)
			}

			// windows.classification_included -> windows.classifications_included
			if blk := wbody.FirstMatchingBlock("windows", nil); blk != nil {
				fixInBlock(blk)
			}

			// operating_system removed
			wbody.RemoveAttribute("operating_system")
		},
	},
}
