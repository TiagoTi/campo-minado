// Code generated by MockGen. DO NOT EDIT.
// Source: ./aplicacao/dominio/sorteador.go

// Package mock_dominio is a generated GoMock package.
package mock_dominio

import (
	dominio "campo-minado/aplicacao/dominio"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockGeradorPosicoes is a mock of GeradorPosicoes interface.
type MockGeradorPosicoes struct {
	ctrl     *gomock.Controller
	recorder *MockGeradorPosicoesMockRecorder
}

// MockGeradorPosicoesMockRecorder is the mock recorder for MockGeradorPosicoes.
type MockGeradorPosicoesMockRecorder struct {
	mock *MockGeradorPosicoes
}

// NewMockGeradorPosicoes creates a new mock instance.
func NewMockGeradorPosicoes(ctrl *gomock.Controller) *MockGeradorPosicoes {
	mock := &MockGeradorPosicoes{ctrl: ctrl}
	mock.recorder = &MockGeradorPosicoesMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockGeradorPosicoes) EXPECT() *MockGeradorPosicoesMockRecorder {
	return m.recorder
}

// Tabuleiro mocks base method.
func (m *MockGeradorPosicoes) Tabuleiro(tamanho dominio.Posicao, qntBomba int, posicaoInicial dominio.Posicao) (dominio.Tabuleiro, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Tabuleiro", tamanho, qntBomba, posicaoInicial)
	ret0, _ := ret[0].(dominio.Tabuleiro)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Tabuleiro indicates an expected call of Tabuleiro.
func (mr *MockGeradorPosicoesMockRecorder) Tabuleiro(tamanho, qntBomba, posicaoInicial interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Tabuleiro", reflect.TypeOf((*MockGeradorPosicoes)(nil).Tabuleiro), tamanho, qntBomba, posicaoInicial)
}
