package cli_test

import (
	"fmt"
	"testing"

	"github.com/filipegorges/hex/adapters/cli"
	mock_application "github.com/filipegorges/hex/application/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	name := "Product Test"
	price := 25.99
	status := "enabled"
	id := "abc"

	mock := mock_application.NewMockProductInterface(ctrl)

	mock.EXPECT().GetName().Return(name).AnyTimes()
	mock.EXPECT().GetPrice().Return(price).AnyTimes()
	mock.EXPECT().GetStatus().Return(status).AnyTimes()
	mock.EXPECT().GetID().Return(id).AnyTimes()

	svc := mock_application.NewMockProductServiceInterface(ctrl)
	svc.EXPECT().Create(name, price).Return(mock, nil).AnyTimes()
	svc.EXPECT().Get(id).Return(mock, nil).AnyTimes()
	svc.EXPECT().Enable(mock).Return(mock, nil).AnyTimes()
	svc.EXPECT().Disable(mock).Return(mock, nil).AnyTimes()

	expCreate := fmt.Sprintf("Product ID %s with the name %s has the price of %f and is %s", id, name, price, status)
	expEnable := fmt.Sprintf("Product %s has been enabled", name)
	expDisable := fmt.Sprintf("Product %s has been disabled", name)
	expDefault := fmt.Sprintf("Product ID %s\nName %s\nPrice %f\nStatus %s\n", id, name, price, status)

	result, err := cli.Run(svc, "create", "", name, price)
	require.Nil(t, err)
	require.Equal(t, expCreate, result)

	result, err = cli.Run(svc, "enable", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expEnable, result)

	result, err = cli.Run(svc, "disable", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expDisable, result)

	result, err = cli.Run(svc, "", id, "", 0)
	require.Nil(t, err)
	require.Equal(t, expDefault, result)
}
