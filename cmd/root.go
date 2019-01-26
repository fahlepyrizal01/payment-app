//menunjukan bahwa file ini berada di package mana
package cmd

//melakukan import package yang diperlukan
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	_ "github.com/lib/pq"

	pengguna "github.com/fahlepyrizal01/payment-app/pengguna"
	transaksi "github.com/fahlepyrizal01/payment-app/transaksi"
)

//variabel global
var (
	//variabel untuk database
	dbPool *sql.DB
	//variabel untuk file configuration
	cfgFile string
)

//variabel untuk running command service server dengan package cobra
var rootCmd = &cobra.Command{
	//mengunakan nama payment
	Use: "payment",
	//fungsi untuk menjalankan service server
	Run: func(cmd *cobra.Command, args []string) {
		//menambahkan cancel pada context
		ctx, cancel := context.WithCancel(SignalContext(context.Background()))
		defer cancel()
		//menjalankan fungsi initDB
		InitDB()
		//transfer db ke package transaksi
		transaksi.Init(dbPool)
		pengguna.Init(dbPool)
		//menjalankan server
		lis, err := net.Listen("tcp", fmt.Sprint(":", viper.GetInt("app.port")))
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
		}
		s := grpc.NewServer()

		//menjalankan server transaksi
		pengguna.RegisterPenggunaServiceServer(s, &pengguna.Pengguna_server{})
		transaksi.RegisterTransaksiServiceServer(s, &transaksi.Transaksi_server{})
		reflection.Register(s)

		//melakukan check pada proses yang berjalan
		go func() {
			if err := s.Serve(lis); err != nil {
				log.Fatalf("failed to serve: %v", err)
			}
			cancel()
		}()

		//memberitahu bahwa proses telah selesai
		<-ctx.Done()

		//menghentikan service
		s.GracefulStop()

	},
}

//fungsi untuk menginisialisasi configuration dengan package cobra dengan parameter fungsi
//initConfig
func init() {
	cobra.OnInitialize(initConfig)
}

//fungsi untuk mengeksekuasi package cmd
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

//fungsi untuk menginisialisasi configuration
func initConfig() {
	//membuat type configuration toml
	viper.SetConfigType("toml")
	//mengatur configuration dengan variabel cfgFile jika tidak kosong
	//membuat configuration path jika variabel cfgFile kosong
	if cfgFile != "" {

		viper.SetConfigFile(cfgFile)
	} else {

		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		viper.AddConfigPath(".")
		viper.AddConfigPath(home)
		viper.AddConfigPath("/etc/payment")
		viper.SetConfigName(".payment")
	}

	viper.AutomaticEnv()

	//mealakukan check error pada configuration
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

//fungsi untuk menginisialisasi database
func InitDB() {
	var db, err = sql.Open("postgres", fmt.Sprintf("postgresql://%s@localhost:%d/%s?sslmode=disable", viper.GetString("database.username"), viper.GetInt("database.port"), viper.GetString("database.name")))
	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	dbPool = db
}

//fungsi untuk memberitahu bahwa proses masih berjalan atau tidak
func SignalContext(ctx context.Context) context.Context {
	//menmbahkan cancel pada context
	ctx, cancel := context.WithCancel(ctx)
	//memberitahu proses yang beralan
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	//fungsi untuk menghentikan signal
	go func() {

		<-sigs

		signal.Stop(sigs)
		close(sigs)
		cancel()
	}()

	return ctx
}
