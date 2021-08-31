$('#formulario-cadastro').on("submit", criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault()  // previni o comportamento default de carregar a pagina automatica do formulario

    if ($('#senha').val() != $('#confirmar-senha').val()) {
        alert('As senhas não conhecidem')
        return
    }

    // Criando uma requisição para ser enviado a API
    $.ajax({
        url: "/usuarios", // vai procurar dentro do app web uma rota com uri ou url '/usuarios'
        method: "POST",
        data : { // campo de dados que será enviados
            nome : $('#nome').val(),
            email : $('#email').val(),
            nick : $('#nick').val(),
            senha : $('#senha').val()
        }
    }).done(function(){ // executa quando o status da api retornar 200, 201, 204 status de sucesso
        alert("Usuario cadastrado com sucesso!")
    }).fail(function(erro) { // executa quando o status 400 404 401 403 500
        console.log(erro)
        alert("Erro ao cadastrar!")
    });
}