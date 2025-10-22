# log

## Usage

```lua
local log = require("log")

log.debug("message")
log.info("message")
log.warn("message")
log.error("message")

-- format
log.debugf("%s %s", "message", "message")
log.infof("%s %s", "message", "message")
log.warnf("%s %s", "message", "message")
log.errorf("%s %s", "message", "message")
```