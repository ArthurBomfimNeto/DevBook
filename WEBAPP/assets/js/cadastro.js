$('#formulario-cadastro').on("submit", criarUsuario)

function criarUsuario(evento) {
    evento.preventDefault()
    console.log("Dentro da funct ")

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
            senha : $('#senha').val(),
        }
    })
}