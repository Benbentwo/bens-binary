package common

import (
	"github.com/Benbentwo/utils/util"
	"github.com/go-errors/errors"
)

var (
	UnimplementedText = "Reached Unimplemented code"
)

func ErrorUnimplemented() error {
	util.Logger().Errorf(util.ColorInfo(Shrug) + ": " + UnimplementedText)
	return errors.New("Unimplemented Error")
}
