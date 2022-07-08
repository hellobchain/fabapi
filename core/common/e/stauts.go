package e

const (
	ALLREADY_CREATE  = iota + 1001 // 已创建
	ALLREADY_BUILD                 // 已构建
	ALLREADY_RUN                   //已行中
	ALLREADY_STOP                  //已停止
	ALLREADY_RELEASE               //已释放
	ALLREADY_INIT
)

const (
	FAILURE_CREATE  = iota + 2001 //创建失败
	FAILURE_BUILD                 //构建失败
	FAILURE_RUN                   //运行失败
	FAILURE_STOP                  //停止失败
	FAILURE_RELEASE               //释放失败
	FAILURE_INIT
)

const (
	WITING_SCOPE    = iota + 3000 //等待范围
	WAITING_CREATE                //等待创建
	WAITING_BUILD                 //等待构建
	WAITING_RUN                   //等待运行
	WAITING_STOP                  //等待停止
	WAITING_RELEASE               //等待释放
	WAITING_INIT

	WAITING_JOIN
	WAITING_LEAVE
	WAITING_UPGRADE
	WAITING_DOWNGRADE
	WAITING_MINNER
	WAITING_DEFAULT
)

const (
	CHIAN_NOT_ANCHOR = 0
	CHAIN_ANCHOR     = 1
)
const ( //链类型
	CHAIN_MAINCHAIN   = iota + 1 //本地主链
	CHAIN_SUBCHAIN               //本地子链
	CHAIN_REMOTECHAIN            //远程链
)

const (
	//链及节点status 状态
	NOT_INIT = iota + 1 //待初始化
	ALLOW               //允许
	FORBID              //禁止
	ROLE_NOTJOIN
	ROLE_ADMIN
	ROLE_MEMBER
	//节点身份
	PEERTYPE_DEFAULT
	PEERTYPE_MINNER

	//共识类型
	RAFT = "raft"
	POA  = "poa"

	TRUE  = "true"
	FALSE = "false"

	FABRIC = "fabric"
	SIPC   = "sipc"
)

type Role int

var ChianStatusFlag = map[int]string{}

var PodStatusFlag = map[string]int{
	RunningStatus:                Running,
	PendingInitializedStatus:     Pending_Initialized,     //等待中
	PendingReadyStatus:           Pending_Ready,           //等待中
	PendingContainersReadyStatus: Pending_ContainersReady, //等待中
	PendingPodScheduledStatus:    Pending_PodScheduled,    //等待中

	RunningInitializedStatus:     Running_Initialized,     //运行中
	RunningReadyStatus:           Running_Ready,           //运行中
	RunningContainersReadyStatus: Running_ContainersReady, //运行中
	RunningPodScheduledStatus:    Running_PodScheduled,    //运行中

	SucceededInitializedStatus:     Succeeded_Initialized,     //正常终止
	SucceededReadyStatus:           Succeeded_Ready,           //正常终止
	SucceededContainersReadyStatus: Succeeded_ContainersReady, //正常终止
	SucceededPodScheduledStatus:    Succeeded_PodScheduled,    //正常终止

	FailedInitializedStatus:     Failed_Initialized,     //异常停止
	FailedReadyStatus:           Failed_Ready,           //异常停止
	FailedContainersReadyStatus: Failed_ContainersReady, //异常停止
	FailedPodScheduledStatus:    Failed_PodScheduled,    //异常停止

	UnkonwnInitializedStatus:      Unkonwn_Initialized,     //未知状态
	UnkonwnReadyStatus:            Unkonwn_Ready,           //未知状态
	UnkonwnContainersReadyStatus:  Unkonwn_ContainersReady, //未知状态
	UnkonwnPodScheduledStatus:     Unkonwn_PodScheduled,    //未知状态
	PodHostIPIsEmptyStatus:        Pod_HostIPIsEmpty,
	PodIPIsEmptyStatus:            Pod_IPIsEmpty,
	PodContainersRunningStatus:    Pod_ContainersRunning,
	PodContainersWaitingStatus:    Pod_ContainersWaiting,
	PodContainersTerminatedStatus: Pod_ContainersTerminated,
	OfflineStatus:                 Offline,
}

const (
	Running                 = iota + 100
	Pending_Initialized     //等待中
	Pending_Ready           //等待中
	Pending_ContainersReady //等待中
	Pending_PodScheduled
	Running_Initialized       //运行中
	Running_Ready             //运行中
	Running_ContainersReady   //运行中
	Running_PodScheduled      //运行中
	Succeeded_Initialized     //正常终止
	Succeeded_Ready           //正常终止
	Succeeded_ContainersReady //正常终止
	Succeeded_PodScheduled    //正常终止
	Failed_Initialized        //异常停止
	Failed_Ready              //异常停止
	Failed_ContainersReady    //异常停止
	Failed_PodScheduled       //异常停止
	Unkonwn_Initialized       //未知状态
	Unkonwn_Ready             //未知状态
	Unkonwn_ContainersReady   //未知状态
	Unkonwn_PodScheduled      //未知状态
	Pod_HostIPIsEmpty
	Pod_IPIsEmpty
	Pod_ContainersRunning
	Pod_ContainersWaiting
	Pod_ContainersTerminated
	Offline // not online
)

const (
	RunningStatus                = "Running"
	PendingInitializedStatus     = "PendingInitialized"     //等待中
	PendingReadyStatus           = "PendingReady"           //等待中
	PendingContainersReadyStatus = "PendingContainersReady" //等待中
	PendingPodScheduledStatus    = "PendingPodScheduled"    //等待中

	RunningInitializedStatus     = "RunningInitialized"     //运行中
	RunningReadyStatus           = "RunningReady"           //运行中
	RunningContainersReadyStatus = "RunningContainersReady" //运行中
	RunningPodScheduledStatus    = "RunningPodScheduled"    //运行中

	SucceededInitializedStatus     = "SucceededInitialized"     //正常终止
	SucceededReadyStatus           = "SucceededReady"           //正常终止
	SucceededContainersReadyStatus = "SucceededContainersReady" //正常终止
	SucceededPodScheduledStatus    = "SucceededPodScheduled"    //正常终止

	FailedInitializedStatus     = "FailedInitialized"     //异常停止
	FailedReadyStatus           = "FailedReady"           //异常停止
	FailedContainersReadyStatus = "FailedContainersReady" //异常停止
	FailedPodScheduledStatus    = "FailedPodScheduled"    //异常停止

	UnkonwnInitializedStatus     = "UnkonwnInitialized"     //未知状态
	UnkonwnReadyStatus           = "UnkonwnReady"           //未知状态
	UnkonwnContainersReadyStatus = "UnkonwnContainersReady" //未知状态
	UnkonwnPodScheduledStatus    = "UnkonwnPodScheduled"    //未知状态

	PodHostIPIsEmptyStatus        = "EmptyPodHostIP"
	PodIPIsEmptyStatus            = "EmptyPodIP"
	PodContainersRunningStatus    = "RunningPodContainers"
	PodContainersWaitingStatus    = "WaitingPodContainers"
	PodContainersTerminatedStatus = "TerminatedPodContainers"

	OfflineStatus = "Offline"
)
