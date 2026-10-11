$('#cadastro').on('submit', criarUsuario);

function criarUsuario(evento) {
    evento.preventDefault();

    $.ajax({
        url: "/usuarios",
        method: "POST",
        contentType: "application/json",
        data: JSON.stringify({
            nome: $('#nome').val(),
            nick: $('#nick').val(),
            email: $('#email').val(),
            senha: $('#senha').val()
        })
    }).done(function () {
        alert("Usuário cadastrado com sucesso!");
        window.location = "/login";
    }).fail(function (erro) {
        var mensagem = erro.responseJSON && erro.responseJSON.erro;
        alert(mensagem || "Erro ao cadastrar o usuário!");
    });
}
