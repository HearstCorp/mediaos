package mediaos

import (
	l5g "github.com/neocortical/log5go"
)

const LOGGING_PREFIX = "MEDIAOS"

var log = l5g.Logger(l5g.LogAll)

func InjectLogger(l l5g.Log5Go) {
  log = l
}
