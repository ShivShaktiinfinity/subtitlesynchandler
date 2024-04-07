package srthandler

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
	"subtitlesynchandler/timeutils"
	"time"
)

type SRTHandlerClient struct {
	InputFileName string   `json:"input_file_name"`
	FrameRate     string   `json:"frame_rate"`
	Adbreaks      []string `json:"adbreaks"`
}

func GetSRTHandler(InputFileName string, FrameRate string, Adbreaks []string) *SRTHandlerClient {
	return &SRTHandlerClient{InputFileName: InputFileName, FrameRate: FrameRate, Adbreaks: Adbreaks}
}

func (c *SRTHandlerClient) ReadSubs() (subs []Subtitle, err error) {
	file, err := os.Open(c.InputFileName)
	if err != nil {
		log.Fatal(err)
		return
	}
	scanner := NewScanner(bufio.NewReader(file))
	for scanner.Scan() {
		subs = append(subs, scanner.Subtitle())
	}
	err = scanner.Err()
	return
}

func (c *SRTHandlerClient) WriteSubs(w io.Writer, subs []Subtitle) error {
	for _, sub := range subs {
		_, err := sub.WriteTo(w)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *SRTHandlerClient) ProcessSubs() error {

	srt, err := c.ReadSubs()
	if err != nil {
		log.Fatal(err)
		return err
	}

	fps, err := strconv.ParseFloat(c.FrameRate, 64)
	if err != nil {
		log.Fatal(err)
		return err
	}

	for i := 0; i < len(c.Adbreaks); i++ {
		dropFrameFlag := false
		if strings.Contains(c.Adbreaks[i], ";") {
			dropFrameFlag = true
		}

		admilliseconds, err := timeutils.DropFrameTimecodeToSeconds(strings.ReplaceAll(c.Adbreaks[i], ";", ":"), fps, dropFrameFlag)
		if err != nil {
			log.Fatal(err)
			return err
		}
		log.Println(admilliseconds)
		var finalSrt []Subtitle
		counter := 1
		for _, st := range srt {
			adSecs := (admilliseconds)
			if adSecs >= st.Start.Seconds() && adSecs <= st.End.Seconds() {
				fmt.Println(adSecs, st.Start.Seconds(), st.End.Seconds())
				tmpEnd := st.End
				st.End = time.Duration(adSecs * float64(time.Second))
				counter = counter + 1
				tmp := Subtitle{
					Number: counter,
					Start:  time.Duration((adSecs + 1.000) * float64(time.Second)),
					End:    tmpEnd,
					Text:   st.Text,
					Bounds: st.Bounds,
				}
				//fmt.Printf("%.3f %.3f\n", timeutils.ConvertToSecs(st.Start), timeutils.ConvertToSecs(st.End))
				finalSrt = append(finalSrt, st)
				finalSrt = append(finalSrt, tmp)
				fmt.Println(st, tmp)
			} else {
				st.Number = counter
				finalSrt = append(finalSrt, st)
			}
			counter = counter + 1
		}
		srt = finalSrt
	}

	//for _, tst := range finalSrt {
	//	fmt.Println(tst)
	//}
	err = os.Mkdir("Updated", 0750)
	if err != nil {
		return err
	}
	fileWrite, err := os.Create("Updated/" + path.Base(c.InputFileName))
	if err != nil {
		log.Fatal(err)
		return err
	}

	err = c.WriteSubs(fileWrite, srt)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
