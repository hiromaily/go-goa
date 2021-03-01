<tech>
<div class="ui container" style="margin-bottom: 50px;">
    <h3 class="ui header">Tech Model</h3>

    <table class="ui celled striped table">
        <thead>
        <tr>
            <th>ID</th>
            <th>Name</th>
            <th>Update</th>
            <th>Delete</th>
        </tr>
        </thead>
        <tbody>
        <tr each="{ this.parent.items }">
            <td>{ id }</td>
            <td>{ name }</td>
            <td class="collapsing">
                <button class="ui teal button" onclick="{ updateTech }">Update</button>
            </td>
            <td class="collapsing">
                <button class="ui red button" onclick="{ deleteTech }">Delete</button>
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
                <button class="ui blue active button" onclick="{ addTech }">
                  <i class="laptop icon"></i>
                  Add Tech
                </button>
            </th>
        </tr>
        </tfoot>
    </table>

</div>

<!-- add tech modal window  -->
<div class="ui modal">
  <div id="modal_header" class="header">Add Tech</div>
  <div class="content">
    <form class="ui fluid form">
        <div class="ui column middle aligned very relaxed stackable grid">
            <div class="column">
                <div class="field">
                    <label>Name</label>
                    <div class="ui left icon input">
                        <input name="name" type="text" placeholder="Name">
                        <i class="laptop icon"></i>
                    </div>
                </div>

                <button id="save_btn" class="ui blue active button" data-id="0" data-mode="add" onclick="{ saveTech }">
                  Save
                </button>

                <div class="ui error message"></div>

            </div>
        </div>
    </form>
  </div>
</div>
<!-- End add tech modal window  -->

<script>
  //console.log("[tech.tag]")

  this.on('mount', () => {
    // right after the tag is mounted on the page
    $('.ui.form')
        .form({
            fields: {
                name: {
                    identifier  : 'name',
                    rules: [
                        {
                            type   : 'empty',
                            prompt : 'Please enter tech name'
                        },
                        {
                            type   : 'minLength[1]',
                            prompt : 'Tech name must be at least 1 characters'
                        },
                        {
                            type   : 'maxLength[40]',
                            prompt : 'Tech name must be at most 40 characters'
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

  //show modal for delete tech
  deleteTech(e) {
    let res = confirm('Are you sure to delete this tech?')
    if(!res) return

    let url = `/api/tech/${e.item.id}`
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
            self.parent.callAPI({element:'tech', url:'/api/tech'})
        }
    }
    rg.callAPI(url, payload, 'DELETE', fn, null)
  }

  //show modal for update tech
  updateTech(e) {
    //TODO:update form
    $('input[name="name"]').val(e.item.name)
    $('#modal_header').html('Update Tech')
    //$('#save_btn').data('mode','update')
    $('#save_btn').attr("data-mode", "update")
    $('#save_btn').attr("data-id", e.item.id)

    $('.ui.modal').modal('show')
  }

  //show modal for add tech
  addTech(e) {
    $('input[name="name"]').val('')
    $('#modal_header').html('Add Tech')
    $('#save_btn').attr("data-mode", "add")

    $('.ui.modal').modal('show')
  }

  //call tech api (POST or PUT)
  saveTech(e) {
    //validate
    if(!$('.ui.form').form('is valid')){
      $('.ui.form').form('validate form')
      return
    }

    if (e.target.dataset.mode == "add"){
      //add
      let url = '/api/tech';
      let payload = {
        name: $('input[name="name"]').val()
      };
      let fn = (json) => {
        if ((json.status && json.status != 200) || !json.id) {
            //error
            console.log("error: ", json)

            $('.ui.form').form('add prompt', 'name')
            $('.ui.form .error.message').html('Please enter a valid value').show()
        }else{
            //success
            console.log("add OK")
            $('.ui.form .error.message').hide()
            alert("added!")
            $('.ui.modal').modal('hide')
            //refresh
            self.parent.callAPI({element:'tech', url:'/api/tech'})
        }
      }

      rg.callAPI(url, payload, 'POST', fn, null)
    }else{
      //update
      let url = `/api/tech/${e.target.dataset.id}`
      let payload = {
        name: $('input[name="name"]').val()
      };
      let fn = (json) => {
        if ((json.status && json.status != 200) || !json.id) {
            //error
            console.log("error: ", json)

            $('.ui.form').form('add prompt', 'name')
            $('.ui.form .error.message').html('Please enter a valid value').show()
        }else{
            //success
            console.log("add OK");
            $('.ui.form .error.message').hide()
            alert("updated!")
            $('.ui.modal').modal('hide')
            //refresh
            self.parent.callAPI({element:'tech', url:'/api/tech'})
        }
      }

      rg.callAPI(url, payload, 'PUT', fn, null)
    }

  }

</script>

</tech>