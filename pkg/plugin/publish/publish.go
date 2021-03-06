package publish

import (
	"github.com/drone/drone/pkg/build/buildfile"
)

// Publish stores the configuration details
// for publishing build artifacts when
// a Build has succeeded
type Publish struct {
	S3 *S3 `yaml:"s3,omitempty"`
}

func (p *Publish) Write(f *buildfile.Buildfile) {
	if p.S3 != nil {
		p.S3.Write(f)
	}
}
