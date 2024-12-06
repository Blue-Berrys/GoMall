package utlis

type SessionUserIdKey string

const (
	SessionUserId SessionUserIdKey = "user_id"
	ServiceName                    = "frontend" //因为hertz没有ServiceName，所以自己定义一个
)
