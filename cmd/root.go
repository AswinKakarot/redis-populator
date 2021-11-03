package cmd

import (
	"flag"
	"fmt"
	"os"
	"sync"

	"github.com/AswinKakarot/redis-populator/pkg/config"
	"github.com/AswinKakarot/redis-populator/pkg/errors"
	"github.com/AswinKakarot/redis-populator/pkg/item"
	"github.com/AswinKakarot/redis-populator/pkg/redisutil"
	"github.com/AswinKakarot/redis-populator/pkg/util"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

var rootCmd = &cobra.Command{
	Use:   "redis-populator",
	Short: "Populates and validates data on redis standalone instance and clusters",
	Long: `Redis Populator is highly customizable random data injector for redis
It supports both redis standalone instance and redis clusters
`,
	Run: func(cmd *cobra.Command, args []string) {
		var (
			fs  *util.LocalStore
			err error
			is  *item.ItemStore
		)
		if util.LocalStoreEnabled() {
			fs, err = util.GetLocalStore()
			if err != nil {
				fmt.Printf("unable to initilize local file store with error: %s\n", err)
				os.Exit(errors.FS_ERROR)
			}
		}
		if redisutil.GetRedisConfig().IsCluster {
			o, err := redisutil.NewCluster()
			if err != nil {
				fmt.Printf("unable to establish redis cluster connection with error: %s\n", err)
				os.Exit(errors.REDIS_ERROR)
			}
			is = item.NewItemStore(fs, o)
		} else {
			o, err := redisutil.NewStandalone()
			if err != nil {
				fmt.Printf("unable to establish redis connection with error: %s\n", err)
				os.Exit(errors.REDIS_ERROR)
			}
			is = item.NewItemStore(fs, o)
		}
		wg := new(sync.WaitGroup)
		wg.Add(2)
		var errChan = make(chan errors.ErrorChannel, 1)

		go func() {
			defer wg.Done()
			err := is.Write()
			errChan <- errors.ErrorChannel{
				Type:   "write",
				Errors: err,
			}
		}()

		go func() {
			defer wg.Done()
			err := is.Read(false)
			errChan <- errors.ErrorChannel{
				Type:   "read",
				Errors: err,
			}
		}()

		chanErr := <-errChan
		if chanErr.Type == "write" {
			if chanErr.Errors != nil {
				fmt.Printf("redis population failed with error: %s\n", chanErr.Errors)
				is.PrintWriteStatus()
			}
		}
		if chanErr.Type == "read" {
			if chanErr.Errors != nil {
				fmt.Printf("redis validation failed with error: %s\n", chanErr.Errors)
				is.PrintReadStatus()
			}

		}
		wg.Wait()
		is.PrintWriteStatus()
		is.PrintReadStatus()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

}

func init() {
	// Command Line Flags initialization
	pflag.CommandLine.AddFlagSet(config.FlagSet())
	pflag.CommandLine.AddFlagSet(util.FlagSet())
	pflag.CommandLine.AddFlagSet(redisutil.FlagSet())
	pflag.CommandLine.AddFlagSet(item.FlagSet())
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	rootCmd.AddCommand(populateCmd)
	rootCmd.AddCommand(validateCmd)
}
