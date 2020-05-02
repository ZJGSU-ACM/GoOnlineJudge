package config

import (
	"os"
)

//server
const CookieExpires = 1800

var Datapath = os.Getenv("DATA_PATH")
var JudgeHost = os.Getenv("JUDGE_HOST")

//CONSTANT
const (
	ProblemPerPage  = 50
	ContestPerPage  = 100
	SolutionPerPage = 30
	UserPerPage     = 50
)

const (
	PageHeadLimit = 1
	PageTailLimit = 1
	PageMidLimit  = 2
)

const (
	JudgePD  = 0  //Pending
	JudgeRJ  = 1  //Running & judging
	JudgeCE  = 2  //Compile Error
	JudgeAC  = 3  //Accepted
	JudgeRE  = 4  //Runtime Error
	JudgeWA  = 5  //Wrong Answer
	JudgeTLE = 6  //Time Limit Exceeded
	JudgeMLE = 7  //Memory Limit Exceeded
	JudgeOLE = 8  //Output Limit Exceeded
	JudgePE  = 9  //Presentation Error
	JudgeNA  = 10 //System Error
	JudgeRPD = 11 //Rejudge Pending
)

const (
	FlagNA = uint8(iota) // None
	FLagAC
	FLagER
)
const (
	LanguageNA   = 0 //None
	LanguageC    = 1 //C
	LanguageCPP  = 2 //C++
	LanguageJAVA = 3 //Java
	LanguagePy2 = 4 //Python2
	LanguagePy3 = 5 //Python3
)

const (
	ModuleNA = 0 //None
	ModuleP  = 1 //Problem
	ModuleC  = 2 //Contest
)

const (
	PrivilegeNA = 0 //None
	PrivilegePU = 1 //Primary User
	PrivilegeTC = 2 //Teacher
	PrivilegeAD = 3 //Admin
)

const (
	EncryptNA = 0 //None
	EncryptPB = 1 //Public
	EncryptPT = 2 //Private
	EncryptPW = 3 //Password
)

const (
	StatusReverse   = 0 //不可用
	StatusIncon     = 1 //正在比赛中
	StatusAvailable = 2 //可用
	StatusPending   = 3 //等待
	StatusRunning   = 4 //进行中
	StatusEnding    = 5 //结束
)

//remote OJ status
const (
	StatusOk          = 0
	StatusUnavailable = 1
)

// 权限分离
const (
	AccessAdmin = 1 << iota //管理员页面
	ViewReverse             //查看保留问题、新闻、竞赛
	Notice                  //通知消息

	AddProblem    //添加问题
	DeleteProblem //删除问题
	Testcase      //测试数据管理
	ReJudge       //重判

	AddContest    //添加竞赛
	DeleteContest //删除竞赛
	AddNews       //添加新闻
	DeleteNews    //删除新闻

	ViewCode //查看代码
	ViewSim  //查看相似度

	UserControl  //用户控制
	GenerateUser //生成用户
)

const (
	ProFull  = AddProblem | DeleteProblem | Testcase | ReJudge
	ConFull  = AddContest | DeleteContest
	NewsFull = AddNews | DeleteNews
	CodeFull = ViewCode | ViewSim
	UserFull = UserControl | GenerateUser
)
