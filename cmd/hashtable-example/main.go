package main

import (
	"log/slog"

	hashtable "github.com/Marlliton/temp/pkg/hash-table"
	"github.com/Marlliton/temp/pkg/logger"
)

func main() {
	logger.Init()
	slog.Info("--- Executando Exemplo de Tabela Hash ---")

	ht := hashtable.New()
	slog.Info("Tabela Hash vazia criada.")

	slog.Info("--- Testando Put ---")
	ht.Put("nome", "Marlliton")
	ht.Put("sobrenome", "Ferreira")
	ht.Put("cidade", "São Paulo")

	// NOTE: Este item deve colidir com "cidade" se o hash for simples
	ht.Put("idadec", "30")
	slog.Info("Adicionados 4 itens à tabela.")

	slog.Info("--- Testando Get ---")
	nome, ok := ht.Get("nome")
	slog.Info("Buscando 'nome'", "valor", nome, "encontrado", ok)

	cidade, ok := ht.Get("cidade")
	slog.Info("Buscando 'cidade'", "valor", cidade, "encontrado", ok)

	chaveInexistente, ok := ht.Get("pais")
	slog.Info("Buscando 'pais' (chave inexistente)", "valor", chaveInexistente, "encontrado", ok)

	slog.Info("--- Testando Delete ---")
	deleted := ht.Delete("sobrenome")
	slog.Info("Tentando deletar 'sobrenome'", "sucesso", deleted)

	sobrenome, ok := ht.Get("sobrenome")
	slog.Info("Buscando 'sobrenome' após deleção", "valor", sobrenome, "encontrado", ok)

	deleted = ht.Delete("pais")
	slog.Info("Tentando deletar 'pais' (chave inexistente)", "sucesso", deleted)

	slog.Info("--- Exemplo de Tabela Hash Finalizado ---")
}
