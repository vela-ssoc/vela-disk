package disk

import (
	"github.com/vela-ssoc/vela-kit/vela"
)

var (
	xEnv vela.Environment
)

/*
	local disk = vela.disk.C

*/

func WithEnv(env vela.Environment) {
	xEnv = env
	sum := New()
	xEnv.Set("disk", sum)
}
