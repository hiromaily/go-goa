<main>
  <user if={ tag==='user' } />
  <company if={ tag==='company' } />
  <tech if={ tag==='tech' } />
  <admin if={ tag==='admin' } />

  <script>
    var self = this
    self.data = {
      user: "User",
      company: "Company",
      tech: "Tech"
    }

    var r = riot.route.create()
    r('*', function(id) {
      console.log("id:",id)
      self.tag = id
      self.update()
    })
    r(function() {
      //default
      self.tag = "admin"
      self.update()
    })
  </script>

</main>