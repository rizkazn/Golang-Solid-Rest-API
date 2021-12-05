package repos

import (
	"github.com/rizkazn/gorestfull/models"
	"github.com/stretchr/testify/mock"
)

type HistoryMockRepos struct {
	Mock mock.Mock
}

var dataHist = models.History{
	Id:         1,
	No_invoice: "No Invoice for Testing",
	Cashier:    "Cashier for Testing",
	Date:       "Date for Testing",
	Orders:     1,
	Amount:     50000,
}

var datumHist = models.Histories{
	dataHist,
}

func (mock *HistoryMockRepos) FindAll() (*models.Histories, error) {
	args := mock.Mock.Called()
	// result := args.Get(0)
	// return result.(*models.Histories), args.Error(1)
	if args == nil {
		return nil, nil
	}

	return &datumHist, nil
}

func (mock *HistoryMockRepos) Save(history *models.History) error {
	args := mock.Mock.Called(history)

	if args == nil {
		return nil
	}

	return nil
}

func (mock *HistoryMockRepos) Update(history *models.History, id string) error {
	args := mock.Mock.Called(history, id)

	if args == nil {
		return nil
	}

	return nil
}

func (mock *HistoryMockRepos) Remove(id string) error {
	args := mock.Mock.Called(id)

	if args == nil {
		return nil
	}

	return nil
}
