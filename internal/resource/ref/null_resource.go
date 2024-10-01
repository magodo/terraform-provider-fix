package ref

import (
	"context"

	"github.com/hashicorp/hcl/v2"
	"github.com/magodo/terrafix-sdk/tfxsdk"
)

var NullResource = tfxsdk.ReferenceConfigUpgraders{
	0: tfxsdk.ReferenceConfigUpgrader{
		ReferenceConfigUpgrader: func(ctx context.Context, req tfxsdk.UpgradeReferenceConfigRequest, resp *tfxsdk.UpgradeReferenceConfigResponse) {
			tvs := req.Traversals
			ntvs := make([]hcl.Traversal, 0, len(tvs))
			for _, tv := range tvs {
				ntv, err := tfxsdk.TraversalReplace(
					tv,
					append(tv[:2], hcl.TraverseAttr{Name: "triggers"}),
					hcl.Traversal{hcl.TraverseAttr{Name: "the_triggers"}},
				)
				if err != nil {
					resp.Error = err
					return
				}
				ntvs = append(ntvs, ntv)
			}
			resp.Traversals = ntvs
		},
	},
}
