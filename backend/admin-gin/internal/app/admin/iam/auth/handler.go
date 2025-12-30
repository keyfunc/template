package auth

import (
	"net/http"

	"mall-api/internal/pkg/cookie"
	pkghttp "mall-api/internal/pkg/http"

	"github.com/gin-gonic/gin"
)

type handler struct {
	se service
	cm *cookie.CookieManager
}

func newHandler(se service, cm *cookie.CookieManager) *handler {
	return &handler{se: se, cm: cm}
}

// @Summary 登录
// @Description 用户名/密码登录
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body loginReq true "登录信息"
// @Success 200 {object} pkghttp.HttpResponse[loginRes] "成功"
// @Router /admin/auth/login [post]
func (h *handler) login(c *gin.Context) {

	// 1.读取接口传参
	var req loginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest)
		return
	}

	// 2.调用 service 层的 login 业务
	res, err := h.se.login(c.Request.Context(), &req)
	if err != nil {
		pkghttp.Fail(c, http.StatusUnauthorized, err.Error())
		return
	}

	// 3. gin 设置 refresh cookie
	h.cm.Set(c, res.RefreshToken)

	pkghttp.OK(c, res)
}

// @Summary 注册
// @Description 用户注册
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body registerReq true "注册信息"
// @Success 200 {object} pkghttp.HttpResponse[pkghttp.Empty] "成功"
// @Router /admin/auth/register [post]
func (h *handler) register(c *gin.Context) {

	// 1. 读取接口传参
	var req registerReq
	if err := c.ShouldBindJSON(&req); err != nil {
		pkghttp.Fail(c, http.StatusBadRequest)
		return
	}

	// 2. 调用 service 层的用户注册
	if err := h.se.register(&req); err != nil {
		pkghttp.Fail(c, http.StatusConflict, err.Error())
		return
	}

	// 3. 返回成功
	pkghttp.OK(c, pkghttp.Empty{})

}

// @Summary 注销
// @Description 用户注销
// @Tags Auth
// @Produce json
// @Success 200 {object} pkghttp.HttpResponse[pkghttp.Empty] "成功"
// @Router /admin/auth/logout [post]
func (h *handler) logout(c *gin.Context) {

	// 1.参数 refresh_token, 从 cookie 中取值
	refreshToken, err := h.cm.Get(c)
	if err != nil {
		pkghttp.Fail(c, http.StatusBadRequest)
		return
	}

	// 2. 调用 service 层注销业务
	code, err := h.se.logout(c.Request.Context(), refreshToken)
	if err != nil {
		pkghttp.Fail(c, code, err.Error())
		return
	}

	// 3. 清除 refresh_token cookie
	h.cm.Remove(c)

	// 4. 成功
	pkghttp.OK(c, pkghttp.Empty{})

}

// @Summary 刷新Token
// @Description 使用RefreshToken刷新AccessToken
// @Tags Auth
// @Produce json
// @Success 200 {object} pkghttp.HttpResponse[loginRes] "成功"
// @Router /admin/auth/refresh [post]
func (h *handler) refresh(c *gin.Context) {

	// 1.参数 refresh_token, 从 cookie 中取值
	refreshToken, err := h.cm.Get(c)
	if err != nil {
		pkghttp.Fail(c, http.StatusBadRequest)
		return
	}

	// 2.调用 service 层的 refresh 业务
	res, err := h.se.refresh(c.Request.Context(), refreshToken)
	if err != nil {
		pkghttp.Fail(c, http.StatusInternalServerError, err.Error())
		return
	}

	// 3.返回 token 对
	pkghttp.OK(c, res)
}
