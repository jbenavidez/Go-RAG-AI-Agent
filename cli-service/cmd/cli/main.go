package main

import (
	"bufio"
	pb "clipService/proto/generated"
	"fmt"
	"os"
	"strings"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Application struct {
	GRPCClient pb.AIAgentServiceClient
}

func main() {
	//set grpc client
	var app Application
	//set grp client
	conn, err := grpc.Dial("rag-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		fmt.Println("unable to connected to grpc server")
		panic(err)
	}
	client := pb.NewAIAgentServiceClient(conn)
	app.GRPCClient = client
	//init cli
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to the NYC Capital Project RAG AI CLI")
	fmt.Println("Ask any question related to nyc capitals projects from 2023-2025.")
	fmt.Println("The AI agent will provide answers based on the data provided by data.cityofnewyork.us")
	fmt.Println("Type your question and press Enter. Type 'exit' to quit.\n")

	for {
		fmt.Print("\nEnter question: ")
		question, _ := reader.ReadString('\n')
		question = strings.TrimSpace(question)
		if question == "exit" {
			break
		}

		fmt.Println("\nAnswer:")

	}
}
