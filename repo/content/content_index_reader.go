package content

import (
	"context"

	"github.com/kopia/kopia/repo/blob"
)

// IndexBlobReader defines content read API.
type IndexBlobReader interface {
	ParseIndexBlob(ctx context.Context, blobID blob.ID) ([]Info, error)
	IndexBlobs(ctx context.Context, includeInactive bool) ([]IndexBlobInfo, error)
}
