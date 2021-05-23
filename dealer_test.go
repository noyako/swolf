package swolf

import (
	"testing"

	"github.com/golang/mock/gomock"
	mock "github.com/noyako/swolf/mock_swolf"
	"github.com/stretchr/testify/assert"
)

func TestDealerGet(t *testing.T) {
	assert := assert.New(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	id := "id"
	db := "id_db"

	mockMaser := mock.NewMockMasterController(mockCtrl)
	mockTenant := mock.NewMockTenantController(mockCtrl)
	mockPool := mock.NewMockPool(mockCtrl)

	mockPool.EXPECT().Get(id).Return(nil, false).Times(1)
	mockMaser.EXPECT().Get(id).Return(db, nil).Times(1)
	mockTenant.EXPECT().Get(db).Return(nil, nil).Times(1)
	mockPool.EXPECT().Put(id, nil).Return().Times(1)

	d := Dealer{mockMaser, mockTenant, mockPool}
	k, e := d.Get(id)
	assert.Nil(k)
	assert.Nil(e)
}

func TestDealerCreate(t *testing.T) {
	assert := assert.New(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	id := "id"
	db := "id_db"

	mockMaser := mock.NewMockMasterController(mockCtrl)
	mockTenant := mock.NewMockTenantController(mockCtrl)
	mockPool := mock.NewMockPool(mockCtrl)

	mockMaser.EXPECT().Create(id).Return(db, nil).Times(1)
	mockTenant.EXPECT().Create(db).Return(nil, nil).Times(1)

	d := Dealer{mockMaser, mockTenant, mockPool}
	k, e := d.Create(id)
	assert.Nil(k)
	assert.Nil(e)
}

func TestDealerDelete(t *testing.T) {
	assert := assert.New(t)
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	id := "id"
	db := "id_db"

	mockMaser := mock.NewMockMasterController(mockCtrl)
	mockTenant := mock.NewMockTenantController(mockCtrl)
	mockPool := mock.NewMockPool(mockCtrl)

	mockMaser.EXPECT().Delete(id).Return(db, nil).Times(1)
	mockTenant.EXPECT().Delete(db).Return(nil).Times(1)

	d := Dealer{mockMaser, mockTenant, mockPool}
	e := d.Delete(id)
	assert.Nil(e)
}
