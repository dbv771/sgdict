package cmd

import (
	"time"

	"github.com/mritd/sgdict/pkg/wordrank"

	"github.com/spf13/cobra"
)

var wordrankCmd = &cobra.Command{
	Use:   "wordrank",
	Short: "update word rank(use baidu)",
	Run: func(cmd *cobra.Command, args []string) {
		wordrank.BaiduWorkRank()
	},
}

func init() {
	rootCmd.AddCommand(wordrankCmd)
	wordrankCmd.PersistentFlags().StringVar(&wordrank.FilePath, "file", "", "rime dict file")
	wordrankCmd.PersistentFlags().StringVar(&wordrank.Proxy, "proxy", "", "http client proxy")
	wordrankCmd.PersistentFlags().IntVar(&wordrank.PoolSize, "poolsize", 5, "goroutine pool size")
	wordrankCmd.PersistentFlags().DurationVar(&wordrank.Timeout, "timeout", 10*time.Second, "http client timeout")
	wordrankCmd.PersistentFlags().IntVar(&wordrank.RetryCount, "retry", 5, "auto retry count")
	wordrankCmd.PersistentFlags().DurationVar(&wordrank.RetryMaxWaitTime, "retrywaittime", 5*time.Second, "retry max wait time")
}
