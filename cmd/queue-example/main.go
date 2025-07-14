package main

import (
	"log/slog"

	"github.com/Marlliton/temp/pkg/logger"
	"github.com/Marlliton/temp/pkg/queue"
)

func main() {
	logger.Init()
	q := queue.New("z@gmail.com")
	slog.Warn("Iniciando Enqueue", "itens iniciais", q.Items())
	q.Enqueue("c@mail.com")
	q.Enqueue("d@mail.com")

	it, ok := q.Peek()
	slog.Info("Função Peek", "primeiro elemento", it, "sucesso", ok)
	slog.Info("Status da fila", "Todos os itens", q.Items())

	slog.Warn("Testando a função Len", "Tamanho da fila", q.Len())
	slog.Warn("Testando a função IsEmpty", "Fila vazia", q.IsEmpty())

	slog.Warn("Iniciando Dequeue")
	q.Dequeue()
	q.Dequeue()
	slog.Info("Status da fila", "Depois do Dequeue()", q.Items())
	it, ok = q.Dequeue()
	slog.Info("Informações da removação", "Item", it, "Sucesso", ok)

	slog.Warn("Verificando se fila está fazia", "Fila vazia", q.IsEmpty())
}
