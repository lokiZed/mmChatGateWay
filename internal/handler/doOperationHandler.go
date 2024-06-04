package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"mmChat/internal/logic"
	"mmChat/internal/svc"
)

func DoOperationHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewDoOperationLogic(r.Context(), svcCtx)
		err := l.DoOperation()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
