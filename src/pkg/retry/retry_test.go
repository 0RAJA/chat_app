package retry

import (
	"fmt"
	"log"
	"testing"
	"time"

	"github.com/0RAJA/Rutils/pkg/utils"
	"github.com/stretchr/testify/require"
)

func TestRetry(t *testing.T) {
	t.Parallel()
	end := int(utils.RandomInt(1, 5))
	start := 0
	goodFunc := func() error {
		start++
		if start == end {
			return nil
		}
		return fmt.Errorf("%d", start)
	}
	log.Println("start")
	report := <-NewTry("test", goodFunc, 100*time.Millisecond, end).Run()
	require.True(t, report.Result)
	require.Equal(t, report.Times, end)
	require.Len(t, report.Errs, end-1)
	for i, v := range report.Errs {
		require.Equal(t, fmt.Sprintf("%d", i+1), v.Error())
	}
}
