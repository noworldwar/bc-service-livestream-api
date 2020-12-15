package srvclient

import (
	ppf "github.com/PGITAb/bc-proto-entity-playerprofile-go"
	wal "github.com/PGITAb/bc-proto-wallet-go"
	"github.com/micro/go-micro/client/grpc"
)

var (
	PlayerClient ppf.Service
	WalletClient wal.Service
)

func Init() {
	PlayerClient = ppf.NewService("pg.srv.entity.playerprofile", grpc.NewClient())
	WalletClient = wal.NewService("pg.srv.wallet", grpc.NewClient())
}
