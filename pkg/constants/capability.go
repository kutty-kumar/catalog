package constants

import commons "github.com/kutty-kumar/commons/set"

type Capability int

const (
	Chat = 1
	Audio = 2
	Video = 4
	Physical = 8
)

func (capability Capability) deriveCapabilities(bitMask int) commons.Set {
	result := commons.NewThreadUnsafeSet()
	 masks := []int{Chat, Audio, Video, Physical}
	  for _, mask := range masks {
	  	if mask & bitMask > 0 {
	  		result.Add(mask)
		}
	  }
	return result
}

func (capability Capability) deriveBitMask(capabilities commons.Set) int {
	bitMask := 0
	for capability := range capabilities.Iter() {
		capabilityInt := capability.(int)
		bitMask |= capabilityInt
	}
	return bitMask
}
