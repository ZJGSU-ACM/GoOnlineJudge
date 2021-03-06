package controller

import (
	"github.com/ZJGSU-ACM/GoOnlineJudge/class"
	"github.com/ZJGSU-ACM/restweb"
)

type FAQController struct {
	class.Controller
} //@Controller

//faq 页面
//@URL: /faq @method: GET
func (fc *FAQController) Index() {
	restweb.Logger.Debug("FAQ Page")

	fc.Output["Title"] = "FAQ"
	fc.Output["IsFAQ"] = true
	fc.RenderTemplate("view/layout.tpl", "view/faq.tpl")
}
