<user>
<div class="ui container" style="margin-bottom: 50px;">
    <h3 class="ui header">User Model</h3>

    <table class="ui celled striped table">
        <thead>
        <tr>
            <th>ID</th>
            <th>Name</th>
            <th>E-mail address</th>
            <th>Delete</th>
        </tr>
        </thead>
        <tbody>
        <tr each="{ this.parent.items }">
            <td>{ id }</td>
            <td>{ user_name }</td>
            <td>{ email }</td>
            <td class="collapsing">
                <button class="ui button">Delete</button>
            </td>
        </tr>
        </tbody>
        <tfoot class="full-width">
        <tr>
            <th></th>
            <th>
                <div class="ui pagination menu">
                    <a class="item">1</a>
                    <div class="disabled item">...</div>
                    <a class="item active">10</a>
                    <a class="item">11</a>
                    <a class="item">12</a>
                </div>
            </th>
            <th colspan="2">
                <button class="ui blue active button" onclick="{ addUser }">
                  <i class="user icon"></i>
                  Add User
                </button>
            </th>
        </tr>
        </tfoot>
    </table>

</div>

<!-- add user modal window  -->
<div class="ui modal adduser">
  <div class="header">Add User</div>
  <div class="content">
    <form class="ui fluid form">
        <div class="ui two column middle aligned very relaxed stackable grid">
            <div class="column">
                <div class="field">
                    <label>Username</label>
                    <div class="ui left icon input">
                        <input name="username" type="text" placeholder="Username">
                        <i class="user icon"></i>
                        <!--
                        <div class="ui left pointing label">
                            That name is taken!
                        </div>
                        -->
                    </div>
                </div>

                <div class="field">
                    <label>Email</label>
                    <div class="ui left icon input">
                        <input name="email" type="text" placeholder="Email">
                        <i class="mail icon"></i>
                    </div>
                </div>

                <div class="field">
                    <label>Password</label>
                    <div class="ui left icon input">
                        <input name="password" type="text" placeholder="Password">
                        <i class="lock icon"></i>
                    </div>
                </div>
                <button class="ui blue active button" onclick="{ saveUser }">
                  Save
                </button>

                <div class="ui error message"></div>

            </div>
        </div>
    </form>
  </div>
</div>
<!-- End add user modal window  -->

<script>
  console.log("[user.tag]")

  this.on('mount', function() {
    // right after the tag is mounted on the page
    $('.ui.form')
        .form({
            fields: {
                username: {
                    identifier  : 'username',
                    rules: [
                        {
                            type   : 'empty',
                            prompt : 'Please enter your username'
                        },
                        {
                            type   : 'minLength[4]',
                            prompt : 'Your username must be at least 4 characters'
                        },
                        {
                            type   : 'maxLength[20]',
                            prompt : 'Your username must be at most 20 characters'
                        }
                    ]
                },
                email: {
                    identifier  : 'email',
                    rules: [
                        {
                            type   : 'empty',
                            prompt : 'Please enter your e-mail'
                        },
                        {
                            type   : 'email',
                            prompt : 'Please enter a valid e-mail'
                        },
                        {
                            type   : 'maxLength[50]',
                            prompt : 'Your e-mail must be at most 50 characters'
                        }
                    ]
                },
                password: {
                    identifier  : 'password',
                    rules: [
                        {
                            type   : 'empty',
                            prompt : 'Please enter your password'
                        },
                        {
                            type   : 'minLength[6]',
                            prompt : 'Your password must be at least 6 characters'
                        },
                        {
                            type   : 'maxLength[20]',
                            prompt : 'Your password must be at most 20 characters'
                        }
                    ]
                }
            },
            onSuccess: function(event, fields) {
                console.log('success')
                event.preventDefault();
            //},
            //onFailure: function(error, fields) {
            //    console.log('failure')
            //    event.preventDefault();
            }
        })
  })

  self = this

  //show modal for add user
  addUser(e) {
     //$('#userform').removeClass('hid')
     $('.ui.modal').modal('show')
  }

  //call user api
  saveUser(e) {
    if( $('.ui.form').form('is valid') ) {
      self.callAPI()
    }else{
      $('.ui.form').form('validate form')
    }
  }

  self.callAPI = function(){
    console.log('call API')

    let key = sessionStorage.getItem('jwt');
    let url = '/api/user';
    let payload = {
        username: $('input[name="username"]').val(),
        email: $('input[name="email"]').val(),
        password: $('input[name="password"]').val()
    };

    fetch(url, {
        method: 'POST',
        headers: {
            'Authorization': 'Bearer '+key,
            "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
    }).then(function(response) {
        return response.json();
    }).then(function(json) {
        console.log("res:", json);
        if ((json.status && json.status != 200) || !json.id) {
            //error
            console.log("error: ", json);

            $('.ui.form').form('add prompt', 'usrname');
            $('.ui.form .error.message').html('Please enter a valid value').show();
        }else{
            //success
            console.log("add OK");
            $('.ui.form .error.message').hide();

            //show popup or reload list??
        }
    });

    return
  }

</script>

</user>