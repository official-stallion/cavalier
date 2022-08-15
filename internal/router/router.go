package router

type Router struct {
	defaultGateWay string

	routingTable map[string]string
}

func New(cfg Config) Router {
	return Router{
		defaultGateWay: cfg.DefaultGateWay,

		routingTable: make(map[string]string, cfg.NumberOfInterface),
	}
}

func (r *Router) RegisterIp(routerInterface string, ip string) {
	r.routingTable[ip] = routerInterface
}

func (r *Router) Next(ip string) string {
	return r.match(ip)
}

func (r *Router) match(ip string) string {
	for key := range r.routingTable {
		if key == ip {
			return r.routingTable[key]
		}
	}

	return r.defaultGateWay
}
