package timeutils

import (
	"testing"
	"time"
)

func TestConvertToSecs(t *testing.T) {
	type args struct {
		duration time.Duration
	}
	tests := []struct {
		name string
		args args
		want float64
	}{
		{
			name: "test different time code conversion",
			args: args{duration: time.Duration(72000000)},
			want: 0.21599999999999997,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ConvertToSecs(tt.args.duration); got != tt.want {
				t.Errorf("ConvertToSecs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropFrameTimecodeToMilliseconds(t *testing.T) {
	type args struct {
		tc        string
		frameRate float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "test different time code conversion",
			args:    args{tc: "00:10:09:18", frameRate: 25.00},
			want:    609720,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DropFrameTimecodeToMilliseconds(tt.args.tc, tt.args.frameRate)
			if (err != nil) != tt.wantErr {
				t.Errorf("DropFrameTimecodeToMilliseconds() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DropFrameTimecodeToMilliseconds() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDropFrameTimecodeToSeconds(t *testing.T) {
	type args struct {
		tc            string
		frameRate     float64
		dropFrameFlag bool
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "test different time code conversion",
			args:    args{tc: "00:10:09:18", frameRate: 25.00, dropFrameFlag: true},
			want:    609.6089999999999,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DropFrameTimecodeToSeconds(tt.args.tc, tt.args.frameRate, tt.args.dropFrameFlag)
			if (err != nil) != tt.wantErr {
				t.Errorf("DropFrameTimecodeToMillisecondsTest() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("DropFrameTimecodeToMillisecondsTest() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSMPTEPart(t *testing.T) {
	type args struct {
		part string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseSMPTEPart(tt.args.part); got != tt.want {
				t.Errorf("ParseSMPTEPart() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParseSMPTEWithRate(t *testing.T) {
	type args struct {
		tc        string
		frameRate float64
	}
	tests := []struct {
		name    string
		args    args
		want    float64
		wantErr bool
	}{
		{
			name:    "test different time code conversion",
			args:    args{tc: "00:10:09:18", frameRate: 25.00},
			want:    609.720,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseSMPTEWithRate(tt.args.tc, tt.args.frameRate)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseSMPTEWithRate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ParseSMPTEWithRate() got = %v, want %v", got, tt.want)
			}
		})
	}
}
