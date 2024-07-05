package handler

import (
	"bytes"
	"compress/flate"
	"compress/gzip"
	"context"
	net_http "net/http"
	"strings"

	utils_err "github.com/Rochakb/go-stater-project/utils/err"

	"github.com/go-kit/kit/transport/http"
	"github.com/pkg/errors"
)

var (
	errBadRequest     = errors.New("bad request")
	errInternalServer = errors.New("internal server error")
)

func errorWriter(cx context.Context, err error, code int, w net_http.ResponseWriter) {
	e := utils_err.NewError(
		err, code, errors.Cause(err).Error(),
	)

	w.WriteHeader(int(e.Code()))

	writejs := func(e *utils_err.Error, w net_http.ResponseWriter) {
		// write default
		js, er := e.JSON()
		if er != nil {
			panic(er)
		}

		w.Write(js)
	}

	v, ok := cx.Value(http.ContextKeyRequestAccept).(string)
	if !ok {
		writejs(e, w)
		return
	}

	switch {
	case v == "":
		writejs(e, w)
	case strings.Contains(v, "application/json"):
		writejs(e, w)
	default:
		writejs(e, w)
	}
}

func errorEncoder(
	cx context.Context, e error, w net_http.ResponseWriter,
) {
	switch errors.Cause(e) {
	case errBadRequest:
		errorWriter(cx, e, net_http.StatusBadRequest, w)
	case errInternalServer:
		errorWriter(cx, e, net_http.StatusInternalServerError, w)
	default:
		errorWriter(cx, e, net_http.StatusInternalServerError, w)
	}
}

func gzipCompress(b []byte) ([]byte, error) {
	var buf bytes.Buffer
	gw := gzip.NewWriter(&buf)

	_, err := gw.Write(b)
	if err != nil {
		return nil, err
	}

	if err := gw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func deflateCompress(b []byte) ([]byte, error) {
	var buf bytes.Buffer

	fw, err := flate.NewWriter(&buf, 5)
	if err != nil {
		return nil, err
	}

	_, err = fw.Write(b)
	if err != nil {
		return nil, err
	}

	if err := fw.Close(); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
