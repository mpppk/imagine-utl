package cmd

import (
	"encoding/json"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/mpppk/imagine/domain/model"
	"github.com/mpppk/imagine/util"

	"github.com/mpppk/imagine-utl/cmd/option"
	"github.com/spf13/afero"

	"github.com/spf13/cobra"
)

func pathToTagNames(p string, depth int) []string {
	list := strings.Split(filepath.ToSlash(filepath.Dir(p)), "/")
	if len(list) == 0 {
		return []string{}
	}
	var tagNames []string
	ignoreList := []string{".", "..", "/"}

	isTagName := func(tagName string, ignoreList []string) bool {
		for _, ignoreC := range ignoreList {
			if ignoreC == tagName {
				return false
			}
		}
		return true
	}

	for _, tagName := range list[len(list)-depth:] {
		if isTagName(tagName, ignoreList) {
			tagNames = append(tagNames, tagName)
		}
	}
	return tagNames
}

func newLoadCmd(fs afero.Fs) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "load",
		Short: "Load assets",
		Long:  ``,
		RunE: func(cmd *cobra.Command, args []string) error {
			conf, err := option.NewLoadCmdConfigFromViper(args)
			if err != nil {
				return err
			}
			for p := range util.LoadImagesFromDir(conf.Dir, 10000) {
				//parentDir, ok := toParentDirName(p)
				asset := model.NewImportAssetFromFilePath(p)
				p, err := filepath.Rel(conf.Dir, asset.Path)
				if err != nil {
					return fmt.Errorf("failed to extract asset path: %w", err)
				}
				asset.Path = p

				tagNames := pathToTagNames(asset.Path, int(conf.Depth))

				var boxes []*model.ImportBoundingBox
				for _, name := range tagNames {
					boxes = append(boxes, &model.ImportBoundingBox{
						TagName: name,
					})
				}
				asset.BoundingBoxes = boxes

				j, err := json.Marshal(asset)
				if err != nil {
					return fmt.Errorf("failed to marshal asset to json: %w", err)
				}

				cmd.Println(string(j))
			}

			return nil
		},
	}

	registerFlags := func(cmd *cobra.Command) error {
		flags := []option.Flag{
			&option.StringFlag{
				BaseFlag: &option.BaseFlag{
					Name:       "dir",
					IsRequired: true,
					Usage:      "target directory",
				},
				IsDirName: true,
			},
			&option.UintFlag{
				BaseFlag: &option.BaseFlag{
					Name:  "depth",
					Usage: "depth of directory name used as tag name",
				},
			},
		}
		return option.RegisterFlags(cmd, flags)
	}

	if err := registerFlags(cmd); err != nil {
		return nil, err
	}

	return cmd, nil
}

func init() {
	cmdGenerators = append(cmdGenerators, newLoadCmd)
}
