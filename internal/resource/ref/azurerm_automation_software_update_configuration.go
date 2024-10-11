package ref

import (
	"context"

	"github.com/hashicorp/hcl/v2"
	"github.com/magodo/terrafix-sdk/tfxsdk"
	"github.com/zclconf/go-cty/cty"
)

var AutomationSoftwareUpdateConfiguration = tfxsdk.ReferenceConfigUpgraders{
	0: tfxsdk.ReferenceConfigUpgrader{
		ReferenceConfigUpgrader: func(ctx context.Context, req tfxsdk.UpgradeReferenceConfigRequest, resp *tfxsdk.UpgradeReferenceConfigResponse) {
			tvs := req.Traversals
			ntvs := make([]hcl.Traversal, 0, len(tvs))

			for _, tv := range tvs {
				var err error
				tv, err = tfxsdk.TraversalReplace(
					tv,
					append(
						append(hcl.Traversal{}, tv[:2]...),
						hcl.TraverseAttr{Name: "linux"},
						hcl.TraverseIndex{Key: cty.NumberIntVal(0)},
						hcl.TraverseAttr{Name: "classification_included"},
					),
					append(
						hcl.Traversal{},
						hcl.TraverseAttr{Name: "classifications_included"},
						hcl.TraverseIndex{Key: cty.NumberIntVal(0)},
					),
				)
				if err != nil {
					resp.Error = err
					return
				}

				tv, err = tfxsdk.TraversalReplace(
					tv,
					append(
						append(hcl.Traversal{}, tv[:2]...),
						hcl.TraverseAttr{Name: "windows"},
						hcl.TraverseIndex{Key: cty.NumberIntVal(0)},
						hcl.TraverseAttr{Name: "classification_included"},
					),
					append(
						hcl.Traversal{},
						hcl.TraverseAttr{Name: "classifications_included"},
						hcl.TraverseIndex{Key: cty.NumberIntVal(0)},
					),
				)
				if err != nil {
					resp.Error = err
					return
				}

				ntvs = append(ntvs, tv)
			}

			resp.Traversals = ntvs
		},
	},
}
