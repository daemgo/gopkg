package id

import (
	"github.com/daemgo/gopkg/pkg/utils"

	"github.com/pkg/errors"
	"github.com/sony/sonyflake"
)

type snowflake sonyflake.Sonyflake

// NextID generates an ID.
func (s *snowflake) NextID() ID {
	uid, err := (*sonyflake.Sonyflake)(s).NextID()
	if err != nil {
		panic("get sony flake uid failed:" + err.Error())
	}
	return ID(uid)
}

// NewIDGenerator returns an IDGenerator object.
func NewIDGenerator() (IDGenerator, error) {
	ips, err := utils.GetLocalIPs()
	if err != nil {
		panic(err)
	}
	sf := (*snowflake)(sonyflake.NewSonyflake(sonyflake.Settings{
		MachineID: func() (u uint16, e error) {
			return utils.SumIPs(ips), nil
		},
	}))
	if sf == nil {
		return nil, errors.New("failed to new snoyflake object")
	}
	return sf, nil
}
