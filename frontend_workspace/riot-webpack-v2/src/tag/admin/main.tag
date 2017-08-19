<main>
  <user if={ tag==='user' } />
  <company if={ tag==='company' } />
  <tech if={ tag==='tech' } />
  <admin if={ tag==='admin' } />

  <script>
    var self = this
    self.data = {
      user: {element:'user', url:'/api/user'},
      company: {element:'company', url:'/api/company'},
      tech: {element:'tech', url:'/api/company'}
    }
    if (window.debugMode == 1){
        self.data.user.url='/json/userlist.json';
        self.data.company.url='/json/companylist.json';
        self.data.tech.url='/json/techlist.json';
    }

    var r = riot.route.create()
    r('*', function(id) {
      //console.log("id:",id)
      self.tag = id

      //TODO:Ajax
      self.callAPI(self.data[id])
    })
    r(function() {
      //default
      self.tag = "admin"
      self.update()
    })

    self.callAPI = function(obj) {
        let key = sessionStorage.getItem('jwt');
        //fetch
        fetch(obj.url, {
            headers: {
                'Authorization': 'Bearer '+key,
                "Content-Type": "application/json"
            }
        }).then(function(response) {
            return response.json();
        }).then(function(json) {
            if(json.status && json.status != 200){
                sessionStorage.removeItem('jwt');
                sessionStorage.removeItem('id');
                location.href = '/admin.html';
                return;
            }

            //Success
            console.log("OK:", json)
            self.items = json

            self.update()
        });
    };

  </script>

</main>