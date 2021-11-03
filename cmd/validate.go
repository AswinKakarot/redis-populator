package cmd

import (
	"fmt"
	"os"

	"github.com/AswinKakarot/redis-populator/pkg/errors"
	"github.com/AswinKakarot/redis-populator/pkg/item"
	"github.com/AswinKakarot/redis-populator/pkg/redisutil"
	"github.com/AswinKakarot/redis-populator/pkg/util"
	"github.com/spf13/cobra"
)

var validateCmd = &cobra.Command{
	Use:   "validate",
	Short: "Validate redis resource",
	Long: `Validate redis standalone/redis cluster with random keys and values
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
		if err := is.Read(true); err != nil {
			is.PrintReadStatus()
			fmt.Printf("redis validation failed with error: %s\n", err)
			os.Exit(errors.POPULATE_ERROR)
		}
		is.PrintReadStatus()
	},
}
