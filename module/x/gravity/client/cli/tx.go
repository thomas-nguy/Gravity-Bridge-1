package cli

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	"github.com/spf13/cobra"

	"github.com/Gravity-Bridge/Gravity-Bridge/module/x/gravity/types"
)

func GetTxCmd(storeKey string) *cobra.Command {
	//nolint: exhaustivestruct
	gravityTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "Gravity transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	gravityTxCmd.AddCommand([]*cobra.Command{
		CmdSendToEth(),
		CmdRequestBatch(),
		CmdSetOrchestratorAddress(),
		CmdGovIbcMetadataProposal(),
		CmdGovAirdropProposal(),
		CmdGovUnhaltBridgeProposal(),
	}...)

	return gravityTxCmd
}

func CmdGovIbcMetadataProposal() *cobra.Command {
	//nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:   "gov-ibc-metadata [path-to-proposal-json] [initial-deposit]",
		Short: "Creates a governance proposal to set the Metadata of the given IBC token. Once the metadata is set this token can be moved to Ethereum using Gravity Bridge",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cosmosAddr := cliCtx.GetFromAddress()

			initialDeposit, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "bad initial deposit amount")
			}

			if len(initialDeposit) > 1 {
				return fmt.Errorf("coin amounts too long, expecting just 1 coin amount for both amount and bridgeFee")
			}

			proposalFile := args[0]

			contents, err := os.ReadFile(proposalFile)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to read proposal json file")
			}

			proposal := &types.IBCMetadataProposal{}
			err = json.Unmarshal(contents, proposal)
			if err != nil {
				return sdkerrors.Wrap(err, "proposal json file is not valid json")
			}

			proposalAny, err := codectypes.NewAnyWithValue(proposal)
			if err != nil {
				return sdkerrors.Wrap(err, "invalid metadata or proposal details!")
			}

			// Make the message
			msg := govtypes.MsgSubmitProposal{
				Proposer:       cosmosAddr.String(),
				InitialDeposit: initialDeposit,
				Content:        proposalAny,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// Send it
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdGovAirdropProposal() *cobra.Command {
	//nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:   "gov-airdrop [path-to-proposal-json] [initial-deposit]",
		Short: "Creates a governance proposal for an airdrop",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cosmosAddr := cliCtx.GetFromAddress()

			initialDeposit, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "bad initial deposit amount")
			}

			if len(initialDeposit) > 1 {
				return fmt.Errorf("coin amounts too long, expecting just 1 coin amount for both amount and bridgeFee")
			}

			proposalFile := args[0]

			contents, err := os.ReadFile(proposalFile)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to read proposal json file")
			}

			proposal := &types.AirdropProposal{}
			err = json.Unmarshal(contents, proposal)
			if err != nil {
				return sdkerrors.Wrap(err, "proposal json file is not valid json")
			}

			proposalAny, err := codectypes.NewAnyWithValue(proposal)
			if err != nil {
				return sdkerrors.Wrap(err, "invalid metadata or proposal details!")
			}

			// Make the message
			msg := govtypes.MsgSubmitProposal{
				Proposer:       cosmosAddr.String(),
				InitialDeposit: initialDeposit,
				Content:        proposalAny,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// Send it
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdGovUnhaltBridgeProposal() *cobra.Command {
	//nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:   "gov-unhalt-bridge [path-to-proposal-json] [initial-deposit]",
		Short: "Creates a governance proposal to unhalt the Ethereum bridge after an oracle dispute",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cosmosAddr := cliCtx.GetFromAddress()

			initialDeposit, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "bad initial deposit amount")
			}

			if len(initialDeposit) > 1 {
				return fmt.Errorf("coin amounts too long, expecting just 1 coin amount for both amount and bridgeFee")
			}

			proposalFile := args[0]

			contents, err := os.ReadFile(proposalFile)
			if err != nil {
				return sdkerrors.Wrap(err, "failed to read proposal json file")
			}

			proposal := &types.UnhaltBridgeProposal{}
			err = json.Unmarshal(contents, proposal)
			if err != nil {
				return sdkerrors.Wrap(err, "proposal json file is not valid json")
			}

			proposalAny, err := codectypes.NewAnyWithValue(proposal)
			if err != nil {
				return sdkerrors.Wrap(err, "invalid metadata or proposal details!")
			}

			// Make the message
			msg := govtypes.MsgSubmitProposal{
				Proposer:       cosmosAddr.String(),
				InitialDeposit: initialDeposit,
				Content:        proposalAny,
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// Send it
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdSendToEth() *cobra.Command {
	//nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:   "send-to-eth [eth-dest] [amount] [bridge-fee]",
		Short: "Adds a new entry to the transaction pool to withdraw an amount from the Ethereum bridge contract",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cosmosAddr := cliCtx.GetFromAddress()

			amount, err := sdk.ParseCoinsNormalized(args[1])
			if err != nil {
				return sdkerrors.Wrap(err, "amount")
			}
			bridgeFee, err := sdk.ParseCoinsNormalized(args[2])
			if err != nil {
				return sdkerrors.Wrap(err, "bridge fee")
			}

			ethAddr, err := types.NewEthAddress(args[0])
			if err != nil {
				return sdkerrors.Wrap(err, "invalid eth address")
			}

			if len(amount) > 1 || len(bridgeFee) > 1 {
				return fmt.Errorf("coin amounts too long, expecting just 1 coin amount for both amount and bridgeFee")
			}

			// Make the message
			msg := types.MsgSendToEth{
				Sender:    cosmosAddr.String(),
				EthDest:   ethAddr.GetAddress(),
				Amount:    amount[0],
				BridgeFee: bridgeFee[0],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// Send it
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdRequestBatch() *cobra.Command {
	//nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:   "build-batch [token_contract_address]",
		Short: "Build a new batch on the cosmos side for pooled withdrawal transactions",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			cosmosAddr := cliCtx.GetFromAddress()

			// TODO: better denom searching
			msg := types.MsgRequestBatch{
				Sender: cosmosAddr.String(),
				Denom:  fmt.Sprintf("gravity%s", args[0]),
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// Send it
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}

func CmdSetOrchestratorAddress() *cobra.Command {
	//nolint: exhaustivestruct
	cmd := &cobra.Command{
		Use:   "set-orchestrator-address [validator-address] [orchestrator-address] [ethereum-address]",
		Short: "Allows validators to delegate their voting responsibilities to a given key.",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}
			msg := types.MsgSetOrchestratorAddress{
				Validator:    args[0],
				Orchestrator: args[1],
				EthAddress:   args[2],
			}
			if err := msg.ValidateBasic(); err != nil {
				return err
			}
			// Send it
			return tx.GenerateOrBroadcastTxCLI(cliCtx, cmd.Flags(), &msg)
		},
	}
	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
