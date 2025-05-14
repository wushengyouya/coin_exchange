package handler

import (
	"net/http"

	coincommon "github.com/wushengyouya/coin_exchange/coin-common"
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
	// if err := httpx.Parse(r, &req); err != nil {
	// 	httpx.ErrorCtx(r.Context(), w, err)
	// 	return
	// }

	l := logic.NewRegisterLogic(r.Context(), h.svcCtx)
	resp, err := l.Register(&req)
	result := coincommon.NewResult().Deal(resp, err)
	result.Data = resp.Message
	httpx.OkJsonCtx(r.Context(), w, result)

}
