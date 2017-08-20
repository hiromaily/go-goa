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
    r('*', (id) => {
      self.header = self.data[id] || 'Admin'
      self.update()
    })
    r(() => {
      //default
      self.header = "Admin"
      self.update()
    })
  </script>
</header>