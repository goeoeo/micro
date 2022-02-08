package app

import "fmt"

var (
	AppName      string // 应用名称
	AppVersion   string // 应用版本
	BuildVersion string // 编译版本
	BuildTime    string // 编译时间
	GitRevision  string // Git版本
	GitBranch    string // Git分支
	GoVersion    string // Golang信息
)

func Version() (v string) {
	v += fmt.Sprintf("_________________________________\n")
	v += fmt.Sprintf("App Name:\t%s\n", AppName)
	v += fmt.Sprintf("App Version:\t%s\n", AppVersion)
	v += fmt.Sprintf("Build version:\t%s\n", BuildVersion)
	v += fmt.Sprintf("Build time:\t%s\n", BuildTime)
	v += fmt.Sprintf("Git revision:\t%s\n", GitRevision)
	v += fmt.Sprintf("Git branch:\t%s\n", GitBranch)
	v += fmt.Sprintf("Golang Version: %s\n", GoVersion)
	v += fmt.Sprintf("_________________________________\n")
	return
}

func VersionColumns() (columns [][]string) {
	return [][]string{
		{"应用名称", AppName},
		{"应用版本", AppVersion},
		{"编译版本", BuildVersion},
		{"编译时间", BuildTime},
		{"Git版本", GitRevision},
		{"Git分支", GitBranch},
		{"Golang信息", GoVersion},
	}
}
