package config

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"go.mongodb.org/mongo-driver/x/bsonx"
	"log"
	"my-tracking-list-backend/core/domain"
	"my-tracking-list-backend/infrastructure/repository"
	"os"
	"reflect"
	"strings"
	"time"
)

func DatabaseConfig() *mongo.Database {
	credentials := options.Credential{
		Username: os.Getenv("MY_TRACKING_LIST_DATABASE_USERNAME"),
		Password: os.Getenv("MY_TRACKING_LIST_DATABASE_PASSWORD"),
	}
	uri := fmt.Sprintf(
		"mongodb://%s:%s",
		os.Getenv("MY_TRACKING_LIST_DATABASE_HOST"),
		os.Getenv("MY_TRACKING_LIST_DATABASE_PORT"),
	)

	client, err := mongo.NewClient(
		options.Client().
			ApplyURI(uri).
			SetAuth(credentials),
	)
	if err != nil {
		log.Fatalf("Erro ao criar cliente do banco de dados: %v", err)
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	// todo: ver como fechar conexao
	// n pode chamar aqui pq metodo eh chamado e finalizado aqui
	// solucao: criar contexto no main() e ir passando pra baixo
	//defer client.Disconnect(ctx)

	if err != nil {
		log.Fatalf("Erro ao conectar com banco de dados: %v", err)
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatalf("Erro ao pingar banco de dados: %v", err)
	}

	db := client.Database(os.Getenv("MY_TRACKING_LIST_DATABASE_NAME"))

	createIndex(db)

	return db
}

func createIndex(db *mongo.Database) {
	userEmailUniqueIndex(db)
}

func userEmailUniqueIndex(db *mongo.Database) {
	emailField, ok := reflect.TypeOf(&domain.User{}).Elem().FieldByName("Email")
	if !ok {
		log.Fatal("Nao foi possivel obter o field email para criar o index")
	}

	emailTag := string(emailField.Tag.Get("bson"))
	if strings.TrimSpace(emailTag) == "" {
		log.Fatal("Nao foi possivel obter tag email do field Email da struct domain.User{}")
	}

	_, err := db.Collection(repository.UserCollectionName).Indexes().CreateOne(
		context.Background(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: emailTag, Value: bsonx.String("text")}},
			Options: options.Index().SetUnique(true),
		},
	)
	if err != nil {
		log.Fatalf("Erro ao criar index para o campo 'email': %v", err)
	}
}
