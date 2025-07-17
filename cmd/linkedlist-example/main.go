package main

import (
	"log/slog"

	"github.com/Marlliton/temp/pkg/linkedlist"
	"github.com/Marlliton/temp/pkg/logger"
)

func main() {
	logger.Init()
	slog.Info("--- Executando Exemplo de Lista Ligada ---")

	ll := linkedlist.LinkedList[string]{}
	slog.Info("Inicializada uma lista ligada vazia.", "lista", ll.String())

	ll.Prepend("souza")
	ll.Prepend("marlliton")
	slog.Info("Adicionado 'marlliton' e 'souza' no início.", "lista", ll.String(), "tamanho", ll.Len())

	ll.Append("ferreira")
	slog.Info("Adicionado 'ferreira' no final.", "lista", ll.String(), "tamanho", ll.Len())

	it, ok := ll.GetAt(1)
	if ok {
		slog.Info("Item recuperado no índice 1.", "item", it)
	}

	ll.InsertAt("de", 2)
	slog.Info("Inserido 'de' no índice 2.", "lista", ll.String(), "tamanho", ll.Len())

	removedItem, ok := ll.RemoveAt(1)
	if ok {
		slog.Info("Item removido do índice 1.", "item_removido", removedItem, "lista", ll.String(), "tamanho", ll.Len())
	}

	foundItem, ok := ll.Find(func(s string) bool { return s == "ferreira" })
	if ok {
		slog.Info("Encontrado o item 'ferreira'.", "item", foundItem)
	}

	idx := ll.IndexOf(func(s string) bool { return s == "souza" })
	slog.Info("Índice de 'souza'.", "índice", idx)

	_, ok = ll.GetAt(10)
	slog.Info("Tentando obter um item de um índice fora dos limites (10).", "sucesso", ok)

	slog.Info("--- Exemplo de Lista Ligada Finalizado ---")
}
