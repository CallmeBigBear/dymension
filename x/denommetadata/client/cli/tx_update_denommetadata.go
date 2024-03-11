package cli

import (
	"encoding/json"
	"os"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/tx"
	"github.com/spf13/cobra"

	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	"github.com/dymensionxyz/dymension/v3/utils"
	"github.com/dymensionxyz/dymension/v3/x/denommetadata/types"

	govcli "github.com/cosmos/cosmos-sdk/x/gov/client/cli"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types/v1beta1"
)

// NewCmdSubmitUpdateDenomMetadataProposal broadcasts a CreateMetadataProposal message.
func NewCmdSubmitUpdateDenomMetadataProposal() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update-denom-metadata-proposal metadata.json [flags]",
		Short: "proposal to update new denom metadata for a specific token",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			proposal, deposit, err := utils.ParseProposal(cmd)
			if err != nil {
				return err
			}

			path := args[0]

			//nolint:gosec
			fileContent, err := os.ReadFile(path)
			if err != nil {
				return err
			}

			metadata := banktypes.Metadata{}
			err = json.Unmarshal([]byte(fileContent), &metadata)
			if err != nil {
				return err
			}

			err = metadata.Validate()
			if err != nil {
				return err
			}

			content := types.NewUpdateDenomMetadataProposal(proposal.Title, proposal.Description, metadata)
			msg, err := govtypes.NewMsgSubmitProposal(content, deposit, clientCtx.GetFromAddress())
			if err != nil {
				return err
			}

			txf := tx.NewFactoryCLI(clientCtx, cmd.Flags()).WithTxConfig(clientCtx.TxConfig).WithAccountRetriever(clientCtx.AccountRetriever)
			return tx.GenerateOrBroadcastTxWithFactory(clientCtx, txf, msg)
		},
	}

	cmd.Flags().String(govcli.FlagTitle, "", "The proposal title")
	cmd.Flags().String(govcli.FlagDescription, "", "The proposal description")
	cmd.Flags().String(govcli.FlagDeposit, "", "The proposal deposit")

	return cmd
}
