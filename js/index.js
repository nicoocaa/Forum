document.addEventListener('DOMContentLoaded', function() {
  var postButton = document.getElementById('postButton');
  var closeButton = document.getElementById('closeButton'); 
  var modal = document.getElementById('modal');
  var modalEdit = document.getElementById('modalEdit');
  var fileUploadButton = document.getElementById('fileUploadButton');

  if (postButton) {
    postButton.addEventListener('click', function() {
        modal.style.display = 'flex';
    });
  }

  document.querySelectorAll('.btn-edit').forEach(function(btnEdit) {
    btnEdit.addEventListener('click', function(event) {
      event.preventDefault();
      modalEdit.style.display = 'flex';
    });
  });

  if (closeButton) {
    closeButton.addEventListener('click', function() {
        modal.style.display = 'none';
        modalEdit.style.display = 'none';
    });
  }

  if (modal) {
    modal.addEventListener('click', function(event) {
      var centerDiv = document.getElementById('center');
      if (event.target === modal) {
          modal.style.display = 'none';
      }
    });
  }

  if (fileUploadButton) {
    fileUploadButton.addEventListener('click', function() {
      document.getElementById('imageUpload').click();
    });
  }

  var inputPost = document.getElementById('input-post');
  if (inputPost) {
    inputPost.addEventListener('input', function() {
      var maxCaracteres = 280;
      var caracteresSaisis = this.value.length;
      var caracteresRestants = maxCaracteres - caracteresSaisis;

      var compteur = document.getElementById('compteur');
      if (compteur) {
        compteur.innerText = caracteresRestants + ' caractères restants';
      }
    });
  }
});
