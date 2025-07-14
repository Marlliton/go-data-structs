package main

import (
	"log/slog"

	"github.com/Marlliton/temp/pkg/linkedlist"
	"github.com/Marlliton/temp/pkg/logger"
)

func main() {
	logger.Init()
	slog.Info("--- Executando Exemplo de Lista Ligada ---")

	// 1. Crie uma nova lista ligada
	ll := linkedlist.LinkedList[string]{}
	slog.Info("Inicializada uma lista ligada vazia.", "lista", ll.String())

	// 2. Adicione itens no início
	ll.Prepend("souza")
	ll.Prepend("marlliton")
	slog.Info("Adicionado 'marlliton' e 'souza' no início.", "lista", ll.String(), "tamanho", ll.Len())

	// 3. Adicione um item no final
	ll.Append("ferreira")
	slog.Info("Adicionado 'ferreira' no final.", "lista", ll.String(), "tamanho", ll.Len())

	// 4. Obtenha um item em um índice específico
	it, ok := ll.GetAt(1)
	if ok {
		slog.Info("Item recuperado no índice 1.", "item", it)
	}

	// 5. Insira um item em um índice específico
	ll.InsertAt("de", 2)
	slog.Info("Inserido 'de' no índice 2.", "lista", ll.String(), "tamanho", ll.Len())

	// 6. Remova um item em um índice específico
	removedItem, ok := ll.RemoveAt(1)
	if ok {
		slog.Info("Item removido do índice 1.", "item_removido", removedItem, "lista", ll.String(), "tamanho", ll.Len())
	}

	// 7. Encontre um item
	foundItem, ok := ll.Find(func(s string) bool { return s == "ferreira" })
	if ok {
		slog.Info("Encontrado o item 'ferreira'.", "item", foundItem)
	}

	// 8. Obtenha o índice de um item
	idx := ll.IndexOf(func(s string) bool { return s == "souza" })
	slog.Info("Índice de 'souza'.", "índice", idx)

	// 9. Tentando obter um item de um índice inválido
	_, ok = ll.GetAt(10)
	slog.Info("Tentando obter um item de um índice fora dos limites (10).", "sucesso", ok)

	slog.Info("--- Exemplo de Lista Ligada Finalizado ---")
}