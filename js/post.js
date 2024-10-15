document.addEventListener('DOMContentLoaded', function() {
  var modal = document.getElementById('modal');
  var modalEdit = document.getElementById('modalEdit');
  var modalEditPost = document.getElementById('modalEditPost');

  document.getElementById('postButton').addEventListener('click', function() {
    modal.style.display = 'flex';
  });

  document.querySelectorAll('.btn-edit').forEach(function(btn) {
    btn.addEventListener('click', function(event) {
      event.preventDefault();
      modalEdit.style.display = 'flex';
    });
  });

  document.querySelectorAll('.btnEditCom').forEach(function(btn) {
    btn.addEventListener('click', function(event) {
      event.preventDefault(); 
      modalEditPost.style.display = 'flex';
    });
  });
  
  document.querySelectorAll('.closeButton').forEach(function(btn) {
    btn.addEventListener('click', function() {
      document.querySelectorAll('.modal, .modalEdit, .modalEditPost').forEach(function(modal) {
        modal.style.display = 'none';
      });
    });
  });

  document.getElementById('fileUploadButton').addEventListener('click', function() {
    document.getElementById('imageUpload').click();
  });

  function closeModalOnClickOutside(modal) {
    modal.addEventListener('click', function(event) {
      var centerDiv = modal.querySelector('.center');
      if (!centerDiv.contains(event.target)) {
        modal.style.display = 'none';
      }
    });
  }

  var modals = document.querySelectorAll('.modal, .modalEdit, .modalEditPost');
  modals.forEach(closeModalOnClickOutside);

  function updateCharacterCounter(textarea, counterElement, maxLength) {
    var charactersLeft = maxLength - textarea.value.length;
    counterElement.textContent = charactersLeft + ' caractères restants';
  }

  document.querySelectorAll('.input-post').forEach(function(textarea) {
    var counterElement = textarea.closest('form').querySelector('#compteur');
    var maxLength = 280;
    
    updateCharacterCounter(textarea, counterElement, maxLength);

    textarea.addEventListener('input', function() {
      updateCharacterCounter(textarea, counterElement, maxLength);
    });
  });

  document.querySelectorAll('.btnEditCom').forEach(function(btn) {
    btn.addEventListener('click', function(event) {
      event.preventDefault();

      var commentId = this.getAttribute('value');
      var commentContent = this.closest('.comment').querySelector('.content').textContent;

      var textarea = modalEditPost.querySelector('.input-post');
      var submitButton = modalEditPost.querySelector('.add-post');

      textarea.value = commentContent;
      submitButton.value = commentId;
      submitButton.setAttribute('name', 'action');

      modalEditPost.style.display = 'flex';
    });
  });
});
