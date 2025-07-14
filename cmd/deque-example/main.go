package main

import (
	"log/slog"

	"github.com/Marlliton/temp/pkg/deque"
	"github.com/Marlliton/temp/pkg/logger"
)

func main() {
	logger.Init()
	d := deque.New("z@gmail.com")
	slog.Warn("Iniciando Deque", "itens iniciais", d.Items())
	d.AddFront("c@mail.com")
	d.AddBack("d@mail.com")

	it, ok := d.PeekFront()
	it, ok = d.PeekBack()
	slog.Info("Função PeekFront", "primeiro elemento", it, "sucesso", ok)
	slog.Info("Função PeekBack", "primeiro elemento", it, "sucesso", ok)
	slog.Info("Status do deque", "Todos os itens", d.Items())

	slog.Warn("Testando a função Len", "Tamanho do deque", d.Len())
	slog.Warn("Testando a função IsEmpty", "Deque vazio", d.IsEmpty())

	slog.Warn("Iniciando RemoveFront")
	d.RemoveFront()
	slog.Info("Status do deque", "Depois do RemoveFront()", d.Items())

	slog.Warn("Iniciando RemoveBack")
	item, ok := d.RemoveBack()
	slog.Info("Informações da remoção", "Item", item, "Sucesso", ok)
	slog.Info("Status do deque", "Depois do RemoveBack()", d.Items())

	it, ok = d.RemoveFront()
	slog.Info("Informações da remoção", "Item", it, "Sucesso", ok)

	slog.Warn("Verificando se deque está vazio", "Deque vazio", d.IsEmpty())
}
