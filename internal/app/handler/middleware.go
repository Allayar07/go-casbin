package handler

import (
	"fmt"
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/persist"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) getUserRole(ginCtx *gin.Context) (string, error) {
	//role, exists := ginCtx.Get("userRole")
	//if !exists {
	//	ginCtx.JSON(http.StatusInternalServerError, "userRole error")
	//	return "", errors.New("can not get user !!!")
	//}
	//
	//roleString, exists := role.(string)
	//if !exists {
	//	ginCtx.JSON(http.StatusInternalServerError, "role not existing")
	//	return "", errors.New("can not get user !!!")
	//}
	token := ginCtx.GetHeader("Authorization")
	roleString, err := h.Services.Auth.Parse(token)
	if err != nil {
		return "", err
	}

	return roleString, nil
}

func enforce(sub string, obj string, act string, adapter persist.Adapter) (bool, error) {
	enforcer, err := casbin.NewEnforcer("path/RBAC_model.conf", adapter)
	if err != nil {
		return false, fmt.Errorf("failed to create enforcer: %w", err)
	}

	if err = enforcer.LoadPolicy(); err != nil {
		return false, fmt.Errorf("failed to load policy: %w", err)
	}

	ok, err := enforcer.Enforce(sub, obj, act)
	if err != nil {
		return false, fmt.Errorf("failed enforcing: %w", err)
	}

	return ok, nil
}

func (h *Handler) Authorize(obj string, act string, adapter persist.Adapter) gin.HandlerFunc {
	return func(ginCtx *gin.Context) {
		userRole, err := h.getUserRole(ginCtx)
		fmt.Println("user role is :........./ ", userRole)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, err)
			return
		}
		enforced, err := enforce(userRole, obj, act, adapter)
		if err != nil {
			ginCtx.JSON(http.StatusInternalServerError, err)
			return
		}

		if !enforced {
			ginCtx.JSON(http.StatusInternalServerError, err)
			return
		}

		ginCtx.Next()
	}
}
