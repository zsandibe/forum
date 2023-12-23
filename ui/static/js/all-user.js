var allUsers = {{.}}
var usersContent = '<div class="container">';
// Комментарии описания в HTML
if (allUsers) {
  usersContent += `
    <table class="table">
      <thead>
        <tr>
          <th>Id</th>
          <th>Name</th>
          <th>Email</th>
          <th>Moderator Role Requested</th>
          <th>Type</th>
          <th>Make User</th>
          <th>Make Moderator</th>
        </tr>
      </thead>
      <tbody>`;

  allUsers.forEach(function(user) {
    usersContent += `
      <tr>
        <td>${user.ID}</td>
        <td>${user.Username}</td>
        <td>${user.Email}</td>
        <td>${user.Requested}</td>
        <td>${user.Role}</td>
        <td>`
         if(user.Role == "moderator"){
          usersContent += `
          <form class="userTypeChange" action="/v1/user/type/change" method="POST">
            <input type="hidden" name="user-id" value="${user.ID}" />
            <input type="hidden" name="type" value="user" />
            <button class="btn btn-primary btn-block">Make User</button>
          </form>`
          }
         usersContent += ` 
         </td>
        <td>`
          if(user.userType == "user"){
          usersContent += ` 
          <form class="userTypeChange" action="/v1/user/type/change" method="POST">
            <input type="hidden" name="user-id" value="${user.ID}" />
            <input type="hidden" name="type" value="moderator" />
            <button class="btn btn-primary btn-block">Make Moderator</button>
          </form>`
        }
         usersContent += ` 
         </td>
         </td>
      </tr>`;
  });

  usersContent += `
      </tbody>
    </table>`;
}
usersContent += '</div>';

document.getElementById("all-users").innerHTML = usersContent;