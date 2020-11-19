/*
Copyright © 2020 Anand Babu Periasamy https://twitter.com/abperiasamy

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program. If not, see <http://www.gnu.org/licenses/>.
*/
package cmd

import (
	"strconv"
)

const (
	gWhitePrompt string = "█"
	gBlackPrompt string = "░"
)

// White prompt
func whitePrompt() string {
	// ASCII prompt
	if gNoColor {
		return "W" + " " + strconv.Itoa(gMoveCount) + " "
	}

	// Unicode prompt
	if gLightBg { // Invert on light background.
		return gBlackPrompt + " " + strconv.Itoa(gMoveCount) + " "
	}
	return gWhitePrompt + " " + strconv.Itoa(gMoveCount) + " "
}

// Black prompt
func blackPrompt() string {
	defer func() { gMoveCount += 1 }() // Count the nth move.

	// ASCII prompt
	if gNoColor {
		return "B" + " " + strconv.Itoa(gMoveCount) + " "
	}

	// Unicode prompt
	if gLightBg { // Invert on light background.
		return gWhitePrompt + " " + strconv.Itoa(gMoveCount) + " "
	}
	return gBlackPrompt + " " + strconv.Itoa(gMoveCount) + " "
}

// Engine's shell prompt
func enginePrompt() string {
	if gNoColor {
		if gHumanIsBlack {
			return whitePrompt() + ":] "
		}
		return blackPrompt() + ":] "
	} else {
		if gHumanIsBlack {
			return whitePrompt() + "🤖 "
		}
		return blackPrompt() + "🤖 "
	}
}

// Human's shell prompt
func humanPrompt() string {
	if gNoColor {
		if gHumanIsBlack {
			return blackPrompt() + ":) "
		}
		return whitePrompt() + ":) "
	} else {
		if gHumanIsBlack {
			return blackPrompt() + "🙇 "
		}
		return whitePrompt() + "🙇 "
	}
}
