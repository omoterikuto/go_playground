package main

import (
	"encoding/base64"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/proto"
)

func main() {
	aperr := &AppError{
		code:     CodeMaintenance,
		grpcCode: 14,
		msg:      "message",
	}
	st, _ := appErrorToStatus(aperr)

	if p := st.Proto(); p != nil && len(p.Details) > 0 {
		stBytes, err := proto.Marshal(p)
		if err != nil {
			panic(err)
		}

		fmt.Println(base64.RawStdEncoding.EncodeToString(stBytes))
		// CA4SB21lc3NhZ2UaPwoodHlwZS5nb29nbGVhcGlzLmNvbS9nb29nbGUucnBjLkVycm9ySW5mbxITCgYwMDAwMjASCW5vdGFob3RlbA
	}
}

type AppError struct {
	code     Code
	grpcCode codes.Code
	msg      string
	msgJa    string
}

type Code string

const (
	CodeMaintenance Code = "000020"
)

func appErrorToStatus(apperr *AppError) (*status.Status, error) {
	s := apperr.GRPCStatus()
	s2, err := s.WithDetails(
		&errdetails.ErrorInfo{Reason: string(apperr.code), Domain: "notahotel"},
	)
	if err != nil {
		return nil, err
	}
	return s2, nil
}

func (e *AppError) GRPCStatus() *status.Status {
	return status.New(e.grpcCode, e.msg)
}
