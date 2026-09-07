package repositorios

import (
	"DevBook/src/modelos"
	"database/sql"
)

// Publicacoes encapsula o acesso às publicações no banco de dados.
type Publicacoes struct {
	db *sql.DB
}

// NovoRepositorioDePublicacoes cria um repositório de publicações.
func NovoRepositorioDePublicacoes(db *sql.DB) *Publicacoes {
	return &Publicacoes{db}
}

// Criar insere uma publicação no banco de dados.
func (repositorio Publicacoes) Criar(publicacao modelos.Publicacao) (int64, error) {
	statement, erro := repositorio.db.Prepare(
		"INSERT INTO publicacoes (titulo, conteudo, autor_id) VALUES (?, ?, ?)")
	if erro != nil {
		return 0, erro
	}
	defer statement.Close()

	resultado, erro := statement.Exec(publicacao.Titulo, publicacao.Conteudo, publicacao.AutorID)
	if erro != nil {
		return 0, erro
	}

	ultimoIDInserido, erro := resultado.LastInsertId()
	if erro != nil {
		return 0, erro
	}
	return ultimoIDInserido, nil
}

// Buscar retorna as publicações de um usuário e de quem ele segue.
func (repositorio Publicacoes) Buscar(usuarioID uint64) ([]modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`SELECT p.*, u.nick
		 FROM publicacoes p
		 INNER JOIN usuarios u ON u.id = p.autor_id
		 WHERE p.autor_id = ? OR p.autor_id IN (
			SELECT usuario_id FROM seguidores WHERE seguidor_id = ?
		 )
		 ORDER BY p.criadaEm DESC`, usuarioID, usuarioID)
	if erro != nil {
		return nil, erro
	}
	defer linhas.Close()

	var publicacoes []modelos.Publicacao
	for linhas.Next() {
		var publicacao modelos.Publicacao
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick); erro != nil {
			return nil, erro
		}

		publicacoes = append(publicacoes, publicacao)
	}

	if erro = linhas.Err(); erro != nil {
		return nil, erro
	}

	return publicacoes, nil
}

// BuscarPorID busca uma publicação pelo seu ID.
func (repositorio Publicacoes) BuscarPorID(publicacaoID uint64) (modelos.Publicacao, error) {
	linhas, erro := repositorio.db.Query(
		`SELECT p.*, u.nick
		 FROM publicacoes p
		 INNER JOIN usuarios u ON u.id = p.autor_id
		 WHERE p.id = ?`, publicacaoID)
	if erro != nil {
		return modelos.Publicacao{}, erro
	}
	defer linhas.Close()

	var publicacao modelos.Publicacao
	if linhas.Next() {
		if erro = linhas.Scan(
			&publicacao.ID,
			&publicacao.Titulo,
			&publicacao.Conteudo,
			&publicacao.AutorID,
			&publicacao.Curtidas,
			&publicacao.CriadaEm,
			&publicacao.AutorNick); erro != nil {
			return modelos.Publicacao{}, erro
		}
	}

	return publicacao, nil
}
