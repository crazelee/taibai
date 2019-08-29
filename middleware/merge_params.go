package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	TotalParamsKey = "totalParams"
)

// MergeParams merges query/body into a map
func MergeParams(c *gin.Context) {

	params := getParamsFromBody(c.Request)
	for k, v := range getParamsFromJsonBody(c.Request) {
		params[k] = v
	}
	for k, v := range getParamsFromQuery(c.Request) {
		params[k] = v
	}

	b, err := json.Marshal(params)
	if err != nil {
		panic(err)
	}
	c.Header("Content-Type", "application/json;charset=utf-8")
	c.Set(TotalParamsKey, bytes.NewReader(b))
	c.Next()

}

func getParamsFromBody(r *http.Request) map[string]interface{} {
	var params map[string]interface{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&params); err != nil {
		params = map[string]interface{}{}
	}
	return params
}

func getParamsFromQuery(r *http.Request) map[string]interface{} {
	params := make(map[string]interface{})
	for key, val := range r.URL.Query() {
		if len(val) == 1 {
			var p map[string]interface{}
			if err := json.Unmarshal([]byte(val[0]), &p); err != nil {
				params[key] = val[0]
			} else {
				params[key] = p
			}
		} else if len(val) > 1 {
			params[key] = val
		}
	}
	return params
}

func getParamsFromJsonBody(r *http.Request) map[string]interface{} {
	params := make(map[string]interface{})
	for key, val := range r.Form {
		if len(val) == 1 {
			var p map[string]interface{}
			if err := json.Unmarshal([]byte(val[0]), &p); err != nil {
				params[key] = val[0]
			} else {
				params[key] = p
			}
		} else if len(val) > 1 {
			params[key] = val
		}
	}
	return params
}
