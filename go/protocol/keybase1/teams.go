// Auto-generated by avdl-compiler v1.3.16 (https://github.com/keybase/node-avdl-compiler)
//   Input file: avdl/keybase1/teams.avdl

package keybase1

import (
	"github.com/keybase/go-framed-msgpack-rpc/rpc"
	context "golang.org/x/net/context"
)

type TeamRole int

const (
	TeamRole_NONE   TeamRole = 0
	TeamRole_OWNER  TeamRole = 1
	TeamRole_ADMIN  TeamRole = 2
	TeamRole_WRITER TeamRole = 3
	TeamRole_READER TeamRole = 4
)

func (o TeamRole) DeepCopy() TeamRole { return o }

var TeamRoleMap = map[string]TeamRole{
	"NONE":   0,
	"OWNER":  1,
	"ADMIN":  2,
	"WRITER": 3,
	"READER": 4,
}

var TeamRoleRevMap = map[TeamRole]string{
	0: "NONE",
	1: "OWNER",
	2: "ADMIN",
	3: "WRITER",
	4: "READER",
}

func (e TeamRole) String() string {
	if v, ok := TeamRoleRevMap[e]; ok {
		return v
	}
	return ""
}

type TeamApplication int

const (
	TeamApplication_KBFS     TeamApplication = 1
	TeamApplication_CHAT     TeamApplication = 2
	TeamApplication_SALTPACK TeamApplication = 3
)

func (o TeamApplication) DeepCopy() TeamApplication { return o }

var TeamApplicationMap = map[string]TeamApplication{
	"KBFS":     1,
	"CHAT":     2,
	"SALTPACK": 3,
}

var TeamApplicationRevMap = map[TeamApplication]string{
	1: "KBFS",
	2: "CHAT",
	3: "SALTPACK",
}

func (e TeamApplication) String() string {
	if v, ok := TeamApplicationRevMap[e]; ok {
		return v
	}
	return ""
}

type PerTeamKey struct {
	Gen    int   `codec:"gen" json:"gen"`
	Seqno  Seqno `codec:"seqno" json:"seqno"`
	SigKID KID   `codec:"sigKID" json:"sigKID"`
	EncKID KID   `codec:"encKID" json:"encKID"`
}

func (o PerTeamKey) DeepCopy() PerTeamKey {
	return PerTeamKey{
		Gen:    o.Gen,
		Seqno:  o.Seqno.DeepCopy(),
		SigKID: o.SigKID.DeepCopy(),
		EncKID: o.EncKID.DeepCopy(),
	}
}

type TeamMembers struct {
	Owners  []string `codec:"owners" json:"owners"`
	Admins  []string `codec:"admins" json:"admins"`
	Writers []string `codec:"writers" json:"writers"`
	Readers []string `codec:"readers" json:"readers"`
}

func (o TeamMembers) DeepCopy() TeamMembers {
	return TeamMembers{
		Owners: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Owners),
		Admins: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Admins),
		Writers: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Writers),
		Readers: (func(x []string) []string {
			var ret []string
			for _, v := range x {
				vCopy := v
				ret = append(ret, vCopy)
			}
			return ret
		})(o.Readers),
	}
}

type TeamCreateArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
}

func (o TeamCreateArg) DeepCopy() TeamCreateArg {
	return TeamCreateArg{
		SessionID: o.SessionID,
		Name:      o.Name,
	}
}

type TeamGetArg struct {
	SessionID int    `codec:"sessionID" json:"sessionID"`
	Name      string `codec:"name" json:"name"`
}

func (o TeamGetArg) DeepCopy() TeamGetArg {
	return TeamGetArg{
		SessionID: o.SessionID,
		Name:      o.Name,
	}
}

type TeamsInterface interface {
	TeamCreate(context.Context, TeamCreateArg) error
	TeamGet(context.Context, TeamGetArg) (TeamMembers, error)
}

func TeamsProtocol(i TeamsInterface) rpc.Protocol {
	return rpc.Protocol{
		Name: "keybase.1.teams",
		Methods: map[string]rpc.ServeHandlerDescription{
			"teamCreate": {
				MakeArg: func() interface{} {
					ret := make([]TeamCreateArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamCreateArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamCreateArg)(nil), args)
						return
					}
					err = i.TeamCreate(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
			"teamGet": {
				MakeArg: func() interface{} {
					ret := make([]TeamGetArg, 1)
					return &ret
				},
				Handler: func(ctx context.Context, args interface{}) (ret interface{}, err error) {
					typedArgs, ok := args.(*[]TeamGetArg)
					if !ok {
						err = rpc.NewTypeError((*[]TeamGetArg)(nil), args)
						return
					}
					ret, err = i.TeamGet(ctx, (*typedArgs)[0])
					return
				},
				MethodType: rpc.MethodCall,
			},
		},
	}
}

type TeamsClient struct {
	Cli rpc.GenericClient
}

func (c TeamsClient) TeamCreate(ctx context.Context, __arg TeamCreateArg) (err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamCreate", []interface{}{__arg}, nil)
	return
}

func (c TeamsClient) TeamGet(ctx context.Context, __arg TeamGetArg) (res TeamMembers, err error) {
	err = c.Cli.Call(ctx, "keybase.1.teams.teamGet", []interface{}{__arg}, &res)
	return
}
