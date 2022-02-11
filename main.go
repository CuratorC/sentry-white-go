package main

import (
	"fmt"
	"github.com/curatorc/cngf/config"
	"github.com/curatorc/cngf/console"
	"github.com/curatorc/cngf/logger"
	"os"
	"os/exec"
	"runtime"
	"sentry-white-go/app/cmd"
	"sentry-white-go/app/cmd/make"
	"sentry-white-go/bootstrap"
	btsConfig "sentry-white-go/config"

	"github.com/spf13/cobra"
)

func init() {
	// 加载 config 目录下的配置信息
	btsConfig.Initialize()
}

func main() {

	// 应用的主入口，默认调用 cmd.CmdServe 命令
	var rootCmd = &cobra.Command{
		Use:   config.Get("app.name"),
		Short: "A simple forum project",
		Long:  `Default will run "serve" command, you can use "-h" flag to see all subcommands`,

		// rootCmd 的所有子命令都会执行以下代码
		PersistentPreRun: func(command *cobra.Command, args []string) {

			// 配置初始化，依赖命令行 --env 参数
			config.InitConfig(cmd.Env)

			// 初始化 Logger
			bootstrap.SetupLogger()

			// 初始化缓存
			bootstrap.SetupCache()
		},
	}

	// 注册子命令
	rootCmd.AddCommand(
		cmd.CmdServe,   // 服务运行
		cmd.CmdKey,     // 生成密钥
		cmd.CmdPlay,    // Play 调试
		make.CmdMake,   // make 命令
		cmd.CmdMigrate, // 数据库迁移
		cmd.CmdDBSeed,  // 数据库填充
		cmd.CmdCache,   // 缓存命令
	)

	// 配置默认运行 Web 服务
	cmd.RegisterDefaultCmd(rootCmd, cmd.CmdServe)

	// 注册全局参数，--env
	cmd.RegisterGlobalFlags(rootCmd)

	err := open("./html/index.html")
	logger.LogIf(err)
	// 执行主命令
	if err := rootCmd.Execute(); err != nil {
		console.Exit(fmt.Sprintf("Failed to run app with %v: %s", os.Args, err.Error()))
	}
}

// open opens the specified URL in the default browser of the user.
func open(url string) error {
	var cmd string
	var args []string

	switch runtime.GOOS {
	case "windows":
		cmd = "cmd"
		args = []string{"/c", "start"}
	case "darwin":
		cmd = "open"
	default: // "linux", "freebsd", "openbsd", "netbsd"
		cmd = "xdg-open"
	}
	args = append(args, url)
	return exec.Command(cmd, args...).Start()
}
