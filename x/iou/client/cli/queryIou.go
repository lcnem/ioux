package cli

import (
	"context"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/lcnem/ioux/x/iou/types"
	"github.com/spf13/cobra"
)

func CmdListIou() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list-namespace",
		Short: "list all namespace",
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			pageReq, err := client.ReadPageRequest(cmd.Flags())
			if err != nil {
				return err
			}

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryAllIouNamespaceRequest{
				Pagination: pageReq,
			}

			res, err := queryClient.IouNamespaceAll(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}

func CmdShowIou() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "show-namespace [id]",
		Short: "shows a iou namespace",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)

			queryClient := types.NewQueryClient(clientCtx)

			params := &types.QueryGetIouNamespaceRequest{
				Id: args[0],
			}

			res, err := queryClient.IouNamespace(context.Background(), params)
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	flags.AddQueryFlagsToCmd(cmd)

	return cmd
}
