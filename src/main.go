package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-lambda-go/lambda"
	"time"
)

type Times struct {
	UTC         time.Time
	ANSIC       string
	UnixDate    string
	RubyDate    string
	RFC822      string
	RFC822Z     string
	RFC850      string
	RFC1123     string
	RFC1123Z    string
	RFC3339     string
	RFC3339Nano string
	Kitchen     string
	Stamp       string
	StampMilli  string
	StampMicro  string
	StampNano   string
	Unix        int64
	UnixNano    int64
	ISOWeek     string
}

func HandleRequest(ctx context.Context) (Times, error) {
	t := time.Now().UTC()
	thisYear, thisWeek := t.ISOWeek()
	ts := Times{
		UTC:         t,
		ANSIC:       t.Format(time.ANSIC),
		UnixDate:    t.Format(time.UnixDate),
		RubyDate:    t.Format(time.RubyDate),
		RFC822:      t.Format(time.RFC822),
		RFC822Z:     t.Format(time.RFC822Z),
		RFC850:      t.Format(time.RFC850),
		RFC1123:     t.Format(time.RFC1123),
		RFC1123Z:    t.Format(time.RFC1123Z),
		RFC3339:     t.Format(time.RFC3339),
		RFC3339Nano: t.Format(time.RFC3339Nano),
		Kitchen:     t.Format(time.Kitchen),
		Stamp:       t.Format(time.Stamp),
		StampMilli:  t.Format(time.StampMilli),
		StampMicro:  t.Format(time.StampMicro),
		StampNano:   t.Format(time.StampNano),
		Unix:        t.Unix(),
		UnixNano:    t.UnixNano(),
		ISOWeek:     fmt.Sprintf("%d-W%02d", thisYear, thisWeek),
	}
	return ts, nil
}

func main() {
	lambda.Start(HandleRequest)
}
