<main>
  <admin if={ tag==='admin' } />
  <user if={ tag==='user' } />
  <user_detail if={ tag==='user_detail' } />
  <company if={ tag==='company' } />
  <tech if={ tag==='tech' } />


  <script>
    var self = this
    self.data = {
      user: {element:'user', url:'/api/user'},
      user_detail: {element:'user_detail', url:'/api/user/'},
      company: {element:'company', url:'/api/company'},
      tech: {element:'tech', url:'/api/tech'}
    }
    if (window.debugMode == 1){
        self.data.user.url='/json/userlist.json';
        self.data.user_detail.url='/json/userdetail.json';
        self.data.company.url='/json/companylist.json';
        self.data.tech.url='/json/techlist.json';
    }

    var r = riot.route.create()
    r('*', (collection) => {
      console.log("collection:", collection)
      self.tag = collection

      self.callAPI(self.data[collection])
    })
    r('*/*', (collection, id) => {
      console.log("collection:", collection)
      console.log("id:", id)
      self.tag = `${collection}_detail`
      self.user_id = id
      self.edit = true

      if (window.debugMode != 1){
        self.data[self.tag].url = '/api/user/'+id
      }

      self.callAPI(self.data[self.tag])
    })

    r(() => {
      //default
      self.tag = "admin"
      self.update()
    })

    self.callAPI = (obj) => {
        let key = sessionStorage.getItem('jwt');
        //fetch
        fetch(obj.url, {
            headers: {
                'Authorization': 'Bearer '+key,
                "Content-Type": "application/json"
            }
        }).then((response) => {
            return response.json();
        }).then((json) => {
            if(json.status && json.status != 200){
                sessionStorage.removeItem('jwt');
                sessionStorage.removeItem('id');
                location.href = '/admin/';
                return;
            }

            //Success
            //console.log("OK:", json)
            self.items = json

            self.update()
        });
    };

  </script>

</main>