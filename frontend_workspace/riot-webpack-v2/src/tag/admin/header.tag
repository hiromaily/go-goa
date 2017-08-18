<header>
  <h2 class="ui center aligned header">{ header }</h2>

  <script>
    var self = this
    self.data = {
      user: "User",
      company: "Company",
      tech: "Tech"
    }

    var r = riot.route.create()
    r('*', function(id) {
      self.header = self.data[id] || 'Admin'
      self.update()
    })
    r(function() {
      //default
      self.header = "Admin"
      self.update()
    })
  </script>
</header>