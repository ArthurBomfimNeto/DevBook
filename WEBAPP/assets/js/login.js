$('#login').on('submit', fazerLogin);// quando clicar em submit ele executa o formulario e a função atrelada

function fazerLogin(evento) {
    evento.preventDefault(); // previni o comportamento default de carregar a pagina automatica do formulario

    $.ajax({
        url: "/login",
        method: "POST",
        data : {
            email: $('#email').val(),
            senha: $('#senha').val(),
        }
    }).done(function(){
        window.location= "/home" // assim que ele logar sera direcionado a pagina principal da rede
    }).fail(function(erro){
        alert('Usuario ou senha invalido!')
        console.log(erro)
    })
}

