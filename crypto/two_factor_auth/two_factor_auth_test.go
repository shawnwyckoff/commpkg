package two_factor_auth

import (
	"fmt"
	"github.com/shawnwyckoff/commpkg/apputil/test"
	"testing"
)

func TestTwoFactorAuth(t *testing.T) {
	pwd, secondsRemaining, err := TwoFactorAuth("nzxxiidbebvwk6jb")
	test.Assert(t, err)
	if secondsRemaining < 0 || secondsRemaining > 30 {
		test.PrintlnExit(t, fmt.Sprintf("invalid secondsRemaining %d", secondsRemaining))
	}
	if len(pwd) != 6 {
		test.PrintlnExit(t, fmt.Sprintf("invalid 2FA password %s", pwd))
	}
}
