package controller

import (
	"github.com/ZJGSU-ACM/GoOnlineJudge/class"
	"github.com/ZJGSU-ACM/restweb"
)

type OSCController struct {
	class.Controller
} //@Controller

//@URL: /osc @method: GET
func (oc *OSCController) Index() {
	restweb.Logger.Debug("OSC Page")

	oc.Output["Title"] = "ZJGSU OSC"
	oc.Output["IsOSC"] = true
	oc.RenderTemplate("view/layout.tpl", "view/osc.tpl")
}
