package swolf

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/noyako/mock_swolf"
)

func TestDealerGet(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mock.NewMasterController(mockCtrl)
	// mockTenant := mock.NewTenantController(mockCtrl)

	// d := Dealer{mockMaser, mockTenant}

	// mockMaser.EXPECT().
}
