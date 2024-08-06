$("#parar-de-seguir").on("click", pararDeSeguir);

function pararDeSeguir() {
  const usuarioId = $(this).data("usuario-id");
  $(this).prop("disabled", true);

  $.ajax({
    url: `/usuarios/${usuarioId}/parar-de-seguir`,
    method: "POST",
  }).done(() => {
    window.location = `/usuarios/${usuarioId}`
  }).fail(() => {
    Swal.fire("Ops...", "Erro ao parar de seguir o usuário", "error");
    $("#parar-de-seguir").prop("disabled", false);
  })

}

function seguir() {

}