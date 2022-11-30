package handler

import (
	"net/http"

	"github.com/i-Things/things/shared/result"

	"github.com/i4de/iThings-demo/src/papisvr/internal/logic"
	"github.com/i4de/iThings-demo/src/papisvr/internal/svc"
)

func pingHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := logic.NewPingLogic(r.Context(), svcCtx)
		err := l.Ping()
		result.Http(w, r, nil, err)
	}
}
