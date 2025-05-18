package cmd

import (
	"context"
	"fmt"
	"io"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/google/go-github/v72/github"
	"github.com/koki-develop/clive/internal/cache"

	"github.com/koki-develop/clive/internal/styles"
	"github.com/koki-develop/clive/internal/util"
)

func notifyNewRelease(w io.Writer) error {
	s, err := cache.NewStore(12 * time.Hour)
	if err != nil {
		return err
	}

	c, err := s.Get("release")
	if err != nil {
		return err
	}

	var release github.RepositoryRelease
	if c != nil && !c.Expired() {
		// If the cache exists and has not expired, no notification is given.
		return nil
	} else {
		// Retrieve the latest release and save cache.
		cl := github.NewClient(nil)
		r, _, err := cl.Repositories.GetLatestRelease(context.Background(), "koki-develop", "clive")
		if err != nil {
			return err
		}
		release = *r
		if err := s.Set("release", map[string]string{"name": *release.Name}); err != nil {
			return err
		}
	}

	// If a newer version is released, notify it.
	if util.Version(*release.Name).Newer(util.Version(version)) {
		txt := styles.StyleNotificationText.Render(fmt.Sprintf("A new version (%s) is available!", *release.Name))
		link := fmt.Sprintf("See: %s", styles.StyleLink.Render(fmt.Sprintf("https://github.com/koki-develop/clive/releases/%s", *release.Name)))
		n := lipgloss.JoinVertical(lipgloss.Center, txt, link)
		n = util.Border(n, styles.StyleNotificationBorder)

		_, _ = fmt.Fprintln(w, n)
	}

	return nil
}
