package route

import (
	"net/http"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]HandlerFunc
}

func (r *router) addRoute(method string, path string, handler HandlerFunc) {
	parts := parsePath(path)
	key := method + "-" + path
	_, ok := r.roots[method]
	if !ok {
		//给map的当前方法kv赋值
		r.roots[method] = &node{}
	}
	r.roots[method].insert(path, parts, 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	parts := parsePath(path)
	params := make(map[string]string)
	_, ok := r.roots[method]
	if !ok {
		return nil, nil
	}

	node := r.roots[method].search(parts, 0)
	//if node != nil {
	//	for index, part := range parts {
	//		if part[0] == '?' {
	//		}
	//	}
	//}
	return node, params
}

func (r *router) gteRouters(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

func parsePath(path string) []string {
	split := strings.Split(path, "/")
	parts := make([]string, 0)
	for _, item := range split {
		if item != "" {
			parts = append(parts, item)
			//if item[0] == '*' {
			//	break
			//}
		}
	}
	return parts
}

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.path
		//r.handlers[key](c)
		c.handlers = append(c.handlers, r.handlers[key])
	} else {
		c.handlers = append(c.handlers, func(c *Context) {
			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
		})
	}
	c.Next()
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]HandlerFunc),
	}
}
