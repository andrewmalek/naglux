package collector

import "github.com/andrewmalek/naglux/data"

type ResultQueues map[data.Target]chan Printable
