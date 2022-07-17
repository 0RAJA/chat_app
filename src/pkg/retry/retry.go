package retry

import (
	"fmt"
	"time"
)

// Try 尝试重试函数，如果函数执行失败，则延长时间重试
// name: 重试任务名称, F: 需要重试的函数, Duration 默认延迟重试时间, MaxTimes 最大重试次数
type Try struct {
	Name     string
	F        func() error
	Duration time.Duration
	MaxTimes int
}

func NewTry(name string, f func() error, d time.Duration, maxTimes int) *Try {
	return &Try{
		Name:     name,
		F:        f,
		Duration: d,
		MaxTimes: maxTimes,
	}
}

// Report 尝试重试的报告
// Result为函数执行结果，Times为尝试次数, Duration为函数执行时间，SumDuration为总执行时间, Errs为函数执行错误记录
type Report struct {
	Name        string
	Result      bool
	Times       int
	SumDuration time.Duration
	Errs        []error
}

func (report *Report) Error() string {
	return fmt.Sprintf("[retry]名称:%s,结果:%v,尝试次数:%v,总时间:%v,错误:%v", report.Name, report.Result, report.Times, report.SumDuration, report.Errs)
}

// Run 开始尝试重试，返回chan可以用于接收尝试报告
func (try Try) Run() <-chan Report {
	result := make(chan Report, 1)
	go func() {
		defer close(result)
		start := time.Now()
		var errs []error
		for i := 0; i < try.MaxTimes; i++ {
			time.Sleep(try.Duration)
			err := try.F()
			if err == nil {
				result <- Report{
					Name:        try.Name,
					Result:      true,
					Times:       i + 1,
					SumDuration: time.Since(start),
					Errs:        errs,
				}
				return
			}
			errs = append(errs, err)
		}
		result <- Report{
			Name:        try.Name,
			Result:      false,
			Times:       try.MaxTimes,
			SumDuration: time.Since(start),
			Errs:        errs,
		}
	}()
	return result
}
