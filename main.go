package main

import (
	"flag"
	"log"
	"strings"
	"subtitlesynchandler/srthandler"
)

//// Read file
//func ReadSubs(r io.Reader) (subs []srthandler.Subtitle, err error) {
//	scanner := srthandler.NewScanner(r)
//	for scanner.Scan() {
//		subs = append(subs, scanner.Subtitle())
//	}
//	err = scanner.Err()
//	return
//}
//
//func WriteSubs(w io.Writer, subs []srthandler.Subtitle) error {
//	for _, sub := range subs {
//		_, err := sub.WriteTo(w)
//		if err != nil {
//			return err
//		}
//	}
//	return nil
//}

func runSubtitleHandlerCLI(fileName *string, adbreaks *string, frameRate *string) {
	if strings.Contains(strings.ToLower(*fileName), ".srt") {
		srtH := srthandler.GetSRTHandler(*fileName, *frameRate, strings.Split(*adbreaks, ","))
		err := srtH.ProcessSubs()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

func main() {
	cliMode := flag.Bool("cli", false, "Run in CLI mode")
	fileName := flag.String("subtitlepath", "", "path of the subtitle need to be processed")
	adbreaks := flag.String("ads", "", "comma separated adbreak points in SMPTE format which need to be split")
	frameRate := flag.String("fps", "", "framerate of the subtitle")
	flag.Parse()

	if *cliMode {
		runSubtitleHandlerCLI(fileName, adbreaks, frameRate)
	}

}
