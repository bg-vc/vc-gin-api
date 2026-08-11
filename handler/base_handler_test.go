package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/bg-vc/vc-gin-api/pkg/errno"
	"github.com/gin-gonic/gin"
)

func TestSendResponseUsesHTTPStatusAndEnvelope(t *testing.T) {
	gin.SetMode(gin.TestMode)
	recorder := httptest.NewRecorder()
	context, _ := gin.CreateTestContext(recorder)
	SendResponse(context, errno.ErrParam, nil)

	if recorder.Code != http.StatusBadRequest {
		t.Fatalf("expected 400, got %d", recorder.Code)
	}
	var response Response
	if err := json.Unmarshal(recorder.Body.Bytes(), &response); err != nil {
		t.Fatal(err)
	}
	if response.Code != errno.ErrParam.Code || response.Msg != errno.ErrParam.Msg {
		t.Fatalf("unexpected response: %+v", response)
	}
}
