package app

import (
	pb "DeNet-test_Task/docs/grpc/gen"
	"DeNet-test_Task/internal/domain/entity"
	"DeNet-test_Task/internal/domain/repository"
	"DeNet-test_Task/internal/domain/service"
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
)

type AccountService interface {
	GetAccount(address, signature string) (*entity.Account, error)
	GetAccounts(addresses []string, tokenAddress string) ([]entity.ERC20Balance, error)
}

type AccountApp struct {
	accountServ AccountService
}

type grpcServer struct {
	pb.UnimplementedAccountServiceServer
	app *AccountApp
}

func New() *AccountApp {
	return &AccountApp{}
}

func (app *AccountApp) Run() error {
	lis, err := net.Listen("tcp", "50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	repo := repository.NewAccountRepository()
	serv := service.NewAccountService(repo)
	app.accountServ = serv
	// accountApp := NewAccountApp(serv)

	// grpcServer := &grpcServer{app: accountApp}

	s := grpc.NewServer()
	pb.RegisterAccountServiceServer(s, &grpcServer{})
	log.Println("Server is running on port 50051")
	return s.Serve(lis)
}

func (g *grpcServer) GetAccount(ctx context.Context, req *pb.GetAccountRequest) (*pb.GetAccountResponse, error) {
	account, err := g.app.accountServ.GetAccount(req.EthereumAddress, req.CryptoSignature)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &pb.GetAccountResponse{
		GastokenBalance: account.GastokenBalance,
		WalletNonce:     account.WalletNonce,
	}, nil
}

func (g *grpcServer) GetAccounts(req *pb.GetAccountsRequest, stream pb.AccountService_GetAccountsServer) error {
	balances, err := g.app.accountServ.GetAccounts(req.EthereumAddresses, req.Erc20TokenAddress)
	if err != nil {
		log.Println(err)
	}

	// Отправляем результаты в поток
	for _, balance := range balances {
		err = stream.Send(&pb.GetAccountsResponse{
			EthereumAddress: balance.EthereunAddress,
			Erc20Balance:    balance.Balance,
		})
		if err != nil {
			log.Println(err)
		}
	}

	return nil
}
