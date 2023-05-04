package pkg

import (
	"reflect"
	"testing"
)

func TestParseConfig(t *testing.T) {
	configInvalidLocation := []string{"../../configs/no.config.yaml"}
	configInvalidConfigFileLocation := []string{"../../configs/bad.yaml"}
	configSampleLocation := []string{"../../configs/sample.yaml"}
	configMultipleUserLocation := []string{"../../configs/multiple.user.yaml", "../../configs/multiple.user2.yaml"}
	expectedSampleAuth := Authn{
		[]User{
			{
				"Grafana",
				"Loki",
				"tenant-1",
			},
		},
	}
	expectedMultipleUserAuth := Authn{
		[]User{
			{
				"User-a",
				"pass-a",
				"tenant-a",
			},
			{
				"User-b",
				"pass-b",
				"tenant-b",
			},
			{
				"User-c",
				"pass-c",
				"tenant-c",
			},
			{
				"User-d",
				"pass-d",
				"tenant-d",
			},
		},
	}
	type args struct {
		location *[]string
	}
	tests := []struct {
		name    string
		args    args
		want    *Authn
		wantErr bool
	}{
		{
			"Basic",
			args{
				&configSampleLocation,
			},
			&expectedSampleAuth,
			false,
		}, {
			"Multiples users",
			args{
				&configMultipleUserLocation,
			},
			&expectedMultipleUserAuth,
			false,
		}, {
			"Invalid location",
			args{
				&configInvalidLocation,
			},
			nil,
			true,
		}, {
			"Invalid yaml file",
			args{
				&configInvalidConfigFileLocation,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseConfig(tt.args.location)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseConfig() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseConfig() = %v, want %v", got, tt.want)
			}
		})
	}
}
