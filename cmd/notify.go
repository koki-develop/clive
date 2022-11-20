package cmd

import (
	"fmt"
	"io"
	"time"

	"github.com/charmbracelet/lipgloss"
	"github.com/koki-develop/clive/pkg/cache"
	"github.com/koki-develop/clive/pkg/github"
	"github.com/koki-develop/clive/pkg/styles"
	"github.com/koki-develop/clive/pkg/util"
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

	var release github.Release
	if c != nil && !c.Expired() {
		// TBD: Might be too noisy to notify every time.
		// c.Bind(&release)
		return nil
	} else {
		cl := github.New()
		r, err := cl.GetLatestRelease("koki-develop", "clive")
		if err != nil {
			return err
		}
		release = *r
		if err := s.Set("release", release); err != nil {
			return err
		}
	}

	if util.Version(release.Name).Newer(util.Version(version)) {
		txt := styles.StyleNotificationText.Render(fmt.Sprintf("A new version (%s) is available!", release.Name))
		link := fmt.Sprintf("See: %s", styles.StyleLink.Render(fmt.Sprintf("https://github.com/koki-develop/clive/releases/%s", release.Name)))
		n := lipgloss.JoinVertical(lipgloss.Center, txt, link)
		n = util.Border(n, styles.StyleNotificationBorder)

		_, _ = fmt.Fprintln(w, n)
	}

	return nil
}
