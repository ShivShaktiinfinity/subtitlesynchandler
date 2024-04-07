# This code will split the SRT subtitles based on adbreak points
Sample run.
```bash
./subtitlesynchandler --cli=true --subtitlepath=testsamples/test.srt --ads="00:10:11:01,00:15:10:02,00:20:10:02" --fps=29.97
```
We will add more features like increment the cues based on given time in seconds.
