package model

import (
	"GoOnlineJudge/config"
	"GoOnlineJudge/model/class"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"strconv"
)

type User struct {
	Uid string `json:"uid"bson:"uid"`
	Pwd string `json:"pwd"bson:"pwd"`

	Nick   string `json:"nick"bson:"nick"`
	Mail   string `json:"mail"bson:"mail"`
	School string `json:"school"bson:"school"`
	Motto  string `json:"motto"bson:"motto"`

	Privilege int `json:"privilege"bson:"privilege"`

	Solve  int `json:"solve"bson:"solve"`
	Submit int `json:"submit"bson:"submit"`

	Status int    `json:"status"bson:"status"`
	Create string `json:"create"bson:'create'`
}

var uDetailSelector = bson.M{"_id": 0}
var uListSelector = bson.M{"_id": 0, "uid": 1, "nick": 1, "motto": 1, "privilege": 1, "solve": 1, "submit": 1, "status": 1}

type UserModel struct {
	class.Model
}

// POST /User?login
func (this *UserModel) Login(uid, pwd string) (*User, error) {
	log.Println("Server UserModel Login")

	var err error
	pwd, err = this.EncryptPassword(pwd)
	if err != nil {
		return nil, class.EncryptErr
	}

	err = this.OpenDB()
	if err != nil {
		return nil, class.DBErr
	}
	defer this.CloseDB()

	var alt User
	err = this.DB.C("User").Find(bson.M{"uid": uid}).Select(uDetailSelector).One(&alt)
	if err == mgo.ErrNotFound {
		return nil, class.NotFoundErr
	} else if err != nil {
		return nil, class.OpErr
	}

	if pwd == alt.Pwd {
		log.Println("Server UserModel Login Successfully")
		return &User{
			Uid:       alt.Uid,
			Nick:      alt.Nick,
			Privilege: alt.Privilege,
			Status:    alt.Status,
		}, nil
	}
	log.Println("Server UserModel Login Failed")
	return &User{
		Uid:       "",
		Nick:      "",
		Privilege: config.PrivilegeNA,
		Status:    config.StatusReverse,
	}, nil
}

// POST /User?logout
//这个函数貌似没干什么事啊==
func (this *UserModel) Logout() {
	log.Println("Server UserModel Logout")

	// var one User
	// err := this.LoadJson(r.Body, &one)
	// if err != nil {
	// 	http.Error(w, "load error", 400)
	// 	return
	// }

	// w.WriteHeader(200)
}

// POST /User?password/uid?<uid>
func (this *UserModel) Password(uid, pwd string) error {
	log.Println("Server UserModel Password")

	pwd, err := this.EncryptPassword(pwd)
	if err != nil {
		return class.EncryptErr
	}

	alt := make(map[string]interface{})
	alt["pwd"] = pwd

	err = this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /User?privilege/uid?<uid>
func (this *UserModel) Privilege(uid, privilege string) error {
	log.Println("Server UserModel Privilege")

	alt := make(map[string]interface{})
	alt["privilege"] = privilege

	err := this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /User?detail/uid?<uid>
func (this *UserModel) Detail(uid string) (*User, error) {
	log.Println("Server UserModel Detail")

	err := this.OpenDB()
	if err != nil {
		return nil, class.DBErr
	}
	defer this.CloseDB()

	var one User
	err = this.DB.C("User").Find(bson.M{"uid": uid}).Select(uDetailSelector).One(&one)
	if err == mgo.ErrNotFound {
		return nil, class.NotFoundErr
	} else if err != nil {
		return nil, class.OpErr
	}

	return &one, nil
}

// POST /User?delete/uid?<uid>
func (this *UserModel) Delete(uid string) error {
	log.Println("Server UserModel Delete")

	err := this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("User").Remove(bson.M{"uid": uid})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /User?insert
func (this *UserModel) Insert(one User) error {
	log.Println("Server UserModel Insert")

	var err error
	one.Pwd, err = this.EncryptPassword(one.Pwd)
	if err != nil {
		return class.EncryptErr
	}

	err = this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	//one.Privilege = config.PrivilegePU
	one.Solve = 0
	one.Submit = 0
	one.Status = config.StatusAvailable
	one.Create = this.GetTime()

	err = this.DB.C("User").Insert(&one)
	if err != nil {
		return class.OpErr
	}

	// b, err := json.Marshal(map[string]interface{}{
	// 	"uid":       one.Uid,
	// 	"privilege": one.Privilege,
	// 	"status":    one.Status,
	// })

	return nil
}

// POST /User?update/uid?<uid>
func (this *UserModel) Update(uid string, ori User) error {
	log.Println("Server UserModel Update")

	alt := make(map[string]interface{})
	alt["nick"] = ori.Nick
	alt["mail"] = ori.Mail
	alt["school"] = ori.School
	alt["motto"] = ori.Motto

	err := this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$set": alt})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /User?status/uid?<uid>
func (this *UserModel) Status(uid string) error {
	log.Println("Server UserModel Status")

	err := this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$inc": bson.M{"status": 1}})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /User?record/uid?<uid>/action?<solve/submit>
func (this *UserModel) Record(uid, action string) error {
	log.Println("Server UserModel Submit")

	var inc int
	switch action {
	case "solve":
		inc = 1
	case "submit":
		inc = 0
	default:
		return class.ArgsErr
	}

	err := this.OpenDB()
	if err != nil {
		return class.DBErr
	}
	defer this.CloseDB()

	err = this.DB.C("User").Update(bson.M{"uid": uid}, bson.M{"$inc": bson.M{"solve": inc, "submit": 1}})
	if err == mgo.ErrNotFound {
		return class.NotFoundErr
	} else if err != nil {
		return class.OpErr
	}

	return nil
}

// POST /User?list/offset?<offset>/limit?<limit>/uid?<uid>/nick?<nick>
func (this *UserModel) List(args map[string]string) ([]*User, error) {
	log.Println("Server UserModel List")

	query, err := this.CheckQuery(args)
	if err != nil {
		return nil, class.ArgsErr
	}

	err = this.OpenDB()
	if err != nil {
		return nil, class.DBErr
	}
	defer this.CloseDB()

	q := this.DB.C("User").Find(query).Select(uListSelector).Sort("-solve", "submit")

	if v, ok := args["offset"]; ok {
		offset, err := strconv.Atoi(v)
		if err != nil {
			return nil, class.ArgsErr
		}
		q = q.Skip(offset)
	}

	if v, ok := args["limit"]; ok {
		limit, err := strconv.Atoi(v)
		if err != nil {
			return nil, class.ArgsErr
		}
		q = q.Limit(limit)
	}

	var list []*User
	err = q.All(&list)
	if err != nil {
		return nil, class.QueryErr
	}

	return list, nil
}

func (this *UserModel) CheckQuery(args map[string]string) (query bson.M, err error) {
	query = make(bson.M)

	if v, ok := args["uid"]; ok {
		query["uid"] = v
	}
	if v, ok := args["nick"]; ok {
		query["nick"] = bson.M{"$regex": bson.RegEx{v, "i"}}
	}
	return
}