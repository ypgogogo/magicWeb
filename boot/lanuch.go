package boot

import (
	"github.com/yamakiller/magicLibs/args"
	"github.com/yamakiller/magicLibs/logger"
	"github.com/yamakiller/magicWeb/frame"
)

//Lanuch doc
//@Method Lanuch @Summary lanuch web system
//@Param (frame.Spawn) Spawn framework function
func Lanuch(spawn frame.Spawn) {
	frm := spawn()
	var err error
	args.Instance().Parse()
	if err = frm.Start(); err != nil {
		logger.Error(0, "%+v", err)
		goto exit
	}

exit:
	frm.Shutdown()
}
