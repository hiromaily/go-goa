<user>
<div class="ui container" style="margin-bottom: 50px;">
    <h3 class="ui header">User Model</h3>

    <table class="ui celled striped table">
        <thead>
        <tr>
            <th>ID</th>
            <th>Name</th>
            <th>E-mail address</th>
            <th>Update</th>
            <th>Delete</th>
        </tr>
        </thead>
        <tbody>
        <tr each="{ this.parent.items }">
            <td><a href="#user/{ id }">{ id }</a></td>
            <td>{ user_name }</td>
            <td>{ email }</td>
            <td class="collapsing">
                <button class="ui teal button" onclick="{ updateUser }">Update</button>
            </td>
            <td class="collapsing">
                <button class="ui red button" onclick="{ deleteUser }">Delete</button>
            </td>
        </tr>
        </tbody>
        <tfoot class="full-width">
        <tr>
            <th></th>
            <th colspan="2">
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
<div class="ui modal">
  <div id="modal_header" class="header">Add User</div>
  <div class="content">
    <form class="ui fluid form">
        <div class="ui column middle aligned very relaxed stackable grid">
            <div class="column">
                <div class="field">
                    <label>Username</label>
                    <div class="ui left icon input">
                        <input name="user_name" type="text" placeholder="Username">
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
                <button id="save_btn" class="ui blue active button" data-id="0" data-mode="add" onclick="{ saveUser }">
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
  //console.log("[user.tag]")

  this.on('mount', () => {
    // right after the tag is mounted on the page
    $('.ui.form')
        .form({
            fields: {
                user_name: {
                    identifier  : 'user_name',
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
            onSuccess: (event, fields) => {
                console.log('success')
                event.preventDefault();
            //},
            //onFailure: (error, fields) => {
            //    console.log('failure')
            //    event.preventDefault();
            }
        })
  })

  self = this

  //show modal for delete user
  deleteUser(e) {
    let res = confirm('Are you sure to delete this user?')
    if(!res) return

    let url = `/api/user/${e.item.id}`
    let payload = {}
    let fn = (json) => {
        if ((json.status && json.status != 200) || !json.id) {
            //error
            console.log("error: ", json)
            alert("error:", json)
        }else{
            //success
            console.log("delete OK")
            alert("deleted!")
            //refresh
            self.parent.callAPI({element:'user', url:'/api/user'})
        }
    }
    self.callAPI(url, payload, 'DELETE', fn)
  }

  //show modal for update user
  updateUser(e) {
    //TODO:update form
    $('input[name="user_name"]').val(e.item.user_name)
    $('input[name="email"]').val(e.item.email)
    $('input[name="password"]').val('**********')
    $('#modal_header').html('Update User')
    //$('#save_btn').data('mode','update')
    $('#save_btn').attr("data-mode", "update")
    $('#save_btn').attr("data-id", e.item.id)

    $('.ui.modal').modal('show')
  }

  //show modal for add user
  addUser(e) {
    $('input[name="user_name"]').val('')
    $('input[name="email"]').val('')
    $('input[name="password"]').val('')
    $('#modal_header').html('Add User')
    //$('#save_btn').data('mode','add')
    $('#save_btn').attr("data-mode", "add")

    $('.ui.modal').modal('show')
    //$('#userform').removeClass('hid')
  }

  //call user api (POST or PUT)
  saveUser(e) {
    //validate
    if(!$('.ui.form').form('is valid')){
      $('.ui.form').form('validate form')
      return
    }

    if (e.target.dataset.mode == "add"){
      //add
      let url = '/api/user';
      let payload = {
        user_name: $('input[name="user_name"]').val(),
        email: $('input[name="email"]').val(),
        password: $('input[name="password"]').val()
      };
      let fn = (json) => {
        if ((json.status && json.status != 200) || !json.id) {
            //error
            console.log("error: ", json)

            $('.ui.form').form('add prompt', 'user_name')
            $('.ui.form .error.message').html('Please enter a valid value').show()
        }else{
            //success
            console.log("add OK")
            $('.ui.form .error.message').hide()
            alert("added!")
            $('.ui.modal').modal('hide')
            //refresh
            self.parent.callAPI({element:'user', url:'/api/user'})
        }
      }

      self.callAPI(url, payload, 'POST', fn)
    }else{
      //update
      let url = `/api/user/${e.target.dataset.id}`
      let payload = {
        user_name: $('input[name="user_name"]').val(),
        email: $('input[name="email"]').val(),
        password: $('input[name="password"]').val()
      };
      let fn = (json) => {
        if ((json.status && json.status != 200) || !json.id) {
            //error
            console.log("error: ", json)

            $('.ui.form').form('add prompt', 'user_name')
            $('.ui.form .error.message').html('Please enter a valid value').show()
        }else{
            //success
            console.log("add OK");
            $('.ui.form .error.message').hide()
            alert("updated!")
            $('.ui.modal').modal('hide')
            //refresh
            self.parent.callAPI({element:'user', url:'/api/user'})
        }
      }

      self.callAPI(url, payload, 'PUT', fn)
    }

  }

  //common fetch API function
  self.callAPI = (url, payload, mtd, fn) => {
    let key = sessionStorage.getItem('jwt')

    fetch(url, {
        method: mtd,
        headers: {
            'Authorization': 'Bearer '+key,
            "Content-Type": "application/json"
        },
        body: JSON.stringify(payload)
    }).then((response) => {
        return response.json()
    }).then((json) => {
        console.log("res:", json)
        fn(json)
    });

    return
  }

</script>

</user>