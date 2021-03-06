package holder

import (
	"github.com/ecnuvj/vhoj_common/pkg/common/constants/remote_oj"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/querier"
	"github.com/ecnuvj/vhoj_submitter/pkg/remote/querier/oj"
)

var Queriers map[remote_oj.RemoteOJ]querier.IQuerier

func init() {
	Queriers = make(map[remote_oj.RemoteOJ]querier.IQuerier)
	Queriers[remote_oj.HDU] = oj.HduQuerier
	Queriers[remote_oj.POJ] = oj.PojQuerier
}
