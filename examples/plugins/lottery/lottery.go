package lottery

import (
	diadem "github.com/diademnetwork/go-diadem"
	"github.com/diademnetwork/go-diadem/builtin/types/coin"
	"github.com/diademnetwork/go-diadem/plugin"
	contract "github.com/diademnetwork/go-diadem/plugin/contractpb"
	"github.com/diademnetwork/go-diadem/types"
)

type Lottery struct {
}

var coinContractKey = []byte("coincontract")

func transfer(ctx contract.Context, to diadem.Address, amount *diadem.BigUInt) error {
	req := &coin.TransferRequest{
		To:     to.MarshalPB(),
		Amount: &types.BigUInt{Value: *amount},
	}

	coinAddr, err := ctx.Resolve("coin")
	if err != nil {
		return err
	}
	return contract.Call(ctx, coinAddr, req, nil)
}

func (c *Lottery) Meta() (plugin.Meta, error) {
	return plugin.Meta{
		Name:    "lottery",
		Version: "1.0.0",
	}, nil
}

func (c *Lottery) Init(ctx contract.Context, req *LotteryInit) {
	winnerAddr := diadem.UnmarshalAddressPB(req.Winner)
	transfer(ctx, winnerAddr, diadem.NewBigUIntFromInt(1000))
}

var Contract plugin.Contract = contract.MakePluginContract(&Lottery{})

func main() {
	plugin.Serve(Contract)
}
