$('#login').on('submit', fazerLogin);

function fazerLogin(event) {
  event.preventDefault();

  $.ajax({
    url: "/login",
    method: "POST",
    data: {
      email: $('#email').val(),
      senha: $('#senha').val(),
    }
  }).done(function() {
    window.location = "/home";
  }).fail(function() {
    Swal.fire("Ops...", "Usuário ou senha inválidos!", "error");
  })
}
