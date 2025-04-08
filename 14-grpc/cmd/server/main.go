package main

import (
	"database/sql"
	"log"
	"net"

	_ "github.com/mattn/go-sqlite3" // SQLite driver
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"grpc/internal/pb"
	"grpc/internal/service"
)

func main() {
	// Initialize database connection for SQLite
	db, err := sql.Open("sqlite3", "./database.db")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Create tables if they don't exist
	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS categories (
		id TEXT PRIMARY KEY,
		name TEXT NOT NULL,
		description TEXT
	)`)
	if err != nil {
		log.Fatalf("Failed to create tables: %v", err)
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Failed to ping database: %v", err)
	}

	// Create gRPC server
	grpcServer := grpc.NewServer()

	// Register services
	categoryService := service.NewCategoryService(db)
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	// Enable reflection
	reflection.Register(grpcServer)

	// Start server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	log.Println("gRPC Server running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
