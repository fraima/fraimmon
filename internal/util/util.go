package util

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"fraima.io/fraimmon/internal/types"
)

func UrlTreatment(uri string) (interface{}, int) {

	var i interface{}

	listUri := strings.Split(uri, "/")

	switch listUri[1] {

	case "value":

		re := regexp.MustCompile(`\/value\/(counter|gauge)\/(\w*)`)
		sliceReg := re.FindStringSubmatch(uri)

		switch sliceReg[1] {

		case "counter":
			var c types.Counter
			c.Name = sliceReg[2]
			return c, http.StatusOK

		case "gauge":
			var g types.Gauge
			g.Name = sliceReg[2]
			return g, http.StatusOK

		default:
			return i, http.StatusNotImplemented
		}

	case "update":
		re := regexp.MustCompile(`\/update\/(counter|gauge)\/(\w*)\/(\w.*)`)
		sliceReg := re.FindStringSubmatch(uri)

		switch sliceReg[1] {

		case "counter":
			var c types.Counter

			v, err := strconv.ParseInt(sliceReg[3], 10, 64)

			if err != nil {
				return i, http.StatusNotImplemented
			}

			c.Name = sliceReg[2]
			c.Value = v
			return c, http.StatusOK

		case "gauge":
			var g types.Gauge

			v, err := strconv.ParseFloat(sliceReg[3], 64)
			if err != nil {
				return i, http.StatusNotImplemented
			}

			g.Name = sliceReg[2]
			g.Value = v
			return g, http.StatusOK

		default:
			return i, http.StatusNotImplemented
		}

	default:
		return i, http.StatusNotImplemented
	}

}
