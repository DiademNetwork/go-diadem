package commands

import (
	"errors"
	"strconv"

	"github.com/spf13/cobra"

	tgtypes "github.com/diademnetwork/go-diadem/builtin/types/transfer_gateway"
	"github.com/diademnetwork/go-diadem/cli"
)

const (
	GatewayName        = "gateway"
	diademGatewayName    = "diademcoin-gateway"
	BinanceGatewayName = "binance-gateway"
)

func UnsafeResetBlockCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "unsafe-reset-last-eth-block <blockNumber> <gatewayType>",
		Short: "WARNING: Resets the Ethereum block number used by the Gateway to sync with Ethereum",
		Args:  cobra.MinimumNArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			var block uint64
			var err error
			if len(args) == 0 {
				block = uint64(0)
			} else {
				block, err = strconv.ParseUint(args[0], 10, 64)
				if err != nil {
					return err
				}
			}

			var name string
			switch args[1] {
			case GatewayName:
				name = GatewayName
			case diademGatewayName:
				name = diademGatewayName
			case BinanceGatewayName:
				name = BinanceGatewayName
			default:
				return errors.New("invalid gateway name")
			}

			return cli.CallContract(name, "ResetMainnetBlock", &tgtypes.TransferGatewayResetMainnetBlockRequest{
				LastMainnetBlockNum: block,
			}, nil)
		},
	}
}

func AddUnsafeCommands(root *cobra.Command) {
	root.AddCommand(
		UnsafeResetBlockCmd(),
	)
}
