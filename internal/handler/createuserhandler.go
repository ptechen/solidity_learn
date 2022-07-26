package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"solidity_learn/internal/logic"
	"solidity_learn/internal/svc"
	"solidity_learn/internal/types"

	"github.com/zeromicro/go-zero/rest/httpx"
)

func CreateUserHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewCreateUserLogic(r.Context(), svcCtx)
		err := l.CreateUser(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
