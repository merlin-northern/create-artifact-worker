// Copyright 2020 Northern.tech AS
//
//    Licensed under the Apache License, Version 2.0 (the "License");
//    you may not use this file except in compliance with the License.
//    You may obtain a copy of the License at
//
//        http://www.apache.org/licenses/LICENSE-2.0
//
//    Unless required by applicable law or agreed to in writing, software
//    distributed under the License is distributed on an "AS IS" BASIS,
//    WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//    See the License for the specific language governing permissions and
//    limitations under the License.
package cmd

import (
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/mendersoftware/create-artifact-worker/config"
	mlog "github.com/mendersoftware/create-artifact-worker/log"
)

var rootCmd = &cobra.Command{
	Use:   "create-artifact",
	Short: "Artifact generator CLI.",
	Long: "\nSupports the following env vars:\n\n" +
		"CREATE_ARTIFACT_VERBOSE enable verbose logging (default: false)\n",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		mlog.Error(err.Error())
		os.Exit(1)
	}
}

func init() {
	rootCmd.AddCommand(singleFileCmd)

	config.Init()
	mlog.Init(viper.GetBool(config.CfgVerbose))
}
