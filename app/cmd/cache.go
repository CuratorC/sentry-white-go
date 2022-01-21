package cmd

import (
	"fmt"
	"github.com/curatorc/cngf/cache"
	"github.com/curatorc/cngf/console"
	"github.com/curatorc/cngf/logger"

	"github.com/spf13/cobra"
)

var CmdCache = &cobra.Command{
	Use:   "cache",
	Short: "Cache management",
}

var CmdCacheClear = &cobra.Command{
	Use:   "clear",
	Short: "clear cache",
	Run:   runCacheClear,
}

var CmdCacheForget = &cobra.Command{
	Use:   "forget",
	Short: "Delete redis key, example: cache forget cache-key",
	Run:   runCacheForget,
}

// forget 命令的选项
var cacheKey string

func init() {
	// 注册 cache 命令的子命令
	CmdCache.AddCommand(CmdCacheClear, CmdCacheForget)

	// 设置 cache forget 命令的选项
	CmdCacheForget.Flags().StringVarP(&cacheKey, "key", "k", "", "Key of the cache")
	err := CmdCacheForget.MarkFlagRequired("key")
	if err != nil {
		logger.ErrorString("cache.init", "MarkFlagRequired", err.Error())
		return
	}
}

func runCacheClear(cmd *cobra.Command, args []string) {
	cache.Flush()
	console.Success("Cache cleared.")
}

func runCacheForget(cmd *cobra.Command, args []string) {
	cache.Forget(cacheKey)
	console.Success(fmt.Sprintf("Cache key [%s] deleted.", cacheKey))
}
