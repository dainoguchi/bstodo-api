package usecase

import (
	"context"
	"github.com/dainoguchi/bstodo-api/internal/infra/sqlc"
	"github.com/dainoguchi/bstodo-api/internal/usecase/input"
	"github.com/dainoguchi/bstodo-api/test/testutil"
	"github.com/jackc/pgx/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"testing"
	"time"
)

type TodoUsecaseSuite struct {
	suite.Suite
	db  *pgx.Conn
	ctx context.Context
	now time.Time
}

func TestTodoUsecaseSuite(t *testing.T) {
	suite.Run(t, new(TodoUsecaseSuite))
}

func (s *TodoUsecaseSuite) BeforeTest(suiteName string, testName string) {
	s.ctx = context.Background()
	// 自国を固定
	s.now = time.Date(2022, 5, 10, 12, 34, 56, 0, time.UTC)
	s.db = testutil.OpenDBForTest(s.T())
}

func (s *TodoUsecaseSuite) AfterTest(suiteName string, testName string) {
	if err := s.db.Close(context.Background()); err != nil {
		s.T().Fatal(err)
	}
}

func (s *TodoUsecaseSuite) TestCreateTodo() {
	s.T().Helper()
	s.T().Parallel()

	ctx := s.ctx
	tx, err := testutil.OpenDBForTest(s.T()).Begin(ctx)
	// 終了後必ずrollbackすることでdbにデータを残さない
	s.T().Cleanup(func() { _ = tx.Rollback(ctx) })
	if err != nil {
		s.T().Fatalf("failed to create transaction %v", err)
	}

	type args struct {
		ctx   context.Context
		input *input.CreateTodoInput
	}

	tests := []struct {
		name string
		args args
		//want *sqlc.Todo
	}{
		{
			name: "正常系",
			args: args{
				context.Background(),
				&input.CreateTodoInput{
					Auth0ID:     "auth0id",
					Title:       "title",
					Priority:    "high",
					Description: testutil.ToStrP("description"),
					DueDate:     testutil.ToTimeP(s.now),
				},
			},
		},
		{
			name: "正常系: descriptionがnil",
			args: args{
				context.Background(),
				&input.CreateTodoInput{
					Auth0ID:  "auth0id",
					Title:    "title",
					Priority: "high",
					DueDate:  testutil.ToTimeP(s.now),
				},
			},
		},
		{
			name: "正常系: duedateがnil",
			args: args{
				context.Background(),
				&input.CreateTodoInput{
					Auth0ID:     "auth0id",
					Title:       "title",
					Priority:    "high",
					Description: testutil.ToStrP("description"),
					DueDate:     testutil.ToTimeP(s.now),
				},
			},
		},
	}

	for _, tt := range tests {
		s.Run(tt.name, func() {
			uc := &todoUsecase{
				db: tx,
			}

			_, err := uc.CreateTodo(tt.args.ctx, tt.args.input)
			assert.NoError(s.T(), err)
		})
	}

	// 異常系
	invalidTests := []struct {
		name string
		args args
		want *sqlc.Todo
	}{
		{
			name: "異常系: priorityがhigh, mid, low以外",
			args: args{
				context.Background(),
				&input.CreateTodoInput{
					Auth0ID:     "auth0id",
					Title:       "title",
					Priority:    "normal",
					Description: testutil.ToStrP("description"),
					DueDate:     testutil.ToTimeP(s.now),
				},
			},
		},
		{
			name: "異常系: Auth0IDがnil",
			args: args{
				context.Background(),
				&input.CreateTodoInput{
					Title:       "title",
					Priority:    "high",
					Description: testutil.ToStrP("description"),
					DueDate:     testutil.ToTimeP(s.now),
				},
			},
		},
		{
			name: "異常系: titleが１文字未満",
			args: args{
				context.Background(),
				&input.CreateTodoInput{
					Title:       "",
					Priority:    "high",
					Description: testutil.ToStrP("description"),
					DueDate:     testutil.ToTimeP(s.now),
				},
			},
		},
	}

	for _, tt := range invalidTests {
		s.Run(tt.name, func() {
			uc := &todoUsecase{
				db: tx,
			}

			_, err := uc.CreateTodo(tt.args.ctx, tt.args.input)
			assert.Error(s.T(), err)
		})
	}
}
