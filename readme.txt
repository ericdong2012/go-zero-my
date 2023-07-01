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

part3:
goctl

part4:
zrpc


1. 先看懂，跑测试用例，把core跑通， rest最小化跑通
2. 自己写