package util

import (
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"fraima.io/fraimmon/internal/dtype"
)

func URLTreatment(uri string) (interface{}, int) {

	var i interface{}

	listUri := strings.Split(uri, "/")

	switch listUri[1] {

	case "value":

		re := regexp.MustCompile(`\/value\/(counter|gauge)\/(\w*)`)
		sliceReg := re.FindStringSubmatch(uri)

		switch sliceReg[1] {

		case "counter":
			var c dtype.Counter
			c.Name = sliceReg[2]
			return c, http.StatusOK

		case "gauge":
			var g dtype.Gauge
			g.Name = sliceReg[2]
			return g, http.StatusOK

		default:
			return i, http.StatusNotImplemented
		}

	case "update":
		re := regexp.MustCompile(`\/update\/(counter|gauge)\/(\w*)\/(\w.*)`)
		sliceReg := re.FindStringSubmatch(uri)

		// Этот кусок тут чисто что бы прошел тест ибо конфликтует с условием ниже
		// /update/unknown/testCounter/100  -хочет 501
		// /update/counter/testCounter/none -хочет 400
		// хотя оба на мой взгляд бед реквест. В любом случае хотел бы глянуть
		// как планировалось.
		testDone := strings.Split(uri, "/")[2]
		switch testDone {
		case "counter":
		case "gauge":
		default:
			return i, http.StatusNotImplemented
		}

		if len(sliceReg) < 4 {
			return i, http.StatusNotFound
		}

		switch sliceReg[1] {

		case "counter":
			var c dtype.Counter

			v, err := strconv.ParseInt(sliceReg[3], 10, 64)

			if err != nil {
				return c, http.StatusBadRequest
			}

			c.Name = sliceReg[2]
			c.Value = v
			return c, http.StatusOK

		case "gauge":
			var g dtype.Gauge

			v, err := strconv.ParseFloat(sliceReg[3], 64)
			if err != nil {
				return g, http.StatusBadRequest
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
