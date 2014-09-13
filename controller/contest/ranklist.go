package contest

import (
	"GoOnlineJudge/class"
	"GoOnlineJudge/config"
	"GoOnlineJudge/model"
	"net/http"
	"sort"
	"strconv"
)

type RanklistController struct {
	Contest
}

func (this RanklistController) Route(w http.ResponseWriter, r *http.Request) {
	class.Logger.Debug("RankList")
	this.InitContest(w, r)
	qry := make(map[string]string)
	qry["module"] = strconv.Itoa(config.ModuleC)
	qry["mid"] = strconv.Itoa(this.Cid)
	qry["sort"] = "resort"

	solutionModel := model.SolutionModel{}
	solutionList, err := solutionModel.List(qry)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}

	UserMap := make(map[string]*userRank)
	var pro *probleminfo
	var user *userRank
	for _, v := range solutionList {
		user = UserMap[v.Uid]
		if user == nil {
			user = &userRank{}
			UserMap[v.Uid] = user
			user.ProblemList = make([]*probleminfo, len(this.Index), len(this.Index))
		}
		user.Uid = v.Uid
		pid := this.Index[v.Pid]
		pro = user.ProblemList[pid]
		if pro == nil {
			pro = &probleminfo{}
			user.ProblemList[pid] = pro
		}
		if pro.Judge == config.JudgeAC {
			continue
		}
		pro.Pid = pid
		if v.Judge != config.JudgeAC && v.Judge != config.JudgePD && v.Judge != config.JudgeRJ {
			pro.Count++
			pro.Time += 20 * 60 //罚时20分钟
		} else if v.Judge == config.JudgeAC {
			pro.Time += v.Create - this.ContestDetail.Start
			pro.Judge = config.JudgeAC
			user.Time += pro.Time
			user.Solved += 1
		}
	}
	UserList := newSorter(UserMap)
	sort.Sort(UserList)

	this.Data["UserList"] = UserList
	this.Data["IsContestRanklist"] = true
	this.Data["Cid"] = this.Cid
	this.Data["ProblemList"] = this.Index
	this.Execute(w, "view/layout.tpl", "view/contest/ranklist.tpl")
}

type userRank struct {
	Uid         string
	ProblemList []*probleminfo
	Time        int64
	Solved      int
}

type probleminfo struct {
	Pid   int
	Time  int64
	Count int
	Judge int
}

type UserSorter []*userRank

func newSorter(m map[string]*userRank) UserSorter {
	us := make([]*userRank, 0, len(m))
	for _, v := range m {
		us = append(us, v)
	}
	return us
}

func (u UserSorter) Len() int {
	return len(u)
}

func (u UserSorter) Less(i, j int) bool {
	if u[i].Solved > u[j].Solved {
		return true
	} else if u[i].Solved == u[j].Solved {
		if u[i].Time < u[j].Time {
			return true
		} else if u[i].Time >= u[j].Time {
			return false
		}
	}
	return false
}

func (u UserSorter) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}
