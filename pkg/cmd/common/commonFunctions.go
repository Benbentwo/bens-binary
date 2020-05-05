package common

import (
	"github.com/Benbentwo/utils/util"
	"github.com/go-errors/errors"
)

var (
	UnimplementedText = "Reached Unimplemented code"
)

func ErrorUnimplemented() error {
	return errors.Errorf(util.ColorInfo(Shrug) + ": " + UnimplementedText)
}
