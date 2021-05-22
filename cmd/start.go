package cmd

import (
	"log"
	"plotcarrier/app"
	"time"

	"github.com/spf13/cobra"
)

var srcDisk string
var dstDisk string
var interval int32

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start moving src plot files to dst",
	Long: `It will compare plot file size between src and dst when file moving completed.
use "plot-carrier start --src src_disk --dst disk --interval 120`,
	Run: func(cmd *cobra.Command, args []string) {
		for {
			err := app.Start(srcDisk, dstDisk)
			if err != nil {
				log.Println(err)
				break
			}
			log.Printf("Sleep %d seconds...", interval)
			time.Sleep(time.Duration(interval) * time.Second)
		}
	},
}

func init() {
	rootCmd.AddCommand(startCmd)

	startCmd.Flags().StringVar(&srcDisk, "src", "", "Src disk")
	startCmd.Flags().StringVar(&dstDisk, "dst", "", "dst disk")
	startCmd.Flags().Int32VarP(&interval, "interval", "t", 120, "seconds of scan interval, default is 120 seconds")
	startCmd.MarkFlagRequired("src")
	startCmd.MarkFlagRequired("dst")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// startCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
