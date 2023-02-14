package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"

	pb "beverage-grpc/pkg"
)

const sqlStorageTypeName = "sql"

type sql struct {
	conn pgx.Conn
}

func NewSqlStorage() *sql {
	return &sql{}
}

func (s *sql) Add(ctx context.Context, beverage *pb.Beverage) (*pb.Beverage, error) {
	if beverage == nil {
		log.Fatalf("sql add: nil beverage provided")
	}
	err := s.createBeverageTable(ctx)
	if err != nil {
		log.Fatalf("beverage table creation failed: %v", err)
	}

	// transaction is not necessary here
	tx, err := s.conn.Begin(ctx)
	if err != nil {
		log.Fatalf("cannot start transaction: %v", err)
	}

	query := "INSERT INTO beverages(type, name, volume, price) VALUES ($1, $2, $3, $4)"

	_, err = tx.Exec(
		ctx,
		query,
		beverage.Attr.Type,
		beverage.Attr.Name,
		beverage.Attr.Volume,
		beverage.Price,
	)
	if err != nil {
		log.Fatalf("error while inserting beverage into db: %v", err)
	}

	err = tx.Commit(ctx)
	if err != nil {
		log.Fatalf("sql add: error to commit transaction: %v", err)
	}

	return beverage, nil
}

func (s *sql) List(ctx context.Context) (*pb.BeverageList, error) {
	var list = &pb.BeverageList{}

	query := "SELECT * FROM beverages"
	rows, err := s.conn.Query(ctx, query)
	if err != nil {
		log.Fatalf("sql list: error getting list of beverages: %v", err)
	}
	defer rows.Close()

	for rows.Next() {
		b := pb.Beverage{}
		err = rows.Scan(&b)
		if err != nil {
			log.Fatalf("sql list: error while scaning beverage record: %v", err)
		}

		list.Beverages = append(list.Beverages, &b)
	}

	return list, nil
}

// You should always use migrations mechanism in a production project
// In this project, we will do all the actions with the base inside the code
func (s *sql) createBeverageTable(ctx context.Context) error {
	query := `
			CREATE TABLE IF NOT EXISTS beverages(
			    id SERIAL PRIMARY KEY,
				type int,
			    name text,
			    volume int,
			    price int
			);
`
	_, err := s.conn.Exec(ctx, query)
	if err != nil {
		return err
	}

	return nil
}
