package handler

import (
	"errors"
	"net/http"

	coincommon "github.com/wushengyouya/coin_exchange/coin-common"
	"github.com/wushengyouya/coin_exchange/coin-common/tools"
	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/logic"
	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/svc"
	"github.com/wushengyouya/coin_exchange/ucenter-api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

type RegisterHandler struct {
	svcCtx *svc.ServiceContext
}

func NewRegisterHandler(svcCtx *svc.ServiceContext) *RegisterHandler {
	return &RegisterHandler{
		svcCtx: svcCtx,
	}
}
func (h *RegisterHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req types.Request
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}
	newResult := coincommon.NewResult()
	if req.Captcha == "" {
		httpx.OkJsonCtx(r.Context(), w, newResult.Deal(nil, errors.New("人机校验不通过")))
		return
	}
	// 获取Ip
	req.Ip = tools.GetRemoteClientIp(r)
	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.Register(&req)
	result := coincommon.NewResult().Deal(resp, err)
	result.Data = resp.Message
	httpx.OkJsonCtx(r.Context(), w, result)

}

func (h *RegisterHandler) SendCode(w http.ResponseWriter, r *http.Request) {
	var req types.CodeRequest
	if err := httpx.ParseJsonBody(r, &req); err != nil {
		httpx.ErrorCtx(r.Context(), w, err)
		return
	}

	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.SendCode(&req)
	result := coincommon.NewResult().Deal(resp, err)
	httpx.OkJsonCtx(r.Context(), w, result)
}
