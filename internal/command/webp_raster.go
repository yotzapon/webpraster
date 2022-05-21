package command

import (
	"errors"
	"github.com/spf13/cobra"
	"webptoraster/internal/services"
)

func configureWebpRasterCommand(command *cobra.Command) {
	rootCommand := &cobra.Command{
		Use:   "webp",
		Short: "convert webp image type to other raster type (jpeg, png)",
	}

	pathSubCommand := &cobra.Command{
		Use:   "jpeg {path of webp images}",
		Short: "convert webp images to jpegs by provided path",
		RunE:  convertToJpeg,
	}

	command.AddCommand(rootCommand)
	rootCommand.AddCommand(pathSubCommand)
}

func convertToJpeg(cmd *cobra.Command, args []string) error {
	if len(args) == 0 || len(args) > 1 {
		return errors.New("too may arguments")
	}

	services.DirWalk(args[0])
	return nil
}
