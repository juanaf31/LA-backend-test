package configs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitDBWithoutConfig(t *testing.T) {
	_, err,_,_ := InitDB()
	if assert.NotNil(t,err){
		assert.Equal(t,"dial tcp: lookup yourhost: no such host",err.Error())
	}else{
		assert.Nil(t,err)
	}
}

func TestCreateRouter(t *testing.T) {
	router := CreatRouter()
	assert.NotNil(t, router)
}