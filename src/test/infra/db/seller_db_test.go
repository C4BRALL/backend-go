package test

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	sellerDb "github.com/backend/src/internal/infra/db"
	"github.com/stretchr/testify/assert"

	"github.com/backend/src/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func TestCreateSeller(t *testing.T) {
	// Crie uma conexão de banco de dados simulada com sqlmock
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Erro ao criar mock de banco de dados: %v", err)
	}
	defer db.Close()

	// Inicialize o GORM com o driver SQL mock
	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})
	if err != nil {
		t.Fatalf("Erro ao criar conexão GORM: %v", err)
	}

	// Configurar as expectativas para a existência da tabela
	mock.ExpectQuery("SELECT count(*) FROM information_schema.tables WHERE table_schema = CURRENT_SCHEMA() AND table_name = 'sellers' AND table_type = 'BASE TABLE'").
		WillReturnRows(sqlmock.NewRows([]string{"count"}).AddRow(0))

	// Configurar as expectativas para criar a tabela
	mock.ExpectExec(`CREATE TABLE "sellers" .*`).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Realizar migrações
	if err := gormDB.AutoMigrate(&entity.Seller{}); err != nil {
		t.Fatalf("Erro ao executar migrações: %v", err)
	}

	// Criar um vendedor simulado
	seller, _ := entity.NewSeller("teste", "teste@mail.com", "12312312332", "passphrase", "85912341234")

	// Configurar o mock para esperar uma chamada para a função Create
	mock.ExpectExec("INSERT INTO sellers").
		WithArgs(seller.ID, seller.Name, seller.Email, seller.Document, seller.Password, seller.Phone, seller.Type, seller.Status).
		WillReturnResult(sqlmock.NewResult(1, 1))

	// Criar uma instância do repositório do vendedor com o GORM mockado
	sDB := sellerDb.NewSeller(gormDB)

	// Chamar a função Create
	err = sDB.Create(seller)

	// Verificar se não há erros
	assert.Nil(t, err)

	// Garantir que todas as expectativas do mock tenham sido satisfeitas
	assert.Nil(t, mock.ExpectationsWereMet())
}
