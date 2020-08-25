package types

import (
	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/params"
	"github.com/ixofoundation/ixo-blockchain/x/did"
)

// Parameter store keys
var (
	KeyIxoDid                       = []byte("IxoDid")
	KeyProjectMinimumInitialFunding = []byte("ProjectMinimumInitialFunding")
)

// project parameters
type Params struct {
	IxoDid                       did.Did `json:"ixo_did" yaml:"ixo_did"`
	ProjectMinimumInitialFunding sdk.Int `json:"project_minimum_initial_funding" yaml:"project_minimum_initial_funding"`
}

// ParamTable for project module.
func ParamKeyTable() params.KeyTable {
	return params.NewKeyTable().RegisterParamSet(&Params{})
}

func NewParams(projectMinimumInitialFunding sdk.Int, ixoDid did.Did) Params {
	return Params{
		IxoDid:                       ixoDid,
		ProjectMinimumInitialFunding: projectMinimumInitialFunding,
	}

}

// default project module parameters
func DefaultParams() Params {
	return Params{
		IxoDid:                       did.Did(""),  // blank
		ProjectMinimumInitialFunding: sdk.OneInt(), // 1
	}
}

func (p Params) String() string {
	return fmt.Sprintf(`Project Params:
  Ixo Did: %s
  Project Minimum Initial Funding: %s

`, p.ProjectMinimumInitialFunding, p.IxoDid)
}

func validateIxoDid(i interface{}) error {
	v := i.(sdk.Dec)

	if v.IsNil() {
		return fmt.Errorf("ixo did must be positive: %s", v)
	}

	return nil
}

func validateProjectMinimumInitialFunding(i interface{}) error {
	v := i.(sdk.Dec)

	if v.IsNil() {
		return fmt.Errorf("project minimum initial funding must be positive: %s", v)
	}

	return nil
}

// Implements params.ParamSet
func (p *Params) ParamSetPairs() params.ParamSetPairs {
	return params.ParamSetPairs{
		{KeyIxoDid, &p.IxoDid, validateIxoDid},
		{KeyProjectMinimumInitialFunding, &p.ProjectMinimumInitialFunding, validateProjectMinimumInitialFunding},
	}
}
