package timeutils

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func DropFrameTimecodeToMilliseconds(tc string, frameRate float64) (float64, error) {
	parts := strings.Split(tc, ":")
	if len(parts) != 4 {
		return 0, fmt.Errorf("invalid timecode format: %s", tc)
	}

	hours := ParseSMPTEPart(parts[0])
	minutes := ParseSMPTEPart(parts[1])
	seconds := ParseSMPTEPart(parts[2])
	frames := ParseSMPTEPart(parts[3])

	// Calculate total frames
	totalFrames := (hours * 3600 * int(frameRate)) + (minutes * 60 * int(frameRate)) + (seconds * int(frameRate)) + frames
	//totalFrames := (hours * 3600) + (minutes * 60) + (seconds) + frames

	// Adjust for drop frame
	droppedFrames := 0
	if frameRate == 29.97 {
		droppedFrames = 2 * (minutes / 10)
	}
	totalFrames -= droppedFrames

	// Convert frames to milliseconds
	milliseconds := float64(totalFrames) * (1000.0 / frameRate)

	return milliseconds, nil
}

func DropFrameTimecodeToSeconds(tc string, frameRate float64, dropFrameFlag bool) (float64, error) {
	parts := strings.Split(tc, ":")
	if len(parts) != 4 {
		return 0, fmt.Errorf("invalid timecode format: %s", tc)
	}

	hours := ParseSMPTEPart(parts[0])
	minutes := ParseSMPTEPart(parts[1])
	seconds := ParseSMPTEPart(parts[2])
	frames := ParseSMPTEPart(parts[3])

	// Calculate total frames
	tmpminutes := float64((hours * 60) + minutes)
	minutesCounted := tmpminutes - math.Floor(tmpminutes/10.00)
	tmpSeconds := tmpminutes*60 + float64(seconds)
	tmpframes := float64(frames) + math.Ceil(tmpSeconds*frameRate)
	//// Adjust for drop frame

	if dropFrameFlag {
		if frames < 2 && minutesCounted > 0 && seconds == 0 {
			tmpframes = tmpframes + (2 - float64(frames))
		}
		tmpframes = tmpframes - (minutesCounted * 2)
	}
	//totalFrames -= droppedFrames

	// Convert frames to milliseconds
	//milliseconds := float64(totalFrames) * (1000.0 / frameRate)
	res := (tmpframes / frameRate) * 1.001
	return res, nil
}

func ParseSMPTEWithRate(tc string, frameRate float64) (float64, error) {
	// Regular expression to match SMPTE timecode
	re := regexp.MustCompile(`(\d+):(\d+):(\d+):(\d+)`)
	matches := re.FindStringSubmatch(tc)
	if len(matches) != 5 {
		return 0, fmt.Errorf("invalid SMPTE timecode format: %s", tc)
	}

	// Parse hours, minutes, seconds, and frames
	hours, _ := strconv.Atoi(matches[1])
	minutes, _ := strconv.Atoi(matches[2])
	seconds, _ := strconv.Atoi(matches[3])
	frames, _ := strconv.Atoi(matches[4])

	// Calculate total frames
	totalFrames := (hours * 3600 * int(frameRate)) + (minutes * 60 * int(frameRate)) + (seconds * int(frameRate)) + frames

	// Convert frames to seconds
	secondsValue := float64(totalFrames) / frameRate

	return secondsValue, nil
}

func ParseSMPTEPart(part string) int {
	var value int
	_, err := fmt.Sscanf(part, "%02d", &value)
	if err != nil {
		return -1
	}
	return value
}

func ConvertToSecs(duration time.Duration) float64 {
	return duration.Hours()*3600 + duration.Minutes()*60 + duration.Seconds()
}
