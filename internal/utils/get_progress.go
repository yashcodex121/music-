/*
  - This file is part of YukkiMusic.
  - YukkiMusic — A Telegram bot that streams music into group voice chats with seamless playback and control.
  - Copyright (C) 2025 TheTeamVivek
  - Licensed under GNU General Public License v3.0
*/

package utils

import (
	"fmt"
	"math"
	"strings"
	"time"

	"github.com/amarnathcjd/gogram/telegram"
)

func GetProgress(mystic *telegram.NewMessage) *telegram.ProgressManager {
	pm := telegram.NewProgressManager(2)

	if mystic == nil {
		return pm
	}

	var opts *telegram.SendOptions
	if replyMarkup := mystic.ReplyMarkup(); replyMarkup != nil {
		opts = &telegram.SendOptions{ReplyMarkup: *replyMarkup}
	}

	pm.WithCallback(func(pi *telegram.ProgressInfo) {
		text := fmt.Sprintf(
			`🎵 **Downloading Track** 🎵

╭─❥ **Progress**: %.1f%%
├─❥ **Speed**: %s
├─❥ **ETA**: %s
╰─❥ **Elapsed**: %s`,
			pi.Percentage,
			pi.SpeedString(),
			pi.ETAString(),
			pi.ElapsedString(),
		)

		mystic.Edit(text, opts)
	})

	return pm
}

func GetProgressBar(playedSec, durationSec int) string {
	if durationSec == 0 || playedSec <= 0 {
		return "▱▱▱▱▱▱▱▱▱▱"
	}

	percentage := float64(playedSec) / float64(durationSec)
	filled := int(math.Floor(percentage * 10))
	empty := 10 - filled

	filledBar := strings.Repeat("▰", filled)
	emptyBar := strings.Repeat("▱", empty)
	playhead := "🔘"

	return fmt.Sprintf("%s%s%s", filledBar, playhead, emptyBar)
}
