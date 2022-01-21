package google

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/xmapst/alertingwebhook/cmd"
	"github.com/xmapst/alertingwebhook/dingtalk"
	"github.com/xmapst/alertingwebhook/handlers"
	"net/http"
	"strconv"
	"time"
)

// Notification
// @Summary Notification
// @description post notification webhook
// @Tags Google
// @Param content body NotificationStruct true "content"
// @Success 200 {object} handlers.JSONResult{}
// @Failure 500 {object} handlers.JSONResult{}
// @Router /google/notification [post]
func Notification(c *gin.Context) {
	render := handlers.Gin{Context: c}
	accessToken, found := c.GetQuery("access_token")
	if !found {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}
	if _, ok := cmd.RoboterMap[accessToken]; !ok {
		c.JSON(http.StatusUnauthorized, map[string]interface{}{
			"code":    http.StatusUnauthorized,
			"message": "unauthorized",
		})
		return
	}
	var req = new(NotificationStruct)
	if err := render.ShouldBindJSON(req); err != nil {
		logrus.Error(err)
		c.JSON(http.StatusBadRequest, map[string]interface{}{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	roboter := dingtalk.NewRobot(
		accessToken,
		cmd.RoboterMap[accessToken],
	)
	err := roboter.SendMarkdown(req.Incident.PolicyName, req.markdownText(), []string{}, false)
	if err != nil {
		logrus.Error(err)
	}
	render.SetJson("success")
}

func (r *NotificationStruct) markdownText() string {
	body := fmt.Sprintf("## %s  ", r.Incident.ConditionName)
	if r.Incident.StartedAt != 0 {
		body += fmt.Sprintf("\n- 开始时间: %s ", r.parseTime(r.Incident.StartedAt))
	}
	if r.Incident.EndedAt != 0 {
		body += fmt.Sprintf("\n- 结束时间: %s ", r.parseTime(r.Incident.EndedAt))
	}
	body += fmt.Sprintf("\n- 项目: [%s](https://console.cloud.google.com/?project=%s)  ", r.Incident.ScopingProjectID, r.Incident.ScopingProjectID)
	body += fmt.Sprintf("\n- 资源类型: %s  ", r.Incident.ResourceTypeDisplayName)
	body += fmt.Sprintf("\n- 资源名称: %s  ", r.Incident.ResourceDisplayName)
	body += fmt.Sprintf("\n- 告警策略: %s  ", r.Incident.PolicyName)
	body += fmt.Sprintf("\n- 阈值%%: %s  ", r.parseFloat(r.Incident.ThresholdValue))
	body += fmt.Sprintf("\n- 触发%%: %s  ", r.parseFloat(r.Incident.ObservedValue))
	body += fmt.Sprintf("\n> %s [详情](%s)", r.Incident.Summary, r.Incident.URL)
	return body
}

func (r *NotificationStruct) parseFloat(s string) string {
	f, err := strconv.ParseFloat("0.1", 32)
	if err != nil {
		logrus.Error(err)
		return s
	}
	return fmt.Sprintf("%.2f", f*100)
}

func (r *NotificationStruct) parseTime(t int64) string {
	tm := time.Unix(t, 0)
	return tm.Format(time.RFC3339)
}
