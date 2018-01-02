package build

import (
	"fmt"
	"strings"

	"github.com/gobuffalo/buffalo/generators/assets/webpack"
	"github.com/gobuffalo/envy"
	pack "github.com/gobuffalo/packr/builder"
	"github.com/pkg/errors"
)

func (b *Builder) buildAssets() error {
	var err error
	if b.WithWebpack && b.Options.WithAssets {
		envName := envy.Get("GO_ENV", "development")
		fmt.Println("compiling", envName)
		envy.Set("NODE_ENV", envName)
		if envName == "development" {
			err = b.exec(webpack.BinPath)
		} else {
			err = b.exec(webpack.BinPath, "--config", "webpack."+strings.ToLower(envName)+".config.js")
		}

		if err != nil {
			return errors.WithStack(err)
		}
	}

	p := pack.New(b.ctx, b.Root)
	p.Compress = b.Compress

	if !b.Options.WithAssets {
		p.IgnoredBoxes = append(p.IgnoredBoxes, "../public/assets")
	}

	if b.ExtractAssets && b.Options.WithAssets {
		p.IgnoredBoxes = append(p.IgnoredBoxes, "../public/assets")
		err := b.buildExtractedAssets()
		if err != nil {
			return errors.WithStack(err)
		}
	}

	return p.Run()
}
