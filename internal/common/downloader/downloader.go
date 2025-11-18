package downloader

import (
	"context"
	"io"

	"github.com/Minecraft-Unified-Hub-Team/RedstonePanel/internal/domain"
)

// Downloader fetches content via HTTP(S).
type Downloader interface {
	// Get returns the response body stream. Callers MUST close the reader.
	Get(ctx context.Context, url domain.URL) (io.ReadCloser, error)
}
