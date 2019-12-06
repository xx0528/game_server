package main

type RobotMgr struct {
	RobotMap map[string]*Robot
}

func NewRobotMgr() *RobotMgr {
	mgr := new(RobotMgr)
	mgr.Init()
	return mgr
}

func (mgr *RobotMgr) Init() {
	mgr.RobotMap = make(map[string]*Robot)
}

func (mgr *RobotMgr) Get(rid string) *Robot {
	return mgr.RobotMap[rid]
}

func (mgr *RobotMgr) Add(robot *Robot) {
	mgr.RobotMap[robot.objid] = robot
}

func (mgr *RobotMgr) Del(objid string) {
	delete(mgr.RobotMap, objid)
}

func (mgr *RobotMgr) Close() {
	for _, robot := range mgr.RobotMap {
		robot.Logout()
	}
}
