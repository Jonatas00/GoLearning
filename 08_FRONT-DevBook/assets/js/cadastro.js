$('#formulario-cadastro').on('submit', criarUsuario);

function criarUsuario(event) {
  event.preventDefault();
  
  if ($("#senha").val() != $("#confirmarSenha").val()) {
    Swal.fire("Ops...", "As senhas não cincidem!", "error");

    return
  }

  $.ajax({
    url: "/usuarios",
    method: "POST",
    data: {
      nome: $("#nome").val(),
      email: $("#email").val(),
      nick: $("#nick").val(),
      senha: $("#senha").val(),
    }
  }).done(() => {
    Swal.fire("Sucesso!", "Usuário cadastrado com sucesso", "success")
    .then(() => {
      $.ajax({
        url: "/login",
        method: "POST",
        data: {
          email: $("#email").val(),
          senha: $("#senha").val()
        }
      }).done(() => {
        window.location = "/home";
      }).fail(() => {
        Swal.fire("Ops...", "Erro ao cadastrar usuário", "error");
      })
    })
  }).fail((erro) => {
    Swal.fire("Ops...", "Erro ao cadastrar usuário!", "error");
  });
}
