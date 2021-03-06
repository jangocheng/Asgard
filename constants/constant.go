package constants

import "github.com/dalonghahaha/avenger/structs"

const (
	TYPE_AGENT  = int64(1)
	TYPE_APP    = int64(2)
	TYPE_JOB    = int64(3)
	TYPE_TIMING = int64(4)
	TYPE_GROUP  = int64(5)
	TYPE_USER   = int64(6)

	ACTION_CREATE  = int64(1)
	ACTION_UPDATE  = int64(2)
	ACTION_DELETE  = int64(3)
	ACTION_COPY    = int64(4)
	ACTION_START   = int64(5)
	ACTION_RESTART = int64(6)
	ACTION_PAUSE   = int64(7)

	USER_ROLE_ADMIN  = "Administrator"
	USER_ROLE_NORMAL = "User"

	USER_STATUS_UNVERIFIED = int64(0)
	USER_STATUS_NORMAL     = int64(1)
	USER_STATUS_FORBIDDEN  = int64(-1)

	AGENT_OFFLINE   = int64(0)
	AGENT_ONLINE    = int64(1)
	AGENT_FORBIDDEN = int64(-1)

	GROUP_STATUS_DELETED = int64(-1)
	GROUP_STATUS_UNUSAGE = int64(0)
	GROUP_STATUS_USAGE   = int64(1)

	APP_STATUS_DELETED = int64(-1)
	APP_STATUS_UNKNOWN = int64(-2)
	APP_STATUS_FAILED  = int64(-3)
	APP_STATUS_STOP    = int64(0)
	APP_STATUS_RUNNING = int64(1)
	APP_STATUS_PAUSE   = int64(2)

	JOB_STATUS_DELETED = int64(-1)
	JOB_STATUS_UNKNOWN = int64(-2)
	JOB_STATUS_STOP    = int64(0)
	JOB_STATUS_RUNNING = int64(1)
	JOB_STATUS_PAUSE   = int64(2)

	TIMING_STATUS_DELETED  = int64(-1)
	TIMING_STATUS_UNKNOWN  = int64(-2)
	TIMING_STATUS_STOP     = int64(0)
	TIMING_STATUS_RUNNING  = int64(1)
	TIMING_STATUS_PAUSE    = int64(2)
	TIMING_STATUS_FINISHED = int64(3)
)

var ACTION_NAME = map[int64]string{
	ACTION_CREATE:  "创建",
	ACTION_UPDATE:  "更新",
	ACTION_DELETE:  "删除",
	ACTION_COPY:    "复制",
	ACTION_START:   "启动",
	ACTION_RESTART: "重启",
	ACTION_PAUSE:   "停止",
}

var TYPE_NAME = map[int64]string{
	TYPE_AGENT:  "实例",
	TYPE_APP:    "应用",
	TYPE_JOB:    "计划任务",
	TYPE_TIMING: "定时任务",
	TYPE_GROUP:  "分组",
	TYPE_USER:   "用户",
}

var USER_STATUS = []structs.M{
	{
		"ID":   USER_STATUS_UNVERIFIED,
		"Name": "未审核",
	},
	{
		"ID":   USER_STATUS_NORMAL,
		"Name": "正常",
	},
	{
		"ID":   USER_STATUS_FORBIDDEN,
		"Name": "禁用",
	},
}

var AGENT_STATUS = []structs.M{
	{
		"ID":   AGENT_ONLINE,
		"Name": "在线",
	},
	{
		"ID":   AGENT_OFFLINE,
		"Name": "离线",
	},
	{
		"ID":   AGENT_FORBIDDEN,
		"Name": "禁用",
	},
}

var GROUP_STATUS = []structs.M{
	{
		"ID":   GROUP_STATUS_UNUSAGE,
		"Name": "未启用",
	},
	{
		"ID":   GROUP_STATUS_USAGE,
		"Name": "启用",
	},
	{
		"ID":   GROUP_STATUS_DELETED,
		"Name": "已删除",
	},
}

var APP_STATUS = []structs.M{
	{
		"ID":   APP_STATUS_STOP,
		"Name": "停止",
	},
	{
		"ID":   APP_STATUS_RUNNING,
		"Name": "运行中",
	},
	{
		"ID":   APP_STATUS_PAUSE,
		"Name": "暂停",
	},
	{
		"ID":   APP_STATUS_UNKNOWN,
		"Name": "未知",
	},
	{
		"ID":   APP_STATUS_DELETED,
		"Name": "已删除",
	},
}

var JOB_STATUS = []structs.M{
	{
		"ID":   JOB_STATUS_STOP,
		"Name": "停止",
	},
	{
		"ID":   JOB_STATUS_RUNNING,
		"Name": "运行中",
	},

	{
		"ID":   JOB_STATUS_PAUSE,
		"Name": "暂停",
	},
	{
		"ID":   JOB_STATUS_UNKNOWN,
		"Name": "未知",
	},
	{
		"ID":   JOB_STATUS_DELETED,
		"Name": "已删除",
	},
}

var TIMING_STATUS = []structs.M{
	{
		"ID":   TIMING_STATUS_STOP,
		"Name": "停止",
	},
	{
		"ID":   TIMING_STATUS_RUNNING,
		"Name": "运行中",
	},
	{
		"ID":   TIMING_STATUS_PAUSE,
		"Name": "暂停",
	},
	{
		"ID":   TIMING_STATUS_UNKNOWN,
		"Name": "未知",
	},
	{
		"ID":   TIMING_STATUS_FINISHED,
		"Name": "已完成",
	},
	{
		"ID":   TIMING_STATUS_DELETED,
		"Name": "已删除",
	},
}
