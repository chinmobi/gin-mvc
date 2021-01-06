// Copyright 2014 Manu Martinez-Almeida.  All rights reserved.
// Use of this source code is governed by a MIT style license that can be found
// at https://github.com/gin-gonic/gin/blob/master/LICENSE

package internal

type TopicTree struct {
	topic string
	root  *node
}

type TopicTreeS []TopicTree

func (trees TopicTreeS) get(topic string) *node {
	for _, tree := range trees {
		if tree.topic == topic {
			return tree.root
		}
	}
	return nil
}

func (trees TopicTreeS) AddRoute(topic, path string, handlers HandlersChain) TopicTreeS {
	root := trees.get(topic)

	if root == nil {
		root = new(node)
		root.fullPath = "/"
		trees = append(trees, TopicTree{ topic: topic, root: root })
	}

	root.addRoute(path, handlers)

	return trees
}

func (trees TopicTreeS) Handle(c *Context, event *Event) {
	unescape := false

	topic := event.Topic
	rPath := event.RoutingPath

	c.resetWithEvent(event)

	t := trees
	for i, tl := 0, len(t); i < tl; i++ {
		if t[i].topic != topic {
			continue
		}

		root := t[i].root
		// Find route in tree
		value := root.getValue(rPath, c.params, unescape)
		if value.params != nil {
			c.Params = *value.params
		}
		if value.handlers != nil {
			c.handlers = value.handlers
			c.fullPath = value.fullPath
			c.Next()
			return
		}

		break
	}
}
