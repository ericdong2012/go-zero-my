 part1:
    lang
    stringx
    mathx
    iox
    timex
    cmdline
    errorx
    filex
    jsonx
    jsontype
    netx
    naming

    sysx
    mapping todo 实现了json的 tag 匹配
    logx    todo 实现了链路追踪
    codec   加密

    hash   todo 一致性hash ring
    conf   和 mapping 有很强的关联性
    contextx
    fs     todo go build     	logx.rotatelogger   fs.CloseOnExec(l.fp)
    utils

    syncx  SingleFlight
    collection  滑动窗口，时间轮，ring, fifo, cache（lru）

    rescue      recover gorunsafe()
    threading   启动task
    search      route路由tree

    prof     tablewriter 生成报表
    proc     goroutines, profile(cpu, mem, mutex, block, threading ...)
             shutdown中有几个变量，否则会报错
                	wrapUpListeners          = new(listenerManager)
                	shutdownListeners        = new(listenerManager)
                	delayTimeBeforeForceQuit = waitTime

                	syscall（go 系统包， 有很多底层的汇编文件）

    stat     cgroup   cpu  mem  topK(container/heap)  gock   (会用到executors)  metrics 可以给出压测报告


    queue  consumer, producer, multiProducer, balancedProducer
    storex  用到了timewheel
        cache,
        mongo,
        pg,
        clickhouse,
        kv,
        redis,
        sqlc,
        sqlx  具体的查询语句等   bulkinsert
    bloom      布隆过滤器
    breaker    熔断器
    executers  执行一些任务， 跟threading 还是很有关联性的
    discover  etcd服务发现
    limit     限流(redis )

    load
        自适应负载  todo

    service
        serviceConf
            metricsurl
            prometheus.StartAgent
            trace.StartAgent
            stat.SetReportWriter(stat.NewRemoteWriter(sc.MetricsUrl))


        Starter interface {
        		Start()
        	}

        	// Stopper is the interface wraps the Stop method.
        	Stopper interface {
        		Stop()
        	}

        	// Service is the interface that groups Start and Stop methods.
        	Service interface {
        		Starter
        		Stopper
        	}

        	// A ServiceGroup is a group of services.
        	// Attention: the starting order of the added services is not guaranteed.
        	ServiceGroup struct {
        		services []Service
        		stopOnce func()
        	}

    metrics
        各种prometheus 格式指标
    prometheus
        连接prometheus

    trace
        opentelemetry(jaeger, zipkin)
        支持grpc
        attribute 中对grpc 做了解析，方便对接openxxx

    fx
       retry
       timeout & recover
       parallel
       stream  集成了很多方法： match, filter, concat, join, group， map, reduce, first, count(), distinct() 等等

    mr



part2:
rest
    token   jwt
    pathvar  context 里面加东西
    httpx
        router 定义了一些接口， 在router/patrouter 中实现,  在 server 中调用
        request  封装了很多请求方法体  ParseHeaders， ParseForm， ParseJsonBody， ParsePath
        response 封装了很多返回方法体，包括 error, ok, json 等
        vars     常量

    handler
        能直接包装 http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})
        在engine 中调用

    interval 和handler 类似，更多是中间件
        cors
            在server 中会调用到
        security  对请求内容加密
            用到了codec 中的rsa 加密，解密等
        log
        starter   里面有一些 proc 的监听器， 在 engine 中会调用到

    router
        是httpx 中router的实现

        type patRouter struct {
        	trees      map[string]*search.Tree
        	notFound   http.Handler
        	notAllowed http.Handler
        }

        ServeHTTP    server 中的tomiddle 中有用到, （应该是很多地方有用到，需要调试）

    server
        Server struct {
        		ngin   *engine
        		router httpx.Router
        	}
        )

        NewServer
        AddRoute

        Start           调用engine 中的 start, (最终会调用starter中 StartHttp --> start )  --> bindRoutes -->  stat createMetrics,  bindFeaturedRoutes --> signatureVerifier, bindRoute --> 加载各种handler, appendAuthHandler, convertMiddleware,  router.Handle -->

        use
        withxxx          RouteOption
        WithMiddleware   []Route

    engine
        start
        bindRoutes

    config 定义很多yaml 等配置文件的结构体
    type

part3:
goctl

part4:
zrpc


1. 先看懂，跑测试用例，把core跑通， rest最小化跑通
2. 自己写